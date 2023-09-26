package component

// Code generated by scrapper/entityBp.ts;

// Allows an entity to deal damage through a melee attack.  
type Behavior_MeleeAttack struct {
  // Allows the entity to use this attack behavior, only once EVER.
  AttackOnce bool `json:"attack_once,omitempty"`
  // Defines the entity types this entity will attack.
  AttackTypes string `json:"attack_types,omitempty"`
  // If the entity is on fire, this allows the entity's target to catch on fire after being hit.
  CanSpreadOnFire bool `json:"can_spread_on_fire,omitempty"`
  // Cooldown time (in seconds) between attacks.
  CooldownTime float64 `json:"cooldown_time,omitempty"`
  // Time (in seconds) to add to attack path recalculation when the target is beyond the "path_inner_boundary".
  InnerBoundaryTimeIncrease float64 `json:"inner_boundary_time_increase,omitempty"`
  // Unused. No effect on "minecraft:behavior.melee_attack".
  MaxDist float64 `json:"max_dist,omitempty"`
  // Maximum base time (in seconds) to recalculate new attack path to target (before increases applied).
  MaxPathTime float64 `json:"max_path_time,omitempty"`
  // Field of view (in degrees) when using the sensing component to detect an attack target.
  MeleeFov float64 `json:"melee_fov,omitempty"`
  // Minimum base time (in seconds) to recalculate new attack path to target (before increases applied).
  MinPathTime float64 `json:"min_path_time,omitempty"`
  // [Trigger] Defines the event to trigger when this entity successfully attacks.
  OnAttack interface{} `json:"on_attack,omitempty"`
  // Time (in seconds) to add to attack path recalculation when the target is beyond the "path_outer_boundary".
  OuterBoundaryTimeIncrease float64 `json:"outer_boundary_time_increase,omitempty"`
  // Time (in seconds) to add to attack path recalculation when this entity cannot move along the current path.
  PathFailTimeIncrease float64 `json:"path_fail_time_increase,omitempty"`
  // Distance at which to increase attack path recalculation by "inner_boundary_tick_increase".
  PathInnerBoundary float64 `json:"path_inner_boundary,omitempty"`
  // Distance at which to increase attack path recalculation by "outer_boundary_tick_increase".
  PathOuterBoundary float64 `json:"path_outer_boundary,omitempty"`
  // This entity will have a 1 in N chance to stop it's current attack, where N = "random_stop_interval".
  RandomStopInterval int `json:"random_stop_interval,omitempty"`
  // Used with the base size of the entity to determine minimum target-distance before trying to deal attack damage.
  ReachMultiplier float64 `json:"reach_multiplier,omitempty"`
  // Toggles (on/off) the need to have a full path from the entity to the target when using this melee attack behavior.
  RequireCompletePath bool `json:"require_complete_path,omitempty"`
  // Allows the actor to be set to persist upon targeting a player
  SetPersistent bool `json:"set_persistent,omitempty"`
  // This multiplier modifies the attacking entity's speed when moving toward the target.
  SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
  // Unused. No effect on "minecraft:behavior.melee_attack".
  TargetDist float64 `json:"target_dist,omitempty"`
  // Allows the entity to track the attack target, even if the entity has no sensing.
  TrackTarget bool `json:"track_target,omitempty"`
  // Maximum rotation (in degrees), on the X-axis, this entity can rotate while trying to look at the target.
  XMaxRotation float64 `json:"x_max_rotation,omitempty"`
  // Maximum rotation (in degrees), on the Y-axis, this entity can rotate its head while trying to look at the target.
  YMaxHeadRotation float64 `json:"y_max_head_rotation,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
