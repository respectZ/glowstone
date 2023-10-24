package item

import (
	"fmt"
	"reflect"

	g_util "github.com/respectZ/glowstone/util"
	util_component "github.com/respectZ/glowstone/util/component"
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
	componentName := util_component.GetComponentName(name)
	if component, ok := i.Item.Components[componentName]; ok {
		// switch componentType.Kind() {
		// case reflect.Ptr:
		// 	if nameType == reflect.TypeOf(reflect.ValueOf(component).Elem().Interface()) {
		// 		reflect.ValueOf(component).Elem().Set(reflect.ValueOf(name))
		// 		return component, nil
		// 	}
		// case reflect.Struct:
		// 	if nameType.Kind() == reflect.Ptr {
		// 		if componentType == reflect.TypeOf(reflect.ValueOf(name).Elem().Interface()) {
		// 			reflect.ValueOf(name).Elem().Set(reflect.ValueOf(component))
		// 			return component, nil
		// 		}
		// 	} else {
		// 		if componentType == reflect.TypeOf(reflect.ValueOf(name).Interface()) {
		// 			reflect.ValueOf(name).Set(reflect.ValueOf(component))
		// 			return component, nil
		// 		}
		// 	}
		// }
		// Convert map to struct
		// util_component.ConvertMapToStruct(component.(map[string]interface{}), name)
		// Assign component to the struct
		r, err := util_component.Get(component, name)
		if err != nil {
			return nil, err
		}
		i.Item.Components[componentName] = r

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
		componentName := util_component.GetComponentName(component)
		c := reflect.TypeOf(component)
		if c.Kind() == reflect.String && c.Name() == "string" {
			i.Item.Components[componentName] = &struct{}{}
		} else if reflect.TypeOf(component).Kind() == reflect.Ptr {
			i.Item.Components[componentName] = component
		} else {
			i.Item.Components[componentName] = &component
		}
	}
}

func (i *item) RemoveComponent(name string) {
	delete(i.Item.Components, name)
}
