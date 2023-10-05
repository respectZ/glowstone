package component

// Code generated by scrapper/entityBp.ts;

// Allows the mob to randomly look around.  
type Behavior_RandomLookAround struct {
  // The range of time in seconds the mob will stay looking in a random direction before looking elsewhere
  LookTime *float64 `json:"look_time,omitempty"`
  // The rightmost angle a mob can look at on the horizontal plane with respect to its initial facing direction.
  MaxAngleOfViewHorizontal *int `json:"max_angle_of_view_horizontal,omitempty"`
  // The leftmost angle a mob can look at on the horizontal plane with respect to its initial facing direction.
  MinAngleOfViewHorizontal *int `json:"min_angle_of_view_horizontal,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
