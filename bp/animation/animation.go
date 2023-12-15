package animation

import (
	g_util "github.com/respectZ/glowstone/util"
)

type AnimationMap map[string]*animationData

type IAnimationMap interface {
	UnmarshalJSON([]byte) error
	Add(string, *animationData)
	Get(string) (*animationData, bool)
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]*animationData

	// Create a new animation.
	//
	// name: The name of the animation.
	//
	// return: The new animation.
	//
	// Example:
	//
	//     animation := animation.Animations.New("animation.zombie.walk")
	New(string) *animationData
}

func (m *AnimationMap) UnmarshalJSON(data []byte) error {
	var temp map[string]*animationData
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*m = temp
	return nil
}

func (m *AnimationMap) Add(name string, animation *animationData) {
	if *m == nil {
		*m = make(AnimationMap)
	}
	(*m)[name] = animation
}

func (m *AnimationMap) Get(name string) (*animationData, bool) {
	if *m == nil {
		*m = make(AnimationMap)
	}
	animation, ok := (*m)[name]
	return animation, ok
}

func (m *AnimationMap) Clear() {
	if *m == nil {
		*m = make(AnimationMap)
	}
	*m = make(AnimationMap)
}

func (m *AnimationMap) Remove(name string) {
	if *m == nil {
		*m = make(AnimationMap)
	}
	delete(*m, name)
}

func (m *AnimationMap) IsEmpty() bool {
	return m.Size() == 0
}

func (m *AnimationMap) Size() int {
	if *m == nil {
		*m = make(AnimationMap)
	}
	return len(*m)
}

func (m *AnimationMap) All() map[string]*animationData {
	if *m == nil {
		*m = make(AnimationMap)
	}
	return *m
}

func (m *AnimationMap) New(name string) *animationData {
	m.Add(name, &animationData{
		AnimationLength: 0.0,
		Loop:            false,
		Timeline:        ITimeline(&Timeline{}),
	})
	return (*m)[name]
}

func (a *Animation) Encode() ([]byte, error) {
	bp, err := g_util.MarshalJSON(a)
	if err != nil {
		return nil, err
	}
	return bp, nil
}
