package rp

import (
	rp "github.com/respectZ/glowstone/rp/render_controller"
)

type RenderControllerFile struct {
	Data *rp.RenderControllerDefinition

	// Other stuff
}

// Create a new render controller file
//
// Example:
//
//	render_controller := rp.NewRenderController("weapon/zwei.render_controllers.json")
func NewRenderController(dest string) *RenderControllerFile {
	return &RenderControllerFile{
		Data: rp.New(),
	}
}

func LoadRenderController(src string) (*rp.RenderControllerDefinition, error) {
	rc, err := rp.Load(src)
	if err != nil {
		return nil, err
	}
	return rc, nil
}

func (e *RenderControllerFile) Encode() ([]byte, error) {
	rc, err := e.Data.Encode()
	if err != nil {
		return nil, err
	}
	return rc, nil
}

// Returns all render controllers that exist in the render controller file.
//
// Example:
//
//	render_controllers := e.GetRenderControllers() // []string{"controller.render.weapon", "controller.render.shield"}
func (e *RenderControllerFile) GetRenderControllers() []string {
	var render_controllers []string
	for k := range e.Data.RenderControllers.All() {
		render_controllers = append(render_controllers, k)
	}
	return render_controllers
}
