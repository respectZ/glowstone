package attachable

import (
	g_util "github.com/respectZ/glowstone/util"
)

func (a *Attachable) Encode(minify ...bool) ([]byte, error) {
	if a.Attachable.Description.Animations.IsEmpty() {
		a.Attachable.Description.Animations = nil
	}
	if a.Attachable.Description.Item.IsEmpty() {
		a.Attachable.Description.Item = nil
	}
	return g_util.MarshalJSON(a, minify...)
}

func (a *Attachable) GetIdentifier() string {
	return a.Attachable.Description.Identifier
}
