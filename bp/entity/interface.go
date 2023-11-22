package entity

// TODO: Scripts / Animate
// Components, ComponentGroups, Events
type IEntity interface {
	Encode() ([]byte, error)
	// UnmarshalJSON([]byte) error

	// GetIdentifier returns the namespace:identifier of the entity
	//
	// Example:
	//
	//     id := entity.GetIdentifier() // returns "minecraft:zombie"
	GetIdentifier() string
	// GetRuntimeIdentifier returns the runtime identifier of the entity
	GetRuntimeIdentifier() (string, error)
	// SetRuntimeIdentifier sets the runtime identifier of the entity
	SetRuntimeIdentifier(string)

	GetSpawnable() bool
	SetSpawnable(bool)
	GetSummonable() bool
	SetSummonable(bool)

	/** Properties **/
	GetProperties() []entityProperties

	/** Component Groups **/

	/** Components **/

	// GetComponent returns the component of the entity
	GetComponent(interface{}) (interface{}, error)
	// GetComponents returns the components of the entity
	GetComponents() map[string]interface{}
	// AddComponent adds the component to the entity
	AddComponent(...interface{})
	// RemoveComponent removes the component from the entity
	RemoveComponent(string)

	// GetEvent returns the event of the entity
	GetEvent(string) (*EntityEvent, error)
	// GetEvents returns the events of the entity
	GetEvents() map[string]*EntityEvent
	// SetEvent sets the event of the entity
	SetEvent(string, EntityEvent)
	// RemoveEvent removes the event from the entity
	RemoveEvent(string)
}

type cg map[string]interface{}

type component_group interface {
	GetComponent(interface{}) (interface{}, error)
	GetComponents() [](interface{})
	AddComponent(...interface{})
	RemoveComponent(string)
}
