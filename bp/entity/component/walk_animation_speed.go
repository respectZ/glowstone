package component

// Code generated by scrapper/entityBp.ts;

// Sets the speed multiplier for this entity's walk animation speed.  
type WalkAnimationSpeed struct {
  // The higher the number, the faster the animation for walking plays. A value of 1.0 means normal speed, while 2.0 means twice as fast.
  Value float64 `json:"value,omitempty"`
}