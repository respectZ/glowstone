package rp

import (
	"strings"

	rp "github.com/respectZ/glowstone/rp/model/entity"
)

type GeometryFile struct {
	Data *rp.Geometry

	// Other stuff
	Subdir   string
	Filename string
}

// Create a new geometry file
//
// Example:
//
//	geometry := rp.NewGeometry("glostone:diamond_sword")
func NewGeometry(identifier string) *GeometryFile {
	return &GeometryFile{
		Data: rp.New(identifier),
	}
}

// Load loads geometry
//
// Example:
//
//	geometry, err := rp.LoadGeometry("./rp/geometry/geometry.json")
func LoadGeometry(src string) (*rp.Geometry, error) {
	geometry, err := rp.Load(src)
	if err != nil {
		return nil, err
	}
	return geometry, nil
}

// Encode returns []byte of the geometry.
//
// Example:
//
//	geometry, err := e.Encode()
func (e *GeometryFile) Encode() ([]byte, error) {
	geometry, err := e.Data.Encode()
	if err != nil {
		return nil, err
	}
	return geometry, nil
}

// GetIdentifier returns the identifier of the entity without namespace
//
// Example:
//
//	identifier := e.GetIdentifier()
func (e *GeometryFile) GetIdentifier() string {
	return strings.Split(e.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the entity with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (e *GeometryFile) GetNamespaceIdentifier() string {
	return e.Data.GetIdentifier()
}

func (e *GeometryFile) GetFilename() string {
	if e.Filename == "" {
		return e.GetIdentifier() + ".geo.json"
	}
	return e.Filename
}
