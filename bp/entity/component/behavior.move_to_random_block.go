package component

// Code generated by scrapper/entityBp.ts;

// Allows mob to move towards a random block.  
type Behavior_MoveToRandomBlock struct {
  // Defines the distance from the mob, in blocks, that the block to move to will be chosen.
  BlockDistance float64 `json:"block_distance,omitempty"`
  // Defines the distance in blocks the mob has to be from the block for the movement to be finished.
  WithinRadius float64 `json:"within_radius,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
