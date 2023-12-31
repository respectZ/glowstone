package glowstone

import (
	"fmt"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/item"
	"github.com/respectZ/glowstone/item"
	g_util "github.com/respectZ/glowstone/util"
)

type ItemBP map[string]*item.Item

type IItemBP interface {
	Add(string, bp.Item)
	Get(string) (bp.Item, bool)
	Remove(string)
	Clear()
	All() map[string]*item.Item
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// New creates a new BPItem.
	//
	// Example:
	//
	//	a := New("glowstone:chair")
	//	b := New("glowstone:table")
	New(string) bp.Item

	Save(string) error

	LoadAll(string) error
}

func (e *ItemBP) Add(key string, value bp.Item) {
	if *e == nil {
		*e = make(map[string]*item.Item)
	}
	p, ok := (*e)[key]
	if ok {
		p.BP = value
	} else {
		(*e)[key] = &item.Item{
			BP: value,
		}
	}
}

func (e *ItemBP) Get(key string) (bp.Item, bool) {
	if *e == nil {
		*e = make(map[string]*item.Item)
	}
	value, ok := (*e)[key]
	return value.BP, ok
}

func (e *ItemBP) Remove(key string) {
	if *e == nil {
		*e = make(map[string]*item.Item)
	}
	delete(*e, key)
}

func (e *ItemBP) Clear() {
	if *e == nil {
		*e = make(map[string]*item.Item)
	}
	*e = make(map[string]*item.Item)
}

func (e *ItemBP) All() map[string]*item.Item {
	if *e == nil {
		*e = make(map[string]*item.Item)
	}
	return *e
}

func (e *ItemBP) IsEmpty() bool {
	if *e == nil {
		*e = make(map[string]*item.Item)
	}
	return len(*e) == 0
}

func (e *ItemBP) Size() int {
	if *e == nil {
		*e = make(map[string]*item.Item)
	}
	return len(*e)
}

func (e *ItemBP) UnmarshalJSON(data []byte) error {
	if *e == nil {
		*e = make(map[string]*item.Item)
	}
	return g_util.UnmarshalJSON(data, e)
}

func (e *ItemBP) New(identifier string) bp.Item {
	f := strings.Split(identifier, ":")
	a := bp.New(f[0], f[1])
	e.Add(identifier, a)
	return a
}

func (e *ItemBP) Save(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]*item.Item)
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
		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.Item, v.Subdir, fmt.Sprintf("%s.json", v.GetIdentifier())), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *ItemBP) LoadAll(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]*item.Item)
	}
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.Item))
	if err != nil {
		return err
	}
	for _, file := range files {
		a, err := item.Load(file)
		if err != nil {
			return err
		}
		// Strip the pathToBP from the file path
		file := strings.TrimPrefix(file, pathToBP+string(filepath.Separator)+destDirectory.Item+string(filepath.Separator))
		// Get all the directories
		subdir := filepath.Dir(file)
		// Set subdir
		a.Subdir = subdir
		e.Add(a.GetNamespaceIdentifier(), a.BP)
	}
	return nil
}
