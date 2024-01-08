package rp

import (
	"fmt"
	"strings"

	rp "github.com/respectZ/glowstone/rp/animation"
)

type AnimationFile struct {
	Data *rp.AnimationDefinition

	// Other stuff
}

func LoadAnimation(src string) (*rp.AnimationDefinition, error) {
	animation, err := rp.Load(src)
	if err != nil {
		return nil, err
	}
	return animation, nil
}

func (e *AnimationFile) Encode() ([]byte, error) {
	animation, err := e.Data.Encode()
	if err != nil {
		return nil, err
	}
	return animation, nil
}

// Returns all animations that exist in the animation file.
//
// Example:
//
//	animations := e.GetAnimations() // []string{"animation.player.move", "animation.player.idle"}
func (e *AnimationFile) GetAnimations() []string {
	var animations []string
	for k := range e.Data.Animations.All() {
		animations = append(animations, k)
	}
	return animations
}

// Returns all animations that exist in the animation file as a map.
// The key is the animation name without "animation." prefix.
//
// Example:
//
//	animations := e.GetAnimationsAsMap() // map[string]string{"player.move": "animation.player.move", "player.idle": "animation.player.idle"}
func (e *AnimationFile) GetAnimationsAsMap() map[string]string {
	animations := make(map[string]string)
	for k := range e.Data.Animations.All() {
		animations[strings.TrimPrefix(k, "animation.")] = k
	}
	return animations
}

// Remap animations to match the animation name with the file name.
//
//	"animation.asdfghijkl.move" -> "animation.player.move"
//
// Example:
//
//	e.RemapAnimations("player")
func (e *AnimationFile) RemapAnimations(filename string) {
	for k, v := range e.Data.Animations.All() {
		s := strings.Split(k, ".")
		animName := strings.Join(s[2:], ".")
		newKey := fmt.Sprintf("animation.%s.%s", filename, animName)
		if newKey == k {
			continue
		}

		e.Data.Animations.Add(fmt.Sprintf("animation.%s.%s", filename, animName), v)
		e.Data.Animations.Remove(k)
	}
}
