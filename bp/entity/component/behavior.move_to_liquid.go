package component

// Code generated by scrapper/entityBp.ts;

// Allows the mob to move into a liquid when on land.  
type Behavior_MoveToLiquid struct {
  // Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
  GoalRadius float64 `json:"goal_radius,omitempty"`
  // The material type of the liquid block to find. Valid values are "Any", "Water", and "Lava".
  MaterialType string `json:"material_type,omitempty"`
  // The number of blocks each tick that the mob will check within its search range and height for a valid block to move to. A value of 0 will have the mob check every block within range in one tick
  SearchCount int `json:"search_count,omitempty"`
  // Height in blocks the mob will look for the liquid block to move towards
  SearchHeight int `json:"search_height,omitempty"`
  // The distance in blocks it will look for the liquid block to move towards
  SearchRange int `json:"search_range,omitempty"`
  // Movement speed multiplier of the mob when using this AI Goal
  SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}