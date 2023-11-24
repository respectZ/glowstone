package animation_controller

import (
	g_util "github.com/respectZ/glowstone/util"
)

type animationController struct {
	FormatVersion       string                              `json:"format_version"`
	AnimationController map[string]*animationControllerData `json:"animation_controllers"`
}

type animationControllerData struct {
	InitialState string                               `json:"initial_state,omitempty"`
	States       map[string]*animationControllerState `json:"states"`
}

type animationControllerState struct {
	BlendTransition      float64             `json:"blend_transition,omitempty"`
	BlendViaShortestPath *bool               `json:"blend_via_shortest_path,omitempty"`
	OnEntry              []string            `json:"on_entry,omitempty"`
	OnExit               []string            `json:"on_exit,omitempty"`
	Animations           []interface{}       `json:"animations,omitempty"`
	Transitions          []map[string]string `json:"transitions,omitempty"`

	ParticleEffects []*particleEffect     `json:"particle_effects,omitempty"`
	Parameters      []interface{}         `json:"parameters,omitempty"` // TODO: implement
	Variables       map[string]*variables `json:"variables,omitempty"`
	SoundEffects    []*soundEffect        `json:"sound_effects,omitempty"`
}

type particleEffect struct {
	BindToActor     *bool  `json:"bind_to_actor,omitempty"`
	Effect          string `json:"effect,omitempty"`
	Locator         string `json:"locator,omitempty"`
	PreEffectScript string `json:"pre_effect_script,omitempty"`
}

type variables struct {
	Input      string             `json:"input,omitempty"`
	RemapCurve map[string]float64 `json:"remap_curve,omitempty"`
}

type soundEffect struct {
	Effect string `json:"effect,omitempty"`
}

// New returns a new animation controller
//
// Example:
//
//	a := animation_controller.New()
func New() AnimationController {
	return &animationController{
		FormatVersion:       "1.12.0",
		AnimationController: make(map[string]*animationControllerData),
	}
}

// Load loads an animation controller from a file
//
// Example:
//
//	a, err := animation_controller.Load("player.animation_controller.json")
func Load(src string) (AnimationController, error) {
	a := New()
	err := g_util.LoadJSON(src, a)
	return a, err
}
