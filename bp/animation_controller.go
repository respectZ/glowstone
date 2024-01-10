package bp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/animation_controller"
	g_util "github.com/respectZ/glowstone/util"
)

// We don't need options atm, so we use bp/animation_controller.go

type AnimationControllerBP map[string]*AnimationControllerFile

type IAnimationControllerBP interface {
	Add(string, *bp.AnimationControllerFile)

	// Get returns the AnimationController of the given path.
	//
	// Example:
	//
	//  animation_controller, ok := project.BP.AnimationController.Get("player.animation_controller.json")
	Get(string) (*bp.AnimationControllerFile, bool)

	// GetFile returns the AnimationControllerFile of the given path.
	//
	// Example:
	//
	//  animation_controller, ok := project.BP.AnimationController.GetFile("player.animation_controller.json")
	GetFile(string) (*AnimationControllerFile, bool)

	// Remove removes the AnimationController of the given path.
	//
	// Example:
	//
	//  project.BP.AnimationController.Remove("player.animation_controller.json")
	Remove(string)
	Clear()

	// All returns all AnimationControllers.
	//
	// Example:
	//
	//  animation_controllers := project.BP.AnimationController.All()
	All() map[string]*AnimationControllerFile
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

	// Load a single animation_controller file, without subdir and filename.
	//
	// Last parameter is whether the file should be added to the project.
	//
	// Default is true.
	//
	// Example:
	//
	//	animation_controller, err := project.BP.AnimationController.Load(filepath.Join(project.BP.Path, "animation_controllers", "player.animation_controller.json"))
	Load(string, ...bool) (*AnimationControllerFile, error)
}

func (m *AnimationControllerBP) Add(key string, value *bp.AnimationControllerFile) {
	p, ok := (*m)[key]
	if ok {
		p.Data = value
	} else {
		(*m)[key] = &AnimationControllerFile{
			Data: value,
		}
	}
}

func (m *AnimationControllerBP) Get(key string) (*bp.AnimationControllerFile, bool) {
	value, ok := (*m)[key]
	if !ok {
		return nil, false
	}
	return value.Data, true
}

func (m *AnimationControllerBP) GetFile(key string) (*AnimationControllerFile, bool) {
	value, ok := (*m)[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func (m *AnimationControllerBP) Remove(key string) {
	delete(*m, key)
}

func (m *AnimationControllerBP) Clear() {
	*m = make(map[string]*AnimationControllerFile)
}

func (m *AnimationControllerBP) All() map[string]*AnimationControllerFile {
	return *m
}

func (m *AnimationControllerBP) IsEmpty() bool {
	return len(*m) == 0
}

func (m *AnimationControllerBP) Size() int {
	return len(*m)
}

func (m *AnimationControllerBP) UnmarshalJSON(data []byte) error {
	return g_util.UnmarshalJSON(data, m)
}

func (m *AnimationControllerBP) New(dest string) *bp.AnimationControllerFile {
	a := bp.New()
	m.Add(dest, a)
	return a
}

func (m *AnimationControllerBP) Save(pathToBP string) error {
	for dest, v := range m.All() {
		data, err := v.Encode()
		if err != nil {
			return err
		}
		err = g_util.WriteFile(filepath.Join(pathToBP, destDirectory.AnimationController, dest), data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *AnimationControllerBP) LoadAll(pathToBP string) error {
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.AnimationController))
	if err != nil && !os.IsNotExist(err) {
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

func (m *AnimationControllerBP) Load(path string, add ...bool) (*AnimationControllerFile, error) {
	a := bp.New()
	err := g_util.LoadJSON(path, &a)
	if err != nil {
		return nil, err
	}
	if len(add) > 0 && add[0] || len(add) == 0 {
		file := filepath.Base(path)
		subdirs := strings.Split(filepath.Dir(path), string(filepath.Separator))

		// Add subdir until "animation_controllers" is reached
		// Reverse loop
		for i := len(subdirs) - 1; i >= 0; i-- {
			if subdirs[i] == destDirectory.AnimationController {
				break
			}
			file = subdirs[i] + string(filepath.Separator) + file
		}
		m.Add(file, a)
	}
	return &AnimationControllerFile{
		Data: a,
	}, nil
}
