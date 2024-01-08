package rp

import (
	"os"
	"path/filepath"
	"strings"

	rp "github.com/respectZ/glowstone/rp/animation_controller"
	g_util "github.com/respectZ/glowstone/util"
)

type AnimationControllers map[string]*AnimationControllerFile

type IAnimationControllers interface {
	Add(string, *rp.AnimationControllerFile)
	Get(string) (*rp.AnimationControllerFile, bool)
	Remove(string)
	Clear()
	All() map[string]*AnimationControllerFile
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// Create a new animation controller.
	//
	// Example:
	//
	//	a := New("vanilla/player.controller_animation.json")
	New(string) *rp.AnimationControllerFile

	// Saves all AnimationControllerFile to the given path.
	//
	// Example:
	//
	//	project := glowstone.NewProject()
	// 	ac := project.RP.AnimationController
	// 	ac.New("vanilla/player.controller_animation.json")
	// 	ac.New("furniture/table.controller_animation.json")
	// 	ac.Save(path.Join("packs", "RP"))
	Save(string) error

	LoadAll(string) error
}

func (e *AnimationControllers) Add(key string, value *rp.AnimationControllerFile) {
	p, ok := (*e)[key]
	if ok {
		p.Data = value
	} else {
		(*e)[key] = &AnimationControllerFile{
			Data: value,
		}
	}
}

func (e *AnimationControllers) Get(key string) (*rp.AnimationControllerFile, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value.Data, true
}

func (e *AnimationControllers) Remove(key string) {
	delete(*e, key)
}

func (e *AnimationControllers) Clear() {
	*e = make(map[string]*AnimationControllerFile)
}

func (e *AnimationControllers) All() map[string]*AnimationControllerFile {
	return *e
}

func (e *AnimationControllers) IsEmpty() bool {
	return len(*e) == 0
}

func (e *AnimationControllers) Size() int {
	return len(*e)
}

func (e *AnimationControllers) UnmarshalJSON(b []byte) error {
	return g_util.UnmarshalJSON(b, e)
}

func (e *AnimationControllers) New(dest string) *rp.AnimationControllerFile {
	ac := rp.New()
	e.Add(dest, ac)
	return ac
}

func (e *AnimationControllers) Save(pathToRP string) error {
	for dest, ac := range *e {
		b, err := ac.Encode()
		if err != nil {
			return err
		}

		err = g_util.WriteFile(filepath.Join(pathToRP, destDirectory.AnimationController, dest), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *AnimationControllers) LoadAll(pathToRP string) error {
	files, err := g_util.Walk(filepath.Join(pathToRP, destDirectory.AnimationController))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, file := range files {
		ac, err := rp.Load(file)
		if err != nil {
			return err
		}
		// Strip to get subdir
		file = strings.TrimPrefix(file, pathToRP+string(filepath.Separator)+destDirectory.AnimationController+string(filepath.Separator))
		e.Add(file, ac)
	}
	return nil
}
