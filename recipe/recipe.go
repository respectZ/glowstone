package recipe

import (
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
