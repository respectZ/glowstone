package bp

import (
	"strings"

	bp "github.com/respectZ/glowstone/bp/recipe"
)

type RecipeFile struct {
	Data bp.RecipeInterface

	// Other stuff
	Subdir   string
	Filename string
}

// Create a new brewing container recipe file
//
// Example:
//
//	recipe := bp.NewRecipeBrewingContainer("glostone:diamond_sword")
func NewRecipeBrewingContainer(identifier string) *RecipeFile {
	return &RecipeFile{
		Data: bp.NewBrewingContainer(identifier),
	}
}

// Create a new brewing mix recipe file
//
// Example:
//
//	recipe := bp.NewRecipeBrewingMix("glostone:diamond_sword")
func NewRecipeBrewingMix(identifier string) *RecipeFile {
	return &RecipeFile{
		Data: bp.NewBrewingMix(identifier),
	}
}

// Create a new furnace recipe file
//
// Example:
//
//	recipe := bp.NewRecipeFurnace("glostone:diamond_sword")
func NewRecipeFurnace(identifier string) *RecipeFile {
	return &RecipeFile{
		Data: bp.NewFurnace(identifier),
	}
}

// Create a new shaped recipe file
//
// Example:
//
//	recipe := bp.NewRecipeShaped("glostone:diamond_sword")
func NewRecipeShaped(identifier string) *RecipeFile {
	return &RecipeFile{
		Data: bp.NewShaped(identifier),
	}
}

// Create a new shapeless recipe file
//
// Example:
//
//	recipe := bp.NewRecipeShapeless("glostone:diamond_sword")
func NewRecipeShapeless(identifier string) *RecipeFile {
	return &RecipeFile{
		Data: bp.NewShapeless(identifier),
	}
}

// Create a new smithing recipe file
//
// Example:
//
//	recipe := bp.NewRecipeSmithing("glostone:diamond_sword")
func NewRecipeSmithing(identifier string) *RecipeFile {
	return &RecipeFile{
		Data: bp.NewSmithingTransform(identifier),
	}
}

// Load an recipe file from the given path
//
// Example:
//
//	recipe, err := bp.LoadRecipe(filepath.Join("packs", "BP", "recipe", "sword.json")
func LoadRecipe(src string) (bp.RecipeInterface, error) {
	recipe, err := bp.Load(src)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

// Encode returns []byte of the recipe.
//
// Example:
//
//	recipe, err := e.Encode()
func (e *RecipeFile) Encode() ([]byte, error) {
	recipe, err := e.Data.Encode()
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

// GetIdentifier returns the identifier of the recipe without namespace
//
// Example:
//
//	identifier := e.GetIdentifier()
func (e *RecipeFile) GetIdentifier() string {
	return strings.Split(e.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the recipe with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (e *RecipeFile) GetNamespaceIdentifier() string {
	return e.Data.GetIdentifier()
}

func (e *RecipeFile) GetFilename() string {
	if e.Filename == "" {
		return e.GetIdentifier() + ".json"
	}
	return e.Filename
}
