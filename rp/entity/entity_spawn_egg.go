package entity

func (e *clientEntitySpawnEgg) GetBaseColor() string {
	return e.BaseColor
}

func (e *clientEntitySpawnEgg) SetBaseColor(baseColor string) {
	e.BaseColor = baseColor
}

func (e *clientEntitySpawnEgg) GetOverlayColor() string {
	return e.OverlayColor
}

func (e *clientEntitySpawnEgg) SetOverlayColor(overlayColor string) {
	e.OverlayColor = overlayColor
}

func (e *clientEntitySpawnEgg) GetTexture() string {
	return e.Texture
}

func (e *clientEntitySpawnEgg) SetTexture(texture string) {
	e.Texture = texture
}

func (e *clientEntitySpawnEgg) GetTextureIndex() int {
	return e.TextureIndex
}

func (e *clientEntitySpawnEgg) SetTextureIndex(textureIndex int) {
	e.TextureIndex = textureIndex
}
