package component

// Code generated by scrapper/entityBp.ts;

// Allows the actor to break doors assuming that that flags set up for the component to use in navigation  
type Annotation_BreakDoor struct {
  // The time in seconds required to break through doors.
  BreakTime float64 `json:"break_time,omitempty"`
  // The minimum difficulty that the world must be on for this entity to break doors.
  MinDifficulty string `json:"min_difficulty,omitempty"`
}