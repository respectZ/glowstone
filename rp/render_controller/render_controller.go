package rendercontroller

import (
	"github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

type RenderController struct {
	Arrays                   *RenderControllerArrays `json:"arrays,omitempty"`
	Color                    *Color                  `json:"color,omitempty"`
	FilterLighting           interface{}             `json:"filter_lighting,omitempty"`
	Geometry                 string                  `json:"geometry,omitempty"`
	IgnoreLighting           interface{}             `json:"ignore_lighting,omitempty"`
	IsHurtColor              *Color                  `json:"is_hurt_color,omitempty"`
	LightColorMultiplier     float64                 `json:"light_color_multiplier,omitempty"`
	Materials                IObjects                `json:"materials,omitempty"`
	OnFireColor              *Color                  `json:"on_fire_color,omitempty"`
	OverlayColor             *Color                  `json:"overlay_color,omitempty"`
	PartVisibility           IObjects                `json:"part_visibility,omitempty"`
	RebuildAnimationMatrices bool                    `json:"rebuild_animation_matrices,omitempty"`
	Textures                 types.IStringArray      `json:"textures,omitempty"`
	UvAnim                   *UvAnim                 `json:"uv_anim,omitempty"`
}

type renderController_parse struct {
	Arrays                   *RenderControllerArrays `json:"arrays,omitempty"`
	Color                    *Color                  `json:"color,omitempty"`
	FilterLighting           interface{}             `json:"filter_lighting,omitempty"`
	Geometry                 string                  `json:"geometry,omitempty"`
	IgnoreLighting           interface{}             `json:"ignore_lighting,omitempty"`
	IsHurtColor              *Color                  `json:"is_hurt_color,omitempty"`
	LightColorMultiplier     float64                 `json:"light_color_multiplier,omitempty"`
	Materials                *Objects                `json:"materials,omitempty"`
	OnFireColor              *Color                  `json:"on_fire_color,omitempty"`
	OverlayColor             *Color                  `json:"overlay_color,omitempty"`
	PartVisibility           *Objects                `json:"part_visibility,omitempty"`
	RebuildAnimationMatrices bool                    `json:"rebuild_animation_matrices,omitempty"`
	Textures                 *types.StringArray      `json:"textures,omitempty"`
	UvAnim                   *UvAnim                 `json:"uv_anim,omitempty"`
}

type RenderControllers map[string]*RenderController

type IRenderControllers interface {
	Add(string, *RenderController)
	Get(string) (*RenderController, bool)
	Remove(string)
	Clear()
	All() map[string]*RenderController
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// Create a new render controller.
	//
	// Example:
	//
	//     renderController := rendercontroller.New("controller.render.weapon")
	New(string) *RenderController
}

func (r *RenderControllers) Add(key string, value *RenderController) {
	(*r)[key] = value
}

func (r *RenderControllers) Get(key string) (*RenderController, bool) {
	value, ok := (*r)[key]
	return value, ok
}

func (r *RenderControllers) Remove(key string) {
	delete(*r, key)
}

func (r *RenderControllers) Clear() {
	*r = make(map[string]*RenderController)
}

func (r *RenderControllers) All() map[string]*RenderController {
	return *r
}

func (r *RenderControllers) IsEmpty() bool {
	return len(*r) == 0
}

func (r *RenderControllers) Size() int {
	return len(*r)
}

func (r *RenderControllers) UnmarshalJSON(b []byte) error {
	var temp map[string]*renderController_parse
	if err := g_util.UnmarshalJSON(b, &temp); err != nil {
		return err
	}
	*r = make(map[string]*RenderController)
	for k, v := range temp {
		(*r)[k] = &RenderController{
			Arrays:                   v.Arrays,
			Color:                    v.Color,
			FilterLighting:           v.FilterLighting,
			Geometry:                 v.Geometry,
			IgnoreLighting:           v.IgnoreLighting,
			IsHurtColor:              v.IsHurtColor,
			LightColorMultiplier:     v.LightColorMultiplier,
			Materials:                v.Materials,
			OnFireColor:              v.OnFireColor,
			OverlayColor:             v.OverlayColor,
			PartVisibility:           v.PartVisibility,
			RebuildAnimationMatrices: v.RebuildAnimationMatrices,
			Textures:                 v.Textures,
			UvAnim:                   v.UvAnim,
		}
	}
	return nil
}

func (r *RenderControllers) New(key string) *RenderController {
	renderController := &RenderController{
		Arrays: &RenderControllerArrays{
			Geometries: &types.MapStringArray{},
			Materials:  &types.MapStringArray{},
			Textures:   &types.MapStringArray{},
		},
		Materials:      &Objects{},
		PartVisibility: &Objects{},
		Textures:       &types.StringArray{},
		UvAnim: &UvAnim{
			Offset: &ArrayMisc{},
			Scale:  &ArrayMisc{},
		},
	}
	(*r)[key] = renderController
	return renderController
}
