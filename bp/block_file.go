package bp

import (
	"strings"

	bp "github.com/respectZ/glowstone/bp/block"
)

type BlockFile struct {
	Data bp.Block

	// Other stuff
	Subdir   string
	Lang     string
	Filename string
}

// Create a new block file
//
// Example:
//
//	block := bp.NewBlock("glostone:diamond_sword")
func NewBlock(identifier string) *BlockFile {
	return &BlockFile{
		Data: bp.New(identifier),
	}
}

// Load loads block
//
// Example:
//
//	block, err := bp.LoadBlock("./bp/block/block.json")
func LoadBlock(src string) (bp.Block, error) {
	block, err := bp.Load(src)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// Encode returns []byte of the block.
//
// Example:
//
//	block, err := e.Encode()
func (e *BlockFile) Encode() ([]byte, error) {
	block, err := e.Data.Encode()
	if err != nil {
		return nil, err
	}
	return block, nil
}

// GetIdentifier returns the identifier of the entity without namespace
//
// Example:
//
//	identifier := e.GetIdentifier()
func (e *BlockFile) GetIdentifier() string {
	return strings.Split(e.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the entity with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (e *BlockFile) GetNamespaceIdentifier() string {
	return e.Data.GetIdentifier()
}

func (e *BlockFile) GetFilename() string {
	if e.Filename == "" {
		return e.GetIdentifier() + ".json"
	}
	return e.Filename
}
