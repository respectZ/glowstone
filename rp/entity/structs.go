package entity

import (
	"fmt"

	g_util "github.com/respectZ/glowstone/util"
)

type entity struct {
	FormatVersion string        `json:"format_version"`
	Entity        *ClientEntity `json:"minecraft:client_entity"`
}

type ClientEntity struct {
	Description *ClientEntityDescription `json:"description"`
}

type ClientEntityDescription struct {
	Identifier              string                          `json:"identifier"`
	Materials               map[string]string               `json:"materials"`
	Textures                map[string]string               `json:"textures"`
	Geometry                map[string]string               `json:"geometry"`
	Animations              map[string]string               `json:"animations,omitempty"`
	Scripts                 *clientEntityDescriptionScripts `json:"scripts,omitempty"`
	SpawnEgg                *clientEntitySpawnEgg           `json:"spawn_egg,omitempty"`
	RenderControllers       []interface{}                   `json:"render_controllers,omitempty"`
	EnableAttachables       bool                            `json:"enable_attachables,omitempty"`
	HeldItemIgnoresLighting bool                            `json:"held_item_ignores_lighting,omitempty"`
	HideArmor               bool                            `json:"hide_armor,omitempty"`
	ParticleEffects         map[string]string               `json:"particle_effects,omitempty"`
	SoundEffects            map[string]string               `json:"sound_effects,omitempty"`
	ParticleEmitters        map[string]string               `json:"particle_emitters,omitempty"`
}

type clientEntityDescriptionScripts struct {
	ParentSetup  string            `json:"parent_setup,omitempty"`
	Variables    map[string]string `json:"variables,omitempty"`
	ScaleX       string            `json:"scalex,omitempty"`
	ScaleY       string            `json:"scaley,omitempty"`
	ScaleZ       string            `json:"scalez,omitempty"`
	Scale        string            `json:"scale,omitempty"`
	Initialize   []string          `json:"initialize,omitempty"`
	PreAnimation []string          `json:"pre_animation,omitempty"`
	Animate      []interface{}     `json:"animate,omitempty"`
}

type clientEntitySpawnEgg struct {
	BaseColor    string `json:"base_color,omitempty"`
	OverlayColor string `json:"overlay_color,omitempty"`

	Texture      string `json:"texture,omitempty"`
	TextureIndex int    `json:"texture_index,omitempty"`
}

func New(namespace string, identifier string) Entity {
	e := &entity{
		FormatVersion: "1.12.0",
		Entity: &ClientEntity{
			Description: &ClientEntityDescription{
				Identifier: fmt.Sprintf("%s:%s", namespace, identifier),
				Materials: map[string]string{
					"default": "entity_alphatest",
				},
				Textures: map[string]string{
					"default": fmt.Sprintf("textures/entity/%s", identifier),
				},
				Geometry: map[string]string{
					"default": fmt.Sprintf("geometry.%s", identifier),
				},
				RenderControllers: []interface{}{
					"controller.render.default",
				},
			},
		},
	}
	return e
}

func Load(dir string) (Entity, error) {
	e := &entity{}
	err := g_util.LoadJSON(dir, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
