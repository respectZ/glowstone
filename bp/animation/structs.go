package animation

import (
	g_util "github.com/respectZ/glowstone/util"
)

type Animation struct {
	FormatVersion string        `json:"format_version"`
	Animations    IAnimationMap `json:"animations"`
}

type animationData struct {
	AnimationLength float64   `json:"animation_length"`
	Loop            bool      `json:"loop"`
	AnimUpdateTime  string    `json:"anim_time_update,omitempty"`
	Timeline        ITimeline `json:"timeline"`
}

// New returns a new animation
//
// Example:
//
//	a := animation.New()
func New() *Animation {
	anim := &Animation{
		FormatVersion: "1.12.0",
		Animations:    IAnimationMap(&AnimationMap{}),
	}
	return anim
}

// Load loads an animation from a file
//
// Example:
//
//	a, err := animation.Load("player.animation.json")
func Load(src string) (*Animation, error) {
	a := New()
	err := g_util.LoadJSON(src, a)
	return a, err
}
