package glowstone

import (
	"fmt"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/entity"
	"github.com/respectZ/glowstone/entity"
	g_util "github.com/respectZ/glowstone/util"
)

type EntityBP map[string]*entity.Entity

type IEntityBP interface {
	Add(string, *bp.Entity)
	Get(string) (*bp.Entity, bool)
	Remove(string)
	Clear()
	All() map[string]*bp.Entity
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
	if *e == nil {
		*e = make(map[string]*entity.Entity)
	}
	p, ok := (*e)[key]
	if ok {
		p.BP = value
	} else {
		(*e)[key] = &entity.Entity{
			BP: value,
		}
	}
}

func (e *EntityBP) Get(key string) (*bp.Entity, bool) {
	if *e == nil {
		*e = make(map[string]*entity.Entity)
	}
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value.BP, ok
}

func (e *EntityBP) Remove(key string) {
	if *e == nil {
		*e = make(map[string]*entity.Entity)
	}
	delete(*e, key)
}

func (e *EntityBP) Clear() {
	if *e == nil {
		*e = make(map[string]*entity.Entity)
	}
	*e = make(map[string]*entity.Entity)
}

func (e *EntityBP) All() map[string]*bp.Entity {
	if *e == nil {
		*e = make(map[string]*entity.Entity)
	}
	temp := make(map[string]*bp.Entity)
	for k, v := range *e {
		temp[k] = v.BP
	}
	return temp
}

func (e *EntityBP) IsEmpty() bool {
	if *e == nil {
		*e = make(map[string]*entity.Entity)
	}
	return len(*e) == 0
}

func (e *EntityBP) Size() int {
	if *e == nil {
		*e = make(map[string]*entity.Entity)
	}
	return len(*e)
}

func (e *EntityBP) UnmarshalJSON(data []byte) error {
	if *e == nil {
		*e = make(map[string]*entity.Entity)
	}
	return g_util.UnmarshalJSON(data, e)
}

func (e *EntityBP) New(dest string) *bp.Entity {
	f := strings.Split(dest, ":")
	a := bp.New(f[0], f[1])
	e.Add(dest, a)
	return a
}

func (e *EntityBP) Save(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]*entity.Entity)
	}
	for _, v := range *e {
		data := v.BP
		if data == nil {
			continue
		}
		b, err := data.Encode()
		if err != nil {
			return err
		}
		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.Entity, v.Subdir, fmt.Sprintf("%s.json", v.GetIdentifier())), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *EntityBP) LoadAll(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]*entity.Entity)
	}
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.Entity))
	if err != nil {
		return err
	}
	for _, file := range files {
		a, err := entity.LoadBP(file)
		if err != nil {
			return err
		}
		r := entity.Entity{}
		// Strip the pathToBP from the file path
		file = strings.TrimPrefix(file, pathToBP+string(filepath.Separator)+destDirectory.Entity+string(filepath.Separator))
		// Get all the directories
		subdir := filepath.Dir(file)
		// Set subdir
		r.BP = a
		r.Subdir = subdir
		e.Add(r.GetNamespaceIdentifier(), r.BP)
	}
	return nil
}
