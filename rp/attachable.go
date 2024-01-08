package rp

import (
	"os"
	"path/filepath"
	"strings"

	rp "github.com/respectZ/glowstone/rp/attachable"
	g_util "github.com/respectZ/glowstone/util"
)

type Attachables map[string]*AttachableFile

type IAttachables interface {
	Add(string, *rp.Attachable)

	// Get returns the Attachable of the given identifier.
	//
	// Example:
	//
	//  attachable, ok := project.RP.Attachable.Get("glowstone:chair")
	Get(string) (*rp.Attachable, bool)

	// GetFile returns the AttachableFile of the given identifier.
	//
	// Example:
	//
	//  attachable, ok := project.RP.Attachable.GetFile("glowstone:chair")
	GetFile(string) (*AttachableFile, bool)
	Remove(string)
	Clear()

	// All returns all Attachables.
	//
	// Example:
	//
	//  attachables := project.RP.Attachable.All()
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

	// Load a single attachable file.
	//
	// Last parameter is whether the file should be added to the project.
	//
	// Default is true.
	//
	// Example:
	//
	//	attachable, err := project.RP.Attachable.Load(filepath.Join(project.RP.Path, "attachables", "glowstone.json"))
	Load(string, ...bool) (*AttachableFile, error)
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

func (e *Attachables) GetFile(key string) (*AttachableFile, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value, true
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

		err = g_util.WriteFile(filepath.Join(pathToRP, destDirectory.Attachable, attachable.Subdir, attachable.GetFilename()), b)
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

func (e *Attachables) Load(src string, add ...bool) (*AttachableFile, error) {
	a, err := rp.Load(src)
	if err != nil {
		return nil, err
	}

	filename := filepath.Base(src)

	// Get subdir
	subdirs := strings.Split(src, string(filepath.Separator))
	subdir := ""
	// Reverse loop
	for i := len(subdirs) - 1; i >= 0; i-- {
		if subdirs[i] == destDirectory.Attachable {
			break
		}
		subdir = subdirs[i] + string(filepath.Separator) + subdir
	}
	if len(add) > 0 && add[0] || len(add) == 0 {
		e.Add(a.GetIdentifier(), a)
	}

	data, ok := e.GetFile(a.GetIdentifier())
	if !ok {
		data = &AttachableFile{
			Data: a,
		}
	}
	data.Subdir = subdir
	data.Filename = filename
	return data, nil
}
