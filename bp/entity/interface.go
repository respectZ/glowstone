package entity

// TODO: Scripts / Animate
// Components, ComponentGroups, Events
type Entity interface {
	Encode() ([]byte, error)

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

	GetAnimation(string) (string, error)
	GetAnimations() map[string]string
	SetAnimation(string, string)
	RemoveAnimation(string)

	// GetProperty returns the property of the entity
	GetProperty(string) (EntityProperties, error)
	// GetProperties returns the properties of the entity
	GetProperties() map[string]EntityProperties
	// SetPoperty sets the property of the entity
	SetProperty(string, EntityProperties)

	// GetComponentGroup returns the component group of the entity
	GetComponentGroup(string) (component_group, error)
	// GetComponentGroups returns the component groups of the entity
	GetComponentGroups() []string
	// SetComponentGroup sets the component group of the entity
	SetComponentGroup(string, ComponentGroup)
	// RemoveComponentGroup removes the component group of the entity
	RemoveComponentGroup(string)
	// AddComponentGroup adds the component group to the entity
	//
	// Example:
	//
	//     entity.AddComponentGroup("despawn", &component.InstantDespawn{})
	AddComponentGroup(string, ...interface{})

	// GetComponent returns the component of the entity
	GetComponent(interface{}) (interface{}, error)
	// GetComponents returns the components of the entity
	GetComponents() map[string]interface{}
	// AddComponent adds the component to the entity
	AddComponent(...interface{})
	// RemoveComponent removes the component from the entity
	RemoveComponent(string)

	// GetEvent returns the event of the entity
	GetEvent(string) (EntityEvent, error)
	// GetEvents returns the events of the entity
	GetEvents() map[string]EntityEvent
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
