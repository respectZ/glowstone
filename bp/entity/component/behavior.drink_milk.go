package component

// Code generated by scrapper/entityBp.ts;

import (
  f "github.com/respectZ/glowstone/bp/types"
)

// Allows the mob to drink milk based on specified environment conditions.  
type Behavior_DrinkMilk struct {
  // Time (in seconds) that the goal is on cooldown before it can be used again.
  CooldownSeconds *float64 `json:"cooldown_seconds,omitempty"`
  // Conditions that need to be met for the behavior to start.
  Filters *f.Filter `json:"filters,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
