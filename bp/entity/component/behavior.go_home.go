package component

// Code generated by scrapper/entityBp.ts;

// Allows the mob to move back to the position they were spawned.  
type Behavior_GoHome struct {
  // Distance in blocks that the mob is considered close enough to the end of the current path. A new path will then be calculated to continue toward home.
  CalculateNewPathRadius float64 `json:"calculate_new_path_radius,omitempty"`
  // Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot.
  GoalRadius float64 `json:"goal_radius,omitempty"`
  // A random value to determine when to randomly move somewhere. This has a 1/interval chance to choose this goal.
  Interval int `json:"interval,omitempty"`
  // [Trigger] Event(s) to run when this goal fails.
  OnFailed interface{} `json:"on_failed,omitempty"`
  // [Trigger] Event(s) to run when this mob gets home.
  OnHome interface{} `json:"on_home,omitempty"`
  // Movement speed multiplier of the mob when using this AI Goal.
  SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
