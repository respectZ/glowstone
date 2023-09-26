package component

// Code generated by scrapper/entityBp.ts;

// Allows Equine mobs to be Horse Traps and be triggered like them, spawning a lightning bolt and a bunch of horses when a player is nearby. Can only be used by Horses, Mules, Donkeys and Skeleton Horses.  
type Behavior_SkeletonHorseTrap struct {
  // Amount of time in seconds the trap exists. After this amount of time is elapsed, the trap is removed from the world if it hasn't been activated
  Duration float64 `json:"duration,omitempty"`
  // Distance in blocks that the player has to be within to trigger the horse trap
  WithinRadius float64 `json:"within_radius,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}