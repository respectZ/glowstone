package component

// Code generated by scrapper/entityBp.ts;

// Allow slimes to float in water / lava. Can only be used by Slime and Magma Cubes.  
type Behavior_SlimeFloat struct {
  // Percent chance a slime or magma cube has to jump while in water / lava.
  JumpChancePercentage float64 `json:"jump_chance_percentage,omitempty"`
  // Determines the multiplier the entity's speed is modified by when moving through water / lava.
  SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}