package component

// Code generated by scrapper/entityBp.ts;

// [Not a component] The list of positions and number of riders for each position for entities riding this entity
type RideableSeats struct {
  // Angle in degrees that a rider is allowed to rotate while riding this entity. Omit this property for no limit
  LockRiderRotation float64 `json:"lock_rider_rotation,omitempty"`
  // Defines the maximum number of riders that can be riding this entity for this seat to be valid
  MaxRiderCount int `json:"max_rider_count,omitempty"`
  // Defines the minimum number of riders that need to be riding this entity before this seat can be used
  MinRiderCount int `json:"min_rider_count,omitempty"`
  // [Vector3 [a,b,c]] Position of this seat relative to this entity's position
  Position []float64 `json:"position,omitempty"`
  // [Molang String] Offset to rotate riders by
  RotateRiderBy string `json:"rotate_rider_by,omitempty"`
}

// Determines whether this entity can be ridden. Allows specifying the different seat positions and quantity.  
type Rideable struct {
  // The seat that designates the driver of the entity. This is only observed by the horse/boat styles of riding; minecarts/entities with "minecraft:controlled_by_player" give control to any player in any seat.
  ControllingSeat int `json:"controlling_seat,omitempty"`
  // If true, this entity can't be interacted with if the entity interacting with it is crouching
  CrouchingSkipInteract bool `json:"crouching_skip_interact,omitempty"`
  // List of entities that can ride this entity
  FamilyTypes []interface{} `json:"family_types,omitempty"`
  // The text to display when the player can interact with the entity when playing with Touch-screen controls
  InteractText string `json:"interact_text,omitempty"`
  // The max width a mob can be to be a passenger. A value of 0 ignores this parameter.
  PassengerMaxWidth float64 `json:"passenger_max_width,omitempty"`
  // This field may exist in old data but isn't used by minecraft:rideable.
  Priority int `json:"priority,omitempty"`
  // If true, this entity will pull in entities that are in the correct family_types into any available seats
  PullInEntities bool `json:"pull_in_entities,omitempty"`
  // If true, this entity will be picked when looked at by the rider
  RiderCanInteract bool `json:"rider_can_interact,omitempty"`
  // The number of entities that can ride this entity at the same time
  SeatCount int `json:"seat_count,omitempty"`
  // The list of positions and number of riders for each position for entities riding this entity
  Seats []RideableSeats `json:"seats,omitempty"`
}
