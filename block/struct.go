package block

import (
	"strings"

	bp "github.com/respectZ/glowstone/bp/block"
)

type Block struct {
	BP bp.Block

	// Other stuff
	Subdir string
	Lang   string
}

func New(namespace string, identifier string, subdir ...string) *Block {
	block := &Block{
		BP: bp.New(namespace, identifier),
	}
	if len(subdir) > 0 {
		block.Subdir = subdir[0]
	}
	return block
}

// LoadBP loads the BP from the given path
//
// Example:
//
//	e, err := block.LoadBP("./bp/block/super_stone.json")
func LoadBP(src string) (bp.Block, error) {
	bp, err := bp.Load(src)
	if err != nil {
		return nil, err
	}
	return bp, nil
}

// Encode returns []byte of the BP.
//
// Example:
//
//	bp, err := e.Encode()
func (b *Block) Encode() ([]byte, error) {
	bp, err := b.BP.Encode()
	if err != nil {
		return nil, err
	}
	return bp, nil
}

// GetIdentifier returns the identifier of the entity without namespace
//
// Example:
//
//	identifier := e.GetIdentifier()
func (b *Block) GetIdentifier() string {
	return strings.Split(b.BP.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the entity with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (b *Block) GetNamespaceIdentifier() string {
	return b.BP.GetIdentifier()
}

// SetLang sets the lang of the item name.
//
// Example:
//
//	e.SetLang("Super Stone")
func (b *Block) SetLang(lang string) {
	b.Lang = lang
}
