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
type Animations map[string]*rp.AnimationDefinition

type IAnimations interface {
	UnmarshalJSON([]byte) error
	Add(string, *rp.AnimationDefinition)
	Get(string) *rp.AnimationDefinition
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]*rp.AnimationDefinition

	Save(string) error

	LoadAll(string) error
}

func (a *Animations) UnmarshalJSON(data []byte) error {
	var temp map[string]*rp.AnimationDefinition
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}

func (a *Animations) Add(s string, v *rp.AnimationDefinition) {
	(*a)[s] = v
}

func (a *Animations) Get(s string) *rp.AnimationDefinition {
	return (*a)[s]
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

func (a *Animations) All() map[string]*rp.AnimationDefinition {
	return *a
}

func (a *Animations) Save(pathToRP string) error {
	for k, animation := range *a {
		b, err := animation.Encode()
		if err != nil {
			return err
		}
		if err := g_util.Writefile(filepath.Join(pathToRP, destDirectory.Animation, k), b); err != nil {
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
		(*a)[file] = animation
	}
	return nil
}
