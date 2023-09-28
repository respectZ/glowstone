package animation

import (
	g_util "github.com/respectZ/glowstone/util"
)

func (a *animation) Encode() ([]byte, error) {
	bp, err := g_util.MarshalJSON(a)
	if err != nil {
		return nil, err
	}
	return bp, nil
}

func (a *animation) NewAnimation(name string) *animationData {
	a.Animations[name] = &animationData{
		AnimationLength: 0.0,
		Loop:            false,
		Timeline:        make(map[string][]string),
	}
	return a.Animations[name]
}

func (a *animation) GetAnimation(name string) *animationData {
	return a.Animations[name]
}

func (a *animation) RemoveAnimation(name string) {
	delete(a.Animations, name)
}
