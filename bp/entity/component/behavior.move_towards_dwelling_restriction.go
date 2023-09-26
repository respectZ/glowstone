package component

// Code generated by scrapper/entityBp.ts;

// Allows mobs with the dweller component to move toward their Village area that the mob should be restricted to.  
type Behavior_MoveTowardsDwellingRestriction struct {
  // This multiplier modifies the entity's speed when moving towards it's restriction.
  SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
