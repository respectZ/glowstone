package entity

import (
	"fmt"

	glowstone "github.com/respectZ/glowstone/util"
)

func (e *Entity) Encode() ([]byte, error) {
	if e.Entity.Description.Animations.IsEmpty() {
		e.Entity.Description.Animations = nil
	}
	if e.Entity.Description.Properties.IsEmpty() {
		e.Entity.Description.Properties = nil
	}
	if e.Entity.Components.IsEmpty() {
		e.Entity.Components = nil
	}
	if e.Entity.ComponentGroups.IsEmpty() {
		e.Entity.ComponentGroups = nil
	}
	if e.Entity.Events != nil {
		events := e.Entity.Events.All()
		for _, event := range events {
			if event.Add != nil && event.Add.ComponentGroups.IsEmpty() {
				event.Add = nil
			}
			if event.Remove != nil && event.Remove.ComponentGroups.IsEmpty() {
				event.Remove = nil
			}
			if event.SetProperty != nil && event.SetProperty.IsEmpty() {
				event.SetProperty = nil
			}
		}
	}
	return glowstone.MarshalJSON(e)
}

func (e *Entity) GetIdentifier() string {
	return e.Entity.Description.Identifier
}

func (e *Entity) GetRuntimeIdentifier() (string, error) {
	if e.Entity.Description.RuntimeIdentifier != "" {
		return e.Entity.Description.RuntimeIdentifier, nil
	}
	return "", fmt.Errorf("runtime identifier not found")
}

//	func (e *Entity) AddComponent(components ...interface{}) {
//		if e.Entity.Components == nil {
//			e.Entity.Components = make(map[string]IComponent)
//		}
//		for _, component := range components {
//			componentName := util_component.GetComponentName(component)
//			c := reflect.TypeOf(component)
//			if c.Kind() == reflect.String && c.Name() == "string" {
//				e.Entity.Components[componentName] = struct{}{}
//			} else {
//				e.Entity.Components[componentName] = component
//			}
//		}
//	}
