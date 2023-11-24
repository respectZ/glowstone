package entity

import (
	g_util "github.com/respectZ/glowstone/util"
)

type EntityProperties map[string]*entityProperties

type entityProperties struct {
	ClientSync bool        `json:"client_sync"`
	Type       string      `json:"type"`
	Range      []int       `json:"range,omitempty"`  // only for float
	Values     []string    `json:"values,omitempty"` // only for enum
	Default    interface{} `json:"default"`
}

type IEntityProperties interface {
	UnmarshalJSON([]byte) error
	// Create a new bool property.
	//
	// name: The name of the property.
	//
	// clientSync: Whether the property should be synced to the client.
	//
	// defaultValue: The default value of the property.
	//
	// return: The new property.
	//
	// Example:
	//
	//     prop := entity.Properties.NewBool("minecraft:is_baby", true, false)
	NewBool(string, bool, ...bool) *entityProperties

	// Create a new int property.
	//
	// name: The name of the property.
	//
	// clientSync: Whether the property should be synced to the client.
	//
	// min: The minimum value of the property.
	//
	// max: The maximum value of the property.
	//
	// def: The default value of the property.
	//
	// Example:
	//
	//     prop := entity.Properties.NewInt("minecraft:age", true, 0, 100, 0)
	NewInt(string, bool, int, int, int) *entityProperties

	// Create a new enum property.
	//
	// name: The name of the property.
	//
	// clientSync: Whether the property should be synced to the client.
	//
	// values: The possible values of the property.
	//
	// def: The default value of the property.
	//
	// Example:
	//
	//     prop := entity.Properties.NewEnum("minecraft:color", true, []string{"red", "green", "blue"}, "red")
	NewEnum(string, bool, []string, string) *entityProperties

	// Get all properties.
	//
	// return: All properties.
	//
	// Example:
	//
	//     for _, prop := range entity.Properties.All() {
	//         fmt.Println(prop)
	//     }
	All() []*entityProperties

	// Get a property by name.
	//
	// name: The name of the property.
	//
	// return: The property, or nil if not found.
	//
	// Example:
	//
	//     prop, ok := entity.Properties.Get("minecraft:is_baby")
	Get(string) (*entityProperties, bool)
	IsEmpty() bool
}

func (e *EntityProperties) UnmarshalJSON(data []byte) error {
	var temp map[string]*entityProperties
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*e = temp
	return nil
}

func (e *EntityProperties) NewBool(name string, clientSync bool, defaultValue ...bool) *entityProperties {
	prop := &entityProperties{
		ClientSync: clientSync,
		Type:       "bool",
		Default:    false,
	}
	if len(defaultValue) > 0 {
		prop.Default = defaultValue[0]
	}
	(*e)[name] = prop
	return prop
}

func (e *EntityProperties) NewInt(name string, clientSync bool, min int, max int, def int) *entityProperties {
	prop := &entityProperties{
		ClientSync: clientSync,
		Type:       "int",
		Range:      []int{min, max},
		Default:    def,
	}

	(*e)[name] = prop
	return (*e)[name]
}

func (e *EntityProperties) NewEnum(name string, clientSync bool, values []string, def string) *entityProperties {
	prop := &entityProperties{
		ClientSync: clientSync,
		Type:       "enum",
		Values:     values,
		Default:    def,
	}
	(*e)[name] = prop
	return (*e)[name]
}

func (e *EntityProperties) All() []*entityProperties {
	properties := make([]*entityProperties, 0)
	for _, property := range *e {
		properties = append(properties, property)
	}
	return properties
}

func (e *EntityProperties) Get(name string) (*entityProperties, bool) {
	property, ok := (*e)[name]
	return property, ok
}

func (e *EntityProperties) IsEmpty() bool {
	return len(*e) == 0
}
