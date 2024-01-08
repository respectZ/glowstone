package recipe

import g_util "github.com/respectZ/glowstone/util"

type RecipeSmithingTransform struct {
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

type IRecipeSmithingTransform interface {
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

func NewSmithingTransform(identifier string) IRecipeSmithingTransform {
	return &RecipeSmithingTransform{
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

func (r *RecipeSmithingTransform) Encode() ([]byte, error) {
	return g_util.MarshalJSON(r)
}

func (r *RecipeSmithingTransform) GetIdentifier() string {
	return r.RecipeSmithingTransform.Description.Identifier
}

func (r *RecipeSmithingTransform) SetIdentifier(identifier string) {
	r.RecipeSmithingTransform.Description.Identifier = identifier
}

func (r *RecipeSmithingTransform) GetTags() []string {
	return r.RecipeSmithingTransform.Tags
}

func (r *RecipeSmithingTransform) AddTag(tag string) {
	r.RecipeSmithingTransform.Tags = append(r.RecipeSmithingTransform.Tags, tag)
}

func (r *RecipeSmithingTransform) GetTemplate() string {
	return r.RecipeSmithingTransform.Template
}

func (r *RecipeSmithingTransform) SetTemplate(template string) {
	r.RecipeSmithingTransform.Template = template
}

func (r *RecipeSmithingTransform) GetBase() string {
	return r.RecipeSmithingTransform.Base
}

func (r *RecipeSmithingTransform) SetBase(base string) {
	r.RecipeSmithingTransform.Base = base
}

func (r *RecipeSmithingTransform) GetAddition() string {
	return r.RecipeSmithingTransform.Addition
}

func (r *RecipeSmithingTransform) SetAddition(addition string) {
	r.RecipeSmithingTransform.Addition = addition
}
