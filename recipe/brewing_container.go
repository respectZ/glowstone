package recipe

import (
	"fmt"
	"strings"

	"github.com/respectZ/glowstone/bp/recipe"
)

type RecipeBrewingContainer struct {
	Data recipe.RecipeBrewingContainer

	// Other stuffs
	Subdir string
}

// Create a new BrewingContainer recipe
//
// Example:
//
//	r := recipe.New("minecraft", "stick")
func NewBrewingContainer(namespace string, identifier string, subdir ...string) *RecipeBrewingContainer {
	r := &RecipeBrewingContainer{
		Data: recipe.NewBrewingContainer(fmt.Sprintf("%s:%s", namespace, identifier)),
	}
	if len(subdir) > 0 {
		r.Subdir = subdir[0]
	}
	return r
}

func (r *RecipeBrewingContainer) Encode() ([]byte, error) {
	return r.Data.Encode()
}

// GetIdentifier returns the identifier of the recipe without namespace
//
// Example:
//
//	identifier := r.GetIdentifier()
func (r *RecipeBrewingContainer) GetIdentifier() string {
	return strings.Split(r.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the recipe with namespace
//
// Example:
//
//	identifier := r.GetNamespaceIdentifier()
func (r *RecipeBrewingContainer) GetNamespaceIdentifier() string {
	return r.Data.GetIdentifier()
}

// GetSubdir returns the subdir of the recipe
//
// Example:
//
//	subdir := r.GetSubdir()
func (r *RecipeBrewingContainer) GetSubdir() string {
	return r.Subdir
}
