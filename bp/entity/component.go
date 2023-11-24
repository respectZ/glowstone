package entity

import (
	"fmt"
	"reflect"

	util "github.com/respectZ/glowstone/util"
	util_component "github.com/respectZ/glowstone/util/component"
)

type Component map[string]interface{}

type IComponent interface {
	UnmarshalJSON([]byte) error

	// Add a component to entity.
	//
	// - Example:
	//
	//    var instantDespawn component.InstantDespawn
	//    entity.Entity.Components.Add(&instantDespawn)
	Add(...interface{}) error
	// Get a component
	//
	// - Example:
	//
	//    var instantDespawn component.InstantDespawn
	//    instantDespawnInterface, err := entity.GetComponent(&instantDespawn)
	Get(interface{}) (interface{}, error)
	Set(interface{})
	Clear()
	Remove(string)
	IsEmpty() bool
}

func (c *Component) UnmarshalJSON(data []byte) error {
	var temp map[string]interface{}
	if err := util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*c = temp
	return nil
}

func (c *Component) Add(components ...interface{}) error {
	if (*c) == nil {
		(*c) = make(map[string]interface{})
	}
	for _, component := range components {
		componentName := util_component.GetComponentName(component)
		(*c)[componentName] = component
		typeOf := reflect.TypeOf(component)
		if typeOf.Kind() == reflect.String && typeOf.Name() == "string" {
			(*c)[componentName] = struct{}{}
		} else {
			(*c)[componentName] = component
		}
	}
	return nil
}

func (c *Component) Get(new interface{}) (interface{}, error) {
	if reflect.TypeOf(new).Kind() == reflect.String {
		if component, ok := (*c)[new.(string)]; ok {
			return component, nil
		}
	}
	componentName := util_component.GetComponentName(new)
	if component, ok := (*c)[componentName]; ok {
		r, err := util_component.Get(component, new)
		if err != nil {
			return nil, err
		}
		// Assign component to the struct
		(*c)[componentName] = r

		return component, nil
	}
	return nil, fmt.Errorf("component %s not found", new)
}

func (c *Component) Set(component interface{}) {
	componentName := util_component.GetComponentName(component)
	(*c)[componentName] = component
}

func (c *Component) Clear() {
	if *c == nil {
		*c = make(map[string]interface{})
	}
	*c = make(map[string]interface{})
}

func (c *Component) Remove(name string) {
	delete(*c, name)
}

func (c *Component) IsEmpty() bool {
	if *c == nil {
		*c = make(map[string]interface{})
	}
	return len(*c) == 0
}
