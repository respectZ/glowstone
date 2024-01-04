package rp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rp "github.com/respectZ/glowstone/rp/entity"
	g_util "github.com/respectZ/glowstone/util"
)

type Entities map[string]*EntityFile

type IEntities interface {
	Add(string, *rp.Entity)
	Get(string) (*rp.Entity, bool)
	Remove(string)
	Clear()
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
		err = g_util.Writefile(filepath.Join(pathToRP, destDirectory.Entity, entity.Subdir, fmt.Sprintf("%s.entity.json", entity.GetIdentifier())), b)
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
