package component

// Code generated by scrapper/entityBp.ts;

import (
  f "github.com/respectZ/glowstone/bp/types"
)

// A component that applies a mob effect to entities that get within range.  
type MobEffect struct {
  // Time in seconds to wait between each application of the effect.
  CooldownTime int `json:"cooldown_time,omitempty"`
  // How close a hostile entity must be to have the mob effect applied.
  EffectRange *float64 `json:"effect_range,omitempty"`
  // How long the applied mob effect lasts in seconds.
  EffectTime *int `json:"effect_time,omitempty"`
  // The set of entities that are valid to apply the mob effect to.
  EntityFilter *f.Filter `json:"entity_filter,omitempty"`
  // The mob effect that is applied to entities that enter this entities effect range.
  MobEffect string `json:"mob_effect,omitempty"`
}
