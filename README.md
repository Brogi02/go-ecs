#GO-ECS
A dead simple ECS system for golang


OVERVIEW

go-ecs has three main pieces:
- Entity: a UUID string, just an ID, no data of its own
- Component: a plain struct holding data, attached to an entity
- ECS_MANAGER: holds everything and gives you spawn/get/delete functions


1. INITIALIZING

````
gcs := ecs.Init_ecs_manager()
````

Returns a pointer to an ECS_MANAGER. Do this once at startup.


2. CREATING A COMPONENT

A component is just a struct that has a component() method (this method
can be empty, it only exists so the struct satisfies the COMPONENT
interface).
````
type Position struct {
    X, Y float64
}
func (c *Position) component() {}
````

3. SETTING THE COMPONENT CREATOR FUNCTION

The manager needs to know how to build a component from its name string.
Set this before spawning anything, or Spawn will panic.
````
gcs.create_component = func(ecs *ecs.ECS_MANAGER, component_name string) ecs.COMPONENT {
    switch component_name {
    case "POSITION":
        return &Position{}
    default:
        panic("unknown component: " + component_name)
    }
}
````

4. SPAWNING AN ENTITY
````
entity_id := gcs.Spawn([]string{"POSITION"})
````
Pass a list of component names. Returns the new entity's UUID as a
string. Save this, you need it for every later lookup or delete.


5. GETTING A COMPONENT

Copy (read only):

````
copy := gcs.Get_component(entity_id, "POSITION")

Pointer (to actually change the data):
pos := gcs.Get_component(entity_id, "POSITION").(*Position)
pos.X = 10

````
Note: Get_component panics if the component name was never registered,
or if the entity does not own that component.


6. DELETING AN ENTITY
````
gcs.Delete(entity_id)
````
Removes the entity and all of its components from every internal map.

NOTES / GOTCHAS

- Component names are plain strings, nothing checks them at compile
  time. Typos will panic at runtime. Consider using constants.
- create_component must be set before the first Spawn call.
- Get_component panics instead of returning nil/false, so guard calls
  you're not sure about, or make sure the entity actually has that
  component before asking for it.
- Always type-assert to a pointer (e.g. .(*Position)) when you need to
  modify the component, not just read it.
