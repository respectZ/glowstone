package entity

import (
	"fmt"
	"strings"

	types "github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

type Entity struct {
	FormatVersion string        `json:"format_version"`
	Entity        *ClientEntity `json:"minecraft:client_entity"`
}

type ClientEntity struct {
	Description *ClientEntityDescription `json:"description"`
}

type ClientEntityDescription struct {
	Identifier              string                          `json:"identifier"`
	Materials               types.IMapStringString          `json:"materials"`
	Textures                types.IMapStringString          `json:"textures"`
	Geometry                types.IMapStringString          `json:"geometry"`
	Animations              types.IMapStringString          `json:"animations,omitempty"`
	Scripts                 *clientEntityDescriptionScripts `json:"scripts,omitempty"`
	SpawnEgg                *clientEntitySpawnEgg           `json:"spawn_egg,omitempty"`
	RenderControllers       types.IStringArrayConditional   `json:"render_controllers,omitempty"`
	EnableAttachables       bool                            `json:"enable_attachables,omitempty"`
	HeldItemIgnoresLighting bool                            `json:"held_item_ignores_lighting,omitempty"`
	HideArmor               bool                            `json:"hide_armor,omitempty"`
	ParticleEffects         types.IMapStringString          `json:"particle_effects,omitempty"`
	SoundEffects            types.IMapStringString          `json:"sound_effects,omitempty"`
	ParticleEmitters        types.IMapStringString          `json:"particle_emitters,omitempty"`
}

type clientEntitySpawnEgg struct {
	BaseColor    string `json:"base_color,omitempty"`
	OverlayColor string `json:"overlay_color,omitempty"`

	Texture      string `json:"texture,omitempty"`
	TextureIndex int    `json:"texture_index,omitempty"`
}

func New(identifier string) *Entity {
	s := strings.Split(identifier, ":")
	namespace := s[0]
	identifier = s[1]

	e := &Entity{
		FormatVersion: "1.12.0",
		Entity: &ClientEntity{
			Description: &ClientEntityDescription{
				Identifier: fmt.Sprintf("%s:%s", namespace, identifier),
				Materials: &types.MapStringString{
					"default": "entity_alphatest",
				},
				Textures: &types.MapStringString{
					"default": fmt.Sprintf("textures/entity/%s", identifier),
				},
				Geometry: &types.MapStringString{
					"default": fmt.Sprintf("geometry.%s", identifier),
				},
				Animations: &types.MapStringString{},
				Scripts: &clientEntityDescriptionScripts{
					Variables:    &types.MapStringString{},
					Initialize:   &types.StringArray{},
					PreAnimation: &types.StringArray{},
					Animate:      &types.StringArrayConditional{},
				},
				SoundEffects:     &types.MapStringString{},
				ParticleEffects:  &types.MapStringString{},
				ParticleEmitters: &types.MapStringString{},
				RenderControllers: &types.StringArrayConditional{
					"controller.render.default",
				},
				SpawnEgg: &clientEntitySpawnEgg{},
			},
		},
	}
	return e
}

func Load(dir string) (*Entity, error) {
	e := &Entity{
		Entity: &ClientEntity{
			Description: &ClientEntityDescription{
				Materials:  &types.MapStringString{},
				Textures:   &types.MapStringString{},
				Geometry:   &types.MapStringString{},
				Animations: &types.MapStringString{},
				Scripts: &clientEntityDescriptionScripts{
					Variables:    &types.MapStringString{},
					Initialize:   &types.StringArray{},
					PreAnimation: &types.StringArray{},
					Animate:      &types.StringArrayConditional{},
				},
				SoundEffects:      &types.MapStringString{},
				ParticleEffects:   &types.MapStringString{},
				ParticleEmitters:  &types.MapStringString{},
				RenderControllers: &types.StringArrayConditional{},
				SpawnEgg:          &clientEntitySpawnEgg{},
			},
		},
	}
	err := g_util.LoadJSON(dir, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
