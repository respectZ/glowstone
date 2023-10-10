package component

// Code generated by scrapper/entityBp.ts;

// Allows the creeper to swell up when a player is nearby. It can only be used by Creepers.  
type Behavior_Swell struct {
  // This mob starts swelling when a target is at least this many blocks away
  StartDistance *float64 `json:"start_distance,omitempty"`
  // This mob stops swelling when a target has moved away at least this many blocks
  StopDistance *float64 `json:"stop_distance,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
