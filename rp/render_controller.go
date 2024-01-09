package rp

import (
	"os"
	"path/filepath"
	"strings"

	rp "github.com/respectZ/glowstone/rp/render_controller"
	g_util "github.com/respectZ/glowstone/util"
)

type RenderControllers map[string]*RenderControllerFile

type IRenderControllers interface {
	Add(string, *rp.RenderControllerDefinition)

	// Get returns the RenderController of the given path.
	//
	// Example:
	//
	//  render_controller, ok := project.RP.RenderController.Get("vanilla/player.render_controllers.json")
	Get(string) (*rp.RenderControllerDefinition, bool)

	// GetFile returns the RenderControllerDefinition of the given path.
	//
	// Example:
	//
	//  render_controller, ok := project.RP.RenderController.GetFile("vanilla/player.render_controllers.json")
	GetFile(string) (*RenderControllerFile, bool)

	// Remove removes the RenderController of the given path.
	//
	// Example:
	//
	//  project.RP.RenderController.Remove("vanilla/player.render_controllers.json")
	Remove(string)
	Clear()
	All() map[string]*RenderControllerFile
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// Create a new animation controller.
	//
	// Example:
	//
	//	a := New("vanilla/player.render_controllers.json")
	New(string) *rp.RenderControllerDefinition

	// Saves all RenderControllerDefinition to the given path.
	//
	// Example:
	//
	//	project := glowstone.NewProject()
	// 	ac := project.RP.RenderController
	// 	ac.New("vanilla/player.render_controllers.json")
	// 	ac.New("furniture/table.render_controllers.json")
	// 	ac.Save(path.Join("packs", "RP"))
	Save(string) error

	LoadAll(string) error

	// Load a single render_controller file, without subdir and filename.
	//
	// Last parameter is whether the file should be added to the project.
	//
	// Default is true.
	//
	// Example:
	//
	//	render_controller, err := project.RP.RenderController.Load(filepath.Join(project.BP.Path, "render_controllers", "player.render_controller.json"))
	Load(string, ...bool) (*RenderControllerFile, error)
}

func (e *RenderControllers) Add(key string, value *rp.RenderControllerDefinition) {
	p, ok := (*e)[key]
	if ok {
		p.Data = value
	} else {
		(*e)[key] = &RenderControllerFile{
			Data: value,
		}
	}
}

func (e *RenderControllers) Get(key string) (*rp.RenderControllerDefinition, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value.Data, true
}

func (e *RenderControllers) GetFile(key string) (*RenderControllerFile, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func (e *RenderControllers) Remove(key string) {
	delete(*e, key)
}

func (e *RenderControllers) Clear() {
	*e = make(map[string]*RenderControllerFile)
}

func (e *RenderControllers) All() map[string]*RenderControllerFile {
	return *e
}

func (e *RenderControllers) IsEmpty() bool {
	return len(*e) == 0
}

func (e *RenderControllers) Size() int {
	return len(*e)
}

func (e *RenderControllers) UnmarshalJSON(b []byte) error {
	return g_util.UnmarshalJSON(b, e)
}

func (e *RenderControllers) New(dest string) *rp.RenderControllerDefinition {
	ac := rp.New()
	e.Add(dest, ac)
	return ac
}

func (e *RenderControllers) Save(pathToRP string) error {
	for dest, ac := range *e {
		b, err := ac.Encode()
		if err != nil {
			return err
		}

		err = g_util.WriteFile(filepath.Join(pathToRP, destDirectory.RenderController, dest), b)
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: check

func (e *RenderControllers) LoadAll(pathToRP string) error {
	files, err := g_util.Walk(filepath.Join(pathToRP, destDirectory.RenderController))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, file := range files {
		ac, err := rp.Load(file)
		if err != nil {
			return err
		}
		// Strip to get subdir
		file = strings.TrimPrefix(file, pathToRP+string(filepath.Separator)+destDirectory.RenderController+string(filepath.Separator))
		e.Add(file, ac)
	}
	return nil
}

// TODO: check

func (m *RenderControllers) Load(path string, add ...bool) (*RenderControllerFile, error) {
	a := rp.New()
	err := g_util.LoadJSON(path, &a)
	if err != nil {
		return nil, err
	}
	if len(add) > 0 && add[0] || len(add) == 0 {
		subdirs := strings.Split(path, string(filepath.Separator))
		file := filepath.Base(path)

		// Add subdir until "render_controllers" is reached
		// Reverse loop
		for i := len(subdirs) - 1; i >= 0; i-- {
			if subdirs[i] == destDirectory.RenderController {
				break
			}
			file = subdirs[i] + string(filepath.Separator) + file
		}
		m.Add(file, a)
	}
	data, ok := m.GetFile(path)
	if !ok {
		return &RenderControllerFile{
			Data: a,
		}, nil
	}
	return data, nil
}
