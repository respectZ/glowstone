package attachable

import (
	g_util "github.com/respectZ/glowstone/util"
)

type attachable struct {
	FormatVersion string               `json:"format_version"`
	Attachable    *MinecraftAttachable `json:"minecraft:attachable"`
}

type MinecraftAttachable struct {
	Description *MinecraftAttachableDescription `json:"description,omitempty"`
}

type MinecraftAttachableDescription struct {
	Identifier        string            `json:"identifier"`
	Materials         map[string]string `json:"materials"`
	Textures          map[string]string `json:"textures"`
	Geometry          map[string]string `json:"geometry"`
	Animations        map[string]string `json:"animations,omitempty"`
	Item              map[string]string `json:"item,omitempty"`
	QueryableGeometry string            `json:"queryable_geometry,omitempty"`
	RenderControllers []string          `json:"render_controllers"`
}

const (
	FORMAT_VERSION string = "1.20.30"
)

func New(namespace string, identifier string) Attachable {
	return &attachable{
		FormatVersion: FORMAT_VERSION,
		Attachable: &MinecraftAttachable{
			Description: &MinecraftAttachableDescription{
				Identifier:        namespace + ":" + identifier,
				Materials:         make(map[string]string),
				Textures:          make(map[string]string),
				Geometry:          make(map[string]string),
				RenderControllers: make([]string, 0),
			},
		},
	}
}

func Load(src string) (Attachable, error) {
	a := &attachable{}
	err := g_util.LoadJSON(src, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
