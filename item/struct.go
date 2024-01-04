package item

import (
	"strings"

	bp "github.com/respectZ/glowstone/bp/item"
)

type Item struct {
	BP bp.Item

	// Other stuff
	Subdir string
	Lang   string
}

func New(identifier string, subdir ...string) *Item {
	item := &Item{
		BP: bp.New(identifier),
	}
	if len(subdir) > 0 {
		item.Subdir = subdir[0]
	}
	return item
}

func Load(src string) (*Item, error) {
	bp, err := bp.Load(src)
	if err != nil {
		return nil, err
	}
	return &Item{
		BP: bp,
	}, nil
}

// LoadBP loads the BP from the given path
//
// Example:
//
//	e, err := item.LoadBP("./bp/item/stick.json")
func LoadBP(src string) (bp.Item, error) {
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
func (e *Item) Encode() ([]byte, error) {
	bp, err := e.BP.Encode()
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
func (e *Item) GetIdentifier() string {
	return strings.Split(e.BP.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the entity with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (e *Item) GetNamespaceIdentifier() string {
	return e.BP.GetIdentifier()
}

// SetLang sets the lang of the item name.
//
// Example:
//
//	e.SetLang("Super Stick")
func (e *Item) SetLang(lang string) {
	e.Lang = lang
}
