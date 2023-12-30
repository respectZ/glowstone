package glowstone

import (
	"fmt"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/animation_controller"
	g_util "github.com/respectZ/glowstone/util"
)

// We don't need options atm, so we use bp/animation_controller.go

type AnimationControllerBP map[string]*bp.AnimationControllerFile

type IAnimationControllerBP interface {
	Add(string, *bp.AnimationControllerFile)
	Get(string) (*bp.AnimationControllerFile, bool)
	Remove(string)
	Clear()
	All() map[string]*bp.AnimationControllerFile
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// New creates a new BPAnimationController. Dest doesn't need to start with "animation_controllers/"
	//
	// Example:
	//
	//	a := New("player.animation_controller.json")
	//	b := New("player/effect.animation_controller.json")
	New(string) *bp.AnimationControllerFile

	Save(string) error

	LoadAll(string) error
}

func (m *AnimationControllerBP) Add(key string, value *bp.AnimationControllerFile) {
	if *m == nil {
		*m = make(map[string]*bp.AnimationControllerFile)
	}
	(*m)[key] = value
}

func (m *AnimationControllerBP) Get(key string) (*bp.AnimationControllerFile, bool) {
	if *m == nil {
		*m = make(map[string]*bp.AnimationControllerFile)
	}
	value, ok := (*m)[key]
	return value, ok
}

func (m *AnimationControllerBP) Remove(key string) {
	if *m == nil {
		*m = make(map[string]*bp.AnimationControllerFile)
	}
	delete(*m, key)
}

func (m *AnimationControllerBP) Clear() {
	if *m == nil {
		*m = make(map[string]*bp.AnimationControllerFile)
	}
	*m = make(map[string]*bp.AnimationControllerFile)
}

func (m *AnimationControllerBP) All() map[string]*bp.AnimationControllerFile {
	if *m == nil {
		*m = make(map[string]*bp.AnimationControllerFile)
	}
	return *m
}

func (m *AnimationControllerBP) IsEmpty() bool {
	if *m == nil {
		*m = make(map[string]*bp.AnimationControllerFile)
	}
	return len(*m) == 0
}

func (m *AnimationControllerBP) Size() int {
	if *m == nil {
		*m = make(map[string]*bp.AnimationControllerFile)
	}
	return len(*m)
}

func (m *AnimationControllerBP) UnmarshalJSON(data []byte) error {
	if *m == nil {
		*m = make(map[string]*bp.AnimationControllerFile)
	}
	return g_util.UnmarshalJSON(data, m)
}

func (m *AnimationControllerBP) New(dest string) *bp.AnimationControllerFile {
	a := bp.New()
	m.Add(dest, a)
	return a
}

func (m *AnimationControllerBP) Save(pathToBP string) error {
	if *m == nil {
		*m = make(map[string]*bp.AnimationControllerFile)
	}
	for dest, v := range m.All() {
		data, err := v.Encode()
		if err != nil {
			return err
		}
		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.AnimationController, dest), data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *AnimationControllerBP) LoadAll(pathToBP string) error {
	if *m == nil {
		*m = make(map[string]*bp.AnimationControllerFile)
	}
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.AnimationController))
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
		file := strings.TrimPrefix(file, fmt.Sprintf("%s%s%s%s", pathToBP, string(filepath.Separator), destDirectory.AnimationController, string(filepath.Separator)))
		m.Add(file, a)
	}
	return nil
}
