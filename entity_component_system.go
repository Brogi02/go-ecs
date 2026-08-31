package ecs

import "github.com/google/uuid"

var PresetComponents PresetComponentsFunction = func(component_name string) COMPONENT {
	return nil
}

type ECS_MANAGER struct {
	objects          map[string]map[string]COMPONENT
	create_component ComponentCreateFunction
}
type COMPONENT interface {
	component()
}
type ComponentCreateFunction func(*ECS_MANAGER, string) COMPONENT

type PresetComponentsFunction func(string) COMPONENT

func INIT_ECS_MANAGER() *ECS_MANAGER {
	return &ECS_MANAGER{
		objects: make(map[string]map[string]COMPONENT),
	}
}
func (ecs *ECS_MANAGER) GetComponent(object_uuid string, component string) COMPONENT {
	// IF A POINTER TO A COMPONENT IS NEEDED .(*COMPONENT_STRUCT_IDENTIFIER) is necessary to be appended to the function call
	component_map := ecs.objects[component]
	if component_map == nil {
		panic("Component has not been defined or has no map assigned")
	}
	return_value := component_map[object_uuid]
	if return_value != nil {
		return return_value
	} else {
		panic("tried reaching non existing component(NON OWNED or NON EXISTING)")
	}
}
func (ecs *ECS_MANAGER) Delete(object_uuid string) {
	for _, each_map := range ecs.objects {
		delete(each_map, object_uuid)
	}
}
func (ecs *ECS_MANAGER) Spawn(components_list_to_create []string) string {
	uuid_created_for_the_entity := uuid.New().String()
	for _, each_comcomponent_to_create := range components_list_to_create {
		component_map := ecs.objects[each_comcomponent_to_create]

		if component_map == nil {
			component_map = make(map[string]COMPONENT)
			ecs.objects[each_comcomponent_to_create] = component_map
		}
		PresetComponents_return := PresetComponents(each_comcomponent_to_create)
		if PresetComponents_return != nil {
			component_map[uuid_created_for_the_entity] = PresetComponents_return
		} else {
			component_map[uuid_created_for_the_entity] = ecs.create_component(ecs, each_comcomponent_to_create)
		}
	}
	return uuid_created_for_the_entity
}

func (ecs *ECS_MANAGER) GetEntityWithComponents(components []string) []string {
	if len(components) == 0 {
		return []string{}
	}
	common_entities := make(map[string]struct{})
	first_component_map := ecs.objects[components[0]]
	if first_component_map == nil {
		return []string{}
	}
	for entity_uuid := range first_component_map {
		common_entities[entity_uuid] = struct{}{}
	}
	for _, each_component := range components[1:] {
		component_map := ecs.objects[each_component]
		if component_map == nil {
			return []string{}
		}
		for entity_uuid := range common_entities {
			if _, exists := component_map[entity_uuid]; !exists {
				delete(common_entities, entity_uuid)
			}
		}
	}
	entitys := make([]string, 0, len(common_entities))
	for entity_uuid := range common_entities {
		entitys = append(entitys, entity_uuid)
	}
	return entitys
}

/*
	EXAMPLE:


	ecs := init_ecs_manager()
	ecs.create_component = func(ecs *ECS_MANAGER, component_name string) COMPONENT {
		switch component_name {
		case "ENTITY":
			return &ENTITY{}

		case "VELOCITY":
			return &VELOCITY{}

		default:
			panic("unknown component: " + component_name)
		}
	}
*/
