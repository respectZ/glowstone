package item

type Item interface {
	Encode() ([]byte, error)

	// GetIdentifier returns the identifier of the item
	//
	// Example:
	//
	//     id := item.GetIdentifier() // returns "minecraft:stick"
	GetIdentifier() string

	// GetCategory returns the category of the item
	//
	// Example:
	//
	//     category := item.GetCategory() // returns CATEGORY_NATURE
	GetCategory() Category

	// SetCategory sets the category of the item
	//
	// Example:
	//
	//     item.SetCategory(CATEGORY_ITEMS)
	SetCategory(Category)

	// GetGroup returns the group of the item
	//
	// Example:
	//
	//     group := item.GetGroup() // returns GROUP_ANVIL
	GetGroup() Group

	// SetGroup sets the group of the item
	//
	// Example:
	//
	//     item.SetGroup(GROUP_ANVIL)
	SetGroup(Group)

	// GetIsHiddenInCommands returns the is_hidden_in_commands of the item
	//
	// Example:
	//
	//     isHiddenInCommands := item.GetIsHiddenInCommands() // returns false
	GetIsHiddenInCommands() bool

	// SetIsHiddenInCommands sets the is_hidden_in_commands of the item
	//
	// Example:
	//
	//     item.SetIsHiddenInCommands(true)
	SetIsHiddenInCommands(bool)

	// GetComponent returns the component of the item
	//
	// Example:
	//    var iDamage ItemComponent.Damage
	//    component, err := item.GetComponent(&iDamage) // Do something with the iDamage
	GetComponent(interface{}) (interface{}, error)

	// GetComponents returns the components of the item
	//
	// Example:
	//
	//     components := item.GetComponents() // returns []interface{}
	GetComponents() []interface{}

	// AddComponent adds the component to the item
	//
	// Example:
	//
	//     item.AddComponent(&ItemComponent.Damage{Value: 1})
	AddComponent(...interface{})
}
