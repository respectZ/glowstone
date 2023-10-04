package attachable

import (
	g_util "github.com/respectZ/glowstone/util"
)

func (a *attachable) Encode() ([]byte, error) {
	return g_util.MarshalJSON(a)
}

func (a *attachable) GetIdentifier() string {
	return a.Attachable.Description.Identifier
}

func (a *attachable) SetIdentifier(identifier string) {
	a.Attachable.Description.Identifier = identifier
}

func (a *attachable) GetMaterials() map[string]string {
	return a.Attachable.Description.Materials
}

func (a *attachable) SetMaterials(materials map[string]string) {
	a.Attachable.Description.Materials = materials
}

func (a *attachable) GetTextures() map[string]string {
	return a.Attachable.Description.Textures
}

func (a *attachable) SetTextures(textures map[string]string) {
	a.Attachable.Description.Textures = textures
}

func (a *attachable) GetGeometry() map[string]string {
	return a.Attachable.Description.Geometry
}

func (a *attachable) SetGeometry(geometry map[string]string) {
	a.Attachable.Description.Geometry = geometry
}

func (a *attachable) GetAnimations() map[string]string {
	return a.Attachable.Description.Animations
}

func (a *attachable) SetAnimations(animations map[string]string) {
	a.Attachable.Description.Animations = animations
}

func (a *attachable) GetItem() map[string]string {
	return a.Attachable.Description.Item
}

func (a *attachable) SetItem(item map[string]string) {
	a.Attachable.Description.Item = item
}

func (a *attachable) GetQueryableGeometry() string {
	return a.Attachable.Description.QueryableGeometry
}

func (a *attachable) SetQueryableGeometry(queryableGeometry string) {
	a.Attachable.Description.QueryableGeometry = queryableGeometry
}

func (a *attachable) GetRenderControllers() []string {
	return a.Attachable.Description.RenderControllers
}

func (a *attachable) SetRenderControllers(renderControllers []string) {
	a.Attachable.Description.RenderControllers = renderControllers
}
