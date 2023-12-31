package component

// Code generated by scrapper/entityBp.ts;

// Allows the entity to be controlled by the player using an item in the item_controllable property (required). Also requires the minecraft:movement property, and the minecraft:rideable property. On every tick, the entity will attempt to rotate towards where the player is facing with the control item whilst simultaneously moving forward.   
type Behavior_ControlledByPlayer struct {
  // The entity will attempt to rotate to face where the player is facing each tick. The entity will target this percentage of their difference in their current facing angles each tick (from 0.0 to 1.0 where 1.0 = 100%). This is limited by FractionalRotationLimit. A value of 0.0 will result in the entity no longer turning to where the player is facing.
  FractionalRotation *float64 `json:"fractional_rotation,omitempty"`
  // Limits the total degrees the entity can rotate to face where the player is facing on each tick.
  FractionalRotationLimit *float64 `json:"fractional_rotation_limit,omitempty"`
  // Speed multiplier of mount when controlled by player.
  MountSpeedMultiplier *float64 `json:"mount_speed_multiplier,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
