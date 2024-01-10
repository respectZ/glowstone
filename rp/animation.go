package rp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rp "github.com/respectZ/glowstone/rp/animation"
	g_util "github.com/respectZ/glowstone/util"
)

// Contains all animations that exist in the animation file.
//
// Key is the destination path of the animation file.
//
// Value is the animation definition.
//
// k -> v = "player.animation.json" -> ...
type Animations map[string]*AnimationFile

type IAnimations interface {
	UnmarshalJSON([]byte) error
	Add(string, *rp.AnimationDefinition)

	// Get returns the Animation of the given path.
	//
	// Example:
	//
	//  animation, ok := project.RP.Animation.Get("player.animation.json")
	Get(string) (*rp.AnimationDefinition, bool)

	// GetFile returns the AnimationFile of the given path.
	//
	// Example:
	//
	//  animation, ok := project.RP.Animation.GetFile("player.animation.json")
	GetFile(string) (*AnimationFile, bool)
	Clear()

	// Remove removes the Animation of the given path.
	//
	// Example:
	//
	//  project.RP.Animation.Remove("player.animation.json")
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]*AnimationFile

	Save(string) error

	LoadAll(string) error

	// Load a single animation file.
	//
	// Last parameter is whether the file should be added to the project.
	//
	// Default is true.
	//
	// Example:
	//
	//	animation, err := project.RP.Animation.Load(filepath.Join(project.RP.Path, "animations", "player.animation.json"))
	Load(string, ...bool) (*AnimationFile, error)
}

func (a *Animations) UnmarshalJSON(data []byte) error {
	var temp map[string]*AnimationFile
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}

func (a *Animations) Add(s string, v *rp.AnimationDefinition) {
	p, ok := (*a)[s]
	if ok {
		p.Data = v
	} else {
		(*a)[s] = &AnimationFile{
			Data: v,
		}
	}
}

func (a *Animations) Get(s string) (*rp.AnimationDefinition, bool) {
	animation, ok := (*a)[s]
	if !ok {
		return nil, false
	}
	return animation.Data, ok
}

func (a *Animations) GetFile(s string) (*AnimationFile, bool) {
	animation, ok := (*a)[s]
	if !ok {
		return nil, false
	}
	return animation, ok
}

func (a *Animations) Clear() {
	*a = make(Animations, 0)
}

func (a *Animations) Remove(s string) {
	delete(*a, s)
}

func (a *Animations) IsEmpty() bool {
	return len(*a) == 0
}

func (a *Animations) Size() int {
	return len(*a)
}

func (a *Animations) All() map[string]*AnimationFile {
	return *a
}

func (a *Animations) Save(pathToRP string) error {
	for k, animation := range *a {
		b, err := animation.Encode()
		if err != nil {
			return err
		}
		if err := g_util.WriteFile(filepath.Join(pathToRP, destDirectory.Animation, k), b); err != nil {
			return err
		}
	}
	return nil
}

func (a *Animations) LoadAll(pathToRP string) error {
	files, err := g_util.Walk(filepath.Join(pathToRP, destDirectory.Animation))
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	for _, file := range files {
		animation, err := rp.Load(file)
		if err != nil {
			return err
		}
		// Remove the path to the resource pack.
		file := strings.TrimPrefix(file, fmt.Sprintf("%s%s%s%s", pathToRP, string(filepath.Separator), destDirectory.Animation, string(filepath.Separator)))
		(*a)[file] = &AnimationFile{
			Data: animation,
		}
	}
	return nil
}

func (a *Animations) Load(path string, add ...bool) (*AnimationFile, error) {
	r, err := rp.Load(path)
	if err != nil {
		return nil, err
	}

	if len(add) > 0 && add[0] || len(add) == 0 {
		file := filepath.Base(path)
		subdirs := strings.Split(filepath.Dir(path), string(filepath.Separator))

		// Add subdir until "animations" is reached
		// Reverse loop
		for i := len(subdirs) - 1; i >= 0; i-- {
			if subdirs[i] == destDirectory.Animation {
				break
			}
			file = subdirs[i] + string(filepath.Separator) + file
		}
		a.Add(file, r)
	}

	data, ok := a.GetFile(path)
	if !ok {
		data = &AnimationFile{
			Data: r,
		}
	}
	return data, nil
}
