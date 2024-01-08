package bp

// TODO: Rework this since *block.Block is a pointer to interface.

import (
	"os"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/block"
	g_util "github.com/respectZ/glowstone/util"
)

type BlockBP map[string]*BlockFile

type IBlockBP interface {
	Add(string, bp.Block)

	// Get returns the Block of the given identifier.
	//
	// Example:
	//
	//  block, ok := project.BP.Block.Get("glowstone:chair")
	Get(string) (bp.Block, bool)

	// GetFile returns the BlockFile of the given identifier.
	//
	// Example:
	//
	//  block, ok := project.BP.Block.GetFile("glowstone:chair")
	GetFile(string) (*BlockFile, bool)

	// Remove removes the Block of the given identifier.
	//
	// Example:
	//
	//  project.BP.Block.Remove("glowstone:chair")
	Remove(string)
	Clear()

	// All returns all Blocks.
	//
	// Example:
	//
	//  blocks := project.BP.Block.All()
	All() map[string]*BlockFile
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// New creates a new BPBlock.
	//
	// Example:
	//
	//	a := New("glowstone:chair")
	//	b := New("glowstone:table")
	New(string, ...struct {
		Subdir string
		Lang   string
	}) bp.Block

	Save(string) error

	LoadAll(string) error

	// Load a single block file.
	//
	// Last parameter is whether the file should be added to the project.
	//
	// Default is true.
	//
	// Example:
	//
	//	block, err := project.BP.block.Load(filepath.Join(project.BP.Path, "blocks", "glowstone.json"))
	Load(string, ...bool) (*BlockFile, error)
}

func (e *BlockBP) Add(key string, value bp.Block) {
	p, ok := (*e)[key]
	if ok {
		p.Data = value
	} else {
		(*e)[key] = &BlockFile{
			Data: value,
		}
	}
}

func (e *BlockBP) Get(key string) (bp.Block, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value.Data, true
}

func (e *BlockBP) GetFile(key string) (*BlockFile, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func (e *BlockBP) Remove(key string) {
	delete(*e, key)
}

func (e *BlockBP) Clear() {
	*e = make(map[string]*BlockFile)
}

func (e *BlockBP) All() map[string]*BlockFile {
	return *e
}

func (e *BlockBP) IsEmpty() bool {
	return len(*e) == 0
}

func (e *BlockBP) Size() int {
	return len(*e)
}

func (e *BlockBP) UnmarshalJSON(data []byte) error {
	return g_util.UnmarshalJSON(data, e)
}

func (e *BlockBP) New(identifier string, option ...struct {
	Subdir string
	Lang   string
},
) bp.Block {
	var subdir string
	var lang string
	if len(option) > 0 {
		subdir = option[0].Subdir
		lang = option[0].Lang
	}

	bl := bp.New(identifier)
	e.Add(identifier, bl)

	if subdir != "" {
		(*e)[identifier].Subdir = subdir
	}
	if lang != "" {
		(*e)[identifier].Lang = lang
	}

	return bl
}

func (e *BlockBP) Save(pathToBP string) error {
	for _, v := range *e {
		data := v.Data
		if data == nil {
			continue
		}
		b, err := data.Encode()
		if err != nil {
			return err
		}

		err = g_util.WriteFile(filepath.Join(pathToBP, destDirectory.Block, v.Subdir, v.GetFilename()), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *BlockBP) LoadAll(pathToBP string) error {
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.Block))
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
		file = strings.TrimPrefix(file, pathToBP+string(filepath.Separator)+destDirectory.Block+string(filepath.Separator))
		// Get all the directories
		subdir := filepath.Dir(file)
		// Filename
		filename := filepath.Base(file)

		(*e)[a.GetIdentifier()].Subdir = subdir
		(*e)[a.GetIdentifier()].Filename = filename
	}
	return nil
}

func (e *BlockBP) Load(src string, add ...bool) (*BlockFile, error) {
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
		if subdirs[i] == destDirectory.Block {
			break
		}
		subdir = subdirs[i] + string(filepath.Separator) + subdir
	}
	if len(add) > 0 && add[0] || len(add) == 0 {
		e.Add(a.GetIdentifier(), a)
	}
	data, ok := e.GetFile(a.GetIdentifier())
	if !ok {
		data = &BlockFile{
			Data: a,
		}
	}
	data.Subdir = subdir
	data.Filename = filename
	return data, nil
}
