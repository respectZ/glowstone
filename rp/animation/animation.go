package animation

import (
	g_util "github.com/respectZ/glowstone/util"
)

type Animation map[string]*animation

type IAnimation interface {
	UnmarshalJSON([]byte) error
	Add(string, *animation)
	Get(string) *animation
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]*animation

	// Create a new animation.
	//
	// Example:
	//
	//	a := New("animation.weapon.run")
	New(string) *animation
}

func (a *Animation) UnmarshalJSON(data []byte) error {
	var temp map[string]*animation
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}

func (a *Animation) Add(s string, v *animation) {
	(*a)[s] = v
}

func (a *Animation) Get(s string) *animation {
	return (*a)[s]
}

func (a *Animation) Clear() {
	*a = make(Animation, 0)
}

func (a *Animation) Remove(s string) {
	delete(*a, s)
}

func (a *Animation) IsEmpty() bool {
	return len(*a) == 0
}

func (a *Animation) Size() int {
	return len(*a)
}

func (a *Animation) All() map[string]*animation {
	return *a
}

func (a *Animation) New(identifier string) *animation {
	animation := &animation{
		Bones: make(map[string]*animationBone),
	}
	a.Add(identifier, animation)
	return animation
}
