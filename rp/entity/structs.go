package entity

import (
	glowstone "github.com/respectZ/glowstone/util"
)

type Entity_struct struct {
	FormatVersion                   string       `json:"format_version"`
	Entity                          ClientEntity `json:"minecraft:client_entity"`
	GenerateRenderController        bool         `json:"-"`
	GenerateRenderControllerIndexer string       `json:"-"`

	// Other opts
	Textures         []string      `json:"-"`
	ItemTextures     []string      `json:"-"`
	Geometry         []string      `json:"-"`
	Sound            []string      `json:"-"`
	RenderController []interface{} `json:"-"`
}

type ClientEntity struct {
	Description ClientEntityDescription `json:"description"`
}

type ClientEntityDescription struct {
	Identifier              string                          `json:"identifier"`
	Materials               map[string]string               `json:"materials"`
	Textures                map[string]string               `json:"textures"`
	Geometry                map[string]string               `json:"geometry"`
	Animations              map[string]string               `json:"animations,omitempty"`
	Scripts                 *ClientEntityDescriptionScripts `json:"scripts,omitempty"`
	SpawnEgg                *ClientEntitySpawnEgg           `json:"spawn_egg,omitempty"`
	RenderControllers       []interface{}                   `json:"render_controllers,omitempty"`
	EnableAttachables       bool                            `json:"enable_attachables,omitempty"`
	HeldItemIgnoresLighting bool                            `json:"held_item_ignores_lighting,omitempty"`
	HideArmor               bool                            `json:"hide_armor,omitempty"`
	ParticleEffects         map[string]string               `json:"particle_effects,omitempty"`
	SoundEffects            map[string]string               `json:"sound_effects,omitempty"`
	ParticleEmitters        map[string]string               `json:"particle_emitters,omitempty"`
}

type ClientEntityDescriptionScripts struct {
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

type ClientEntitySpawnEgg struct {
	BaseColor    string `json:"base_color"`
	OverlayColor string `json:"overlay_color"`

	Texture      string `json:"texture"`
	TextureIndex int    `json:"texture_index"`
}

type Entity interface {
	Encode() ([]byte, error)
	Load(dir string) error

	GetIdentifier() string
	GetMaterials() map[string]string
	GetTextures() map[string]string
	GetGeometry() map[string]string
	GetAnimations() map[string]string
	GetScripts() *ClientEntityDescriptionScripts
	GetSpawnEgg() *ClientEntitySpawnEgg
	GetRenderControllers() []interface{}
	GetEnableAttachables() bool
	GetHeldItemIgnoresLighting() bool
	GetHideArmor() bool
	GetParticleEffects() map[string]string
	GetSoundEffects() map[string]string
	GetParticleEmitters() map[string]string

	SetIdentifier(identifier string)
	SetEnableAttachables(enableAttachables bool)
	SetHeldItemIgnoresLighting(heldItemIgnoresLighting bool)
	SetHideArmor(hideArmor bool)

	SetSpawnEggColor(baseColor string, overlayColor string)
	SetSpawnEggTexture(texture string, textureIndex ...int)

	AddMaterial(name string, material string)
	AddTexture(name string, path string)
	AddGeometry(name string, path string)
	AddAnimation(name string, path string)

	// Add a render controller to the entity (unconditional).
	AddRenderController(name string)

	// Add a render controller to the entity (conditional).
	// The condition is the name of the variable / query that will be used to determine if the render controller should be applied.
	//
	// Example:
	// 	- name: "controller.render.default"
	// 	- condition: "q.variant == 100"
	AddRenderControllerConditional(name string, condition string)

	// Add a particle effect to the entity.
	//
	// Example:
	// 	- name: "fire"
	// 	- path: "minecraft:mobflame_emitter"
	AddParticleEffect(name string, path string)

	// Add a sound effect to the entity.
	//
	// Example:
	// 	- name: "fizz"
	// 	- path: "random.fizz"
	AddSoundEffect(name string, path string)

	// Add a particle emitter to the entity.
	//
	// Example:
	// 	- name: "fire"
	// 	- path: "minecraft:mobflame_emitter"
	AddParticleEmitter(name string, path string)

	// Empty the materials.
	ResetMaterials()

	// Empty the textures.
	ResetTextures()

	// Empty the geometry.
	ResetGeometry()

	// Empty the animations.
	ResetAnimations()

	// Empty the render controllers.
	ResetRenderControllers()

	// Set spawn egg color automatically based on the texture.
	//
	// Parameters:
	//  - rpPath: The path to the RP folder.
	//
	// Example:
	//  e.SetSpawnEggColorAuto("./packs/RP/")
	SetSpawnEggColorAuto(string) error

	// Automatically generate a render controller and apply it to the entity.
	// This is useful for entities that have multiple geometry and textures.
	// The indexer is the name of the variable / query that will be used to index the geometry and or texture.
	//
	// Example:
	// 	- geometry: ["geometry.humanoid.custom", "geometry.humanoid.custom2"]
	// 	- textures: ["texture.humanoid.custom", "texture.humanoid.custom2"]
	// 	- indexer: "q.variant"
	//
	// This will generate a render controller that will use the query "q.variant" to index the geometry and texture.
	AutoGenerateRenderController(v bool, indexer string)

	// Add a sound effect to the entity and automatically copy the sound file.
	//
	// Example:
	// 	- path: "./player/cave.ogg"
	// 	- shorthand: "cave"
	AddSoundEffectFile(path string, shorthand string)
}

func Load(dir string) (Entity, error) {
	e := &Entity_struct{}
	err := glowstone.LoadJSON(dir, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
