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
	OnEntry     []string            `json:"on_entry,omitempty"`
	OnExit      []string            `json:"on_exit,omitempty"`
	Animations  []string            `json:"animations,omitempty"`
	Transitions []map[string]string `json:"transitions,omitempty"`
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
