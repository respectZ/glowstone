package attachable

import (
	types "github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

type Attachable struct {
	FormatVersion string               `json:"format_version"`
	Attachable    *MinecraftAttachable `json:"minecraft:attachable"`
}

type MinecraftAttachable struct {
	Description *MinecraftAttachableDescription `json:"description,omitempty"`
}

type MinecraftAttachableDescription struct {
	Identifier        string                        `json:"identifier"`
	Materials         types.IMapStringString        `json:"materials"`
	Textures          types.IMapStringString        `json:"textures"`
	Geometry          types.IMapStringString        `json:"geometry"`
	Animations        types.IMapStringString        `json:"animations,omitempty"`
	Scripts           *MinecraftAttachableScripts   `json:"scripts,omitempty"`
	Item              types.IMapStringString        `json:"item,omitempty"`
	QueryableGeometry string                        `json:"queryable_geometry,omitempty"`
	RenderControllers types.IStringArrayConditional `json:"render_controllers"`
}

type MinecraftAttachableScripts struct {
	ParentSetup                          types.IStringArray            `json:"parent_setup,omitempty"`
	PreAnimation                         types.IStringArray            `json:"pre_animation,omitempty"`
	ShouldUpdateBonesAndEffectsOffscreen bool                          `json:"should_update_bones_and_effects_offscreen,omitempty"`
	ShouldUpdateEffectsOffscreen         bool                          `json:"should_update_effects_offscreen,omitempty"`
	Animate                              types.IStringArrayConditional `json:"animate,omitempty"`
}

const (
	FORMAT_VERSION string = "1.20.30"
)

func New(identifier string) *Attachable {
	return &Attachable{
		FormatVersion: FORMAT_VERSION,
		Attachable: &MinecraftAttachable{
			Description: &MinecraftAttachableDescription{
				Identifier: identifier,
				Materials: &types.MapStringString{
					"default":   "entity_alphatest",
					"enchanted": "entity_alphatest_glint",
				},
				Textures: &types.MapStringString{
					"enchanted": "textures/misc/enchanted_item_glint",
				},
				Geometry:   &types.MapStringString{},
				Animations: &types.MapStringString{},
				Scripts: &MinecraftAttachableScripts{
					ParentSetup:  &types.StringArray{},
					PreAnimation: &types.StringArray{},
					Animate:      &types.StringArrayConditional{},
				},
				Item: &types.MapStringString{},
				RenderControllers: &types.StringArrayConditional{
					"controller.render.item_default",
				},
			},
		},
	}
}

func Load(src string) (*Attachable, error) {
	a := New("")
	err := g_util.LoadJSON(src, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
