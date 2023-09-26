package entity

import (
	"strings"

	bp "github.com/respectZ/glowstone/bp/entity"
	rp "github.com/respectZ/glowstone/rp/entity"
)

type Entity struct {
	BP bp.Entity
	RP rp.Entity

	// Other stuff
	Subdir string
}

func New(namespace string, identifier string, subdir ...string) *Entity {
	entity := &Entity{
		BP: bp.New(namespace, identifier),
		RP: rp.New(namespace, identifier),
	}
	if len(subdir) > 0 {
		entity.Subdir = subdir[0]
	}
	return entity
}

// Load loads the BP and RP from the given paths
//
// Example:
//
//	e, err := entity.Load("./bp/entities/humanoid.json", "./rp/entity/humanoid.json")
func Load(srcBP string, srcRP string) (*Entity, error) {
	bp, err := bp.Load(srcBP)
	if err != nil {
		return nil, err
	}
	rp, err := rp.Load(srcRP)
	if err != nil {
		return nil, err
	}
	return &Entity{
		BP: bp,
		RP: rp,
	}, nil
}

// LoadBP loads the BP from the given path
//
// Example:
//
//	e, err := entity.LoadBP("./bp/entities/humanoid.json")
func LoadBP(src string) (bp.Entity, error) {
	bp, err := bp.Load(src)
	if err != nil {
		return nil, err
	}
	return bp, nil
}

// LoadRP loads the RP from the given path
//
// Example:
//
//	e, err := entity.LoadRP("./rp/entity/humanoid.json")
func LoadRP(src string) (rp.Entity, error) {
	rp, err := rp.Load(src)
	if err != nil {
		return nil, err
	}
	return rp, nil
}

// Encode returns []byte of the BP and RP
//
// Example:
//
//	bp, rp, err := e.Encode()
func (e *Entity) Encode() ([]byte, []byte, error) {
	bp, err := e.BP.Encode()
	if err != nil {
		return nil, nil, err
	}
	rp, err := e.RP.Encode()
	if err != nil {
		return nil, nil, err
	}
	return bp, rp, nil
}

// GetIdentifier returns the identifier of the entity without namespace
//
// Example:
//
//	identifier := e.GetIdentifier()
func (e *Entity) GetIdentifier() string {
	return strings.Split(e.BP.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the entity with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (e *Entity) GetNamespaceIdentifier() string {
	return e.BP.GetIdentifier()
}
