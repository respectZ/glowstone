package rp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rp "github.com/respectZ/glowstone/rp/particle"
	g_util "github.com/respectZ/glowstone/util"
)

type Particles map[string]*ParticleFile

type IParticles interface {
	Add(string, *rp.Particle)
	Get(string) (*rp.Particle, bool)
	Remove(string)
	Clear()
	All() map[string]*ParticleFile
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// Create a new particle.
	//
	// Parameters:
	// 	identifier: Identifier of the particle.
	// 	subdir: Subdirectory of the particle.
	//
	// Example:
	//
	//	a := New("glowstone:diamond_sword", "weapon")
	New(string, ...string) *rp.Particle

	// Saves all ParticleFile to the given path.
	//
	// Example:
	//
	//	project := glowstone.NewProject()
	// 	particles := project.RP.Particle
	// 	particles.New("glowstone:chair")
	// 	particles.New("glowstone:table")
	// 	particles.Save(path.Join("packs", "RP"))
	Save(string) error

	LoadAll(string) error
}

func (e *Particles) Add(key string, value *rp.Particle) {
	p, ok := (*e)[key]
	if ok {
		p.Data = value
	} else {
		(*e)[key] = &ParticleFile{
			Data: value,
		}
	}
}

func (e *Particles) Get(key string) (*rp.Particle, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value.Data, true
}

func (e *Particles) Remove(key string) {
	delete(*e, key)
}

func (e *Particles) Clear() {
	*e = make(map[string]*ParticleFile)
}

func (e *Particles) All() map[string]*ParticleFile {
	return *e
}

func (e *Particles) IsEmpty() bool {
	return len(*e) == 0
}

func (e *Particles) Size() int {
	return len(*e)
}

func (e *Particles) UnmarshalJSON(b []byte) error {
	return g_util.UnmarshalJSON(b, e)
}

func (e *Particles) New(identifier string, subdir ...string) *rp.Particle {
	particle := rp.New(identifier)
	e.Add(identifier, particle)
	if len(subdir) > 0 {
		(*e)[identifier].Subdir = subdir[0]
	}
	return particle
}

func (e *Particles) Save(pathToRP string) error {
	for _, particle := range *e {
		b, err := particle.Encode()
		if err != nil {
			return err
		}

		if particle.Filename == "" {
			particle.Filename = fmt.Sprintf("%s.particle.json", particle.GetIdentifier())
		}

		err = g_util.Writefile(filepath.Join(pathToRP, destDirectory.Particle, particle.Subdir, particle.Filename), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Particles) LoadAll(pathToRP string) error {
	files, err := g_util.Walk(filepath.Join(pathToRP, destDirectory.Particle))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, file := range files {
		a, err := rp.Load(file)
		if err != nil {
			return err
		}
		e.Add(a.GetIdentifier(), a)
		// Strip to get subdir
		file = strings.TrimPrefix(file, pathToRP+string(filepath.Separator)+destDirectory.Particle+string(filepath.Separator))
		subdir := filepath.Dir(file)
		filename := filepath.Base(file)

		(*e)[a.GetIdentifier()].Subdir = subdir
		(*e)[a.GetIdentifier()].Filename = filename
	}
	return nil
}
