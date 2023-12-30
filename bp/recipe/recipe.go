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
	var _brewingContainer recipeBrewingContainer
	var _brewingMix recipeBrewingMix
	var _smithingTransform recipeSmithingTransform
	var _shaped recipeShaped
	var _shapeless recipeShapeless
	var _furnace recipeFurnace

	err := g_util.LoadJSON(src, &_brewingContainer)
	if err == nil {
		if _brewingContainer.RecipeBrewingContainer != nil {
			return &_brewingContainer, nil
		}
	}
	err = g_util.LoadJSON(src, &_brewingMix)
	if err == nil {
		if _brewingMix.RecipeBrewingMix != nil {
			return &_brewingMix, nil
		}
	}
	err = g_util.LoadJSON(src, &_smithingTransform)
	if err == nil {
		if _smithingTransform.RecipeSmithingTransform != nil {
			return &_smithingTransform, nil
		}
	}
	err = g_util.LoadJSON(src, &_shaped)
	if err == nil {
		if _shaped.RecipeShaped != nil {
			return &_shaped, nil
		}
	}
	err = g_util.LoadJSON(src, &_shapeless)
	if err == nil {
		if _shapeless.RecipeShapeless != nil {
			return &_shapeless, nil
		}
	}
	err = g_util.LoadJSON(src, &_furnace)
	if err == nil {
		if _furnace.RecipeFurnace != nil {
			return &_furnace, nil
		}
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
