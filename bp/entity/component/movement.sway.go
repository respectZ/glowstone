package component

// Code generated by scrapper/entityBp.ts;

// This move control causes the mob to sway side to side giving the impression it is swimming.  
type Movement_Sway struct {
  // The maximum number in degrees the mob can turn per tick.
  MaxTurn *float64 `json:"max_turn,omitempty"`
  // Strength of the sway movement.
  SwayAmplitude *float64 `json:"sway_amplitude,omitempty"`
  // Multiplier for the frequency of the sway movement.
  SwayFrequency *float64 `json:"sway_frequency,omitempty"`
}
