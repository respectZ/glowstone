package entity

import (
	"fmt"

	f "github.com/respectZ/glowstone/bp/types"
	glowstone "github.com/respectZ/glowstone/util"
)

var FORMAT_VERSION = "1.20.0"

type entity struct {
	FormatVersion string          `json:"format_version"`
	Entity        MinecraftEntity `json:"minecraft:entity"`
}

type MinecraftEntity struct {
	Description     EntityDescription       `json:"description"`
	ComponentGroups map[string]interface{}  `json:"component_groups,omitempty"` // TODO: implement component groups
	Components      map[string]interface{}  `json:"components,omitempty"`       // TODO: implement components
	Events          map[string]*EntityEvent `json:"events,omitempty"`           // TODO: implement events
}

type EntityDescription struct {
	Identifier        string                      `json:"identifier"`
	IsSpawnable       bool                        `json:"is_spawnable"`
	IsSummonable      bool                        `json:"is_summonable"`
	IsExperimental    bool                        `json:"is_experimental"`
	Animations        map[string]string           `json:"animations,omitempty"` // TODO: implement animations
	Properties        map[string]EntityProperties `json:"properties,omitempty"` // key is the property name
	Scripts           *EntityScripts              `json:"scripts,omitempty"`    // only for animation
	RuntimeIdentifier string                      `json:"runtime_identifier,omitempty"`
}

type EntityProperties struct {
	ClientSync bool        `json:"client_sync"`
	Type       string      `json:"type"`
	Range      []int       `json:"range,omitempty"`  // only for float
	Values     []string    `json:"values,omitempty"` // only for enum
	Default    interface{} `json:"default"`
}

type EntityScripts struct {
	Animate []interface{} `json:"animate,omitempty"`
}

type EntityEvent struct {
	Add         *ComponentGroup        `json:"add,omitempty"`
	Remove      *ComponentGroup        `json:"remove,omitempty"`
	SetProperty map[string]interface{} `json:"set_property,omitempty"` // TODO: implement set_property
	Trigger     interface{}            `json:"trigger,omitempty"`      // TODO: implement trigger
	Randomize   []EntityEvent          `json:"randomize,omitempty"`
	Sequence    []EntityEvent          `json:"sequence,omitempty"`
	Filters     *f.Filter              `json:"filters,omitempty"`
}

type ComponentGroup struct {
	ComponentGroups []string `json:"component_groups,omitempty"`
}

func New(namespace string, identifier string) Entity {
	return &entity{
		FormatVersion: FORMAT_VERSION,
		Entity: MinecraftEntity{
			Description: EntityDescription{
				Identifier:     fmt.Sprintf("%s:%s", namespace, identifier),
				IsSpawnable:    true,
				IsSummonable:   true,
				IsExperimental: false,
			},
		},
	}
}

func Load(src string) (Entity, error) {
	e := &entity{}
	err := glowstone.LoadJSON(src, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

/******************* Short-hand functions *******************/

// Shorthand for crating boolean entity property
func NewProperties() *EntityProperties {
	return &EntityProperties{
		ClientSync: true,
		Type:       "bool",
		Range:      nil,
		Values:     nil,
		Default:    false,
	}
}

// Shorthand for crating integer entity property
func NewPropertiesInt(Range []int, Default int) *EntityProperties {
	return &EntityProperties{
		ClientSync: true,
		Type:       "int",
		Range:      Range,
		Values:     nil,
		Default:    Default,
	}
}

// Shorthand for crating enum entity property
func NewPropertiesEnum(Values []string, Default string) *EntityProperties {
	return &EntityProperties{
		ClientSync: true,
		Type:       "enum",
		Range:      nil,
		Values:     Values,
		Default:    Default,
	}
}
