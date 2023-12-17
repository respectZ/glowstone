package component

// Code generated by scrapper/item/main.ts;

// Modifies use effects, including how long an item takes to use and the player's speed when used in combination with components like Shooter, Throwable or Food.
type UseModifiers struct {
  // Modifier value to scale the players movement speed when item is in use.
  MovementModifier float64 `json:"movement_modifier,omitempty"`

  // How long the item takes to use in seconds.
  UseDuration float64 `json:"use_duration,omitempty"`

}


