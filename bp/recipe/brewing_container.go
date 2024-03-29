package recipe

import g_util "github.com/respectZ/glowstone/util"

type RecipeBrewingContainer struct {
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

type IRecipeBrewingContainer interface {
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

func NewBrewingContainer(identifier string) IRecipeBrewingContainer {
	return &RecipeBrewingContainer{
		FormatVersion: FORMAT_VERSION,
		RecipeBrewingContainer: &recipeBrewingContainerData{
			Description: recipeDescription{
				Identifier: identifier,
			},
			Tags: []string{"brewing_stand"},
		},
	}
}

func (r *RecipeBrewingContainer) Encode() ([]byte, error) {
	return g_util.MarshalJSON(r)
}

func (r *RecipeBrewingContainer) GetIdentifier() string {
	return r.RecipeBrewingContainer.Description.Identifier
}

func (r *RecipeBrewingContainer) SetIdentifier(identifier string) {
	r.RecipeBrewingContainer.Description.Identifier = identifier
}

func (r *RecipeBrewingContainer) GetTags() []string {
	return r.RecipeBrewingContainer.Tags
}

func (r *RecipeBrewingContainer) AddTag(tag string) {
	r.RecipeBrewingContainer.Tags = append(r.RecipeBrewingContainer.Tags, tag)
}

func (r *RecipeBrewingContainer) AddUnlock(unlock recipeUnlock) {
	r.RecipeBrewingContainer.Unlock = append(r.RecipeBrewingContainer.Unlock, unlock)
}

func (r *RecipeBrewingContainer) GetInput() string {
	return r.RecipeBrewingContainer.Input
}

func (r *RecipeBrewingContainer) SetInput(input string) {
	r.RecipeBrewingContainer.Input = input
}

func (r *RecipeBrewingContainer) GetReagent() string {
	return r.RecipeBrewingContainer.Reagent
}

func (r *RecipeBrewingContainer) SetReagent(reagent string) {
	r.RecipeBrewingContainer.Reagent = reagent
}

func (r *RecipeBrewingContainer) GetOutput() string {
	return r.RecipeBrewingContainer.Output
}

func (r *RecipeBrewingContainer) SetOutput(output string) {
	r.RecipeBrewingContainer.Output = output
}
