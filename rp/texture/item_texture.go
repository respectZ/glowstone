package texture

import (
	"strings"

	g_util "github.com/respectZ/glowstone/util"
)

type ItemTexture struct {
	ResourcePackName string                         `json:"resource_pack_name,omitempty"`
	TextureData      map[string]ItemTexture_Texture `json:"texture_data"`
}

type ItemTexture_Texture struct {
	// Can be either string or array of strings
	Textures interface{} `json:"textures,omitempty"`
}

// Creates a new item_texture.json file
func NewItemTexture() *ItemTexture {
	return &ItemTexture{
		TextureData: make(map[string]ItemTexture_Texture),
	}
}

// Loads a item_texture.json file
func LoadItemTexture(src string) (*ItemTexture, error) {
	var texture ItemTexture
	err := g_util.LoadJSON(src, &texture)
	if err != nil {
		return nil, err
	}
	return &texture, nil
}

func (i *ItemTexture) Encode() ([]byte, error) {
	return g_util.MarshalJSON(i)
}

// Add a texture to the item_texture.json file
//
// Example:
//
//	item_texture.AddTexture("stick", "textures/items/stick")
func (i *ItemTexture) AddTexture(name string, texture interface{}) {
	if i.TextureData == nil {
		i.TextureData = make(map[string]ItemTexture_Texture)
	}
	switch v := texture.(type) {
	case string:
		// Replace double backslashes with single backslashes
		texture = strings.Replace(v, "\\\\", "\\", -1)
	}
	i.TextureData[name] = ItemTexture_Texture{
		Textures: texture,
	}
}
