package component

// Code generated by scrapper/entityBp.ts;

// Allows the mob to look at the player they are trading with.  
type Behavior_LookAtTradingPlayer struct {
  // The angle in degrees that the mob can see in the Y-axis (up-down)
  AngleOfViewHorizontal *int `json:"angle_of_view_horizontal,omitempty"`
  // The angle in degrees that the mob can see in the X-axis (left-right)
  AngleOfViewVertical *int `json:"angle_of_view_vertical,omitempty"`
  // The distance in blocks from which the entity will look at
  LookDistance *float64 `json:"look_distance,omitempty"`
  // Time range to look at the entity
  LookTime *float64 `json:"look_time,omitempty"`
  // The probability of looking at the target. A value of 1.00 is 100%
  Probability *float64 `json:"probability,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
