package rendercontroller

import (
	"reflect"

	"github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

var FORMAT_VERSION = "1.19.70"

type RenderControllerDefinition struct {
	FormatVersion     string             `json:"format_version"`
	RenderControllers IRenderControllers `json:"render_controllers"`
}

type RenderControllerArrays struct {
	Geometries types.IMapStringArray `json:"geometries,omitempty"`
	Textures   types.IMapStringArray `json:"textures,omitempty"`
	Materials  types.IMapStringArray `json:"matrices,omitempty"`
}

type Color struct {
	R interface{} `json:"r,omitempty"`
	G interface{} `json:"g,omitempty"`
	B interface{} `json:"b,omitempty"`
	A interface{} `json:"a,omitempty"`
}

func New() *RenderControllerDefinition {
	return &RenderControllerDefinition{
		FormatVersion:     FORMAT_VERSION,
		RenderControllers: &RenderControllers{},
	}
}

func Load(src string) (*RenderControllerDefinition, error) {
	r := New()
	err := g_util.LoadJSON(src, r)
	return r, err
}

func (r *RenderControllerDefinition) Encode(minify ...bool) ([]byte, error) {
	for _, v := range r.RenderControllers.All() {
		if v.Arrays.Geometries.IsEmpty() {
			v.Arrays.Geometries = nil
		}
		if v.Arrays.Textures.IsEmpty() {
			v.Arrays.Textures = nil
		}
		if v.Arrays.Materials.IsEmpty() {
			v.Arrays.Materials = nil
		}
		if reflect.ValueOf(v.Arrays).Elem().IsZero() {
			v.Arrays = nil
		}

		if v.PartVisibility.IsEmpty() {
			v.PartVisibility = nil
		}

		if v.UvAnim.Offset.IsEmpty() {
			v.UvAnim.Offset = nil
		}
		if v.UvAnim.Scale.IsEmpty() {
			v.UvAnim.Scale = nil
		}
		if reflect.ValueOf(v.UvAnim).Elem().IsZero() {
			v.UvAnim = nil
		}
	}

	return g_util.MarshalJSON(r, minify...)
}

// Set "color" to "{}".
func (r *RenderController) SetColor() {
	r.Color = &Color{}
}

// Set "is_hurt_color" to "{}".
func (r *RenderController) SetIsHurtColor() {
	r.IsHurtColor = &Color{}
}

// Set "on_fire_color" to "{}".
func (r *RenderController) SetOnFireColor() {
	r.OnFireColor = &Color{}
}

// Set "overlay_color" to "{}".
func (r *RenderController) SetOverlayColor() {
	r.OverlayColor = &Color{}
}
