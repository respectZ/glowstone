package animation_controller

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
	v := &animationControllerState{}
	a.States[name] = v
	return v
}
