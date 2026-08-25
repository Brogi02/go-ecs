package ecs

import "github.com/google/uuid"

type ECS_MANAGER struct {
	all_objects      map[string]map[string]COMPONENT
	create_component ComponentCreateFunction
}

type ComponentCreateFunction func(*ECS_MANAGER, string) COMPONENT

func Init_ecs_manager() *ECS_MANAGER {
	return &ECS_MANAGER{
		all_objects: make(map[string]map[string]COMPONENT),
	}
}
func (ecs *ECS_MANAGER) Get_component(object_uuid string, component string) COMPONENT {
	// IF A POINTER TO A COMPONENT IS NEEDED .(*COMPONENT_STRUCT_IDENTIFIER) is necessary to be appended to the function call
	component_map := ecs.all_objects[component]
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
	for _, each_map := range ecs.all_objects {
		delete(each_map, object_uuid)
	}
}
func (ecs *ECS_MANAGER) Spawn(components_list_to_create []string) string {
	uuid_created_for_the_entity := uuid.New().String()
	for _, each_comcomponent_to_create := range components_list_to_create {
		component_map := ecs.all_objects[each_comcomponent_to_create]

		if component_map == nil {
			component_map = make(map[string]COMPONENT)
			ecs.all_objects[each_comcomponent_to_create] = component_map
		}

		component_map[uuid_created_for_the_entity] = ecs.create_component(ecs, each_comcomponent_to_create)
	}
	return uuid_created_for_the_entity
}

type COMPONENT interface {
	component()
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
