package recipe

import (
	"strings"

	"github.com/respectZ/glowstone/bp/recipe"
)

type RecipeInterface interface {
	Encode() ([]byte, error)
	GetIdentifier() string
	GetNamespaceIdentifier() string
	GetSubdir() string
}

type GenericRecipe struct {
	Data   recipe.RecipeInterface
	Subdir string
}

func Load(file string) (*GenericRecipe, error) {
	p, err := recipe.Load(file)
	if err != nil {
		return nil, err
	}
	r := &GenericRecipe{
		Data: p,
	}
	return r, nil
}

func (r *GenericRecipe) Encode() ([]byte, error) {
	return r.Data.Encode()
}

// GetIdentifier returns the identifier of the recipe without namespace
//
// Example:
//
//	identifier := r.GetIdentifier()
func (r *GenericRecipe) GetIdentifier() string {
	return strings.Split(r.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the recipe with namespace
//
// Example:
//
//	identifier := r.GetNamespaceIdentifier()
func (r *GenericRecipe) GetNamespaceIdentifier() string {
	return r.Data.GetIdentifier()
}

// GetSubdir returns the subdir of the recipe
//
// Example:
//
//	subdir := r.GetSubdir()
func (r *GenericRecipe) GetSubdir() string {
	return r.Subdir
}
