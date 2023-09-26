package component

// Code generated by scrapper/entityBp.ts;

// Adds a cooldown to a mob. The intention of this cooldown is to be used to prevent the mob from attempting to aquire new attack targets.  
type AttackCooldown struct {
  // [Trigger] Event to be runned when the cooldown is complete.
  AttackCooldownCompleteEvent interface{} `json:"attack_cooldown_complete_event,omitempty"`
  // Amount of time in seconds for the cooldown. Can be specified as a number or a pair of numbers (min and max).
  AttackCooldownTime float64 `json:"attack_cooldown_time,omitempty"`
}
