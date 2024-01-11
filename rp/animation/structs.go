package animation

import (
	g_util "github.com/respectZ/glowstone/util"
)

type AnimationDefinition struct {
	FormatVersion string     `json:"format_version"`
	Animations    IAnimation `json:"animations,omitempty"`
}

type animation struct {
	Loop                      interface{}               `json:"loop,omitempty"`
	Bones                     map[string]*animationBone `json:"bones,omitempty"`
	AnimTimeUpdate            string                    `json:"anim_time_update,omitempty"`
	AnimationLength           float64                   `json:"animation_length,omitempty"`
	BlendWeight               float64                   `json:"blend_weight,omitempty"`
	LoopDelay                 float64                   `json:"loop_delay,omitempty"`
	OverridePreviousAnimation bool                      `json:"override_previous_animation,omitempty"`

	// Not implented. TODO: implement
	ParticleEffects interface{} `json:"particle_effects,omitempty"`
	// Not implented. TODO: implement
	SoundEffects interface{} `json:"sound_effects,omitempty"`

	StartDelay float64     `json:"start_delay,omitempty"`
	Timeline   interface{} `json:"timeline,omitempty"` // Not sure what this is
}

/** Animation Bone **/

type animationBone struct {
	Position interface{} `json:"position,omitempty"`
	Rotation interface{} `json:"rotation,omitempty"`
	Scale    interface{} `json:"scale,omitempty"`
}

/** funcs **/

func New() *AnimationDefinition {
	return &AnimationDefinition{
		FormatVersion: "1.12.0",
		Animations:    &Animation{},
	}
}

func Load(src string) (*AnimationDefinition, error) {
	a := New()
	err := g_util.LoadJSON(src, a)
	return a, err
}

func (a *AnimationDefinition) Encode(minify ...bool) ([]byte, error) {
	return g_util.MarshalJSON(a, minify...)
}
