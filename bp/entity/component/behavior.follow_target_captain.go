package component

// Code generated by scrapper/entityBp.ts;

// Allows mob to move towards its current target captain.  
type Behavior_FollowTargetCaptain struct {
  // Defines the distance in blocks the mob will stay from its target while following.
  FollowDistance float64 `json:"follow_distance,omitempty"`
  // Defines the maximum distance in blocks a mob can get from its target captain before giving up trying to follow it.
  WithinRadius float64 `json:"within_radius,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
