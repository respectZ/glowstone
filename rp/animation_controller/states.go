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

func (a *AnimationControllerStates) UnmarshalJSON(data []byte) error {
	var temp map[string]*animationControllerState
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = temp
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
