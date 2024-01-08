package recipe

import g_util "github.com/respectZ/glowstone/util"

type RecipeFurnace struct {
	FormatVersion string             `json:"format_version"`
	RecipeFurnace *recipeFurnaceData `json:"minecraft:recipe_furnace"`
}

type recipeFurnaceData struct {
	Description recipeDescription `json:"description"`
	Tags        []string          `json:"tags"`
	Unlock      []recipeUnlock    `json:"unlock,omitempty"`
	Input       string            `json:"input"`
	Output      string            `json:"output"`
}

type IRecipeFurnace interface {
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

	// Get the recipe's output
	GetOutput() string

	// Set the recipe's output
	SetOutput(string)
}

func NewFurnace(identifier string) IRecipeFurnace {
	return &RecipeFurnace{
		FormatVersion: FORMAT_VERSION,
		RecipeFurnace: &recipeFurnaceData{
			Description: recipeDescription{
				Identifier: identifier,
			},
			Tags: []string{"furnace"},
		},
	}
}

func (r *RecipeFurnace) Encode() ([]byte, error) {
	return g_util.MarshalJSON(r)
}

func (r *RecipeFurnace) GetIdentifier() string {
	return r.RecipeFurnace.Description.Identifier
}

func (r *RecipeFurnace) SetIdentifier(identifier string) {
	r.RecipeFurnace.Description.Identifier = identifier
}

func (r *RecipeFurnace) GetTags() []string {
	return r.RecipeFurnace.Tags
}

func (r *RecipeFurnace) AddTag(tag string) {
	r.RecipeFurnace.Tags = append(r.RecipeFurnace.Tags, tag)
}

func (r *RecipeFurnace) AddUnlock(unlock recipeUnlock) {
	if r.RecipeFurnace.Unlock == nil {
		r.RecipeFurnace.Unlock = make([]recipeUnlock, 0)
	}
	r.RecipeFurnace.Unlock = append(r.RecipeFurnace.Unlock, unlock)
}

func (r *RecipeFurnace) GetInput() string {
	return r.RecipeFurnace.Input
}

func (r *RecipeFurnace) SetInput(input string) {
	r.RecipeFurnace.Input = input
}

func (r *RecipeFurnace) GetOutput() string {
	return r.RecipeFurnace.Output
}

func (r *RecipeFurnace) SetOutput(output string) {
	r.RecipeFurnace.Output = output
}
