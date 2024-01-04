package rp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rp "github.com/respectZ/glowstone/rp/attachable"
	g_util "github.com/respectZ/glowstone/util"
)

type Attachables map[string]*AttachableFile

type IAttachables interface {
	Add(string, *rp.Attachable)
	Get(string) (*rp.Attachable, bool)
	Remove(string)
	Clear()
	All() map[string]*AttachableFile
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// Create a new attachable.
	//
	// Example:
	//
	//	a := New("glowstone:diamond_sword")
	New(string) *rp.Attachable

	// Saves all AttachableFile to the given path.
	//
	// Example:
	//
	//	project := glowstone.NewProject()
	// 	attachables := project.RP.Attachable
	// 	attachables.New("glowstone:chair")
	// 	attachables.New("glowstone:table")
	// 	attachables.Save(path.Join("packs", "RP"))
	Save(string) error

	LoadAll(string) error
}

func (e *Attachables) Add(key string, value *rp.Attachable) {
	p, ok := (*e)[key]
	if ok {
		p.Data = value
	} else {
		(*e)[key] = &AttachableFile{
			Data: value,
		}
	}
}

func (e *Attachables) Get(key string) (*rp.Attachable, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value.Data, true
}

func (e *Attachables) Remove(key string) {
	delete(*e, key)
}

func (e *Attachables) Clear() {
	*e = make(map[string]*AttachableFile)
}

func (e *Attachables) All() map[string]*AttachableFile {
	return *e
}

func (e *Attachables) IsEmpty() bool {
	return len(*e) == 0
}

func (e *Attachables) Size() int {
	return len(*e)
}

func (e *Attachables) UnmarshalJSON(b []byte) error {
	return g_util.UnmarshalJSON(b, e)
}

func (e *Attachables) New(identifier string) *rp.Attachable {
	attachable := rp.New(identifier)
	e.Add(identifier, attachable)
	return attachable
}

func (e *Attachables) Save(pathToRP string) error {
	for _, attachable := range *e {
		b, err := attachable.Encode()
		if err != nil {
			return err
		}

		if attachable.Filename == "" {
			attachable.Filename = fmt.Sprintf("%s.attachable.json", attachable.GetIdentifier())
		}

		err = g_util.Writefile(filepath.Join(pathToRP, destDirectory.Attachable, attachable.Subdir, attachable.Filename), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Attachables) LoadAll(pathToRP string) error {
	files, err := g_util.Walk(filepath.Join(pathToRP, destDirectory.Attachable))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, file := range files {
		a, err := rp.Load(file)
		if err != nil {
			return err
		}
		e.Add(a.GetIdentifier(), a)
		// Strip to get subdir
		file = strings.TrimPrefix(file, pathToRP+string(filepath.Separator)+destDirectory.Attachable+string(filepath.Separator))
		subdir := filepath.Dir(file)
		filename := filepath.Base(file)

		(*e)[a.GetIdentifier()].Subdir = subdir
		(*e)[a.GetIdentifier()].Filename = filename
	}
	return nil
}
