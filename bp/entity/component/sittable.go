package component

// Code generated by scrapper/entityBp.ts;

// Defines the entity's 'sit' state.  
type Sittable struct {
  // Event to run when the entity enters the 'sit' state
  SitEvent string `json:"sit_event,omitempty"`
  // Event to run when the entity exits the 'sit' state
  StandEvent string `json:"stand_event,omitempty"`
}
