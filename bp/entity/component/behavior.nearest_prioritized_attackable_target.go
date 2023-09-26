package component

// Code generated by scrapper/entityBp.ts;

import (
  f "github.com/respectZ/glowstone/bp/types"
)

// [Not a component] List of entity types that this mob considers valid targets
type BehaviorNearestPrioritizedAttackableTargetEntityTypes struct {
  // The amount of time in seconds that the mob has to wait before selecting a target of the same type again
  Cooldown float64 `json:"cooldown,omitempty"`
  // Conditions that make this entry in the list valid
  Filters *f.Filter `json:"filters,omitempty"`
  // Maximum distance this mob can be away to be a valid choice
  MaxDist float64 `json:"max_dist,omitempty"`
  // If true, the mob has to be visible to be a valid choice
  MustSee bool `json:"must_see,omitempty"`
  // Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
  MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`
  // If true, the mob will stop being targeted if it stops meeting any conditions.
  ReevaluateDescription bool `json:"reevaluate_description,omitempty"`
  // Multiplier for the running speed. A value of 1.0 means the speed is unchanged
  SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`
  // Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
  WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the mob to check for and pursue the nearest valid target.  
type Behavior_NearestPrioritizedAttackableTarget struct {
  // Time in seconds before selecting a target
  AttackInterval int `json:"attack_interval,omitempty"`
  // The amount of time in seconds that the mob has to wait before selecting a target of the same type again
  Cooldown float64 `json:"cooldown,omitempty"`
  // Corresponding Type: BehaviorNearestPrioritizedAttackableTargetEntityTypes

  EntityTypes interface{} `json:"entity_types,omitempty"`
  // If true, only entities that this mob can path to can be selected as targets
  MustReach bool `json:"must_reach,omitempty"`
  // If true, only entities in this mob's viewing range can be selected as targets
  MustSee bool `json:"must_see,omitempty"`
  // Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
  MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`
  // Time in seconds for a valid target to stay targeted when it becomes and invalid target.
  PersistTime float64 `json:"persist_time,omitempty"`
  // Specifies the priority in which filtered enemy types should be attacked. Lower number means higher priority.
  Priority int `json:"priority,omitempty"`
  // If true, the target will change to the current closest entity whenever a different entity is closer
  ReselectTargets bool `json:"reselect_targets,omitempty"`
  // How many ticks to wait between scanning for a target.
  ScanInterval int `json:"scan_interval,omitempty"`
  // Allows the actor to be set to persist upon targeting a player
  SetPersistent bool `json:"set_persistent,omitempty"`
  // Height in blocks to search for a target mob. -1.0f means the height does not matter.
  TargetSearchHeight float64 `json:"target_search_height,omitempty"`
  // Distance in blocks that the target can be within to launch an attack
  WithinRadius float64 `json:"within_radius,omitempty"`
}
