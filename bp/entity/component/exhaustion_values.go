package component

// Code generated by scrapper/entityBp.ts;

// Defines how much exhaustion each player action should take.  
type ExhaustionValues struct {
  // Amount of exhaustion applied when attacking.
  Attack float64 `json:"attack,omitempty"`
  // Amount of exhaustion applied when taking damage.
  Damage float64 `json:"damage,omitempty"`
  // Amount of exhaustion applied when healed through food regeneration.
  Heal float64 `json:"heal,omitempty"`
  // Amount of exhaustion applied when jumping.
  Jump float64 `json:"jump,omitempty"`
  // Amount of exhaustion applied when mining.
  Mine float64 `json:"mine,omitempty"`
  // Amount of exhaustion applied when sprinting.
  Sprint float64 `json:"sprint,omitempty"`
  // Amount of exhaustion applied when sprint jumping.
  SprintJump float64 `json:"sprint_jump,omitempty"`
  // Amount of exhaustion applied when swimming.
  Swim float64 `json:"swim,omitempty"`
  // Amount of exhaustion applied when walking.
  Walk float64 `json:"walk,omitempty"`
}