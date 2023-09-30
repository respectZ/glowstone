package recipe

import g_util "github.com/respectZ/glowstone/util"

type recipeBrewingMix struct {
	FormatVersion    string                `json:"format_version"`
	RecipeBrewingMix *recipeBrewingMixData `json:"minecraft:recipe_brewing_mix"`
}

type recipeBrewingMixData struct {
	Description recipeDescription `json:"description"`
	Tags        []string          `json:"tags"`
	Unlock      []recipeUnlock    `json:"unlock,omitempty"`
	Input       string            `json:"input"`
	Reagent     string            `json:"reagent"`
	Output      string            `json:"output"`
}

type RecipeBrewingMix interface {
	Encode() ([]byte, error)
	// Return the recipe's identifier
	GetIdentifier() string

	// Set the recipe's identifier
	SetIdentifier(string)

	// Return the recipe's tags
	GetTags() []string

	// Add a tag to the recipe
	AddTag(string)

	// Add Unlock to the recipe
	AddUnlock(recipeUnlock)

	// Get the recipe's input
	GetInput() string

	// Set the recipe's input
	SetInput(string)

	// Get the recipe's reagent
	GetReagent() string

	// Set the recipe's reagent
	SetReagent(string)

	// Get the recipe's output
	GetOutput() string

	// Set the recipe's output
	SetOutput(string)
}

func NewBrewingMix(identifier string) RecipeBrewingMix {
	return &recipeBrewingMix{
		FormatVersion: FORMAT_VERSION,
		RecipeBrewingMix: &recipeBrewingMixData{
			Description: recipeDescription{
				Identifier: identifier,
			},
			Tags: []string{"brewing_stand"},
		},
	}
}

func (r *recipeBrewingMix) Encode() ([]byte, error) {
	return g_util.MarshalJSON(r)
}

func (r *recipeBrewingMix) GetIdentifier() string {
	return r.RecipeBrewingMix.Description.Identifier
}

func (r *recipeBrewingMix) SetIdentifier(identifier string) {
	r.RecipeBrewingMix.Description.Identifier = identifier
}

func (r *recipeBrewingMix) GetTags() []string {
	return r.RecipeBrewingMix.Tags
}

func (r *recipeBrewingMix) AddTag(tag string) {
	r.RecipeBrewingMix.Tags = append(r.RecipeBrewingMix.Tags, tag)
}

func (r *recipeBrewingMix) AddUnlock(unlock recipeUnlock) {
	r.RecipeBrewingMix.Unlock = append(r.RecipeBrewingMix.Unlock, unlock)
}

func (r *recipeBrewingMix) GetInput() string {
	return r.RecipeBrewingMix.Input
}

func (r *recipeBrewingMix) SetInput(input string) {
	r.RecipeBrewingMix.Input = input
}

func (r *recipeBrewingMix) GetReagent() string {
	return r.RecipeBrewingMix.Reagent
}

func (r *recipeBrewingMix) SetReagent(reagent string) {
	r.RecipeBrewingMix.Reagent = reagent
}

func (r *recipeBrewingMix) GetOutput() string {
	return r.RecipeBrewingMix.Output
}

func (r *recipeBrewingMix) SetOutput(output string) {
	r.RecipeBrewingMix.Output = output
}
