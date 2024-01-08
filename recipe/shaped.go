package recipe

import (
	"fmt"
	"strings"

	"github.com/respectZ/glowstone/bp/recipe"
)

type RecipeShaped struct {
	Data recipe.IRecipeShaped

	// Other stuffs
	Subdir string
}

// Create a new shaped recipe
//
// Example:
//
//	r := recipe.New("minecraft", "stick")
func NewShaped(namespace string, identifier string, subdir ...string) *RecipeShaped {
	r := &RecipeShaped{
		Data: recipe.NewShaped(fmt.Sprintf("%s:%s", namespace, identifier)),
	}
	if len(subdir) > 0 {
		r.Subdir = subdir[0]
	}
	return r
}

func (r *RecipeShaped) Encode() ([]byte, error) {
	return r.Data.Encode()
}

// GetIdentifier returns the identifier of the recipe without namespace
//
// Example:
//
//	identifier := r.GetIdentifier()
func (r *RecipeShaped) GetIdentifier() string {
	return strings.Split(r.Data.GetIdentifier(), ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the recipe with namespace
//
// Example:
//
//	identifier := r.GetNamespaceIdentifier()
func (r *RecipeShaped) GetNamespaceIdentifier() string {
	return r.Data.GetIdentifier()
}

// GetSubdir returns the subdir of the recipe
//
// Example:
//
//	subdir := r.GetSubdir()
func (r *RecipeShaped) GetSubdir() string {
	return r.Subdir
}
