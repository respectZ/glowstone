package glowstone

import (
	"github.com/respectZ/glowstone/animation"
	g_util "github.com/respectZ/glowstone/util"
)

type AnimationBP map[string]*animation.BPAnimation

type IAnimationBP interface {
	Add(string, *animation.BPAnimation)
	Get(string) (*animation.BPAnimation, bool)
	Remove(string)
	Clear()
	All() map[string]*animation.BPAnimation
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	New(string) *animation.BPAnimation
}

func (m *AnimationBP) Add(key string, value *animation.BPAnimation) {
	if *m == nil {
		*m = make(map[string]*animation.BPAnimation)
	}
	(*m)[key] = value
}

func (m *AnimationBP) Get(key string) (*animation.BPAnimation, bool) {
	if *m == nil {
		*m = make(map[string]*animation.BPAnimation)
	}
	value, ok := (*m)[key]
	return value, ok
}

func (m *AnimationBP) Remove(key string) {
	if *m == nil {
		*m = make(map[string]*animation.BPAnimation)
	}
	delete(*m, key)
}

func (m *AnimationBP) Clear() {
	if *m == nil {
		*m = make(map[string]*animation.BPAnimation)
	}
	*m = make(map[string]*animation.BPAnimation)
}

func (m *AnimationBP) All() map[string]*animation.BPAnimation {
	if *m == nil {
		*m = make(map[string]*animation.BPAnimation)
	}
	return *m
}

func (m *AnimationBP) IsEmpty() bool {
	if *m == nil {
		*m = make(map[string]*animation.BPAnimation)
	}
	return len(*m) == 0
}

func (m *AnimationBP) Size() int {
	if *m == nil {
		*m = make(map[string]*animation.BPAnimation)
	}
	return len(*m)
}

func (m *AnimationBP) UnmarshalJSON(data []byte) error {
	if *m == nil {
		*m = make(map[string]*animation.BPAnimation)
	}
	return g_util.UnmarshalJSON(data, m)
}

func (m *AnimationBP) New(dest string) *animation.BPAnimation {
	a := animation.NewBP(dest)
	m.Add(dest, a)
	return a
}
