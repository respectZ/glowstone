package glowstone

import (
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/animation"
	g_util "github.com/respectZ/glowstone/util"
)

// We don't need options atm, so we use bp/animation_controller.go

type AnimationBP map[string]*bp.Animation

type IAnimationBP interface {
	Add(string, *bp.Animation)
	Get(string) (*bp.Animation, bool)
	Remove(string)
	Clear()
	All() map[string]*bp.Animation
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// New creates a new BPAnimation. Dest doesn't need to start with "animation/"
	//
	// Example:
	//
	//	a := New("player.animation.json")
	//	b := New("player/effect.animation.json")
	New(string) *bp.Animation

	Save(string) error

	LoadAll(string) error
}

func (m *AnimationBP) Add(key string, value *bp.Animation) {
	if *m == nil {
		*m = make(map[string]*bp.Animation)
	}
	(*m)[key] = value
}

func (m *AnimationBP) Get(key string) (*bp.Animation, bool) {
	if *m == nil {
		*m = make(map[string]*bp.Animation)
	}
	value, ok := (*m)[key]
	return value, ok
}

func (m *AnimationBP) Remove(key string) {
	if *m == nil {
		*m = make(map[string]*bp.Animation)
	}
	delete(*m, key)
}

func (m *AnimationBP) Clear() {
	if *m == nil {
		*m = make(map[string]*bp.Animation)
	}
	*m = make(map[string]*bp.Animation)
}

func (m *AnimationBP) All() map[string]*bp.Animation {
	if *m == nil {
		*m = make(map[string]*bp.Animation)
	}
	return *m
}

func (m *AnimationBP) IsEmpty() bool {
	if *m == nil {
		*m = make(map[string]*bp.Animation)
	}
	return len(*m) == 0
}

func (m *AnimationBP) Size() int {
	if *m == nil {
		*m = make(map[string]*bp.Animation)
	}
	return len(*m)
}

func (m *AnimationBP) UnmarshalJSON(data []byte) error {
	if *m == nil {
		*m = make(map[string]*bp.Animation)
	}
	return g_util.UnmarshalJSON(data, m)
}

func (m *AnimationBP) New(dest string) *bp.Animation {
	a := bp.New()
	m.Add(dest, a)
	return a
}

func (m *AnimationBP) Save(pathToBP string) error {
	if *m == nil {
		*m = make(map[string]*bp.Animation)
	}
	for dest, v := range m.All() {
		data, err := v.Encode()
		if err != nil {
			return err
		}
		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.Animation, dest), data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *AnimationBP) LoadAll(pathToBP string) error {
	if *m == nil {
		*m = make(map[string]*bp.Animation)
	}
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.Animation))
	if err != nil {
		return err
	}
	for _, file := range files {
		a := bp.New()
		err := g_util.LoadJSON(file, &a)
		if err != nil {
			return err
		}
		// Remove the pathToBP from the file path
		file := strings.TrimPrefix(file, pathToBP+string(filepath.Separator)+destDirectory.Animation+string(filepath.Separator))
		m.Add(file, a)
	}
	return nil
}
