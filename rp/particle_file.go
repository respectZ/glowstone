package rp

import (
	"strings"

	rp "github.com/respectZ/glowstone/rp/particle"
)

type ParticleFile struct {
	Data *rp.Particle

	// Other stuff
	Subdir   string
	Filename string
}

// Create a new particle file
//
// Example:
//
//	particle := rp.NewParticle("glowstone:slash")
func NewParticle(identifier string) *ParticleFile {
	return &ParticleFile{
		Data: rp.New(identifier),
	}
}

// Load loads particle
//
// Example:
//
//	particle, err := rp.LoadParticle("./rp/particle/particle.json")
func LoadParticle(src string) (*rp.Particle, error) {
	particle, err := rp.Load(src)
	if err != nil {
		return nil, err
	}
	return particle, nil
}

// Encode returns []byte of the particle.
//
// Example:
//
//	particle, err := e.Encode()
func (e *ParticleFile) Encode() ([]byte, error) {
	particle, err := e.Data.Encode()
	if err != nil {
		return nil, err
	}
	return particle, nil
}

// GetIdentifier returns the identifier of the entity without namespace
//
// Example:
//
//	identifier := e.GetIdentifier()
func (e *ParticleFile) GetIdentifier() string {
	return strings.Split(e.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the entity with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (e *ParticleFile) GetNamespaceIdentifier() string {
	return e.Data.GetIdentifier()
}

func (e *ParticleFile) GetFilename() string {
	if e.Filename == "" {
		return e.GetIdentifier() + ".particle.json"
	}
	return e.Filename
}
