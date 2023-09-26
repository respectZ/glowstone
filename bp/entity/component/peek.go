package component

// Code generated by scrapper/entityBp.ts;

// Defines the entity's 'peek' behavior, defining the events that should be called during it.  
type Peek struct {
  // Event to call when the entity is done peeking.
  OnClose string `json:"on_close,omitempty"`
  // Event to call when the entity starts peeking.
  OnOpen string `json:"on_open,omitempty"`
  // Event to call when the entity's target entity starts peeking.
  OnTargetOpen string `json:"on_target_open,omitempty"`
}