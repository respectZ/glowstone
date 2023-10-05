package component

// Code generated by scrapper/entityBp.ts;

// Allows this entity to be leashed and defines the conditions and events for this entity when is leashed.  
type Leashable struct {
  // If true, players can leash this entity even if it is already leashed to another mob.
  CanBeStolen bool `json:"can_be_stolen,omitempty"`
  // Distance in blocks at which the leash stiffens, restricting movement.
  HardDistance *float64 `json:"hard_distance,omitempty"`
  // Distance in blocks at which the leash breaks.
  MaxDistance *float64 `json:"max_distance,omitempty"`
  // Event to call when this entity is leashed.
  OnLeash string `json:"on_leash,omitempty"`
  // Event to call when this entity is unleashed.
  OnUnleash string `json:"on_unleash,omitempty"`
  // Distance in blocks at which the 'spring' effect starts acting to keep this entity close to the entity that leashed it.
  SoftDistance *float64 `json:"soft_distance,omitempty"`
}
