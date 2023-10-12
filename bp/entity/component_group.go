package entity

import (
	util_component "github.com/respectZ/glowstone/util/component"
)

func (c *cg) GetComponent(name interface{}) {
	componentName := util_component.GetComponentName(name)
	value, exists := (*c)[componentName]
	if !exists {
		return
	}
	// Convert map to struct
	util_component.ConvertMapToStruct(value.(map[string]interface{}), name)
	// Assign component to the struct
	(*c)[componentName] = name
}

func (c *cg) GetComponents() []interface{} {
	components := make([]interface{}, 0)
	for component := range *c {
		components = append(components, component)
	}
	return components
}

func (c *cg) AddComponent(components ...interface{}) {
	for _, component := range components {
		componentName := util_component.GetComponentName(component)
		(*c)[componentName] = component
	}
}

func (c *cg) RemoveComponent(name string) {
	delete(*c, name)
}
