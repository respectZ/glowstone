package recipe

import (
	"fmt"
	"strings"

	"github.com/respectZ/glowstone/bp/recipe"
)

type RecipeShapeless struct {
	Data recipe.RecipeShapeless

	// Other stuffs
	Subdir string
}

// Create a new Shapeless recipe
//
// Example:
//
//	r := recipe.New("minecraft", "stick")
func NewShapeless(namespace string, identifier string, subdir ...string) *RecipeShapeless {
	r := &RecipeShapeless{
		Data: recipe.NewShapeless(fmt.Sprintf("%s:%s", namespace, identifier)),
	}
	if len(subdir) > 0 {
		r.Subdir = subdir[0]
	}
	return r
}

func (r *RecipeShapeless) Encode() ([]byte, error) {
	return r.Data.Encode()
}

// GetIdentifier returns the identifier of the recipe without namespace
//
// Example:
//
//	identifier := r.GetIdentifier()
func (r *RecipeShapeless) GetIdentifier() string {
	return strings.Split(r.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the recipe with namespace
//
// Example:
//
//	identifier := r.GetNamespaceIdentifier()
func (r *RecipeShapeless) GetNamespaceIdentifier() string {
	return r.Data.GetIdentifier()
}

// GetSubdir returns the subdir of the recipe
//
// Example:
//
//	subdir := r.GetSubdir()
func (r *RecipeShapeless) GetSubdir() string {
	return r.Subdir
}
