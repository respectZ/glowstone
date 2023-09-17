package entity

import (
	"fmt"
	"path"

	g_util "github.com/respectZ/glowstone/util"
)

var FORMAT_VERSION = "1.12.0"

func New(namespace string, identifier string, subdir ...string) *Entity_struct {
	e := &Entity_struct{
		FormatVersion: FORMAT_VERSION,
		Entity: ClientEntity{
			Description: ClientEntityDescription{
				Identifier:        fmt.Sprintf("%s:%s", namespace, identifier),
				Materials:         make(map[string]string),
				Textures:          make(map[string]string),
				Geometry:          make(map[string]string),
				RenderControllers: []interface{}{},
			},
		},
	}
	sub := ""
	if len(subdir) > 0 {
		sub = subdir[0]
	}
	// Set default values
	e.AddMaterial("default", "entity_alphatest")
	e.AddTexture("default", path.Join("textures", "entity", sub, identifier))
	e.AddGeometry("default", fmt.Sprintf("geometry.%s", identifier))
	return e
}

func (e *Entity_struct) Encode() ([]byte, error) {
	return g_util.MarshalJSON(e)
}

func (e *Entity_struct) Load(dir string) error {
	err := g_util.LoadJSON(dir, e)
	if err != nil {
		return err
	}
	return nil
}

func (e *Entity_struct) GetIdentifier() string {
	return e.Entity.Description.Identifier
}

func (e *Entity_struct) GetMaterials() map[string]string {
	return e.Entity.Description.Materials
}

func (e *Entity_struct) GetTextures() map[string]string {
	return e.Entity.Description.Textures
}

func (e *Entity_struct) GetGeometry() map[string]string {
	return e.Entity.Description.Geometry
}

func (e *Entity_struct) GetAnimations() map[string]string {
	return e.Entity.Description.Animations
}

func (e *Entity_struct) GetScripts() ClientEntityDescriptionScripts {
	return e.Entity.Description.Scripts
}

func (e *Entity_struct) GetSpawnEgg() ClientEntitySpawnEgg {
	return e.Entity.Description.SpawnEgg
}

func (e *Entity_struct) GetRenderControllers() []interface{} {
	return e.Entity.Description.RenderControllers
}

func (e *Entity_struct) GetEnableAttachables() bool {
	return e.Entity.Description.EnableAttachables
}

func (e *Entity_struct) GetHeldItemIgnoresLighting() bool {
	return e.Entity.Description.HeldItemIgnoresLighting
}

func (e *Entity_struct) GetHideArmor() bool {
	return e.Entity.Description.HideArmor
}

func (e *Entity_struct) GetParticleEffects() map[string]string {
	return e.Entity.Description.ParticleEffects
}

func (e *Entity_struct) GetSoundEffects() map[string]string {
	return e.Entity.Description.SoundEffects
}

func (e *Entity_struct) GetParticleEmitters() map[string]string {
	return e.Entity.Description.ParticleEmitters
}

func (e *Entity_struct) SetIdentifier(identifier string) {
	e.Entity.Description.Identifier = identifier
}

func (e *Entity_struct) SetEnableAttachables(enableAttachables bool) {
	e.Entity.Description.EnableAttachables = enableAttachables
}

func (e *Entity_struct) SetHeldItemIgnoresLighting(heldItemIgnoresLighting bool) {
	e.Entity.Description.HeldItemIgnoresLighting = heldItemIgnoresLighting
}

func (e *Entity_struct) SetHideArmor(hideArmor bool) {
	e.Entity.Description.HideArmor = hideArmor
}

func (e *Entity_struct) SetSpawnEggColor(baseColor string, overlayColor string) {
	e.Entity.Description.SpawnEgg.BaseColor = baseColor
	e.Entity.Description.SpawnEgg.OverlayColor = overlayColor
}

func (e *Entity_struct) SetSpawnEggTexture(texture string, textureIndex ...int) {
	e.Entity.Description.SpawnEgg.Texture = texture
	if len(textureIndex) > 0 {
		e.Entity.Description.SpawnEgg.TextureIndex = textureIndex[0]
	}
}

func (e *Entity_struct) AddMaterial(name string, material string) {
	e.Entity.Description.Materials[name] = material
}

func (e *Entity_struct) AddTexture(name string, texture string) {
	e.Entity.Description.Textures[name] = texture
}

func (e *Entity_struct) AddGeometry(name string, geometry string) {
	e.Entity.Description.Geometry[name] = geometry
}

func (e *Entity_struct) AddAnimation(name string, animation string) {
	e.Entity.Description.Animations[name] = animation
}

func (e *Entity_struct) AddRenderController(name string) {
	e.Entity.Description.RenderControllers = append(e.Entity.Description.RenderControllers, name)
}

func (e *Entity_struct) AddRenderControllerConditional(name string, condition string) {
	e.Entity.Description.RenderControllers = append(e.Entity.Description.RenderControllers, map[string]string{
		name: condition,
	})
}

func (e *Entity_struct) AddParticleEffect(name string, particleEffect string) {
	if e.Entity.Description.ParticleEffects == nil {
		e.Entity.Description.ParticleEffects = make(map[string]string)
	}
	e.Entity.Description.ParticleEffects[name] = particleEffect
}

func (e *Entity_struct) AddSoundEffect(name string, soundEffect string) {
	if e.Entity.Description.SoundEffects == nil {
		e.Entity.Description.SoundEffects = make(map[string]string)
	}
	e.Entity.Description.SoundEffects[name] = soundEffect
}

func (e *Entity_struct) AddParticleEmitter(name string, particleEmitter string) {
	if e.Entity.Description.ParticleEmitters == nil {
		e.Entity.Description.ParticleEmitters = make(map[string]string)
	}
	e.Entity.Description.ParticleEmitters[name] = particleEmitter
}

func (e *Entity_struct) ResetMaterials() {
	e.Entity.Description.Materials = make(map[string]string)
}

func (e *Entity_struct) ResetTextures() {
	e.Entity.Description.Textures = make(map[string]string)
}

func (e *Entity_struct) ResetGeometry() {
	e.Entity.Description.Geometry = make(map[string]string)
}

func (e *Entity_struct) ResetAnimations() {
	e.Entity.Description.Animations = make(map[string]string)
}

func (e *Entity_struct) ResetRenderControllers() {
	e.Entity.Description.RenderControllers = []interface{}{}
}

func (e *Entity_struct) AutoGenerateRenderController(v bool, indexer string) {
	e.GenerateRenderController = v
	e.GenerateRenderControllerIndexer = indexer
}

func (e *Entity_struct) AddSoundEffectFile(path string, shorthand string) {
	// TODO: Add sound effect file
	e.Sound = append(e.Sound, path)
	e.AddSoundEffect(path, path)
}
