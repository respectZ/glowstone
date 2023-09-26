package component

// Code generated by scrapper/entityBp.ts;

// If the mob is carrying a food item, the mob will eat it and the effects will be applied to the mob.  
type Behavior_EatCarriedItem struct {
  // Time in seconds the mob should wait before eating the item.
  DelayBeforeEating float64 `json:"delay_before_eating,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}