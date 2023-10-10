package component

// Code generated by scrapper/entityBp.ts;

// Allows the NPC to use the composter POI to convert excess seeds into bone meal.  
type Behavior_WorkComposter struct {
  // The amount of ticks the NPC will stay in their the work location
  ActiveTime int `json:"active_time,omitempty"`
  // The maximum number of times the mob will interact with the composter.
  BlockInteractionMax *int `json:"block_interaction_max,omitempty"`
  // Determines whether the mob can empty a full composter.
  CanEmptyComposter *bool `json:"can_empty_composter,omitempty"`
  // Determines whether the mob can add items to a composter given that it is not full.
  CanFillComposter *bool `json:"can_fill_composter,omitempty"`
  // If true, this entity can work when their jobsite POI is being rained on.
  CanWorkInRain bool `json:"can_work_in_rain,omitempty"`
  // The amount of ticks the goal will be on cooldown before it can be used again
  GoalCooldown int `json:"goal_cooldown,omitempty"`
  // The maximum number of items which can be added to the composter per block interaction.
  ItemsPerUseMax *int `json:"items_per_use_max,omitempty"`
  // Limits the amount of each compostable item the mob can use. Any amount held over this number will be composted if possible
  MinItemCount *int `json:"min_item_count,omitempty"`
  // [Trigger] Event to run when the mob reaches their jobsite.
  OnArrival interface{} `json:"on_arrival,omitempty"`
  // Unused.
  SoundDelayMax *int `json:"sound_delay_max,omitempty"`
  // Unused.
  SoundDelayMin *int `json:"sound_delay_min,omitempty"`
  // Movement speed multiplier of the mob when using this AI Goal
  SpeedMultiplier *float64 `json:"speed_multiplier,omitempty"`
  // The maximum interval in which the mob will interact with the composter.
  UseBlockMax *int `json:"use_block_max,omitempty"`
  // The minimum interval in which the mob will interact with the composter.
  UseBlockMin *int `json:"use_block_min,omitempty"`
  // If "can_work_in_rain" is false, this is the maximum number of ticks left in the goal where rain will not interrupt the goal
  WorkInRainTolerance *int `json:"work_in_rain_tolerance,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
