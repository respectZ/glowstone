package animation_controller

import (
	"github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

type AnimationControllerFile struct {
	FormatVersion       string               `json:"format_version"`
	AnimationController IAnimationController `json:"animation_controllers"`
}

type animationController struct {
	InitialState string                     `json:"initial_state,omitempty"`
	States       IAnimationControllerStates `json:"states"`
}

type animationControllerState struct {
	OnEntry     types.IStringArray            `json:"on_entry,omitempty"`
	OnExit      types.IStringArray            `json:"on_exit,omitempty"`
	Animations  types.IStringArrayConditional `json:"animations,omitempty"`
	Transitions types.IMapStringArray         `json:"transitions,omitempty"`
}

// New returns a new animation controller
//
// Example:
//
//	a := animation_controller.New()
func New() *AnimationControllerFile {
	return &AnimationControllerFile{
		FormatVersion:       "1.12.0",
		AnimationController: &AnimationController{},
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
	for _, v := range a.AnimationController.All() {
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
		}
	}
	if a.AnimationController != nil && a.AnimationController.IsEmpty() {
		a.AnimationController = nil
	}
	return g_util.MarshalJSON(a)
}
