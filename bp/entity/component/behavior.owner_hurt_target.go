package component

// Code generated by scrapper/entityBp.ts;

import (
  f "github.com/respectZ/glowstone/bp/types"
)

// [Not a component] List of entity types that this entity can target if the potential target is hurt by this mob's owner
type BehaviorOwnerHurtTargetEntityTypes struct {
  // The amount of time in seconds that the mob has to wait before selecting a target of the same type again
  Cooldown float64 `json:"cooldown,omitempty"`
  // Conditions that make this entry in the list valid
  Filters *f.Filter `json:"filters,omitempty"`
  // Maximum distance this mob can be away to be a valid choice
  MaxDist *float64 `json:"max_dist,omitempty"`
  // If true, the mob has to be visible to be a valid choice
  MustSee bool `json:"must_see,omitempty"`
  // Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
  MustSeeForgetDuration *float64 `json:"must_see_forget_duration,omitempty"`
  // If true, the mob will stop being targeted if it stops meeting any conditions.
  ReevaluateDescription bool `json:"reevaluate_description,omitempty"`
  // Multiplier for the running speed. A value of 1.0 means the speed is unchanged
  SprintSpeedMultiplier *float64 `json:"sprint_speed_multiplier,omitempty"`
  // Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
  WalkSpeedMultiplier *float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the mob to target a mob that is hurt by their owner.  
type Behavior_OwnerHurtTarget struct {
  // Corresponding Type: BehaviorOwnerHurtTargetEntityTypes

  EntityTypes interface{} `json:"entity_types,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
