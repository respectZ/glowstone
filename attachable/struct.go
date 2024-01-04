package attachable

import (
	"strings"

	attachable "github.com/respectZ/glowstone/rp/attachable"
)

type Attachable struct {
	Data *attachable.Attachable

	// Other stuff
	Subdir string
}

func New(identifier string, subdir ...string) *Attachable {
	attachable := &Attachable{
		Data: attachable.New(identifier),
	}
	if len(subdir) > 0 {
		attachable.Subdir = subdir[0]
	}
	return attachable
}

// Load loads attachable
//
// Example:
//
//	attachable, err := attachable.Load("./rp/attachable/attachable.json")
func Load(src string) (*attachable.Attachable, error) {
	attachable, err := attachable.Load(src)
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
func (e *Attachable) Encode() ([]byte, error) {
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
func (e *Attachable) GetIdentifier() string {
	return strings.Split(e.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the entity with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (e *Attachable) GetNamespaceIdentifier() string {
	return e.Data.GetIdentifier()
}
