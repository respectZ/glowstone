package block

import (
	"fmt"
	"reflect"
)

func (b *blockDescription) GetIdentifier() string {
	return b.Identifier
}

func (b *blockDescription) SetIdentifier(identifier string) {
	b.Identifier = identifier
}

func (b *blockDescription) GetMenuCategory() *blockMenuCategory {
	return b.MenuCategory
}

func (b *blockDescription) SetCategory(category string) {
	if b.MenuCategory == nil {
		b.MenuCategory = &blockMenuCategory{}
	}
	b.MenuCategory.Category = category
}

func (b *blockDescription) SetGroup(group string) {
	b.MenuCategory.Group = group
}

func (b *blockDescription) AddState(name string, value interface{}) {
	if b.States == nil {
		b.States = make(map[string]interface{})
	}
	b.States[name] = value
}

func (b *blockDescription) GetState(name string) interface{} {
	return b.States[name]
}

func (b *blockDescription) AddTrait(name string, values ...string) error {
	if b.Traits == nil {
		b.Traits = &blockTraits{}
	}
	switch name {
	case "minecraft:placement_direction":
		b.Traits.PlacementDirection = &BlockPlacementDirection{
			EnabledStates: values,
		}
	case "minecraft:placement_position":
		b.Traits.PlacementPosition = &BlockPlacementPosition{
			EnabledStates: values,
		}
	default:
		return fmt.Errorf("trait %s not found", name)
	}
	return nil
}

func (b *blockDescription) AddTraitWithStruct(trait interface{}) error {
	if b.Traits == nil {
		b.Traits = &blockTraits{}
	}
	switch v := trait.(type) {
	case *BlockPlacementDirection:
		b.Traits.PlacementDirection = v
	case *BlockPlacementPosition:
		b.Traits.PlacementPosition = v
	case BlockPlacementDirection:
		b.Traits.PlacementDirection = &v
	case BlockPlacementPosition:
		b.Traits.PlacementPosition = &v
	default:
		return fmt.Errorf("trait %s not found", reflect.TypeOf(trait))
	}
	return nil
}
