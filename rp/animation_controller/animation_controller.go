package animation_controller

import (
	g_util "github.com/respectZ/glowstone/util"
)

func (a *animationController) Encode() ([]byte, error) {
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
