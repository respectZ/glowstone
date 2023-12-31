package component

// Code generated by scrapper/entityBp.ts;

// Adds a timer for drying out that will count down and fire 'dried_out_event' or will stop as soon as the entity will get under rain or water and fire 'stopped_drying_out_event'  
type DryingOutTimer struct {
  // Event to fire when the drying out time runs out.
  DriedOutEvent string `json:"dried_out_event,omitempty"`
  // Event to fire when entity was already dried out but received increase in water supply.
  RecoverAfterDriedOutEvent string `json:"recover_after_dried_out_event,omitempty"`
  // Event to fire when entity stopped drying out, for example got into water or under rain.
  StoppedDryingOutEvent string `json:"stopped_drying_out_event,omitempty"`
  // Amount of time in seconds to dry out fully.
  TotalTime float64 `json:"total_time,omitempty"`
  // Optional amount of additional time in seconds given by using splash water bottle on entity.
  WaterBottleRefillTime float64 `json:"water_bottle_refill_time,omitempty"`
}
