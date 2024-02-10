package rp

import (
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/respectZ/glowstone/rp/sound"
	"github.com/respectZ/glowstone/rp/texts"
	"github.com/respectZ/glowstone/rp/texture"
)

type _destDirectory struct {
	AnimationController string
	Animation           string
	Attachable          string
	Entity              string
	Geometry            string
	Particle            string
	RenderController    string
}

var destDirectory = _destDirectory{
	AnimationController: "animation_controllers",
	Attachable:          "attachables",
	Animation:           "animations",
	Entity:              "entity",
	Geometry:            filepath.Join("models", "entity"),
	Particle:            "particles",
	RenderController:    "render_controllers",
}

type GlowstoneRP struct {
	// Path to the resource pack.
	Path string

	// Contains all attachables in the project.
	Attachable IAttachables

	// Contains all animation controllers in the project.
	AnimationController IAnimationControllers

	// Contains all animations in the project.
	Animation IAnimations

	// Contains all entities in the project.
	Entity IEntities

	// Contains all particles in the project.
	Particle IParticles

	// Contains all geometries in the project.
	Geometry IGeometries

	// Contains all render controllers in the project.
	RenderController IRenderControllers

	// Data from the 'textures/item_texture.json' file.
	ItemTexture *texture.ItemTexture

	// Data from the 'textures/terrain_texture.json' file.
	TerrainTexture *texture.TerrainTexture

	// Data from the 'sound_definitions.json' file.
	SoundDefinition *sound.SoundDefinition

	// Data from the 'blocks.json' file.
	Blocks *Blocks

	// Data from the 'texts/en_US.lang' file.
	Lang *texts.Lang

	// Manifest of the resource pack.
	Manifest *Manifest
}

func New(path string) *GlowstoneRP {
	return &GlowstoneRP{
		Path:                path,
		Attachable:          &Attachables{},
		AnimationController: &AnimationControllers{},
		Animation:           &Animations{},
		Entity:              &Entities{},
		Particle:            &Particles{},
		Geometry:            &Geometries{},
		RenderController:    &RenderControllers{},

		ItemTexture:     &texture.ItemTexture{},
		TerrainTexture:  &texture.TerrainTexture{},
		SoundDefinition: &sound.SoundDefinition{FormatVersion: sound.FORMAT_VERSION},
		Blocks:          &Blocks{},
		Lang:            &texts.Lang{},
		Manifest:        &Manifest{},
	}
}

func (g *GlowstoneRP) Preload() error {
	val := reflect.ValueOf(g).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		loadMethod := field.MethodByName("LoadAll")
		if !loadMethod.IsValid() {
			loadMethod = field.MethodByName("Load")
		}
		if loadMethod.IsValid() {
			result := loadMethod.Call([]reflect.Value{reflect.ValueOf(g.Path)})
			if len(result) > 0 && !result[0].IsNil() {
				err := result[0].Interface().(error)
				return fmt.Errorf("failed to load %s: %v", val.Type().Field(i).Name, err)
			}
		}
	}
	return nil
}

func (g *GlowstoneRP) Save() error {
	val := reflect.ValueOf(g).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		// Prevent overwriting files with empty data
		switch fieldName {
		case "Manifest", "ItemTexture", "TerrainTexture", "SoundDefinition", "Blocks", "Lang":
			isEmptyMethod := field.MethodByName("IsEmpty")
			if isEmptyMethod.IsValid() {
				result := isEmptyMethod.Call([]reflect.Value{})
				if len(result) > 0 && result[0].Interface().(bool) {
					continue
				}
			}
		}

		if saveMethod := field.MethodByName("Save"); saveMethod.IsValid() {
			result := saveMethod.Call([]reflect.Value{reflect.ValueOf(g.Path)})
			if len(result) > 0 && !result[0].IsNil() {
				return result[0].Interface().(error)
			}
		}
	}
	return nil
}
