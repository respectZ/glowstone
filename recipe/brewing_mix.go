package recipe

import (
	"fmt"
	"strings"

	"github.com/respectZ/glowstone/bp/recipe"
)

type RecipeBrewingMix struct {
	Data recipe.IRecipeBrewingMix

	// Other stuffs
	Subdir string
}

// Create a new BrewingMix recipe
//
// Example:
//
//	r := recipe.New("minecraft", "stick")
func NewBrewingMix(namespace string, identifier string, subdir ...string) *RecipeBrewingMix {
	r := &RecipeBrewingMix{
		Data: recipe.NewBrewingMix(fmt.Sprintf("%s:%s", namespace, identifier)),
	}
	if len(subdir) > 0 {
		r.Subdir = subdir[0]
	}
	return r
}

func (r *RecipeBrewingMix) Encode() ([]byte, error) {
	return r.Data.Encode()
}

// GetIdentifier returns the identifier of the recipe without namespace
//
// Example:
//
//	identifier := r.GetIdentifier()
func (r *RecipeBrewingMix) GetIdentifier() string {
	return strings.Split(r.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the recipe with namespace
//
// Example:
//
//	identifier := r.GetNamespaceIdentifier()
func (r *RecipeBrewingMix) GetNamespaceIdentifier() string {
	return r.Data.GetIdentifier()
}

// GetSubdir returns the subdir of the recipe
//
// Example:
//
//	subdir := r.GetSubdir()
func (r *RecipeBrewingMix) GetSubdir() string {
	return r.Subdir
}
