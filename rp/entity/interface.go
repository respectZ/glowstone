package entity

type Entity interface {
	Encode(minify ...bool) ([]byte, error)

	GetFormatVersion() string
	SetFormatVersion(formatVersion string)

	GetIdentifier() string
	SetIdentifier(identifier string)

	/**** Materials ****/
	GetMaterials() map[string]string
	SetMaterials(materials map[string]string)
	AddMaterial(name string, material string)

	/**** Textures ****/

	GetTextures() map[string]string
	SetTextures(textures map[string]string)
	AddTexture(name string, path string)

	/**** Geometry ****/

	GetGeometry() map[string]string
	SetGeometry(geometry map[string]string)
	AddGeometry(name string, path string)

	/**** Animations ****/

	GetAnimations() map[string]string
	SetAnimations(animations map[string]string)
	AddAnimation(name string, path string)

	/**** Scripts ****/
	GetScripts() ClientEntityDescriptionScripts

	/**** Spawn Egg ****/
	GetSpawnEgg() ClientEntitySpawnEgg
}

type ClientEntityDescriptionScripts interface {
	GetParentSetup() string
	SetParentSetup(parentSetup string)

	GetVariables() map[string]string
	SetVariables(variables map[string]string)
	AddVariable(name string, value string)

	GetScaleX() string
	SetScaleX(scaleX string)

	GetScaleY() string
	SetScaleY(scaleY string)

	GetScaleZ() string
	SetScaleZ(scaleZ string)

	GetScale() string
	SetScale(scale string)

	GetInitialize() []string
	SetInitialize(initialize []string)
	AddInitialize(initialize string)

	GetPreAnimation() []string
	SetPreAnimation(preAnimation []string)
	AddPreAnimation(preAnimation string)

	GetAnimate() []interface{}
	SetAnimate(animate []interface{})
	AddAnimate(animate interface{})
}

type ClientEntitySpawnEgg interface {
	GetBaseColor() string
	SetBaseColor(baseColor string)

	GetOverlayColor() string
	SetOverlayColor(overlayColor string)

	GetTexture() string
	SetTexture(texture string)

	GetTextureIndex() int
	SetTextureIndex(textureIndex int)
}
