package entity

import (
	types "github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

type clientEntityDescriptionScripts struct {
	ParentSetup  string                        `json:"parent_setup,omitempty"`
	Variables    types.IMapStringString        `json:"variables,omitempty"`
	ScaleX       string                        `json:"scalex,omitempty"`
	ScaleY       string                        `json:"scaley,omitempty"`
	ScaleZ       string                        `json:"scalez,omitempty"`
	Scale        string                        `json:"scale,omitempty"`
	Initialize   types.IStringArray            `json:"initialize,omitempty"`
	PreAnimation types.IStringArray            `json:"pre_animation,omitempty"`
	Animate      types.IStringArrayConditional `json:"animate,omitempty"`
}

func (e *clientEntityDescriptionScripts) UnmarshalJSON(data []byte) error {
	type Alias clientEntityDescriptionScripts

	aux := &struct {
		ParentSetup  string                        `json:"parent_setup,omitempty"`
		Variables    types.IMapStringString        `json:"variables,omitempty"`
		ScaleX       string                        `json:"scalex,omitempty"`
		ScaleY       string                        `json:"scaley,omitempty"`
		ScaleZ       string                        `json:"scalez,omitempty"`
		Scale        string                        `json:"scale,omitempty"`
		Initialize   types.IStringArray            `json:"initialize,omitempty"`
		PreAnimation types.IStringArray            `json:"pre_animation,omitempty"`
		Animate      types.IStringArrayConditional `json:"animate,omitempty"`
	}{
		Variables:    &types.MapStringString{},
		Initialize:   &types.StringArray{},
		PreAnimation: &types.StringArray{},
		Animate:      &types.StringArrayConditional{},
	}

	if err := g_util.UnmarshalJSON(data, (*Alias)(aux)); err != nil {
		return err
	}

	*e = clientEntityDescriptionScripts(*aux)

	return nil
}
