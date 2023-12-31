package component

// Code generated by scrapper/entityBp.ts;

import (
  f "github.com/respectZ/glowstone/bp/types"
)

// Only usable by the Ender Dragon. Adds a trigger to call on this entity's death.  
type OnDeath struct {
  // The event to run when the conditions for this trigger are met.
  Event string `json:"event,omitempty"`
  // The list of conditions for this trigger to execute.
  Filters *f.Filter `json:"filters,omitempty"`
  // The target of the event.
  Target *string `json:"target,omitempty"`
}
