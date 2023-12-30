package recipe

import (
	"fmt"

	g_util "github.com/respectZ/glowstone/util"
)

const (
	FORMAT_VERSION string = "1.20.0"
)

/*
TODO: Proper naming
RecipeBrewingContainer (interface) -> IRecipeBrewingContainer
recipeBrewingContainer (struct) -> RecipeBrewingContainer
*/

const (
	RecipeContextPlayerInWater      string = "PlayerInWater"
	RecipeContextPlayerHasManyItems string = "PlayerHasManyItems"
	RecipeContextAlwaysUnlocked     string = "AlwaysUnlocked"
	RecipeContextNone               string = "None"
)

type recipeDescription struct {
	Identifier string `json:"identifier"`
}

type recipeIngredient struct {
	Item string `json:"item"`
	Data int    `json:"data,omitempty"`
	Tag  string `json:"tag,omitempty"`
}

type recipeUnlock struct {
	Item    string `json:"item,omitempty"`
	Data    int    `json:"data,omitempty"`
	Context string `json:"context,omitempty"`
}

type recipeResult struct {
	Item  string `json:"item"`
	Data  int    `json:"data,omitempty"`
	Count int    `json:"count,omitempty"`
}

type RecipeInterface interface {
	Encode() ([]byte, error)
	GetIdentifier() string
}

func Load(src string) (RecipeInterface, error) {
	var r interface{}
	err := g_util.LoadJSON(src, &r)
	if err != nil {
		return nil, err
	}
	switch v := r.(type) {
	case recipeBrewingContainer:
		return &v, nil
	case recipeBrewingMix:
		return &v, nil
	case recipeSmithingTransform:
		return &v, nil
	case recipeShaped:
		return &v, nil
	case recipeShapeless:
		return &v, nil
	case recipeFurnace:
		return &v, nil
	}
	return nil, fmt.Errorf("recipe: %s is not a recipe", src)
}

// TODO: recipeSmithingTransformData that accepts tag
/*
{
  "format_version": "1.20.10",
  "minecraft:recipe_smithing_trim": {
    "description": {
      "identifier": "minecraft:smithing_armor_trim"
    },
    "tags": [ "smithing_table" ],
    "template": {
      "tag": "minecraft:trim_templates"
    },
    "base": {
      "tag": "minecraft:trimmable_armors"
    },
    "addition": {
      "tag": "minecraft:trim_materials"
    }
  }
}

*/
