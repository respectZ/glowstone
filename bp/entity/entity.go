package entity

import (
	"fmt"
	"strings"

	glowstone "github.com/respectZ/glowstone/util"
)

func (e *entity) Encode() ([]byte, error) {
	return glowstone.MarshalJSON(e)
}

func (e *entity) GetIdentifier() string {
	return strings.Split(e.Entity.Description.Identifier, ":")[1]
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

// TODO: Rework parameter to the struct, so we can cast it.
//
//	func (e *entity) GetComponent(name string) (interface{}, error) {
//		if component, ok := e.Entity.Components[name]; ok {
//			return component, nil
//		}
//		return nil, fmt.Errorf("component %s not found", name)
//	}
func (e *entity) GetComponent(name interface{}) (interface{}, error) {
	if component, ok := e.Entity.Components[GetComponentName(name)]; ok {
		// Convert map to struct
		convertMapToStruct(component.(map[string]interface{}), name)
		// Assign component to the struct
		e.Entity.Components[GetComponentName(name)] = name

		return component, nil
	}
	return nil, fmt.Errorf("component %s not found", name)
}

func (e *entity) GetComponents() []interface{} {
	components := make([]interface{}, 0)
	for component := range e.Entity.Components {
		components = append(components, component)
	}
	return components
}

func (e *entity) AddComponent(components ...interface{}) {
	for _, component := range components {
		e.Entity.Components[GetComponentName(component)] = component
	}
}

func (e *entity) RemoveComponent(name string) {
	delete(e.Entity.Components, name)
}

func (e *entity) GetEvent(name string) (EntityEvent, error) {
	if event, ok := e.Entity.Events[name]; ok {
		return event, nil
	}
	return EntityEvent{}, fmt.Errorf("event %s not found", name)
}

func (e *entity) GetEvents() map[string]EntityEvent {
	return e.Entity.Events
}

func (e *entity) SetEvent(name string, event EntityEvent) {
	e.Entity.Events[name] = event
}

func (e *entity) RemoveEvent(name string) {
	delete(e.Entity.Events, name)
}

//** Component group *//
