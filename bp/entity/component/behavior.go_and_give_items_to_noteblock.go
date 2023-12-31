package component

// Code generated by scrapper/entityBp.ts;

// The entity will attempt to toss the items from its inventory to a nearby recently played noteblock.  
type Behavior_GoAndGiveItemsToNoteblock struct {
  // Sets the time an entity should continue delivering items to a noteblock after hearing it.
  ListenTime *int `json:"listen_time,omitempty"`
  // [Trigger] Event(s) to run when this mob throws items.
  OnItemThrow interface{} `json:"on_item_throw,omitempty"`
  // Sets the desired distance to be reached before throwing the items towards the block.
  ReachBlockDistance *float64 `json:"reach_block_distance,omitempty"`
  // Sets the entity's speed when running toward the block.
  RunSpeed *float64 `json:"run_speed,omitempty"`
  // Sets the throw force.
  ThrowForce *float64 `json:"throw_force,omitempty"`
  // Sound to play when this mob throws an item.
  ThrowSound string `json:"throw_sound,omitempty"`
  // Sets the vertical throw multiplier that is applied on top of the throw force in the vertical direction.
  VerticalThrowMul *float64 `json:"vertical_throw_mul,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
