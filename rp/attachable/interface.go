package attachable

type Attachable interface {
	Encode() ([]byte, error)

	// Returns the namespace:identifier of the attachable
	GetIdentifier() string

	// Sets the namespace:identifier of the attachable
	SetIdentifier(string)

	// Returns the materials of the attachable
	GetMaterials() map[string]string

	// Sets the materials of the attachable
	SetMaterials(map[string]string)

	// Returns the textures of the attachable
	GetTextures() map[string]string

	// Sets the textures of the attachable
	SetTextures(map[string]string)

	// Returns the geometry of the attachable
	GetGeometry() map[string]string

	// Sets the geometry of the attachable
	SetGeometry(map[string]string)

	// Returns the animations of the attachable
	GetAnimations() map[string]string

	// Sets the animations of the attachable
	SetAnimations(map[string]string)

	// Returns the item of the attachable
	GetItem() map[string]string

	// Sets the item of the attachable
	SetItem(map[string]string)

	// Returns the queryable geometry of the attachable
	GetQueryableGeometry() string

	// Sets the queryable geometry of the attachable
	SetQueryableGeometry(string)

	// Returns the render controllers of the attachable
	GetRenderControllers() []string

	// Sets the render controllers of the attachable
	SetRenderControllers([]string)
}
