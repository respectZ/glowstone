package glowstone

// TODO: Rework this since *block.Block is a pointer to interface.

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/respectZ/glowstone/block"
	bp "github.com/respectZ/glowstone/bp/block"
	g_util "github.com/respectZ/glowstone/util"
)

type BlockBP map[string]*block.Block

type IBlockBP interface {
	Add(string, bp.Block)
	Get(string) (bp.Block, bool)
	Remove(string)
	Clear()
	All() map[string]*block.Block
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
	if *e == nil {
		*e = make(map[string]*block.Block)
	}
	p, ok := (*e)[key]
	if ok {
		p.BP = value
	} else {
		(*e)[key] = &block.Block{
			BP: value,
		}
	}
}

func (e *BlockBP) Get(key string) (bp.Block, bool) {
	if *e == nil {
		*e = make(map[string]*block.Block)
	}
	value, ok := (*e)[key]
	return value.BP, ok
}

func (e *BlockBP) Remove(key string) {
	if *e == nil {
		*e = make(map[string]*block.Block)
	}
	delete(*e, key)
}

func (e *BlockBP) Clear() {
	if *e == nil {
		*e = make(map[string]*block.Block)
	}
	*e = make(map[string]*block.Block)
}

func (e *BlockBP) All() map[string]*block.Block {
	if *e == nil {
		*e = make(map[string]*block.Block)
	}
	return *e
}

func (e *BlockBP) IsEmpty() bool {
	if *e == nil {
		*e = make(map[string]*block.Block)
	}
	return len(*e) == 0
}

func (e *BlockBP) Size() int {
	if *e == nil {
		*e = make(map[string]*block.Block)
	}
	return len(*e)
}

func (e *BlockBP) UnmarshalJSON(data []byte) error {
	if *e == nil {
		*e = make(map[string]*block.Block)
	}
	return g_util.UnmarshalJSON(data, e)
}

func (e *BlockBP) New(identifier string, option ...struct {
	Subdir string
	Lang   string
},
) bp.Block {
	f := strings.Split(identifier, ":")
	var subdir string
	var lang string
	if len(option) > 0 {
		subdir = option[0].Subdir
		lang = option[0].Lang
	}

	bl := block.New(f[0], f[1])
	bl.Subdir = subdir
	bl.Lang = lang

	(*e)[identifier] = bl

	return bl.BP
}

func (e *BlockBP) Save(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]*block.Block)
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
		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.Block, v.Subdir, fmt.Sprintf("%s.json", v.GetIdentifier())), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *BlockBP) LoadAll(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]*block.Block)
	}
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.Block))
	if err != nil {
		return err
	}
	for _, file := range files {
		a, err := block.LoadBP(file)
		if err != nil {
			return err
		}
		// Strip the pathToBP from the file path
		file = strings.TrimPrefix(file, pathToBP+string(filepath.Separator)+destDirectory.Block+string(filepath.Separator))
		// Get all the directories
		subdir := path.Dir(file)
		// Set subdir
		a.Subdir = subdir
		e.Add(a.GetNamespaceIdentifier(), a.BP)
	}
	return nil
}
