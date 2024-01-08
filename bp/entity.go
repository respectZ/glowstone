package bp

import (
	"os"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/entity"
	g_util "github.com/respectZ/glowstone/util"
)

type EntityBP map[string]*EntityFile

type IEntityBP interface {
	Add(string, *bp.Entity)

	// Get returns the Entity of the given identifier.
	//
	// Example:
	//
	//  entity, ok := project.BP.Entity.Get("glowstone:chair")
	Get(string) (*bp.Entity, bool)

	// GetFile returns the EntityFile of the given identifier.
	//
	// Example:
	//
	//  entity, ok := project.BP.Entity.GetFile("glowstone:chair")
	GetFile(string) (*EntityFile, bool)

	// Remove removes the Entity of the given identifier.
	//
	// Example:
	//
	//  project.BP.Entity.Remove("glowstone:chair")
	Remove(string)
	Clear()

	// All returns all Entities.
	//
	// Example:
	//
	//  entitys := project.BP.Entity.All()
	All() map[string]*EntityFile
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// New creates a new BPEntity.
	//
	// Example:
	//
	//	a := New("glowstone:chair")
	//	b := New("glowstone:table")
	New(string) *bp.Entity

	// Save saves all BPEntity to the given path.
	//
	// Example:
	//
	//	project := glowstone.NewProject()
	// 	entities := project.BP.Entity
	// 	entities.New("glowstone:chair")
	// 	entities.New("glowstone:table")
	// 	entities.Save(path.Join("packs", "BP"))
	Save(string) error

	LoadAll(string) error

	// Load a single entity file.
	//
	// Last parameter is whether the file should be added to the project.
	//
	// Default is true.
	//
	// Example:
	//
	//	entity, err := project.BP.Entity.Load(filepath.Join(project.BP.Path, "entities", "player.json"))
	Load(string, ...bool) (*EntityFile, error)
}

func (e *EntityBP) Add(key string, value *bp.Entity) {
	p, ok := (*e)[key]
	if ok {
		p.Data = value
	} else {
		(*e)[key] = &EntityFile{
			Data: value,
		}
	}
}

func (e *EntityBP) Get(key string) (*bp.Entity, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value.Data, ok
}

func (e *EntityBP) GetFile(key string) (*EntityFile, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func (e *EntityBP) Remove(key string) {
	delete(*e, key)
}

func (e *EntityBP) Clear() {
	*e = make(map[string]*EntityFile)
}

func (e *EntityBP) All() map[string]*EntityFile {
	return *e
}

func (e *EntityBP) IsEmpty() bool {
	return len(*e) == 0
}

func (e *EntityBP) Size() int {
	return len(*e)
}

func (e *EntityBP) UnmarshalJSON(data []byte) error {
	return g_util.UnmarshalJSON(data, e)
}

func (e *EntityBP) New(identifier string) *bp.Entity {
	a := bp.New(identifier)
	e.Add(identifier, a)
	return a
}

func (e *EntityBP) Save(pathToBP string) error {
	for _, v := range *e {
		data := v.Data
		if data == nil {
			continue
		}
		b, err := data.Encode()
		if err != nil {
			return err
		}

		err = g_util.WriteFile(filepath.Join(pathToBP, destDirectory.Entity, v.Subdir, v.GetFilename()), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *EntityBP) LoadAll(pathToBP string) error {
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.Entity))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, file := range files {
		a, err := bp.Load(file)
		if err != nil {
			return err
		}
		e.Add(a.GetIdentifier(), a)
		// Strip the pathToBP from the file path
		file = strings.TrimPrefix(file, pathToBP+string(filepath.Separator)+destDirectory.Entity+string(filepath.Separator))
		// Get all the directories
		subdir := filepath.Dir(file)
		// Filename
		filename := filepath.Base(file)

		(*e)[a.GetIdentifier()].Filename = filename
		(*e)[a.GetIdentifier()].Subdir = subdir
	}
	return nil
}

func (e *EntityBP) Load(src string, add ...bool) (*EntityFile, error) {
	a, err := bp.Load(src)
	if err != nil {
		return nil, err
	}

	filename := filepath.Base(src)

	// Get subdir
	subdirs := strings.Split(src, string(filepath.Separator))
	subdir := ""
	// Reverse loop
	for i := len(subdirs) - 1; i >= 0; i-- {
		if subdirs[i] == destDirectory.Entity {
			break
		}
		subdir = subdirs[i] + string(filepath.Separator) + subdir
	}
	if len(add) > 0 && add[0] || len(add) == 0 {
		e.Add(a.GetIdentifier(), a)
	}
	data, ok := e.GetFile(a.GetIdentifier())
	if !ok {
		data = &EntityFile{
			Data: a,
		}
	}
	data.Subdir = subdir
	data.Filename = filename
	return data, nil
}
