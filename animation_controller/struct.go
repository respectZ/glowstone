package animation_controller

import (
	bp "github.com/respectZ/glowstone/bp/animation_controller"
)

type BPAnimationController struct {
	Data *bp.AnimationControllerFile

	// Other stuff
	Dest string
}

// New creates a new BPAnimationController. Dest doesn't need to start with "animation_controllers/"
//
// Example:
//
//	a := NewBP("player.animation_controller.json")
//	b := NewBP("player/effect.animation_controller.json")
func NewBP(dest string) *BPAnimationController {
	return &BPAnimationController{
		Data: bp.New(),
		Dest: dest,
	}
}

// LoadBP loads the BP from the given path. This is not setting the Dest.
//
// Example:
//
//	a, err := animation_controller.Load("player.animation_controller.json")
func LoadBP(src string) (*BPAnimationController, error) {
	a, err := bp.Load(src)
	if err != nil {
		return nil, err
	}
	return &BPAnimationController{
		Data: a,
	}, nil
}

// Encode returns []byte of the BP.
//
// Example:
//
//	bp, err := e.Encode()
func (a *BPAnimationController) Encode() ([]byte, error) {
	return a.Data.Encode()
}
