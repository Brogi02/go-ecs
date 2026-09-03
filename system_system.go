package ecs

type SYSTEMS_MANAGER struct {
	linked_ecs *ECS_MANAGER
	systems    map[string]map[string]SYSTEM_ENTRY
}

type SYSTEM func(*ECS_MANAGER) int

type SYSTEM_ENTRY struct {
	system  SYSTEM
	enabled bool
}

func INIT_SYSTEMS_MANAGER() *SYSTEMS_MANAGER {
	return &SYSTEMS_MANAGER{
		systems: make(map[string]map[string]SYSTEM_ENTRY),
	}
}
func (sysman *SYSTEMS_MANAGER) SetECS_MANAGER(ecs *ECS_MANAGER) {
	sysman.linked_ecs = ecs
}
func (sysman *SYSTEMS_MANAGER) AddSystem(SYSTEM_type string, SYSTEM_id string, SYSTEM_func_to_append SYSTEM) {
	system_entry_for_append := SYSTEM_ENTRY{system: SYSTEM_func_to_append, enabled: true}
	_, sysman_inner_map_for_id_exist := sysman.systems[SYSTEM_type][SYSTEM_id]
	if !sysman_inner_map_for_id_exist {
		sysman.systems[SYSTEM_type] = make(map[string]SYSTEM_ENTRY)
	}
	sysman.systems[SYSTEM_type][SYSTEM_id] = system_entry_for_append
}
func (sysman *SYSTEMS_MANAGER) CallSystemByID(SYSTEM_type string, SYSTEM_id string) {
	if SYSTEM_type != "" {
		if sysman.systems[SYSTEM_type][SYSTEM_id].enabled == true {
			func_to_execute := sysman.systems[SYSTEM_type][SYSTEM_id].system
			return_value_check := func_to_execute(sysman.linked_ecs)
			if return_value_check != 0 {
				panic("Err: CallSystemSingle had an error at runtime function execution : " + SYSTEM_id + string(return_value_check))
			}
		}
	} else if SYSTEM_type == "" {
		for _, each_system_type_map := range sysman.systems {
			for key_system_id, each_system_entry := range each_system_type_map {
				if key_system_id == SYSTEM_id {
					if each_system_entry.enabled == true {
						func_to_execute := each_system_entry.system
						return_value_check := func_to_execute(sysman.linked_ecs)
						if return_value_check != 0 {
							panic("Err: CallSystemSingle had an error at runtime function execution : " + SYSTEM_id + string(return_value_check))
						}
					}
				}
			}
		}
	}
}
func (sysman *SYSTEMS_MANAGER) CallSystemsByType(SYSTEM_type string) {
	for key_systems_type_map, each_system_type_map := range sysman.systems {
		if key_systems_type_map == SYSTEM_type {
			for index, each_system_entry_in_type_map := range each_system_type_map {
				if each_system_entry_in_type_map.enabled == true {
					func_to_execute := each_system_entry_in_type_map.system
					return_value_check := func_to_execute(sysman.linked_ecs)
					if return_value_check != 0 {
						panic("Err: CallSystemSingle had an error at runtime function execution : " + string(index) + string(return_value_check))
					}
				}
			}
		}
	}
}
