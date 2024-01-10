package rp

import (
	"os"
	"path/filepath"
	"strings"

	rp "github.com/respectZ/glowstone/rp/model/entity"
	g_util "github.com/respectZ/glowstone/util"
)

type Geometries map[string]*GeometryFile

type IGeometries interface {
	Add(string, *rp.Geometry)

	// Get returns the Geometry of the given identifier.
	//
	// Example:
	//
	//  geometry, ok := project.RP.Geometry.Get("geometry.chair")
	Get(string) (*rp.Geometry, bool)

	// GetFile returns the GeometryFile of the given identifier.
	//
	// Example:
	//
	//  geometry, ok := project.RP.Geometry.GetFile("geometry.chair")
	GetFile(string) (*GeometryFile, bool)

	// Remove removes the Geometry of the given identifier.
	//
	// Example:
	//
	//  project.RP.Geometry.Remove("geometry.chair")
	Remove(string)
	Clear()

	// All returns all Geometries.
	//
	// Example:
	//
	//  geometries := project.RP.Geometry.All()
	All() map[string]*GeometryFile
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// Create a new model/entity.
	//
	// Example:
	//
	//	a := New("glowstone:diamond_sword")
	New(string) *rp.Geometry

	// Saves all GeometryFile to the given path.
	//
	// Example:
	//
	//	project := glowstone.NewProject()
	// 	project.RP.Geometry.New("glowstone:chair")
	// 	project.RP.Geometry.New("glowstone:table")
	// 	project.RP.Geometry.Save(path.Join("packs", "RP"))
	Save(string) error

	LoadAll(string) error

	// Load a single geometry file, without subdir and filename.
	//
	// Last parameter is whether the file should be added to the project.
	//
	// Default is true.
	//
	// Example:
	//
	//	geometry, err := project.RP.Geometry.Load(filepath.Join(project.RP.Path, "models", "entity", "zombie.geo.json"))
	Load(string, ...bool) (*GeometryFile, error)
}

func (e *Geometries) Add(key string, value *rp.Geometry) {
	p, ok := (*e)[key]
	if ok {
		p.Data = value
	} else {
		(*e)[key] = &GeometryFile{
			Data: value,
		}
	}
}

func (e *Geometries) Get(key string) (*rp.Geometry, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value.Data, true
}

func (e *Geometries) GetFile(key string) (*GeometryFile, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func (e *Geometries) Remove(key string) {
	delete(*e, key)
}

func (e *Geometries) Clear() {
	*e = make(map[string]*GeometryFile)
}

func (e *Geometries) All() map[string]*GeometryFile {
	return *e
}

func (e *Geometries) IsEmpty() bool {
	return len(*e) == 0
}

func (e *Geometries) Size() int {
	return len(*e)
}

func (e *Geometries) UnmarshalJSON(b []byte) error {
	return g_util.UnmarshalJSON(b, e)
}

func (e *Geometries) New(identifier string) *rp.Geometry {
	geo := rp.New(identifier)
	e.Add(identifier, geo)
	return geo
}

func (e *Geometries) Save(pathToRP string) error {
	for _, geo := range *e {
		b, err := geo.Encode()
		if err != nil {
			return err
		}

		err = g_util.WriteFile(filepath.Join(pathToRP, destDirectory.Geometry, geo.Subdir, geo.GetFilename()), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Geometries) LoadAll(pathToRP string) error {
	files, err := g_util.Walk(filepath.Join(pathToRP, destDirectory.Geometry))
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
		file = strings.TrimPrefix(file, pathToRP+string(filepath.Separator)+destDirectory.Geometry+string(filepath.Separator))
		subdir := filepath.Dir(file)
		filename := filepath.Base(file)

		(*e)[a.GetIdentifier()].Subdir = subdir
		(*e)[a.GetIdentifier()].Filename = filename
	}
	return nil
}

func (e *Geometries) Load(src string, add ...bool) (*GeometryFile, error) {
	a, err := rp.Load(src)
	if err != nil {
		return nil, err
	}

	filename := filepath.Base(src)

	// Get subdir
	dirs := filepath.Dir(src)
	// Replace destDirectory.Geometry with empty string
	dirs = strings.Replace(dirs, destDirectory.Geometry, "", 1)

	subdirs := strings.Split(dirs, string(filepath.Separator))
	subdir := ""
	// Reverse loop
	for i := len(subdirs) - 1; i >= 0; i-- {
		if subdirs[i] == destDirectory.Geometry {
			break
		}
		subdir = filepath.Join(subdirs[i], subdir)
	}

	if len(add) > 0 && add[0] || len(add) == 0 {
		e.Add(a.GetIdentifier(), a)
	}

	data, ok := e.GetFile(a.GetIdentifier())
	if !ok {
		data = &GeometryFile{
			Data: a,
		}
	}
	data.Subdir = subdir
	data.Filename = filename
	return data, nil
}
