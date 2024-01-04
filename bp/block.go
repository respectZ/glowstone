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
	Get(string) (bp.Block, bool)
	Remove(string)
	Clear()
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

		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.Block, v.Subdir, v.GetFilename()), b)
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
