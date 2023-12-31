package component

// Code generated by scrapper/entityBp.ts;

// Allows this entity to damage a target by using a running attack.  
type Behavior_ChargeAttack struct {
  // A charge attack cannot start if the entity is farther than this distance to the target.
  MaxDistance *float64 `json:"max_distance,omitempty"`
  // A charge attack cannot start if the entity is closer than this distance to the target.
  MinDistance *float64 `json:"min_distance,omitempty"`
  // Modifies the entity's speed when charging toward the target.
  SpeedMultiplier *float64 `json:"speed_multiplier,omitempty"`
  // Percent chance this entity will start a charge attack, if not already attacking (1.0 = 100%)
  SuccessRate *float64 `json:"success_rate,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
