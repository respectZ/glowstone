package component

// Code generated by scrapper/entityBp.ts;

// Defines the rules for a mob to trust players.  
type Trusting struct {
  // The chance of the entity trusting with each item use between 0.0 and 1.0, where 1.0 is 100%.
  Probability float64 `json:"probability,omitempty"`
  // Event to run when this entity becomes trusting.
  TrustEvent string `json:"trust_event,omitempty"`
  // The list of items that can be used to get the entity to trust players.
  TrustItems []interface{} `json:"trust_items,omitempty"`
}
