package component

// Code generated by scrapper/entityBp.ts;

import (
  f "github.com/respectZ/glowstone/bp/types"
)

// Adds a trigger that will run when a nearby entity of the same type as this entity becomes Angry.  
type OnFriendlyAnger struct {
  // The event to run when the conditions for this trigger are met.
  Event string `json:"event,omitempty"`
  // The list of conditions for this trigger to execute.
  Filters *f.Filter `json:"filters,omitempty"`
  // The target of the event.
  Target *string `json:"target,omitempty"`
}
