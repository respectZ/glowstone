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
	Get(string) (*rp.Geometry, bool)
	Remove(string)
	Clear()
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
	// 	model/entitys := project.RP.Geometry
	// 	model/entitys.New("glowstone:chair")
	// 	model/entitys.New("glowstone:table")
	// 	model/entitys.Save(path.Join("packs", "RP"))
	Save(string) error

	LoadAll(string) error
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

		err = g_util.Writefile(filepath.Join(pathToRP, destDirectory.Geometry, geo.Subdir, geo.GetFilename()), b)
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
