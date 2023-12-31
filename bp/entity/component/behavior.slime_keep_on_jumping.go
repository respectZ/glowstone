package component

// Code generated by scrapper/entityBp.ts;

// Allows the entity to continuously jump around like a slime.  
type Behavior_SlimeKeepOnJumping struct {
  // Determines the multiplier this entity's speed is modified by when jumping around.
  SpeedMultiplier *float64 `json:"speed_multiplier,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
