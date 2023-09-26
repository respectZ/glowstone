package component

// Code generated by scrapper/entityBp.ts;

// Allows the mob to lay an egg block on certain types of blocks if the mob is pregnant.  
type Behavior_LayEgg struct {
  // Allows the mob to lay its eggs from below the target if it can't get there. This is useful if the target block is water with air above, since mobs may not be able to get to the air block above water.
  AllowLayingFromBelow bool `json:"allow_laying_from_below,omitempty"`
  // Block type for the egg to lay. If this is a turtle egg, the number of eggs in the block is randomly set.
  EggType string `json:"egg_type,omitempty"`
  // Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
  GoalRadius float64 `json:"goal_radius,omitempty"`
  // Name of the sound event played when laying the egg. Defaults to lay_egg, which is used for Turtles.
  LayEggSound string `json:"lay_egg_sound,omitempty"`
  // Duration of the laying egg process in seconds.
  LaySeconds float64 `json:"lay_seconds,omitempty"`
  // [Trigger] Event to run when this mob lays the egg.
  OnLay interface{} `json:"on_lay,omitempty"`
  // Height in blocks the mob will look for a target block to move towards
  SearchHeight int `json:"search_height,omitempty"`
  // The distance in blocks it will look for a target block to move towards
  SearchRange int `json:"search_range,omitempty"`
  // Movement speed multiplier of the mob when using this AI Goal
  SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
  // Blocks that the mob can lay its eggs on top of.
  TargetBlocks []interface{} `json:"target_blocks,omitempty"`
  // Types of materials that can exist above the target block. Valid types are Air, Water, and Lava.
  TargetMaterialsAboveBlock []interface{} `json:"target_materials_above_block,omitempty"`
  // Specifies if the default lay-egg animation should be played when the egg is placed or not.
  UseDefaultAnimation bool `json:"use_default_animation,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}