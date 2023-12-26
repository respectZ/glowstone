package animation_controller

import (
	rp "github.com/respectZ/glowstone/rp/animation_controller"
)

type RPAnimationController struct {
	Data *rp.AnimationControllerFile

	// Other stuff
	Dest string
}

// New creates a new RPAnimationController. Dest doesn't need to start with "animation_controllers/"
//
// Example:
//
//	a := NewRP("player.animation_controller.json")
//	b := NewRP("player/effect.animation_controller.json")
func NewRP(dest string) *RPAnimationController {
	return &RPAnimationController{
		Data: rp.New(),
		Dest: dest,
	}
}

// LoadRP loads the RP from the given path. This is not setting the Dest.
//
// Example:
//
//	a, err := animation_controller.Load("player.animation_controller.json")
func LoadRP(src string) (*RPAnimationController, error) {
	a, err := rp.Load(src)
	if err != nil {
		return nil, err
	}
	return &RPAnimationController{
		Data: a,
	}, nil
}

// Encode returns []byte of the RP.
//
// Example:
//
//	rp, err := e.Encode()
func (a *RPAnimationController) Encode() ([]byte, error) {
	return a.Data.Encode()
}
