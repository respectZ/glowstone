package recipe

import (
	"fmt"
	"strings"

	"github.com/respectZ/glowstone/bp/recipe"
)

type RecipeSmithingTransform struct {
	Data recipe.RecipeSmithingTransform

	// Other stuffs
	Subdir string
}

// Create a new SmithingTransform recipe
//
// Example:
//
//	r := recipe.New("minecraft", "stick")
func NewSmithingTransform(namespace string, identifier string, subdir ...string) *RecipeSmithingTransform {
	r := &RecipeSmithingTransform{
		Data: recipe.NewSmithingTransform(fmt.Sprintf("%s:%s", namespace, identifier)),
	}
	if len(subdir) > 0 {
		r.Subdir = subdir[0]
	}
	return r
}

func (r *RecipeSmithingTransform) Encode() ([]byte, error) {
	return r.Data.Encode()
}

// GetIdentifier returns the identifier of the recipe without namespace
//
// Example:
//
//	identifier := r.GetIdentifier()
func (r *RecipeSmithingTransform) GetIdentifier() string {
	return strings.Split(r.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the recipe with namespace
//
// Example:
//
//	identifier := r.GetNamespaceIdentifier()
func (r *RecipeSmithingTransform) GetNamespaceIdentifier() string {
	return r.Data.GetIdentifier()
}

// GetSubdir returns the subdir of the recipe
//
// Example:
//
//	subdir := r.GetSubdir()
func (r *RecipeSmithingTransform) GetSubdir() string {
	return r.Subdir
}
