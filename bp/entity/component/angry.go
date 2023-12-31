package component

// Code generated by scrapper/entityBp.ts;

import (
  f "github.com/respectZ/glowstone/bp/types"
)

// Defines the entity's 'angry' state using a timer.  
type Angry struct {
  // The sound event to play when the mob is angry
  AngrySound string `json:"angry_sound,omitempty"`
  // If true, other entities of the same entity definition within the broadcastRange will also become angry
  BroadcastAnger bool `json:"broadcast_anger,omitempty"`
  // If true, other entities of the same entity definition within the broadcastRange will also become angry whenever this mob attacks
  BroadcastAngerOnAttack bool `json:"broadcast_anger_on_attack,omitempty"`
  // If true, other entities of the same entity definition within the broadcastRange will also become angry whenever this mob is attacked
  BroadcastAngerOnBeingAttacked bool `json:"broadcast_anger_on_being_attacked,omitempty"`
  // Conditions that make this entry in the list valid
  BroadcastFilters *f.Filter `json:"broadcast_filters,omitempty"`
  // Distance in blocks within which other entities of the same entity definition will become angry
  BroadcastRange *int `json:"broadcast_range,omitempty"`
  // A list of entity families to broadcast anger to
  BroadcastTargets []interface{} `json:"broadcast_targets,omitempty"`
  // Event to run after the number of seconds specified in duration expires (when the entity stops being 'angry')
  CalmEvent string `json:"calm_event,omitempty"`
  // The amount of time in seconds that the entity will be angry
  Duration *int `json:"duration,omitempty"`
  // Variance in seconds added to the duration [-delta, delta]
  DurationDelta int `json:"duration_delta,omitempty"`
  // Filter out mob types that it should not attack while angry (other Piglins)
  Filters *f.Filter `json:"filters,omitempty"`
  // The range of time in seconds to randomly wait before playing the sound again
  SoundInterval float64 `json:"sound_interval,omitempty"`
}
