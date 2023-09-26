package component

// Code generated by scrapper/entityBp.ts;

import (
  f "github.com/respectZ/glowstone/bp/types"
)

// Defines the behavior when another entity looks at this entity.  
type Lookat struct {
  // If true, invulnerable entities (e.g. Players in creative mode) are considered valid targets.
  AllowInvulnerable bool `json:"allow_invulnerable,omitempty"`
  // Defines the entities that can trigger this component.
  Filters *f.Filter `json:"filters,omitempty"`
  // The range for the random amount of time during which the entity is 'cooling down' and won't get angered or look for a target.
  LookCooldown float64 `json:"look_cooldown,omitempty"`
  // The event identifier to run when the entities specified in filters look at this entity.
  LookEvent string `json:"look_event,omitempty"`
  // Maximum distance this entity will look for another entity looking at it.
  SearchRadius float64 `json:"search_radius,omitempty"`
  // If true, this entity will set the attack target as the entity that looked at it.
  SetTarget bool `json:"set_target,omitempty"`
}