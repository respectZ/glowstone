package component

// Code generated by scrapper/entityBp.ts;

// [Not a component] List of events to send
type BehaviorSendEventSequence struct {
  // Amount of time in seconds before starting this step
  BaseDelay float64 `json:"base_delay,omitempty"`
  // The event to send to the entity
  Event string `json:"event,omitempty"`
  // The sound event to play when this step happens
  SoundEvent string `json:"sound_event,omitempty"`
}

// Allows the mob to send an event to another mob.  
type Behavior_SendEvent struct {
  // Time in seconds for the entire event sending process
  CastDuration *float64 `json:"cast_duration,omitempty"`
  // If true, the mob will face the entity it sends an event to
  LookAtTarget *bool `json:"look_at_target,omitempty"`
  // List of events to send
  Sequence []BehaviorSendEventSequence `json:"sequence,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
