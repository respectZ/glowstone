package component

// Code generated by scrapper/entityBp.ts;

// Move control that causes the mob to jump as it moves with a specified delay between jumps.  
type Movement_Jump struct {
  // Delay after landing when using the slime move control.
  JumpDelay *float64 `json:"jump_delay,omitempty"`
  // The maximum number in degrees the mob can turn per tick.
  MaxTurn *float64 `json:"max_turn,omitempty"`
}
