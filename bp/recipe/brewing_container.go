package recipe

import g_util "github.com/respectZ/glowstone/util"

type recipeBrewingContainer struct {
	FormatVersion          string                      `json:"format_version"`
	RecipeBrewingContainer *recipeBrewingContainerData `json:"minecraft:recipe_brewing_container"`
}

type recipeBrewingContainerData struct {
	Description recipeDescription `json:"description"`
	Tags        []string          `json:"tags"`
	Unlock      []recipeUnlock    `json:"unlock,omitempty"`
	Input       string            `json:"input"`
	Reagent     string            `json:"reagent"`
	Output      string            `json:"output"`
}

type RecipeBrewingContainer interface {
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

func NewBrewingContainer(identifier string) RecipeBrewingContainer {
	return &recipeBrewingContainer{
		FormatVersion: FORMAT_VERSION,
		RecipeBrewingContainer: &recipeBrewingContainerData{
			Description: recipeDescription{
				Identifier: identifier,
			},
			Tags: []string{"brewing_stand"},
		},
	}
}

func (r *recipeBrewingContainer) Encode() ([]byte, error) {
	return g_util.MarshalJSON(r)
}

func (r *recipeBrewingContainer) GetIdentifier() string {
	return r.RecipeBrewingContainer.Description.Identifier
}

func (r *recipeBrewingContainer) SetIdentifier(identifier string) {
	r.RecipeBrewingContainer.Description.Identifier = identifier
}

func (r *recipeBrewingContainer) GetTags() []string {
	return r.RecipeBrewingContainer.Tags
}

func (r *recipeBrewingContainer) AddTag(tag string) {
	r.RecipeBrewingContainer.Tags = append(r.RecipeBrewingContainer.Tags, tag)
}

func (r *recipeBrewingContainer) AddUnlock(unlock recipeUnlock) {
	r.RecipeBrewingContainer.Unlock = append(r.RecipeBrewingContainer.Unlock, unlock)
}

func (r *recipeBrewingContainer) GetInput() string {
	return r.RecipeBrewingContainer.Input
}

func (r *recipeBrewingContainer) SetInput(input string) {
	r.RecipeBrewingContainer.Input = input
}

func (r *recipeBrewingContainer) GetReagent() string {
	return r.RecipeBrewingContainer.Reagent
}

func (r *recipeBrewingContainer) SetReagent(reagent string) {
	r.RecipeBrewingContainer.Reagent = reagent
}

func (r *recipeBrewingContainer) GetOutput() string {
	return r.RecipeBrewingContainer.Output
}

func (r *recipeBrewingContainer) SetOutput(output string) {
	r.RecipeBrewingContainer.Output = output
}
