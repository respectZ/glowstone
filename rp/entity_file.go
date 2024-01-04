package rp

import (
	"strings"

	rp "github.com/respectZ/glowstone/rp/entity"
)

type EntityFile struct {
	Data *rp.Entity

	// Other stuff
	Subdir   string
	Filename string
	Lang     string
}

// Create a new entity file
//
// Example:
//
//	entity := rp.NewEntity("glostone:diamond_sword")
func NewEntity(identifier string) *EntityFile {
	return &EntityFile{
		Data: rp.New(identifier),
	}
}

// Load an entity file from the given path
//
// Example:
//
//	entity, err := rp.LoadEntity(filepath.Join("packs", "RP", "entity", "player.entity.json")
func LoadEntity(src string) (*rp.Entity, error) {
	entity, err := rp.Load(src)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Encode returns []byte of the entity.
//
// Example:
//
//	entity, err := e.Encode()
func (e *EntityFile) Encode() ([]byte, error) {
	entity, err := e.Data.Encode()
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// GetIdentifier returns the identifier of the entity without namespace
//
// Example:
//
//	identifier := e.GetIdentifier()
func (e *EntityFile) GetIdentifier() string {
	return strings.Split(e.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the entity with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (e *EntityFile) GetNamespaceIdentifier() string {
	return e.Data.GetIdentifier()
}

func (e *EntityFile) GetFilename() string {
	if e.Filename == "" {
		return e.GetIdentifier() + ".entity.json"
	}
	return e.Filename
}
