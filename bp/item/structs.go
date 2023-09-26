package item

import glowstone "github.com/respectZ/glowstone/util"

type item struct {
	FormatVersion string         `json:"format_version"`
	Item          minecraft_item `json:"minecraft:item"`
	subdir        string
}

type minecraft_item struct {
	Description item_description       `json:"description"`
	Components  map[string]interface{} `json:"components"`
}

type item_description struct {
	Identifier   string                         `json:"identifier"`
	MenuCategory *item_description_menuCategory `json:"menu_category,omitempty"`
}

type item_description_menuCategory struct {
	Category           Category `json:"category,omitempty"`
	Group              Group    `json:"group,omitempty"`
	IsHiddenInCommands bool     `json:"is_hidden_in_commands,omitempty"`
}

func Load(src string) (Item, error) {
	e := &item{}
	err := glowstone.LoadJSON(src, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
