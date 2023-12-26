package animation_controller

import (
	"github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

type AnimationController map[string]*animationController

type IAnimationController interface {
	UnmarshalJSON([]byte) error
	Add(string, *animationController)
	Get(string) *animationController
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]*animationController

	// Create a new animation controller.
	//
	// name: The name of the animation controller.
	//
	// return: The new animation controller.
	//
	// Example:
	//
	//  animation_controller := animation_controller.New("controller.animation.zombie.dance")
	New(string) *animationController
}

func (a *AnimationController) UnmarshalJSON(data []byte) error {
	var temp map[string]*animationController
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}

func (a *AnimationController) Add(name string, animationController *animationController) {
	(*a)[name] = animationController
}

func (a *AnimationController) Get(name string) *animationController {
	return (*a)[name]
}

func (a *AnimationController) Clear() {
	*a = make(AnimationController)
}

func (a *AnimationController) Remove(name string) {
	delete(*a, name)
}

func (a *AnimationController) IsEmpty() bool {
	return len(*a) == 0
}

func (a *AnimationController) Size() int {
	return len(*a)
}

func (a *AnimationController) All() map[string]*animationController {
	return *a
}

func (a *AnimationController) New(name string) *animationController {
	v := &animationController{
		InitialState: "",
		States: &AnimationControllerStates{
			"default": &animationControllerState{
				OnEntry:     &types.StringArray{},
				OnExit:      &types.StringArray{},
				Animations:  &types.StringArrayConditional{},
				Transitions: &types.MapStringArray{},
			},
		},
	}
	(*a)[name] = v
	return v
}
