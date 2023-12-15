package animation_controller

import (
	g_util "github.com/respectZ/glowstone/util"
)

type AnimationControllerState map[string]*animationControllerState

type IAnimationControllerState interface {
	UnmarshalJSON([]byte) error
	Add(string, *animationControllerState)
	Get(string) (*animationControllerState, bool)
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

func (a *AnimationControllerState) UnmarshalJSON(data []byte) error {
	var temp map[string]*animationControllerState
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}

func (a *AnimationControllerState) Add(name string, state *animationControllerState) {
	if *a == nil {
		*a = make(AnimationControllerState)
	}
	(*a)[name] = state
}

func (a *AnimationControllerState) Get(name string) (*animationControllerState, bool) {
	if *a == nil {
		*a = make(AnimationControllerState)
	}
	state, ok := (*a)[name]
	return state, ok
}

func (a *AnimationControllerState) Clear() {
	if *a == nil {
		*a = make(AnimationControllerState)
	}
	*a = make(AnimationControllerState)
}
