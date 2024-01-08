package texture

import (
	"path/filepath"

	g_util "github.com/respectZ/glowstone/util"
)

type TerrainTexture struct {
	NumMipLevels     int                                `json:"num_mip_levels,omitempty"`
	Padding          int                                `json:"padding,omitempty"`
	ResourcePackName string                             `json:"resource_pack_name,omitempty"`
	TextureData      map[string]*TerrainTexture_Texture `json:"texture_data"`
}

type TerrainTexture_Texture struct {
	// Can be either string or array of strings, or texturedata
	Textures interface{} `json:"textures,omitempty"`
}

type TerrainTexture_TextureData struct {
	OverlayColor string `json:"overlay_color,omitempty"`
	Path         string `json:"path,omitempty"`
}

// Creates a new terrain_texture.json file
func NewTerrainTexture() *TerrainTexture {
	return &TerrainTexture{
		TextureData: make(map[string]*TerrainTexture_Texture),
	}
}

// Loads a terrain_texture.json file
func LoadTerrainTexture(src string) (*TerrainTexture, error) {
	var texture TerrainTexture
	err := g_util.LoadJSON(src, &texture)
	if err != nil {
		return nil, err
	}
	return &texture, nil
}

func (i *TerrainTexture) Encode() ([]byte, error) {
	return g_util.MarshalJSON(i)
}

// Add a texture to the terrain_texture.json file
//
// Example:
//
//	terrain_texture.AddTexture("grass", "textures/blocks/grass")
func (i *TerrainTexture) AddTexture(name string, texture interface{}) {
	if i.TextureData == nil {
		i.TextureData = make(map[string]*TerrainTexture_Texture)
	}
	i.TextureData[name] = &TerrainTexture_Texture{
		Textures: texture,
	}
}

func (i *TerrainTexture) IsEmpty() bool {
	return len(i.TextureData) == 0
}

func (i *TerrainTexture) Save(pathToRp string) error {
	// Prevent overwriting
	d, err := LoadTerrainTexture(filepath.Join(pathToRp, "textures", "terrain_texture.json"))
	if err == nil {
		for k, v := range i.TextureData {
			d.TextureData[k] = v
		}
		i = d
	}

	data, err := i.Encode()
	if err != nil {
		return err
	}
	return g_util.WriteFile(filepath.Join(pathToRp, "textures", "terrain_texture.json"), data)
}
