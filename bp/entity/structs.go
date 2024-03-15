package entity

import (
	"github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

var FORMAT_VERSION = "1.20.70"

type Entity struct {
	FormatVersion string          `json:"format_version"`
	Entity        MinecraftEntity `json:"minecraft:entity"`
}

type MinecraftEntity struct {
	Description     EntityDescription `json:"description"`
	Components      IComponent        `json:"components,omitempty"`
	ComponentGroups IComponentGroups  `json:"component_groups,omitempty"`
	Events          IEntityEventMap   `json:"events,omitempty"`
}

type EntityDescription struct {
	Identifier        string                 `json:"identifier"`
	IsSpawnable       bool                   `json:"is_spawnable"`
	IsSummonable      bool                   `json:"is_summonable"`
	IsExperimental    bool                   `json:"is_experimental"`
	Animations        types.IMapStringString `json:"animations,omitempty"`
	Properties        IEntityProperties      `json:"properties,omitempty"`
	Scripts           *EntityScripts         `json:"scripts,omitempty"`
	RuntimeIdentifier string                 `json:"runtime_identifier,omitempty"`
}

type ComponentGroup struct {
	ComponentGroups types.IStringArray `json:"component_groups,omitempty"`
}

func New(identifier string) *Entity {
	return &Entity{
		FormatVersion: FORMAT_VERSION,
		Entity: MinecraftEntity{
			Description: EntityDescription{
				Identifier:     identifier,
				IsSpawnable:    true,
				IsSummonable:   true,
				IsExperimental: false,
				Properties:     &EntityProperties{},
				Animations:     &types.MapStringString{},
				Scripts: &EntityScripts{
					Animate: &types.StringArrayConditional{},
				},
			},
			Components:      IComponent(&Component{}),
			ComponentGroups: IComponentGroups(&ComponentGroups{}),
			Events:          IEntityEventMap(&EntityEventMap{}),
		},
	}
}

func Load(src string) (*Entity, error) {
	e := &Entity{}

	// Initialize interface field
	e.Entity.Description.Scripts = &EntityScripts{Animate: &types.StringArrayConditional{}}
	e.Entity.Description.Animations = &types.MapStringString{}
	e.Entity.Description.Properties = &EntityProperties{}
	e.Entity.Components = &Component{}
	e.Entity.ComponentGroups = &ComponentGroups{}
	e.Entity.Events = &EntityEventMap{}

	err := g_util.LoadJSON(src, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
