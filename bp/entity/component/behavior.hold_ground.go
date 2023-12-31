package component

// Code generated by scrapper/entityBp.ts;

// The mob freezes and looks at the mob they are targeting.  
type Behavior_HoldGround struct {
  // Whether to broadcast out the mob's target to other mobs of the same type.
  Broadcast bool `json:"broadcast,omitempty"`
  // Range in blocks for how far to broadcast.
  BroadcastRange *float64 `json:"broadcast_range,omitempty"`
  // Minimum distance the target must be for the mob to run this goal.
  MinRadius *float64 `json:"min_radius,omitempty"`
  // Event to run when target is within the radius. This event is broadcasted if broadcast is true.
  WithinRadiusEvent string `json:"within_radius_event,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
