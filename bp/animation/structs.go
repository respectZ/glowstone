package animation

import (
	g_util "github.com/respectZ/glowstone/util"
)

type animation struct {
	FormatVersion string                    `json:"format_version"`
	Animations    map[string]*animationData `json:"animations"`
}

type animationData struct {
	AnimationLength float64             `json:"animation_length"`
	Loop            bool                `json:"loop"`
	AnimUpdateTime  string              `json:"anim_time_update,omitempty"`
	Timeline        map[string][]string `json:"timeline"`
}

// New returns a new animation
//
// Example:
//
//	a := animation.New()
func New() *animation {
	return &animation{
		FormatVersion: "1.12.0",
		Animations:    make(map[string]*animationData),
	}
}

// Load loads an animation from a file
//
// Example:
//
//	a, err := animation.Load("player.animation.json")
func Load(src string) (*animation, error) {
	a := New()
	err := g_util.LoadJSON(src, a)
	return a, err
}