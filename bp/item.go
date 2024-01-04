package bp

import (
	"os"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/item"
	g_util "github.com/respectZ/glowstone/util"
)

type ItemBP map[string]*ItemFile

type IItemBP interface {
	Add(string, bp.Item)
	Get(string) (bp.Item, bool)
	Remove(string)
	Clear()
	All() map[string]*ItemFile
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
		*e = make(map[string]*ItemFile)
	}
	p, ok := (*e)[key]
	if ok {
		p.Data = value
	} else {
		(*e)[key] = &ItemFile{
			Data: value,
		}
	}
}

func (e *ItemBP) Get(key string) (bp.Item, bool) {
	if *e == nil {
		*e = make(map[string]*ItemFile)
	}
	value, ok := (*e)[key]
	return value.Data, ok
}

func (e *ItemBP) Remove(key string) {
	if *e == nil {
		*e = make(map[string]*ItemFile)
	}
	delete(*e, key)
}

func (e *ItemBP) Clear() {
	if *e == nil {
		*e = make(map[string]*ItemFile)
	}
	*e = make(map[string]*ItemFile)
}

func (e *ItemBP) All() map[string]*ItemFile {
	if *e == nil {
		*e = make(map[string]*ItemFile)
	}
	return *e
}

func (e *ItemBP) IsEmpty() bool {
	if *e == nil {
		*e = make(map[string]*ItemFile)
	}
	return len(*e) == 0
}

func (e *ItemBP) Size() int {
	if *e == nil {
		*e = make(map[string]*ItemFile)
	}
	return len(*e)
}

func (e *ItemBP) UnmarshalJSON(data []byte) error {
	if *e == nil {
		*e = make(map[string]*ItemFile)
	}
	return g_util.UnmarshalJSON(data, e)
}

func (e *ItemBP) New(identifier string) bp.Item {
	a := bp.New(identifier)
	e.Add(identifier, a)
	return a
}

func (e *ItemBP) Save(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]*ItemFile)
	}
	for _, v := range *e {
		data := v.Data
		if data == nil {
			continue
		}
		b, err := data.Encode()
		if err != nil {
			return err
		}
		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.Item, v.Subdir, v.GetFilename()), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *ItemBP) LoadAll(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]*ItemFile)
	}
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.Item))
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
		file := strings.TrimPrefix(file, pathToBP+string(filepath.Separator)+destDirectory.Item+string(filepath.Separator))
		// Get all the directories
		subdir := filepath.Dir(file)
		// Filename
		filename := filepath.Base(file)

		(*e)[a.GetIdentifier()].Filename = filename
		(*e)[a.GetIdentifier()].Subdir = subdir
	}
	return nil
}
