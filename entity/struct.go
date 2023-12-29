package entity

import (
	"fmt"
	"strings"

	bp "github.com/respectZ/glowstone/bp/entity"
	rp "github.com/respectZ/glowstone/rp/entity"
)

type Entity struct {
	BP *bp.Entity
	RP *rp.Entity

	// Other stuff
	Subdir          string
	Lang            string
	SpawnLang       string
	RideHint        string
	RideHintConsole string

	// RPStuff

	AutoSpawnEggTexture bool
}

func New(namespace string, identifier string, subdir ...string) *Entity {
	s := ""
	if len(subdir) > 0 {
		s = subdir[0]
	}
	entity := &Entity{
		BP: bp.New(namespace, identifier),
		RP: rp.New(namespace, identifier),
	}
	entity.Subdir = s
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
func LoadBP(src string) (*bp.Entity, error) {
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
func LoadRP(src string) (*rp.Entity, error) {
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
	var bp []byte
	var rp []byte

	if e.BP != nil {
		bp, _ = e.BP.Encode()
	}
	if e.RP != nil {
		rp, _ = e.RP.Encode()
	}

	if bp == nil && rp == nil {
		return nil, nil, fmt.Errorf("no bp or rp")
	}
	return bp, rp, nil
}

// GetIdentifier returns the identifier of the entity without namespace
//
// Example:
//
//	identifier := e.GetIdentifier()
func (e *Entity) GetIdentifier() string {
	return strings.Split(e.GetNamespaceIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the entity with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (e *Entity) GetNamespaceIdentifier() string {
	if e.BP != nil {
		return e.BP.GetIdentifier()
	}
	return e.RP.GetIdentifier()
}

// SetLang sets the lang of the item name.
//
// Will also apply to the spawn egg.
//
// Example:
//
//	e.SetLang("Cold TNT")
func (e *Entity) SetLang(lang string) {
	e.Lang = lang
}

// SetSpawnLang sets the lang of the spawn egg.
//
// By default, it will be "Spawn <lang>"
//
// Example:
//
//	e.SetSpawnLang("Spawn Cold TNT")
func (e *Entity) SetSpawnLang(lang string) {
	e.SpawnLang = lang
}

// SetRideHint sets the ride hint of the entity.
//
// By default, it will be "Tap jump to exit the <entityName>" if there's minecart:rideable component.
//
// Example:
//
//	e.SetRideHint("Tap jump to exit the car")
func (e *Entity) SetRideHint(hint string) {
	e.RideHint = hint
	if e.RideHintConsole == "" {
		e.RideHintConsole = hint
	}
}
