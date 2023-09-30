package recipe

type RecipeInterface interface {
	Encode() ([]byte, error)
	GetIdentifier() string
	GetNamespaceIdentifier() string
	GetSubdir() string
}
