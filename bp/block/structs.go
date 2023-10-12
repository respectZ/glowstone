package block

type block struct {
	// Specifies the version of the game this block was made in. If the version is lower than the current version, any changes made to the entity in the vanilla version will be applied to it.
	FormatVersion string `json:"format_version"`

	// Block definition, which includes the "description" and "components" sections.
	MinecraftBlock *minecraftBlock `json:"minecraft:block"`
}

type minecraftBlock struct {
	// List of block characteristics that are applicable to all permutations of the block. The description MUST contain an identifier; the other fields are optional.
	Description *blockDescription `json:"description"`

	// List of all components used in this block.
	Component map[string]interface{} `json:"components"`

	// List of block permutations based on MoLang queries
	Permutations []*blockPermutation `json:"permutations,omitempty"`

	permutations []BlockPermutation
}

type blockDescription struct {
	// The identifier for this block. The name must include a namespace and must not use the Minecraft namespace unless overriding a Vanilla block.
	Identifier string `json:"identifier"`

	// Specifies the menu category and group for the block, which determine where this block is placed in the inventory and crafting table container screens. If this field is omitted, the block will not appear in the inventory or crafting table container screens.
	MenuCategory *blockMenuCategory `json:"menu_category,omitempty"`

	// Map of key/value pairs that maps the state name (key) to an array of all possible values for that state (value).
	//
	// Example:
	// 	"states": {
	// 		"custom:is_lit": ["true", "false"]
	// 	}
	States map[string]interface{} `json:"states,omitempty"`

	// Block traits are designed to be a shortcut for creators to use Vanilla block states without needing to define and manage a series of events or triggers on custom blocks. While custom states and permutations can be used to set multiple variations of the same block (whether it's on/off, is flammable, etc.), states that are exposed through traits allow you to access the inherent data certain Vanilla blocks hold.
	Traits *blockTraits `json:"traits,omitempty"`
}

type blockMenuCategory struct {
	// Determines which category this block will be placed under in the inventory and crafting table container screens. Options are "construction", "nature", "equipment", "items", and "none". If omitted or "none" is specified, the block will not appear in the inventory or crafting table container screens.
	Category string `json:"category,omitempty"`

	// Specifies the language file key that maps to which expandable/collapsible group this block will be a part of within a category. If this field is omitted, or there is no group whose name matches the loc string, this block will be placed standalone in the given category.
	Group string `json:"group,omitempty"`
}

type blockTraits struct {
	// placement_direction can add states containing information about the player's rotation when the block is placed. For example, if a block using placement_direction is placed while the player is facing south, the state value will be "south". This will allow a data-driven block to replicate the rotation behavior of a furnace, pumpkin, or terracotta block. Note that while the block contains information, permutations will need to be configured to determine how the block is placed/looks/acts.
	PlacementDirection *BlockPlacementDirection `json:"minecraft:placement_direction,omitempty"`

	// placement_position contains information about where the player placed the block. This allows a block to replicate the upside-down placement of slabs and stairs, as well as the attachment behavior of torches and vines.
	PlacementPosition *BlockPlacementPosition `json:"minecraft:placement_position,omitempty"`
}

type BlockPlacementDirection struct {
	// There are two states that can be included with this trait:
	//	- minecraft:cardinal_direction
	//	- minecraft:facing_direction
	EnabledStates []string `json:"enabled_states,omitempty"`

	// This causes the state within the trait to store a rotated value. In other words, with a rotation offset of 90.0, a block placed when the player is facing south would have a state of 'east'. This rotation offset only applies to the horizontal state values (north, south, east, west).
	YRotationOffset float64 `json:"y_rotation_offset,omitempty"`
}

type BlockPlacementPosition struct {
	// There are two states that can be included with this trait:
	//	- minecraft:block_face
	//	- minecraft:vertical_half
	EnabledStates []string `json:"enabled_states,omitempty"`
}

type blockPermutation struct {
	// A Molang expression that evaluates to true or false to determine if this permutation should be used. For permutation conditions you are limited to using one Molang query: "query.block_state()".
	Condition string `json:"condition,omitempty"`

	// List of all components that are used in this permutation.
	Components map[string]interface{} `json:"components"`
}
