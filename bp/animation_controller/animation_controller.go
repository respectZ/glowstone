package animation_controller

import (
	g_util "github.com/respectZ/glowstone/util"
)

func (a *animationController) Encode() ([]byte, error) {
	// Loop controllers
	for _, v := range a.AnimationController {
		// Loop states
		for _, state := range v.States {
			if state.OnEntry.IsEmpty() {
				state.OnEntry = nil
			}
			if state.OnExit.IsEmpty() {
				state.OnExit = nil
			}
			if state.Animations.IsEmpty() {
				state.Animations = nil
			}
			if state.Transitions.IsEmpty() {
				state.Transitions = nil
			}
		}
	}
	return g_util.MarshalJSON(a)
}

func (a *animationController) New(name string) AnimationControllerData {
	v := &animationControllerData{
		States: make(map[string]*animationControllerState),
	}
	// Create a default state
	v.NewState("default")
	a.AnimationController[name] = v
	return v
}

func (a *animationController) Get(name string) AnimationControllerData {
	return a.AnimationController[name]
}

func (a *animationController) Remove(name string) {
	delete(a.AnimationController, name)
}
