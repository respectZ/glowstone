package recipe

import g_util "github.com/respectZ/glowstone/util"

type recipeShapeless struct {
	FormatVersion   string               `json:"format_version"`
	RecipeShapeless *recipeShapelessData `json:"minecraft:recipe_shapeless"`
}

type recipeShapelessData struct {
	Description recipeDescription  `json:"description"`
	Group       string             `json:"group,omitempty"`
	Tags        []string           `json:"tags"`
	Ingredient  []recipeIngredient `json:"ingredients"`
	Unlock      []recipeUnlock     `json:"unlock,omitempty"`

	// Can be array of struct. TODO: implement array of result
	Result recipeResult `json:"result"`
}

type RecipeShapeless interface {
	Encode() ([]byte, error)
	// Return the recipe's identifier
	GetIdentifier() string

	// Set the recipe's identifier
	SetIdentifier(string)

	// Return the recipe's group
	GetGroup() string

	// Set the recipe's group
	SetGroup(string)

	// Return the recipe's tags
	GetTags() []string

	// Add a tag to the recipe
	AddTag(string)

	// Add Unlock to the recipe
	AddUnlock(recipeUnlock)

	// Get the recipe's result
	GetResult() recipeResult

	// Set the recipe's result
	//
	// Example:
	// 	recipe := glowstone.NewRecipeShapeless("minecraft", "stick")
	// 	recipe.SetResult("minecraft:stick", 4) // 4 sticks.
	// 	recipe.SetResult("minecraft:stick", 4, 1) // 4 sticks with data 1.
	SetResult(string, ...int)
}

func NewShapeless(identifier string) RecipeShapeless {
	return &recipeShapeless{
		FormatVersion: FORMAT_VERSION,
		RecipeShapeless: &recipeShapelessData{
			Description: recipeDescription{
				Identifier: identifier,
			},
			Tags:       []string{"crafting_table"},
			Ingredient: make([]recipeIngredient, 0),
			Result:     recipeResult{Item: identifier},
		},
	}
}

func (r *recipeShapeless) Encode() ([]byte, error) {
	return g_util.MarshalJSON(r)
}

func (r *recipeShapeless) GetIdentifier() string {
	return r.RecipeShapeless.Description.Identifier
}

func (r *recipeShapeless) SetIdentifier(identifier string) {
	r.RecipeShapeless.Description.Identifier = identifier
}

func (r *recipeShapeless) GetGroup() string {
	return r.RecipeShapeless.Group
}

func (r *recipeShapeless) SetGroup(group string) {
	r.RecipeShapeless.Group = group
}

func (r *recipeShapeless) GetTags() []string {
	return r.RecipeShapeless.Tags
}

func (r *recipeShapeless) AddTag(tag string) {
	r.RecipeShapeless.Tags = append(r.RecipeShapeless.Tags, tag)
}

func (r *recipeShapeless) AddUnlock(unlock recipeUnlock) {
	if r.RecipeShapeless.Unlock == nil {
		r.RecipeShapeless.Unlock = make([]recipeUnlock, 0)
	}
	r.RecipeShapeless.Unlock = append(r.RecipeShapeless.Unlock, unlock)
}

func (r *recipeShapeless) GetResult() recipeResult {
	return r.RecipeShapeless.Result
}

func (r *recipeShapeless) SetResult(item string, data ...int) {
	r.RecipeShapeless.Result.Item = item
	if len(data) > 0 {
		r.RecipeShapeless.Result.Count = data[0]
	}
	if len(data) > 1 {
		r.RecipeShapeless.Result.Data = data[1]
	}
}
