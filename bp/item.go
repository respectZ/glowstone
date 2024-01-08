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

	// Get returns the Item of the given identifier.
	//
	// Example:
	//
	//  item, ok := project.BP.Item.Get("glowstone:chair")
	Get(string) (bp.Item, bool)

	// GetFile returns the ItemFile of the given identifier.
	//
	// Example:
	//
	//  item, ok := project.BP.Item.GetFile("glowstone:chair")
	GetFile(string) (*ItemFile, bool)

	// Remove removes the Item of the given identifier.
	//
	// Example:
	//
	//  project.BP.Item.Remove("glowstone:chair")
	Remove(string)
	Clear()

	// All returns all Items.
	//
	// Example:
	//
	//  items := project.BP.Item.All()
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

	// Load a single item file.
	//
	// Last parameter is whether the file should be added to the project.
	//
	// Default is true.
	//
	// Example:
	//
	//	item, err := project.BP.Item.Load(filepath.Join(project.BP.Path, "items", "apple.json"))
	Load(string, ...bool) (*ItemFile, error)
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
	if !ok {
		return nil, false
	}
	return value.Data, true
}

func (e *ItemBP) GetFile(key string) (*ItemFile, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value, true
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
		err = g_util.WriteFile(filepath.Join(pathToBP, destDirectory.Item, v.Subdir, v.GetFilename()), b)
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

func (e *ItemBP) Load(src string, add ...bool) (*ItemFile, error) {
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
		if subdirs[i] == destDirectory.Item {
			break
		}
		subdir = subdirs[i] + string(filepath.Separator) + subdir
	}
	if len(add) > 0 && add[0] || len(add) == 0 {
		e.Add(a.GetIdentifier(), a)
	}
	data, ok := e.GetFile(a.GetIdentifier())
	if !ok {
		data = &ItemFile{
			Data: a,
		}
	}
	data.Subdir = subdir
	data.Filename = filename
	return data, nil
}
