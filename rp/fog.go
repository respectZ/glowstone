package rp

import (
	"os"
	"path/filepath"
	"strings"

	rp "github.com/respectZ/glowstone/rp/fog"
	g_util "github.com/respectZ/glowstone/util"
)

type Fogs map[string]*FogFile

type IFogs interface {
	Add(string, *rp.Fog)

	// Get returns the Fog of the given identifier.
	//
	// Example:
	//
	//  fog, ok := project.RP.Fog.Get("glowstone:black_sky")
	Get(string) (*rp.Fog, bool)

	// GetFile returns the FogFile of the given identifier.
	//
	// Example:
	//
	//  fog, ok := project.RP.Fog.GetFile("glowstone:black_sky")
	GetFile(string) (*FogFile, bool)
	Clear()

	// Remove removes the Fog of the given identifier.
	//
	// Example:
	//
	//  project.RP.Fog.Remove("glowstone:black_sky")
	Remove(string)
	IsEmpty() bool
	Size() int

	// All returns all Fogs.
	//
	// Example:
	//
	//  fogs := project.RP.Fog.All()
	All() map[string]*FogFile

	// Create a new fog.
	//
	// Example:
	//
	//	fog := New("glowstone:black_sky")
	New(string) *rp.Fog

	// Create a new simple air fog.
	//
	// Example:
	//
	//  // Create white air fog.
	//	fog := NewAir("glowstone:black_sky", 255, 255, 255, 19.0, 25.0)
	NewAir(string, uint8, uint8, uint8, float64, float64) *rp.Fog

	// Saves all FogFile to the given path.
	//
	// Example:
	//
	//	project := glowstone.NewProject()
	//  project.RP.Fog.New("glowstone:black_sky")
	//  project.RP.Fog.Save(path.Join("packs", "RP"))
	Save(string) error

	// LoadAll loads all FogFile from the given path.
	//
	// Example:
	//
	//	project := glowstone.NewProject()
	//  project.RP.Fog.LoadAll(path.Join("packs", "RP"))
	LoadAll(string) error

	// Load a single fog file.
	//
	// Last parameter is whether the file should be added to the project.
	//
	// Default is true.
	//
	// Example:
	//
	//	fog, err := project.RP.Fog.Load(filepath.Join(project.RP.Path, "fogs", "black_sky.json"))
	Load(string, ...bool) (*FogFile, error)
}

func (f *Fogs) Add(name string, fog *rp.Fog) {
	(*f)[name] = &FogFile{
		Data: fog,
	}
}

func (f *Fogs) Get(name string) (*rp.Fog, bool) {
	fog, ok := (*f)[name]
	if !ok {
		return nil, false
	}
	return fog.Data, ok
}

func (f *Fogs) GetFile(name string) (*FogFile, bool) {
	fog, ok := (*f)[name]
	if !ok {
		return nil, false
	}
	return fog, ok
}

func (f *Fogs) Clear() {
	*f = make(Fogs, 0)
}

func (f *Fogs) Remove(name string) {
	delete(*f, name)
}

func (f *Fogs) IsEmpty() bool {
	return len(*f) == 0
}

func (f *Fogs) Size() int {
	return len(*f)
}

func (f *Fogs) All() map[string]*FogFile {
	return *f
}

func (f *Fogs) New(identifier string) *rp.Fog {
	fog := rp.New(identifier)
	// Add fog to array
	f.Add(identifier, fog)
	return fog
}

func (f *Fogs) NewAir(identifier string, red, green, blue uint8, density, distance float64) *rp.Fog {
	fog := rp.NewSimpleAir(identifier, red, green, blue, density, distance)
	// Add fog to array
	f.Add(identifier, fog)
	return fog
}

func (f *Fogs) Save(pathToRP string) error {
	for _, fogFile := range *f {
		d, err := fogFile.Encode()
		if err != nil {
			return err
		}
		if err := g_util.WriteFile(filepath.Join(pathToRP, destDirectory.Fog, fogFile.Subdir, fogFile.GetFilename()), d); err != nil {
			return err
		}
	}
	return nil
}

func (f *Fogs) LoadAll(pathToRP string) error {
	files, err := g_util.Walk(filepath.Join(pathToRP, destDirectory.Fog))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, file := range files {
		fog, err := rp.Load(file)
		if err != nil {
			return err
		}
		f.Add(fog.GetIdentifier(), fog)
		// Strip to get subdir
		file = strings.TrimPrefix(file, pathToRP+string(filepath.Separator)+destDirectory.Fog+string(filepath.Separator))
		subdir := filepath.Dir(file)
		// Filename
		filename := filepath.Base(file)

		(*f)[fog.GetIdentifier()].Filename = filename
		(*f)[fog.GetIdentifier()].Subdir = subdir
	}
	return nil
}

func (f *Fogs) Load(src string, add ...bool) (*FogFile, error) {
	fog, err := rp.Load(src)
	if err != nil {
		return nil, err
	}

	filename := filepath.Base(src)

	// Get subdir
	subdirs := strings.Split(filepath.Dir(src), string(filepath.Separator))
	subdir := ""
	// Reverse loop
	for i := len(subdirs) - 1; i >= 0; i-- {
		if subdirs[i] == destDirectory.Fog {
			break
		}
		subdir = subdirs[i] + string(filepath.Separator) + subdir
	}

	if len(add) > 0 && add[0] || len(add) == 0 {
		f.Add(fog.GetIdentifier(), fog)
	}

	data, ok := f.GetFile(fog.GetIdentifier())
	if !ok {
		data = &FogFile{
			Data: fog,
		}
	}
	data.Filename = filename
	data.Subdir = subdir

	return data, nil
}
