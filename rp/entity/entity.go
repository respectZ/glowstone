package entity

import (
	g_util "github.com/respectZ/glowstone/util"
)

func (e *entity) Encode(minify ...bool) ([]byte, error) {
	return g_util.MarshalJSON(e, minify...)
}

func (e *entity) GetFormatVersion() string {
	return e.FormatVersion
}

func (e *entity) SetFormatVersion(formatVersion string) {
	e.FormatVersion = formatVersion
}

func (e *entity) GetIdentifier() string {
	return e.Entity.Description.Identifier
}

func (e *entity) SetIdentifier(identifier string) {
	e.Entity.Description.Identifier = identifier
}

func (e *entity) GetMaterials() map[string]string {
	return e.Entity.Description.Materials
}

func (e *entity) SetMaterials(materials map[string]string) {
	e.Entity.Description.Materials = materials
}

func (e *entity) AddMaterial(name string, material string) {
	e.Entity.Description.Materials[name] = material
}

func (e *entity) GetTextures() map[string]string {
	return e.Entity.Description.Textures
}

func (e *entity) SetTextures(textures map[string]string) {
	e.Entity.Description.Textures = textures
}

func (e *entity) AddTexture(name string, path string) {
	e.Entity.Description.Textures[name] = path
}

func (e *entity) GetGeometry() map[string]string {
	return e.Entity.Description.Geometry
}

func (e *entity) SetGeometry(geometry map[string]string) {
	e.Entity.Description.Geometry = geometry
}

func (e *entity) AddGeometry(name string, path string) {
	e.Entity.Description.Geometry[name] = path
}

func (e *entity) GetAnimations() map[string]string {
	return e.Entity.Description.Animations
}

func (e *entity) SetAnimations(animations map[string]string) {
	e.Entity.Description.Animations = animations
}

func (e *entity) AddAnimation(name string, path string) {
	if e.Entity.Description.Animations == nil {
		e.Entity.Description.Animations = make(map[string]string)
	}
	e.Entity.Description.Animations[name] = path
}

func (e *entity) GetScripts() ClientEntityDescriptionScripts {
	return e.Entity.Description.Scripts
}

func (e *entity) GetSpawnEgg() ClientEntitySpawnEgg {
	return e.Entity.Description.SpawnEgg
}
