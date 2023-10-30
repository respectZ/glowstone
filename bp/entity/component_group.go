package entity

import (
	"fmt"

	util_component "github.com/respectZ/glowstone/util/component"
)

func (c *cg) GetComponent(name interface{}) (interface{}, error) {
	componentName := util_component.GetComponentName(name)
	component, exists := (*c)[componentName]
	if !exists {
		return nil, fmt.Errorf("component %s not found", name)
	}
	r, err := util_component.Get(component, name)
	if err != nil {
		return nil, err
	}
	// Assign component to the struct
	(*c)[componentName] = r
	return component, nil
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
