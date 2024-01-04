package rp

import (
	"strings"

	rp "github.com/respectZ/glowstone/rp/attachable"
)

type AttachableFile struct {
	Data *rp.Attachable

	// Other stuff
	Subdir   string
	Filename string
}

// Create a new attachable file
//
// Example:
//
//	attachable := rp.NewAttachable("glostone:diamond_sword")
func NewAttachable(identifier string) *AttachableFile {
	return &AttachableFile{
		Data: rp.New(identifier),
	}
}

// Load loads attachable
//
// Example:
//
//	attachable, err := rp.LoadAttachable("./rp/attachable/attachable.json")
func LoadAttachable(src string) (*rp.Attachable, error) {
	attachable, err := rp.Load(src)
	if err != nil {
		return nil, err
	}
	return attachable, nil
}

// Encode returns []byte of the attachable.
//
// Example:
//
//	attachable, err := e.Encode()
func (e *AttachableFile) Encode() ([]byte, error) {
	attachable, err := e.Data.Encode()
	if err != nil {
		return nil, err
	}
	return attachable, nil
}

// GetIdentifier returns the identifier of the entity without namespace
//
// Example:
//
//	identifier := e.GetIdentifier()
func (e *AttachableFile) GetIdentifier() string {
	return strings.Split(e.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the entity with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (e *AttachableFile) GetNamespaceIdentifier() string {
	return e.Data.GetIdentifier()
}

func (e *AttachableFile) GetFilename() string {
	if e.Filename == "" {
		return e.GetIdentifier() + ".attachable.json"
	}
	return e.Filename
}
