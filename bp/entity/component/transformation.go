package component

// Code generated by scrapper/entityBp.ts;

// [Not a component] List of components to add to the entity after the transformation
type TransformationAdd struct {
  // Names of component groups to add
  ComponentGroups []interface{} `json:"component_groups,omitempty"`
}

// [Not a component] Defines the properties of the delay for the transformation
type TransformationDelay struct {
  // Chance that the entity will look for nearby blocks that can speed up the transformation. Value must be between 0.0 and 1.0
  BlockAssistChance float64 `json:"block_assist_chance,omitempty"`
  // Chance that, once a block is found, will help speed up the transformation
  BlockChance float64 `json:"block_chance,omitempty"`
  // Maximum number of blocks the entity will look for to aid in the transformation. If not defined or set to 0, it will be set to the block radius
  BlockMax int `json:"block_max,omitempty"`
  // Distance in Blocks that the entity will search for blocks that can help the transformation
  BlockRadius int `json:"block_radius,omitempty"`
  // List of blocks that can help the transformation of this entity
  BlockTypes []interface{} `json:"block_types,omitempty"`
  // Time in seconds before the entity transforms
  Value float64 `json:"value,omitempty"`
}

// Defines an entity's transformation from the current definition into another  
type Transformation struct {
  // List of components to add to the entity after the transformation
  Add TransformationAdd `json:"add,omitempty"`
  // Sound to play when the transformation starts
  BeginTransformSound string `json:"begin_transform_sound,omitempty"`
  // Defines the properties of the delay for the transformation
  Delay TransformationDelay `json:"delay,omitempty"`
  // Cause the entity to drop all equipment upon transformation
  DropEquipment bool `json:"drop_equipment,omitempty"`
  // Cause the entity to drop all items in inventory upon transformation
  DropInventory bool `json:"drop_inventory,omitempty"`
  // Entity Definition that this entity will transform into
  Into string `json:"into,omitempty"`
  // If this entity has trades and has leveled up, it should maintain that level after transformation.
  KeepLevel bool `json:"keep_level,omitempty"`
  // If this entity is owned by another entity, it should remain owned after transformation.
  KeepOwner bool `json:"keep_owner,omitempty"`
  // Cause the entity to keep equipment after going through transformation
  PreserveEquipment bool `json:"preserve_equipment,omitempty"`
  // Sound to play when the entity is done transforming
  TransformationSound string `json:"transformation_sound,omitempty"`
}