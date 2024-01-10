package entity

import (
	"fmt"
	"reflect"

	g_util "github.com/respectZ/glowstone/util"
)

func checkEvents(event *EntityEvent) {
	if event.Add != nil && event.Add.ComponentGroups.IsEmpty() {
		event.Add = nil
	}
	if event.Remove != nil && event.Remove.ComponentGroups.IsEmpty() {
		event.Remove = nil
	}
	if event.SetProperty != nil && event.SetProperty.IsEmpty() {
		event.SetProperty = nil
	}
	if event.Randomize != nil {
		if event.Randomize.IsEmpty() {
			event.Randomize = nil
		} else {
			events := event.Randomize.All()
			for _, event := range events {
				checkEvents(event)
			}
		}
	}
	if event.Sequence != nil {
		if event.Sequence.IsEmpty() {
			event.Sequence = nil
		} else {
			events := event.Sequence.All()
			for _, event := range events {
				checkEvents(event)
			}
		}
	}
}

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
			checkEvents(event)
		}
		if e.Entity.Events.IsEmpty() {
			e.Entity.Events = nil
		}
	}
	if e.Entity.Description.Scripts != nil {
		if e.Entity.Description.Scripts.Animate.IsEmpty() {
			e.Entity.Description.Scripts.Animate = nil
		}
	}
	if reflect.ValueOf(e.Entity.Description.Scripts).Elem().IsZero() {
		e.Entity.Description.Scripts = nil
	}
	return g_util.MarshalJSON(e)
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
