package animation_controller

import (
	"github.com/respectZ/glowstone/types"
)

func (a *animationControllerData) GetState(name string) *animationControllerState {
	return a.States[name]
}

func (a *animationControllerData) SetState(name string, state *animationControllerState) {
	a.States[name] = state
}

func (a *animationControllerData) SetInitialState(name string) {
	a.InitialState = name
}

func (a *animationControllerData) NewState(name string) *animationControllerState {
	v := &animationControllerState{
		OnEntry:     &types.StringArray{},
		OnExit:      &types.StringArray{},
		Animations:  &types.StringArrayConditional{},
		Transitions: &types.MapStringArray{},
	}
	a.States[name] = v
	return v
}
