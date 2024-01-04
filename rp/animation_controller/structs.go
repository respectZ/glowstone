package animation_controller

import (
	"github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

type AnimationControllerFile struct {
	FormatVersion        string               `json:"format_version"`
	AnimationControllers IAnimationController `json:"animation_controllers"`
}

type animationController struct {
	InitialState string                     `json:"initial_state,omitempty"`
	States       IAnimationControllerStates `json:"states"`
}

type animationControllerState struct {
	BlendTransition      float64                       `json:"blend_transition,omitempty"`
	BlendViaShortestPath *bool                         `json:"blend_via_shortest_path,omitempty"`
	OnEntry              types.IStringArray            `json:"on_entry,omitempty"`
	OnExit               types.IStringArray            `json:"on_exit,omitempty"`
	Animations           types.IStringArrayConditional `json:"animations,omitempty"`
	Transitions          types.IMapStringArray         `json:"transitions,omitempty"`

	ParticleEffects IParticleEffects `json:"particle_effects,omitempty"` // TODO: fix this since it can be single object or array
	Parameters      []interface{}    `json:"parameters,omitempty"`       // TODO: implement
	Variables       IVariables       `json:"variables,omitempty"`
	SoundEffects    ISoundEffects    `json:"sound_effects,omitempty"` // TODO: fix this since it can be single object or array
}

type particleEffect struct {
	BindToActor     *bool  `json:"bind_to_actor,omitempty"`
	Effect          string `json:"effect,omitempty"`
	Locator         string `json:"locator,omitempty"`
	PreEffectScript string `json:"pre_effect_script,omitempty"`
}

type variables struct {
	Input      string                `json:"input,omitempty"`
	RemapCurve types.IMapStringFloat `json:"remap_curve,omitempty"`
}

type soundEffect struct {
	Effect string `json:"effect,omitempty"`
}

// New returns a new animation controller
//
// Example:
//
//	a := animation_controller.New()
func New() *AnimationControllerFile {
	return &AnimationControllerFile{
		FormatVersion:        "1.12.0",
		AnimationControllers: IAnimationController(&AnimationController{}),
	}
}

// Load loads an animation controller from a file
//
// Example:
//
//	a, err := animation_controller.Load("player.animation_controller.json")
func Load(src string) (*AnimationControllerFile, error) {
	a := New()
	err := g_util.LoadJSON(src, a)
	return a, err
}

func (a *AnimationControllerFile) Encode() ([]byte, error) {
	for _, v := range a.AnimationControllers.All() {
		for _, v := range v.States.All() {
			if v.Animations != nil && v.Animations.IsEmpty() {
				v.Animations = nil
			}
			if v.OnEntry != nil && v.OnEntry.IsEmpty() {
				v.OnEntry = nil
			}
			if v.OnExit != nil && v.OnExit.IsEmpty() {
				v.OnExit = nil
			}
			if v.Transitions != nil && v.Transitions.IsEmpty() {
				v.Transitions = nil
			}
			if v.ParticleEffects != nil && v.ParticleEffects.IsEmpty() {
				v.ParticleEffects = nil
			}
			if v.Variables != nil {
				_variables := v.Variables.All()
				for k, vv := range _variables {
					if vv.Input == "" {
						delete(_variables, k)
					}
				}
				if v.Variables.IsEmpty() {
					v.Variables = nil
				}
			}
			if v.SoundEffects != nil && v.SoundEffects.IsEmpty() {
				v.SoundEffects = nil
			}
		}
	}
	if a.AnimationControllers != nil && a.AnimationControllers.IsEmpty() {
		a.AnimationControllers = nil
	}
	return g_util.MarshalJSON(a)
}
