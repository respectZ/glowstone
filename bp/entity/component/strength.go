package component

// Code generated by scrapper/entityBp.ts;

// Defines the entity's strength to carry items.  
type Strength struct {
  // The maximum strength of this entity
  Max int `json:"max,omitempty"`
  // The initial value of the strength
  Value int `json:"value,omitempty"`
}