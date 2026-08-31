package ecs

import "reflect"

type SYSTEM_HANDLER struct {
	Systems map[string]SYSTEM
}

func INIT_SYSTEM_MANAGER() *SYSTEM_HANDLER {
	return &SYSTEM_HANDLER{
		Systems: make(map[string]SYSTEM),
	}
}

type SYSTEM interface {
	System()
	Start()
	Update()
}

func (sys_handler *SYSTEM_HANDLER) PushUpdate() {
	for _, each_system := range sys_handler.Systems {
		each_system.Update()
	}
}
func (sys_handler *SYSTEM_HANDLER) PushStart() {
	for _, each_system := range sys_handler.Systems {
		each_system.Start()
	}
}
func (sys_handler *SYSTEM_HANDLER) PushSystem() {
	for _, each_system := range sys_handler.Systems {
		each_system.System()
	}
}
func (sys_handler *SYSTEM_HANDLER) InitSystem(system SYSTEM) {
	sys_handler.Systems[reflect.TypeOf(system).Name()] = system
}

/*
type physics struct{}
var Physics = &physics{}

func (physics) System()
func (physics) Start()
func (physics) Update()
*/
