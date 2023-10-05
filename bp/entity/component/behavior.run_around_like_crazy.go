package component

// Code generated by scrapper/entityBp.ts;

// Allows the mob to run around aimlessly.  
type Behavior_RunAroundLikeCrazy struct {
  // Movement speed multiplier of the mob when using this AI Goal
  SpeedMultiplier *float64 `json:"speed_multiplier,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
