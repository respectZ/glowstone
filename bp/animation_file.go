package bp

import (
	"strings"

	bp "github.com/respectZ/glowstone/bp/animation"
)

type AnimationFile struct {
	Data *bp.Animation

	// Other stuff
	Subdir string
}

func LoadAnimation(src string) (*bp.Animation, error) {
	animation, err := bp.Load(src)
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
