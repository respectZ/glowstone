package recipe

import (
	g_util "github.com/respectZ/glowstone/util"
)

type RecipeShaped struct {
	FormatVersion string            `json:"format_version"`
	RecipeShaped  *recipeShapedData `json:"minecraft:recipe_shaped"`
}

type recipeShapedData struct {
	Description recipeDescription           `json:"description"`
	Group       string                      `json:"group,omitempty"`
	Tags        []string                    `json:"tags"`
	Pattern     []string                    `json:"pattern"`
	Key         map[string]recipeIngredient `json:"key"`
	Unlock      []recipeUnlock              `json:"unlock,omitempty"`

	// Can be array of struct. TODO: implement array of result
	Result recipeResult `json:"result"`
}

type IRecipeShaped interface {
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

	// Get the recipe's pattern
	GetPattern() []string

	// Set the recipe's pattern
	SetPattern([]string)

	// Get the recipe's key
	GetKey() map[string]recipeIngredient

	// Add an ingredient to the recipe's key
	//
	// Example:
	// 	recipe := glowstone.NewRecipeShaped("minecraft", "stick")
	//  recipe.SetPattern([]string{"aaa", "aaa"})
	// 	recipe.AddKey("a", "minecraft:stick")
	AddKey(string, string, ...int)

	// Set the recipe's key
	SetKey(map[string]recipeIngredient)

	// Clear the recipe's key
	ClearKey()

	// Add Unlock to the recipe
	AddUnlock(recipeUnlock)

	// Get the recipe's result
	GetResult() recipeResult

	// Set the recipe's result
	//
	// Example:
	// 	recipe := glowstone.NewRecipeShaped("minecraft", "stick")
	// 	recipe.SetResult("minecraft:stick", 4) // 4 sticks.
	// 	recipe.SetResult("minecraft:stick", 4, 1) // 4 sticks with data 1.
	SetResult(string, ...int)
}

func NewShaped(identifier string) IRecipeShaped {
	return &RecipeShaped{
		FormatVersion: FORMAT_VERSION,
		RecipeShaped: &recipeShapedData{
			Description: recipeDescription{
				Identifier: identifier,
			},
			Tags:    []string{"crafting_table"},
			Pattern: make([]string, 0),
			Key:     make(map[string]recipeIngredient),
			Result:  recipeResult{Item: identifier},
		},
	}
}

func (r *RecipeShaped) Encode() ([]byte, error) {
	return g_util.MarshalJSON(r)
}

func (r *RecipeShaped) GetIdentifier() string {
	return r.RecipeShaped.Description.Identifier
}

func (r *RecipeShaped) SetIdentifier(identifier string) {
	r.RecipeShaped.Description.Identifier = identifier
}

func (r *RecipeShaped) GetGroup() string {
	return r.RecipeShaped.Group
}

func (r *RecipeShaped) SetGroup(group string) {
	r.RecipeShaped.Group = group
}

func (r *RecipeShaped) GetTags() []string {
	return r.RecipeShaped.Tags
}

func (r *RecipeShaped) AddTag(tag string) {
	r.RecipeShaped.Tags = append(r.RecipeShaped.Tags, tag)
}

func (r *RecipeShaped) GetPattern() []string {
	return r.RecipeShaped.Pattern
}

func (r *RecipeShaped) SetPattern(pattern []string) {
	r.RecipeShaped.Pattern = pattern
}

func (r *RecipeShaped) GetKey() map[string]recipeIngredient {
	return r.RecipeShaped.Key
}

func (r *RecipeShaped) AddKey(key string, item string, data ...int) {
	ingredient := recipeIngredient{
		Item: item,
	}
	if len(data) > 0 {
		ingredient.Data = data[0]
	}
	r.RecipeShaped.Key[key] = ingredient
}

func (r *RecipeShaped) SetKey(key map[string]recipeIngredient) {
	r.RecipeShaped.Key = key
}

func (r *RecipeShaped) ClearKey() {
	r.RecipeShaped.Key = make(map[string]recipeIngredient)
}

func (r *RecipeShaped) AddUnlock(unlock recipeUnlock) {
	if r.RecipeShaped.Unlock == nil {
		r.RecipeShaped.Unlock = make([]recipeUnlock, 0)
	}
	r.RecipeShaped.Unlock = append(r.RecipeShaped.Unlock, unlock)
}

func (r *RecipeShaped) GetResult() recipeResult {
	return r.RecipeShaped.Result
}

func (r *RecipeShaped) SetResult(item string, data ...int) {
	r.RecipeShaped.Result.Item = item
	if len(data) > 1 {
		r.RecipeShaped.Result.Data = data[0]
	}
	if len(data) > 0 {
		r.RecipeShaped.Result.Count = data[1]
	}
}
