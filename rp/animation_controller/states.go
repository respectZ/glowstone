package animation_controller

import (
	"github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

type AnimationControllerStates map[string]*animationControllerState

type IAnimationControllerStates interface {
	UnmarshalJSON([]byte) error
	Add(string, *animationControllerState)
	Get(string) *animationControllerState
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]*animationControllerState

	// Create a new state.
	//
	// name: The name of the state.
	//
	// return: The new state.
	//
	// Example:
	//
	//     state := animation_controller.States.New("default")
	New(string) *animationControllerState
}

type animationControllerState_parse struct {
	BlendTransition      float64                      `json:"blend_transition,omitempty"`
	BlendViaShortestPath *bool                        `json:"blend_via_shortest_path,omitempty"`
	OnEntry              types.StringArray            `json:"on_entry,omitempty"`
	OnExit               types.StringArray            `json:"on_exit,omitempty"`
	Animations           types.StringArrayConditional `json:"animations,omitempty"`
	Transitions          types.MapStringArray         `json:"transitions,omitempty"`

	ParticleEffects ParticleEffects `json:"particle_effects,omitempty"` // TODO: fix this since it can be single object or array
	Parameters      []interface{}   `json:"parameters,omitempty"`       // TODO: implement
	Variables       Variables       `json:"variables,omitempty"`
	SoundEffects    SoundEffects    `json:"sound_effects,omitempty"` // TODO: fix this since it can be single object or array
}

func (a *AnimationControllerStates) UnmarshalJSON(data []byte) error {
	var temp map[string]*animationControllerState_parse
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = make(AnimationControllerStates)
	for k, v := range temp {
		(*a)[k] = &animationControllerState{
			BlendTransition:      v.BlendTransition,
			BlendViaShortestPath: v.BlendViaShortestPath,
			OnEntry:              &v.OnEntry,
			OnExit:               &v.OnExit,
			Animations:           &v.Animations,
			Transitions:          &v.Transitions,
			ParticleEffects:      &v.ParticleEffects,
			Variables:            &v.Variables,
			SoundEffects:         &v.SoundEffects,
		}
	}
	return nil
}

func (a *AnimationControllerStates) New(name string) *animationControllerState {
	v := &animationControllerState{
		OnEntry:         &types.StringArray{},
		OnExit:          &types.StringArray{},
		Animations:      &types.StringArrayConditional{},
		Transitions:     &types.MapStringArray{},
		ParticleEffects: &ParticleEffects{},
		Variables:       &Variables{},
		SoundEffects:    &SoundEffects{},
	}
	(*a)[name] = v
	return v
}

func (a *AnimationControllerStates) Add(name string, state *animationControllerState) {
	(*a)[name] = state
}

func (a *AnimationControllerStates) Get(name string) *animationControllerState {
	return (*a)[name]
}

func (a *AnimationControllerStates) Clear() {
	*a = make(AnimationControllerStates)
}

func (a *AnimationControllerStates) Remove(name string) {
	delete(*a, name)
}

func (a *AnimationControllerStates) IsEmpty() bool {
	return len(*a) == 0
}

func (a *AnimationControllerStates) Size() int {
	return len(*a)
}

func (a *AnimationControllerStates) All() map[string]*animationControllerState {
	return *a
}
