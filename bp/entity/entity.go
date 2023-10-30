package entity

import (
	"fmt"
	"reflect"

	glowstone "github.com/respectZ/glowstone/util"
	util_component "github.com/respectZ/glowstone/util/component"
)

func (e *entity) Encode() ([]byte, error) {
	return glowstone.MarshalJSON(e)
}

func (e *entity) GetIdentifier() string {
	return e.Entity.Description.Identifier
}

func (e *entity) GetRuntimeIdentifier() (string, error) {
	if e.Entity.Description.RuntimeIdentifier != "" {
		return e.Entity.Description.RuntimeIdentifier, nil
	}
	return "", fmt.Errorf("runtime identifier not found")
}

func (e *entity) SetRuntimeIdentifier(runtimeIdentifier string) {
	e.Entity.Description.RuntimeIdentifier = runtimeIdentifier
}

func (e *entity) GetSpawnable() bool {
	return e.Entity.Description.IsSpawnable
}

func (e *entity) SetSpawnable(spawnable bool) {
	e.Entity.Description.IsSpawnable = spawnable
}

func (e *entity) GetSummonable() bool {
	return e.Entity.Description.IsSummonable
}

func (e *entity) SetSummonable(summonable bool) {
	e.Entity.Description.IsSummonable = summonable
}

func (e *entity) GetAnimation(name string) (string, error) {
	if animation, ok := e.Entity.Description.Animations[name]; ok {
		return animation, nil
	}
	return "", fmt.Errorf("animation %s not found", name)
}

func (e *entity) GetAnimations() map[string]string {
	return e.Entity.Description.Animations
}

func (e *entity) SetAnimation(name string, animation string) {
	e.Entity.Description.Animations[name] = animation
}

func (e *entity) RemoveAnimation(name string) {
	delete(e.Entity.Description.Animations, name)
}

func (e *entity) GetProperty(name string) (EntityProperties, error) {
	if prop, ok := e.Entity.Description.Properties[name]; ok {
		return prop, nil
	}
	return EntityProperties{}, fmt.Errorf("property %s not found", name)
}

func (e *entity) GetProperties() map[string]EntityProperties {
	return e.Entity.Description.Properties
}

func (e *entity) SetProperty(name string, prop EntityProperties) {
	if e.Entity.Description.Properties == nil {
		e.Entity.Description.Properties = make(map[string]EntityProperties)
	}
	e.Entity.Description.Properties[name] = prop
}

func (e *entity) GetComponentGroup(name string) (component_group, error) {
	if group, ok := e.Entity.ComponentGroups[name]; ok {
		// Check if it's a map[string]interface{}
		v, ok := group.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("component group %s is not a map[string]interface{}", name)
		}
		g := cg(v)
		return &g, nil
	}
	return nil, fmt.Errorf("component group %s not found", name)
}

func (e *entity) GetComponentGroups() []string {
	groups := make([]string, 0)
	for group := range e.Entity.ComponentGroups {
		groups = append(groups, group)
	}
	return groups
}

func (e *entity) SetComponentGroup(name string, group ComponentGroup) {
	e.Entity.ComponentGroups[name] = group
}

func (e *entity) RemoveComponentGroup(name string) {
	delete(e.Entity.ComponentGroups, name)
}

func (e *entity) AddComponentGroup(name string, components ...interface{}) {
	if e.Entity.ComponentGroups == nil {
		e.Entity.ComponentGroups = make(map[string]interface{})
	}
	if _, ok := e.Entity.ComponentGroups[name]; !ok {
		e.Entity.ComponentGroups[name] = make(map[string]interface{})
	}
	for _, component := range components {
		componentName := util_component.GetComponentName(component)
		c := reflect.TypeOf(component)
		if c.Kind() == reflect.String && c.Name() == "string" {
			e.Entity.ComponentGroups[name].(map[string]interface{})[componentName] = struct{}{}
		} else {
			e.Entity.ComponentGroups[name].(map[string]interface{})[componentName] = component
		}
	}
}

func (e *entity) GetComponent(name interface{}) (interface{}, error) {
	componentName := util_component.GetComponentName(name)
	if component, ok := e.Entity.Components[componentName]; ok {
		r, err := util_component.Get(component, name)
		if err != nil {
			return nil, err
		}
		// Assign component to the struct
		e.Entity.Components[componentName] = r

		return component, nil
	}
	return nil, fmt.Errorf("component %s not found", name)
}

func (e *entity) GetComponents() map[string]interface{} {
	return e.Entity.Components
}

func (e *entity) AddComponent(components ...interface{}) {
	if e.Entity.Components == nil {
		e.Entity.Components = make(map[string]interface{})
	}
	for _, component := range components {
		componentName := util_component.GetComponentName(component)
		c := reflect.TypeOf(component)
		if c.Kind() == reflect.String && c.Name() == "string" {
			e.Entity.Components[componentName] = struct{}{}
		} else {
			e.Entity.Components[componentName] = component
		}
	}
}

func (e *entity) RemoveComponent(name string) {
	delete(e.Entity.Components, name)
}

func (e *entity) GetEvent(name string) (*EntityEvent, error) {
	if event, ok := e.Entity.Events[name]; ok {
		return event, nil
	}
	return nil, fmt.Errorf("event %s not found", name)
}

func (e *entity) GetEvents() map[string]*EntityEvent {
	return e.Entity.Events
}

func (e *entity) SetEvent(name string, event EntityEvent) {
	if e.Entity.Events == nil {
		e.Entity.Events = make(map[string]*EntityEvent)
	}
	e.Entity.Events[name] = &event
}

func (e *entity) RemoveEvent(name string) {
	delete(e.Entity.Events, name)
}

//** Component group *//
