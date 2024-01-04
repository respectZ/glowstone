package bp

import (
	"fmt"
	"reflect"
)

type _destDirectory struct {
	AnimationController string
	Animation           string
	Block               string
	Entity              string
	Item                string
	LootTable           string
	Recipe              string
}

var destDirectory = _destDirectory{
	AnimationController: "animation_controllers",
	Animation:           "animations",
	Block:               "blocks",
	Entity:              "entities",
	Item:                "items",
	LootTable:           "loot_tables",
	Recipe:              "recipes",
}

type GlowstoneBP struct {
	Path string

	AnimationController IAnimationControllerBP
	Animation           IAnimationBP
	Block               IBlockBP
	Entity              IEntityBP
	Item                IItemBP
	LootTable           ILootTable
	Recipe              IRecipeBP

	Manifest *Manifest
}

func New(path string) *GlowstoneBP {
	return &GlowstoneBP{
		Path:                path,
		AnimationController: &AnimationControllerBP{},
		Animation:           &AnimationBP{},
		Block:               &BlockBP{},
		Entity:              &EntityBP{},
		Item:                &ItemBP{},
		LootTable:           &LootTable{},
		Recipe:              &RecipeBP{},
		Manifest:            &Manifest{},
	}
}

func (g *GlowstoneBP) Initialize() error {
	return nil
}

func (g *GlowstoneBP) Preload() error {
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

func (g *GlowstoneBP) Save() error {
	val := reflect.ValueOf(g).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		// Prevent empty fields from being saved. eg: Manifest
		switch fieldName {
		case "Manifest":
			if g.Manifest.IsEmpty() {
				continue
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
