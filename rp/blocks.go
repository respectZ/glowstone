package rp

import (
	"fmt"

	g_util "github.com/respectZ/glowstone/util"
)

type Blocks map[string]*BlockData

type BlockData struct {
	Sound string `json:"sound,omitempty"`

	// Can be string, or BlockDataTextures
	Textures interface{} `json:"textures,omitempty"`
}

type BlockDataTextures struct {
	// 6 Faces

	Down  string `json:"down,omitempty"`
	Up    string `json:"up,omitempty"`
	North string `json:"north,omitempty"`
	South string `json:"south,omitempty"`
	West  string `json:"west,omitempty"`
	East  string `json:"east,omitempty"`

	// 3 Faces, Down, Up, Side
	Side string `json:"side,omitempty"`
}

// Creates a new blocks.json file
func NewBlocks() *Blocks {
	return &Blocks{}
}

// Loads a blocks.json file
func LoadBlocks(src string) (*Blocks, error) {
	var block Blocks
	var temp map[string]interface{}
	err := g_util.LoadJSON(src, &temp)
	if err != nil {
		return nil, err
	}
	// We remove format_version if exists
	delete(temp, "format_version")
	// Cast to Blocks
	t, err := g_util.MarshalJSON(temp)
	if err != nil {
		return nil, err
	}
	err = g_util.UnmarshalJSON(t, &block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (i *Blocks) Encode() ([]byte, error) {
	return g_util.MarshalJSON(i)
}

// Add a block to the blocks.json file
//
// Example:
//
//	bdata := blocks.AddBlock("allow", "stone", "build_allow")
func (i *Blocks) AddBlock(name string, sound string, texture string) *BlockData {
	if *i == nil {
		*i = make(map[string]*BlockData)
	}
	(*i)[name] = &BlockData{
		Sound:    sound,
		Textures: texture,
	}

	return (*i)[name]
}

// Add a block to the blocks.json file with faces.
//
// Add the faces after calling this function.
//
// Example:
//
//		bdata, bdatatex := blocks.AddBlockFaces("grass", "grass")
//		bdatatex.Down = "grass"
//		bdatatex.Up = "grass"
//	 ...
func (i *Blocks) AddBlockFaces(name string, sound string) (*BlockData, *BlockDataTextures) {
	if *i == nil {
		*i = make(map[string]*BlockData)
	}
	tex := &BlockDataTextures{}
	(*i)[name] = &BlockData{
		Sound:    sound,
		Textures: tex,
	}

	return (*i)[name], tex
}

// Get a block from the blocks.json file
//
// Example:
//
//	block, err := blocks.GetBlock("grass")
func (i *Blocks) GetBlock(name string) (*BlockData, error) {
	if *i == nil {
		return nil, nil
	}
	block, ok := (*i)[name]
	if !ok {
		return nil, nil
	}
	return block, nil
}

// Get a block from the blocks.json file, with faces
//
// Example:
//
//	block, err := blocks.GetBlockFaces("grass")
func (i *Blocks) GetBlockFaces(name string) (*BlockDataTextures, error) {
	if *i == nil {
		return nil, fmt.Errorf("blocks.json is nil")
	}
	block, ok := (*i)[name]
	if !ok {
		return nil, fmt.Errorf("block %s not found", name)
	}
	if block.Textures == nil {
		return nil, fmt.Errorf("block %s has no textures", name)
	}
	textures, ok := block.Textures.(BlockDataTextures)
	if !ok {
		return nil, fmt.Errorf("block %s has no face textures", name)
	}
	return &textures, nil
}
