package block

type Block interface {
	Encode() ([]byte, error)

	GetIdentifier() string
	SetIdentifier(string)

	GetDescription() BlockDescription
	GetComponent(interface{}) error
	GetPermutations() []BlockPermutation

	AddComponent(...interface{})

	// Add a permutation to the block.
	//
	// Example:
	//
	//  mining := component.DestructibleByMining{SecondsToDestroy: 0.6}
	//  AddPermutation("query.block_state('custom:is_lit')", mining)
	AddPermutation(string, ...interface{})
}

type BlockDescription interface {
	GetIdentifier() string
	SetIdentifier(string)

	// Returns a struct containing category and group.
	GetMenuCategory() *blockMenuCategory

	// Set the block's category.
	SetCategory(string)

	// Set the block's group.
	SetGroup(string)

	/*** States ***/

	// Add a state to the block.
	//
	// Example:
	//
	//  AddState("custom:is_lit", []string{"true", "false"})
	//  AddState("custom:power", []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15})
	AddState(string, interface{})

	// Get a state from the block.
	//
	// Example:
	//
	//  is_lit := GetState("custom:is_lit")
	GetState(string) interface{}

	/*** Traits ***/

	// Add a trait to the block.
	//
	// Example:
	//
	//  err := AddTrait("minecraft:placement_direction", "cardinal_direction", "facing_direction")
	//  err := AddTrait("minecraft:block_face", "minecraft:vertical_half")
	AddTrait(string, ...string) error

	// Add a trait to the block using a struct type.
	//
	// Example:
	//
	// AddTraitWithStruct(&BlockPlacementDirection{
	// 	EnabledStates: []string{"cardinal_direction", "facing_direction"},
	// 	YRotationOffset: 90.0,
	// })
	AddTraitWithStruct(interface{}) error
}

type BlockPermutation interface {
	GetCondition() string
	SetCondition(string)
	SetComponent(string, interface{})
}
