package component

// Code generated by scrapper/entityBp.ts;

// Allows the entity to first travel to a random point on the outskirts of the village, and then explore random points within a small distance. This goal requires "minecraft:dweller" and "minecraft:navigation" to execute.  
type Behavior_ExploreOutskirts struct {
  // [Vector3 [a,b,c]] The distance from the boundary the villager must be within in to explore the outskirts.
  DistFromBoundary []float64 `json:"dist_from_boundary,omitempty"`
  // Total distance in blocks the the entity will explore beyond the village bounds when choosing its travel point.
  ExploreDist float64 `json:"explore_dist,omitempty"`
  // This is the maximum amount of time an entity will attempt to reach it's travel point on the outskirts of the village before the goal exits.
  MaxTravelTime float64 `json:"max_travel_time,omitempty"`
  // The wait time in seconds between choosing new explore points will be chosen on a random interval between this value and the minimum wait time. This value is also the total amount of time the entity will explore random points before the goal stops.
  MaxWaitTime float64 `json:"max_wait_time,omitempty"`
  // The entity must be within this distance for it to consider it has successfully reached its target.
  MinDistFromTarget float64 `json:"min_dist_from_target,omitempty"`
  // The minimum perimeter of the village required to run this goal.
  MinPerimeter float64 `json:"min_perimeter,omitempty"`
  // The wait time in seconds between choosing new explore points will be chosen on a random interval between this value and the maximum wait time.
  MinWaitTime float64 `json:"min_wait_time,omitempty"`
  // A new explore point will randomly be chosen within this XZ distance of the current target position when navigation has finished and the wait timer has elapsed.
  NextXz int `json:"next_xz,omitempty"`
  // A new explore point will randomly be chosen within this Y distance of the current target position when navigation has finished and the wait timer has elapsed.
  NextY int `json:"next_y,omitempty"`
  // The multiplier for speed while using this goal. 1.0 maintains the speed.
  SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
  // Each new explore point will be chosen on a random interval between the minimum and the maximum wait time, divided by this value. This does not apply to the first explore point chosen when the goal runs.
  TimerRatio float64 `json:"timer_ratio,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}