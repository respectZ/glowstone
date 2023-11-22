package entity

import jsoniter "github.com/json-iterator/go"

type EntityProperties map[string]entityProperties

type entityProperties struct {
	ClientSync bool        `json:"client_sync"`
	Type       string      `json:"type"`
	Range      []int       `json:"range,omitempty"`  // only for float
	Values     []string    `json:"values,omitempty"` // only for enum
	Default    interface{} `json:"default"`
}

type IEntityProperties interface {
	UnmarshalJSON([]byte) error
	NewBool(string, bool, ...bool)
	NewInt(string, bool, int, int, int)
	NewEnum(string, bool, []string, string)
	All() []entityProperties
	IsEmpty() bool
}

func (e *EntityProperties) UnmarshalJSON(data []byte) error {
	var json = jsoniter.ConfigFastest
	var temp map[string]entityProperties
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*e = temp
	return nil
}

func (e *EntityProperties) NewBool(name string, clientSync bool, defaultValue ...bool) {
	prop := entityProperties{
		ClientSync: clientSync,
		Type:       "bool",
		Default:    false,
	}
	if len(defaultValue) > 0 {
		prop.Default = defaultValue[0]
	}
	(*e)[name] = prop
}

func (e *EntityProperties) NewInt(name string, clientSync bool, min int, max int, def int) {
	(*e)[name] = entityProperties{
		ClientSync: clientSync,
		Type:       "int",
		Range:      []int{min, max},
		Default:    def,
	}
}

func (e *EntityProperties) NewEnum(name string, clientSync bool, values []string, def string) {
	(*e)[name] = entityProperties{
		ClientSync: clientSync,
		Type:       "enum",
		Values:     values,
		Default:    def,
	}
}

func (e *EntityProperties) All() []entityProperties {
	properties := make([]entityProperties, 0)
	for _, property := range *e {
		properties = append(properties, property)
	}
	return properties
}

func (e *EntityProperties) IsEmpty() bool {
	return len(*e) == 0
}
