package attachable

import (
	"reflect"

	g_util "github.com/respectZ/glowstone/util"
)

func (a *Attachable) Encode(minify ...bool) ([]byte, error) {
	if a.Attachable.Description.Animations.IsEmpty() {
		a.Attachable.Description.Animations = nil
	}
	if a.Attachable.Description.Item.IsEmpty() {
		a.Attachable.Description.Item = nil
	}
	if a.Attachable.Description.Scripts.ParentSetup.IsEmpty() {
		a.Attachable.Description.Scripts.ParentSetup = nil
	}
	if a.Attachable.Description.Scripts.Animate.IsEmpty() {
		a.Attachable.Description.Scripts.Animate = nil
	}
	if a.Attachable.Description.Scripts.PreAnimation.IsEmpty() {
		a.Attachable.Description.Scripts.PreAnimation = nil
	}
	if v := reflect.ValueOf(a.Attachable.Description.RenderControllers).Elem(); v.Len() == 0 {
		a.Attachable.Description.RenderControllers = nil
	}
	if reflect.ValueOf(a.Attachable.Description.Scripts).Elem().IsZero() {
		a.Attachable.Description.Scripts = nil
	}
	return g_util.MarshalJSON(a, minify...)
}

func (a *Attachable) GetIdentifier() string {
	return a.Attachable.Description.Identifier
}
