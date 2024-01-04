package bp

import (
	"strings"

	bp "github.com/respectZ/glowstone/bp/item"
)

type ItemFile struct {
	Data bp.Item

	// Other stuff
	Subdir   string
	Filename string
}

// Create a new item file
//
// Example:
//
//	item := bp.NewItem("glostone:diamond_sword")
func NewItem(identifier string) *ItemFile {
	return &ItemFile{
		Data: bp.New(identifier),
	}
}

// Load an item file from the given path
//
// Example:
//
//	item, err := bp.LoadItem(filepath.Join("packs", "BP", "item", "sword.json")
func LoadItem(src string) (bp.Item, error) {
	item, err := bp.Load(src)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Encode returns []byte of the item.
//
// Example:
//
//	item, err := e.Encode()
func (e *ItemFile) Encode() ([]byte, error) {
	item, err := e.Data.Encode()
	if err != nil {
		return nil, err
	}
	return item, nil
}

// GetIdentifier returns the identifier of the item without namespace
//
// Example:
//
//	identifier := e.GetIdentifier()
func (e *ItemFile) GetIdentifier() string {
	return strings.Split(e.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the item with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (e *ItemFile) GetNamespaceIdentifier() string {
	return e.Data.GetIdentifier()
}

func (e *ItemFile) GetFilename() string {
	if e.Filename == "" {
		return e.GetIdentifier() + ".json"
	}
	return e.Filename
}
