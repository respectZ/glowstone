package bp

import (
	"strings"

	bp "github.com/respectZ/glowstone/bp/entity"
)

type EntityFile struct {
	Data *bp.Entity

	// Other stuff
	Subdir   string
	Filename string
}

// Create a new entity file
//
// Example:
//
//	entity := bp.NewEntity("glostone:diamond_sword")
func NewEntity(identifier string) *EntityFile {
	return &EntityFile{
		Data: bp.New(identifier),
	}
}

// Load an entity file from the given path
//
// Example:
//
//	entity, err := bp.LoadEntity(filepath.Join("packs", "BP", "entity", "zombie.json")
func LoadEntity(src string) (*bp.Entity, error) {
	entity, err := bp.Load(src)
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
		return e.GetIdentifier() + ".json"
	}
	return e.Filename
}
