package component

// Code generated by scrapper/entityBp.ts;

// Allows this entity to celebrate surviving a raid by making celebration sounds and jumping.  
type Behavior_Celebrate struct {
  // The sound event to trigger during the celebration.
  CelebrationSound string `json:"celebration_sound,omitempty"`
  // The duration in seconds that the celebration lasts for.
  Duration float64 `json:"duration,omitempty"`
  // Minimum and maximum time between jumping (positive, in seconds).
  JumpInterval float64 `json:"jump_interval,omitempty"`
  // [Trigger] The event to trigger when the goal's duration expires.
  OnCelebrationEndEvent interface{} `json:"on_celebration_end_event,omitempty"`
  // Minimum and maximum time between sound events (positive, in seconds).
  SoundInterval float64 `json:"sound_interval,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
