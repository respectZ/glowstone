package recipe

import (
	"fmt"
	"strings"

	"github.com/respectZ/glowstone/bp/recipe"
)

type RecipeFurnace struct {
	Data recipe.IRecipeFurnace

	// Other stuffs
	Subdir string
}

// Create a new Furnace recipe
//
// Example:
//
//	r := recipe.New("minecraft", "stick")
func NewFurnace(namespace string, identifier string, subdir ...string) *RecipeFurnace {
	r := &RecipeFurnace{
		Data: recipe.NewFurnace(fmt.Sprintf("%s:%s", namespace, identifier)),
	}
	if len(subdir) > 0 {
		r.Subdir = subdir[0]
	}
	return r
}

func (r *RecipeFurnace) Encode() ([]byte, error) {
	return r.Data.Encode()
}

// GetIdentifier returns the identifier of the recipe without namespace
//
// Example:
//
//	identifier := r.GetIdentifier()
func (r *RecipeFurnace) GetIdentifier() string {
	return strings.Split(r.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the recipe with namespace
//
// Example:
//
//	identifier := r.GetNamespaceIdentifier()
func (r *RecipeFurnace) GetNamespaceIdentifier() string {
	return r.Data.GetIdentifier()
}

// GetSubdir returns the subdir of the recipe
//
// Example:
//
//	subdir := r.GetSubdir()
func (r *RecipeFurnace) GetSubdir() string {
	return r.Subdir
}
