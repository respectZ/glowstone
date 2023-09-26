package component

// Code generated by scrapper/entityBp.ts;

// Allows this entity to attack a player by charging at them. The player is chosen by the "minecraft:behavior.dragonscanning". Can only be used by the Ender Dragon.  
type Behavior_Dragonchargeplayer struct {
  // The speed this entity moves when this behavior has started or while it's active.
  ActiveSpeed float64 `json:"active_speed,omitempty"`
  // If the dragon is outside the "target_zone" for longer than "continue_charge_threshold_time" seconds, the charge is canceled.
  ContinueChargeThresholdTime float64 `json:"continue_charge_threshold_time,omitempty"`
  // The speed this entity moves while this behavior is not active.
  FlightSpeed float64 `json:"flight_speed,omitempty"`
  // Minimum and maximum distance, from the target, this entity can use this behavior.
  TargetZone float64 `json:"target_zone,omitempty"`
  // The speed at which this entity turns while using this behavior.
  TurnSpeed float64 `json:"turn_speed,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}