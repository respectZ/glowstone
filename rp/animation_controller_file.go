package rp

import (
	"fmt"
	"strings"

	rp "github.com/respectZ/glowstone/rp/animation_controller"
)

type AnimationControllerFile struct {
	Data *rp.AnimationControllerFile

	// Other stuff
}

// Create a new animation controller file
//
// Example:
//
//	animation_controller := rp.NewAnimationController("vanilla/player.animation.controller.json")
func NewAnimationController(dest string) *AnimationControllerFile {
	return &AnimationControllerFile{
		Data: rp.New(),
	}
}

func LoadAnimationController(src string) (*rp.AnimationControllerFile, error) {
	animation, err := rp.Load(src)
	if err != nil {
		return nil, err
	}
	return animation, nil
}

func (e *AnimationControllerFile) Encode() ([]byte, error) {
	animation, err := e.Data.Encode()
	if err != nil {
		return nil, err
	}
	return animation, nil
}

// Returns all animations that exist in the animation controller file.
//
// Example:
//
//	animations := e.GetAnimations() // []string{"controller.animation.player.move", "controller.animation.player.idle"}
func (e *AnimationControllerFile) GetAnimations() []string {
	var animations []string
	for k := range e.Data.AnimationControllers.All() {
		animations = append(animations, k)
	}
	return animations
}

// Returns all animations that exist in the animation file as a map.
// The key is the animation name, but the "controller.animation." prefix is replaced with "controller.".
//
// Example:
//
//	animations := e.GetAnimationsAsMap() // map[string]string{"controller.player.move": "animation.player.move", "controller.player.idle": "animation.player.idle"}
func (e *AnimationControllerFile) GetAnimationsAsMap() map[string]string {
	animations := make(map[string]string)
	for k := range e.Data.AnimationControllers.All() {
		// controller.animation.player.move -> controller.player.move
		short := strings.TrimPrefix(k, "controller.animation.")
		short = fmt.Sprintf("controller.%s", short)
		animations[short] = k
	}
	return animations
}
