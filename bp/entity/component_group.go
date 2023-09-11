package entity

func (c *cg) GetComponent(name interface{}) {
	value, exists := (*c)[GetComponentName(name)]
	if !exists {
		return
	}
	// Convert map to struct
	convertMapToStruct(value.(map[string]interface{}), name)
	// Assign component to the struct
	(*c)[GetComponentName(name)] = name
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
		(*c)[GetComponentName(component)] = component
	}
}

func (c *cg) RemoveComponent(name string) {
	delete(*c, name)
}
