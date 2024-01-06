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
	Get(string) (*bp.Entity, bool)
	Remove(string)
	Clear()
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

		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.Entity, v.Subdir, v.GetFilename()), b)
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
