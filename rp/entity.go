package rp

import (
	"os"
	"path/filepath"
	"strings"

	rp "github.com/respectZ/glowstone/rp/entity"
	g_util "github.com/respectZ/glowstone/util"
)

type Entities map[string]*EntityFile

type IEntities interface {
	Add(string, *rp.Entity)

	// Get returns the Entity of the given identifier.
	//
	// Example:
	//
	//  entity, ok := project.RP.Entity.Get("glowstone:chair")
	Get(string) (*rp.Entity, bool)

	// GetFile returns the EntityFile of the given identifier.
	//
	// Example:
	//
	//  entity, ok := project.RP.Entity.GetFile("glowstone:chair")
	GetFile(string) (*EntityFile, bool)

	// Remove removes the Entity of the given identifier.
	//
	// Example:
	//
	//  project.RP.Entity.Remove("glowstone:chair")
	Remove(string)
	Clear()

	// All returns all Entitys.
	//
	// Example:
	//
	//  entitys := project.RP.Entity.All()
	All() map[string]*EntityFile
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// Create a new entity.
	//
	// Example:
	//
	//	a := New("glowstone:diamond_sword")
	New(string) *rp.Entity

	// Saves all EntityFile to the given path.
	//
	// Example:
	//
	//	project := glowstone.NewProject()
	// 	entitys := project.RP.Entity
	// 	entities.New("glowstone:chair")
	// 	entities.New("glowstone:table")
	// 	entities.Save(path.Join("packs", "RP"))
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
	//	entity, err := project.RP.Entity.Load(filepath.Join(project.RP.Path, "entity", "player.enitty.json"))
	Load(string, ...bool) (*EntityFile, error)
}

func (e *Entities) Add(key string, value *rp.Entity) {
	p, ok := (*e)[key]
	if ok {
		p.Data = value
	} else {
		(*e)[key] = &EntityFile{
			Data: value,
		}
	}
}

func (e *Entities) Get(key string) (*rp.Entity, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value.Data, true
}

func (e *Entities) GetFile(key string) (*EntityFile, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func (e *Entities) Remove(key string) {
	delete(*e, key)
}

func (e *Entities) Clear() {
	*e = make(map[string]*EntityFile)
}

func (e *Entities) All() map[string]*EntityFile {
	return *e
}

func (e *Entities) IsEmpty() bool {
	return len(*e) == 0
}

func (e *Entities) Size() int {
	return len(*e)
}

func (e *Entities) UnmarshalJSON(b []byte) error {
	return g_util.UnmarshalJSON(b, e)
}

func (e *Entities) New(identifier string) *rp.Entity {
	entity := rp.New(identifier)
	e.Add(identifier, entity)
	return entity
}

func (e *Entities) Save(pathToRP string) error {
	for _, entity := range *e {
		b, err := entity.Encode()
		if err != nil {
			return err
		}
		err = g_util.WriteFile(filepath.Join(pathToRP, destDirectory.Entity, entity.Subdir, entity.GetFilename()), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Entities) LoadAll(pathToRP string) error {
	files, err := g_util.Walk(filepath.Join(pathToRP, destDirectory.Entity))
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
		file = strings.TrimPrefix(file, pathToRP+string(filepath.Separator)+destDirectory.Entity+string(filepath.Separator))
		subdir := filepath.Dir(file)
		// Filename
		filename := filepath.Base(file)

		(*e)[a.GetIdentifier()].Filename = filename
		(*e)[a.GetIdentifier()].Subdir = subdir
	}
	return nil
}

func (e *Entities) Load(src string, add ...bool) (*EntityFile, error) {
	a, err := rp.Load(src)
	if err != nil {
		return nil, err
	}

	filename := filepath.Base(src)

	// Get subdir
	subdirs := strings.Split(filepath.Dir(src), string(filepath.Separator))
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
	data.Filename = filename
	data.Subdir = subdir

	return data, nil
}
