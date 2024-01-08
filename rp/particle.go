package rp

import (
	"os"
	"path/filepath"
	"strings"

	rp "github.com/respectZ/glowstone/rp/particle"
	g_util "github.com/respectZ/glowstone/util"
)

type Particles map[string]*ParticleFile

type IParticles interface {
	Add(string, *rp.Particle)

	// Get returns the Particle of the given identifier.
	//
	// Example:
	//
	//  particle, ok := project.RP.Particle.Get("glowstone:particle_1")
	Get(string) (*rp.Particle, bool)

	// GetFile returns the ParticleFile of the given identifier.
	//
	// Example:
	//
	//  particle, ok := project.RP.Particle.GetFile("glowstone:particle_1")
	GetFile(string) (*ParticleFile, bool)

	// Remove removes the Particle of the given identifier.
	//
	// Example:
	//
	//  project.RP.Particle.Remove("glowstone:particle_1")
	Remove(string)
	Clear()

	// All returns all Particles.
	//
	// Example:
	//
	//  particles := project.RP.Particle.All()
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

	// Load a single particle file.
	//
	// Last parameter is whether the file should be added to the project.
	//
	// Default is true.
	//
	// Example:
	//
	//	particle, err := project.RP.Particle.Load(filepath.Join(project.RP.Path, "particles", "glowstone.particle.json"))
	Load(string, ...bool) (*ParticleFile, error)
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

func (e *Particles) GetFile(key string) (*ParticleFile, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value, true
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

		err = g_util.WriteFile(filepath.Join(pathToRP, destDirectory.Particle, particle.Subdir, particle.GetFilename()), b)
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

func (e *Particles) Load(src string, add ...bool) (*ParticleFile, error) {
	a, err := rp.Load(src)
	if err != nil {
		return nil, err
	}

	filename := filepath.Base(src)

	// Get subdir
	subdirs := strings.Split(src, string(filepath.Separator))
	subdir := ""
	// Reverse loop
	for i := len(subdirs) - 1; i >= 0; i-- {
		if subdirs[i] == destDirectory.Particle {
			break
		}
		subdir = subdirs[i] + string(filepath.Separator) + subdir
	}

	if len(add) > 0 && add[0] || len(add) == 0 {
		e.Add(a.GetIdentifier(), a)
	}

	data, ok := e.GetFile(a.GetIdentifier())
	if !ok {
		data = &ParticleFile{
			Data: a,
		}
	}
	data.Filename = filename
	data.Subdir = subdir
	return data, nil
}
