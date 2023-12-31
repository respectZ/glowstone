package component

// Code generated by scrapper/entityBp.ts;

// Can only be used by Villagers. Allows the villagers to create paths around the village.  
type Behavior_MoveThroughVillage struct {
  // If true, the mob will only move through the village during night time
  OnlyAtNight bool `json:"only_at_night,omitempty"`
  // Movement speed multiplier of the mob when using this AI Goal
  SpeedMultiplier *float64 `json:"speed_multiplier,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
