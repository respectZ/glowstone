package component

// Code generated by scrapper/entityBp.ts;

// Could increase crop growth when entity walks over crop  
type GrowsCrop struct {
  // Value between 0-1. Chance of success per tick.
  Chance float64 `json:"chance,omitempty"`
  // Number of charges
  Charges int `json:"charges,omitempty"`
}