package recipe

import g_util "github.com/respectZ/glowstone/util"

type recipeSmithingTransform struct {
	FormatVersion           string                       `json:"format_version"`
	RecipeSmithingTransform *recipeSmithingTransformData `json:"minecraft:recipe_smithing_transform"`
}

type recipeSmithingTransformData struct {
	Description recipeDescription `json:"description"`
	Tags        []string          `json:"tags"`
	Template    string            `json:"template"`
	Base        string            `json:"base"`
	Addition    string            `json:"addition"`
	Result      string            `json:"result"`
}

type RecipeSmithingTransform interface {
	Encode() ([]byte, error)
	// Return the recipe's identifier
	GetIdentifier() string

	// Set the recipe's identifier
	SetIdentifier(string)

	// Return the recipe's tags
	GetTags() []string

	// Add a tag to the recipe
	AddTag(string)

	// Get the recipe's template
	GetTemplate() string

	// Set the recipe's template
	SetTemplate(string)

	// Get the recipe's base
	GetBase() string

	// Set the recipe's base
	SetBase(string)

	// Get the recipe's addition
	GetAddition() string

	// Set the recipe's addition
	SetAddition(string)
}

func NewSmithingTransform(identifier string) RecipeSmithingTransform {
	return &recipeSmithingTransform{
		FormatVersion: FORMAT_VERSION,
		RecipeSmithingTransform: &recipeSmithingTransformData{
			Description: recipeDescription{
				Identifier: identifier,
			},
			Addition: "minecraft:netherite_ingot",
			Tags:     []string{"smithing_table"},
		},
	}
}

func (r *recipeSmithingTransform) Encode() ([]byte, error) {
	return g_util.MarshalJSON(r)
}

func (r *recipeSmithingTransform) GetIdentifier() string {
	return r.RecipeSmithingTransform.Description.Identifier
}

func (r *recipeSmithingTransform) SetIdentifier(identifier string) {
	r.RecipeSmithingTransform.Description.Identifier = identifier
}

func (r *recipeSmithingTransform) GetTags() []string {
	return r.RecipeSmithingTransform.Tags
}

func (r *recipeSmithingTransform) AddTag(tag string) {
	r.RecipeSmithingTransform.Tags = append(r.RecipeSmithingTransform.Tags, tag)
}

func (r *recipeSmithingTransform) GetTemplate() string {
	return r.RecipeSmithingTransform.Template
}

func (r *recipeSmithingTransform) SetTemplate(template string) {
	r.RecipeSmithingTransform.Template = template
}

func (r *recipeSmithingTransform) GetBase() string {
	return r.RecipeSmithingTransform.Base
}

func (r *recipeSmithingTransform) SetBase(base string) {
	r.RecipeSmithingTransform.Base = base
}

func (r *recipeSmithingTransform) GetAddition() string {
	return r.RecipeSmithingTransform.Addition
}

func (r *recipeSmithingTransform) SetAddition(addition string) {
	r.RecipeSmithingTransform.Addition = addition
}
