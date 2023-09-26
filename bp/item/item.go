package item

import (
	"fmt"

	g_util "github.com/respectZ/glowstone/util"
)

const (
	FORMAT_VERSION = "1.20.30"
)

// Create a new item
//
// Example:
//
//	i := item.New("minecraft", "stick")
func New(namespace string, identifier string, subdir ...string) Item {
	s := ""
	if len(subdir) > 0 {
		s = subdir[0]
	}
	return &item{
		FormatVersion: FORMAT_VERSION,
		Item: minecraft_item{
			Description: item_description{
				Identifier: namespace + ":" + identifier,
			},
			Components: make(map[string]interface{}),
		},
		subdir: s,
	}
}

func (i *item) createMenuCategory() {
	if i.Item.Description.MenuCategory == nil {
		i.Item.Description.MenuCategory = &item_description_menuCategory{}
	}
}

func (i *item) checkMenuCategory() error {
	if i.Item.Description.MenuCategory == nil {
		return fmt.Errorf("item: menu category is nil")
	}
	return nil
}

func (i *item) Encode() ([]byte, error) {
	return g_util.MarshalJSON(i)
}

func (i *item) GetIdentifier() string {
	return i.Item.Description.Identifier
}

func (i *item) GetCategory() Category {
	if err := i.checkMenuCategory(); err != nil {
		return CATEGORY_NATURE
	}
	return i.Item.Description.MenuCategory.Category
}

func (i *item) SetCategory(category Category) {
	i.createMenuCategory()
	i.Item.Description.MenuCategory.Category = category
}

func (i *item) GetGroup() Group {
	if err := i.checkMenuCategory(); err != nil {
		return ""
	}
	return i.Item.Description.MenuCategory.Group
}

func (i *item) SetGroup(group Group) {
	i.createMenuCategory()
	i.Item.Description.MenuCategory.Group = group
}

func (i *item) GetIsHiddenInCommands() bool {
	if err := i.checkMenuCategory(); err != nil {
		return false
	}
	return i.Item.Description.MenuCategory.IsHiddenInCommands
}

func (i *item) SetIsHiddenInCommands(isHiddenInCommands bool) {
	i.createMenuCategory()
	i.Item.Description.MenuCategory.IsHiddenInCommands = isHiddenInCommands
}

func (i *item) GetComponent(name interface{}) (interface{}, error) {
	if component, ok := i.Item.Components[GetComponentName(name)]; ok {
		// Convert map to struct
		convertMapToStruct(component.(map[string]interface{}), name)
		// Assign component to the struct
		i.Item.Components[GetComponentName(name)] = name

		return component, nil
	}
	return nil, fmt.Errorf("component %s not found", name)
}

func (i *item) GetComponents() []interface{} {
	components := make([]interface{}, 0)
	for component := range i.Item.Components {
		components = append(components, component)
	}
	return components
}

func (i *item) AddComponent(components ...interface{}) {
	for _, component := range components {
		i.Item.Components[GetComponentName(component)] = component
	}
}

func (i *item) RemoveComponent(name string) {
	delete(i.Item.Components, name)
}
