package entity

import (
	"fmt"
	"reflect"

	g_util "github.com/respectZ/glowstone/util"
	util_component "github.com/respectZ/glowstone/util/component"
)

type ComponentGroups map[string]Component

type IComponentGroups interface {
	UnmarshalJSON([]byte) error

	// Add a component to specific component group.
	//
	// - Example:
	//
	//    var instantDespawn component.InstantDespawn
	//    entity.Entity.ComponentGroups.Add("despawn", &instantDespawn)
	Add(string, ...interface{}) error

	// Get a component group, returning a map of components.
	//
	// - Example:
	//
	//    components, ok := entity.Entity.ComponentGroups.Get("despawn")
	Get(string) (IComponent, bool)

	// Get a component from specific component group by struct.
	//
	// - Example:
	//
	//    var instantDespawn component.InstantDespawn
	//    instantDespawnInterface, err := entity.Entity.ComponentGroups.GetComponent("despawn", &instantDespawn)
	GetComponent(string, interface{}) (interface{}, error)
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]Component
}

func (m *ComponentGroups) UnmarshalJSON(data []byte) error {
	var temp map[string]Component
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*m = temp
	return nil
}

func (m *ComponentGroups) Add(key string, components ...interface{}) error {
	if *m == nil {
		*m = make(ComponentGroups)
	}
	temp := &map[string]interface{}{}

	for _, component := range components {
		componentName := util_component.GetComponentName(component)
		typeOf := reflect.TypeOf(component)
		if typeOf.Kind() == reflect.String && typeOf.Name() == "string" {
			(*temp)[componentName] = struct{}{}
		} else {
			(*temp)[componentName] = component
		}
	}

	c := Component(*temp)

	(*m)[key] = c

	return nil
}

func (m *ComponentGroups) Get(name string) (IComponent, bool) {
	if *m == nil {
		*m = make(ComponentGroups)
	}
	component, ok := (*m)[name]
	return &component, ok
}

func (m *ComponentGroups) GetComponent(key string, new interface{}) (interface{}, error) {
	if *m == nil {
		*m = make(ComponentGroups)
	}
	oldComponent, ok := (*m)[key]
	if !ok {
		return nil, fmt.Errorf("component group %s not found", key)
	}

	if reflect.TypeOf(new).Kind() == reflect.String {
		if component, ok := oldComponent[new.(string)]; ok {
			return component, nil
		}
	}

	componentName := util_component.GetComponentName(new)
	if component, ok := oldComponent[componentName]; ok {
		r, err := util_component.Get(component, new)
		if err != nil {
			return nil, err
		}
		// Assign component to the struct
		oldComponent[componentName] = r

		return component, nil
	}
	return nil, fmt.Errorf("component %s not found", new)
}

func (m *ComponentGroups) Clear() {
	if *m == nil {
		*m = make(ComponentGroups)
	}
	*m = make(ComponentGroups)
}

func (m *ComponentGroups) Remove(name string) {
	if *m == nil {
		*m = make(ComponentGroups)
	}
	delete(*m, name)
}

func (m *ComponentGroups) IsEmpty() bool {
	if *m == nil {
		*m = make(ComponentGroups)
	}
	return len(*m) == 0
}

func (m *ComponentGroups) Size() int {
	if *m == nil {
		*m = make(ComponentGroups)
	}
	return len(*m)
}

func (m *ComponentGroups) All() map[string]Component {
	if *m == nil {
		*m = make(ComponentGroups)
	}
	return *m
}
