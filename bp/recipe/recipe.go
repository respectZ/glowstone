package recipe

const (
	FORMAT_VERSION string = "1.20.30"
)

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
