package component

// Code generated by scrapper/entityBp.ts;

// Allows the entity to consume a block, replace the eaten block with another block, and trigger an event as a result.  
type Behavior_EatBlock struct {
  // A collection of pairs of blocks; the first ("eat_block")is the block the entity should eat, the second ("replace_block") is the block that should replace the eaten block.
  EatAndReplaceBlockPairs []interface{} `json:"eat_and_replace_block_pairs,omitempty"`
  // [Trigger] The event to trigger when the block eating animation has completed.
  OnEat *interface{} `json:"on_eat,omitempty"`
  // [Molang String] A molang expression defining the success chance the entity has to consume a block.
  SuccessChance *string `json:"success_chance,omitempty"`
  // The amount of time (in seconds) it takes for the block to be eaten upon a successful eat attempt.
  TimeUntilEat *float64 `json:"time_until_eat,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
