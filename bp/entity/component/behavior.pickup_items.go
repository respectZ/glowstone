package component

// Code generated by scrapper/entityBp.ts;

// Allows the mob to pick up items on the ground.  
type Behavior_PickupItems struct {
  // If true, the mob can pickup any item
  CanPickupAnyItem bool `json:"can_pickup_any_item,omitempty"`
  // If true, the mob can pickup items to its hand or armor slots
  CanPickupToHandOrEquipment bool `json:"can_pickup_to_hand_or_equipment,omitempty"`
  // List of items this mob will not pick up
  ExcludedItems []interface{} `json:"excluded_items,omitempty"`
  // Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
  GoalRadius float64 `json:"goal_radius,omitempty"`
  // Maximum distance this mob will look for items to pick up
  MaxDist float64 `json:"max_dist,omitempty"`
  // If true, depending on the difficulty, there is a random chance that the mob may not be able to pickup items
  PickupBasedOnChance bool `json:"pickup_based_on_chance,omitempty"`
  // Movement speed multiplier of the mob when using this AI Goal
  SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
  // If true, this mob will chase after the target as long as it's a valid target
  TrackTarget bool `json:"track_target,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}