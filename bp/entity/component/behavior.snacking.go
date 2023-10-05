package component

// Code generated by scrapper/entityBp.ts;

// Allows the mob to take a load off and snack on food that it found nearby.  
type Behavior_Snacking struct {
  // Items that we are interested in snacking on
  Items []interface{} `json:"items,omitempty"`
  // The cooldown time in seconds before the mob is able to snack again
  SnackingCooldown *float64 `json:"snacking_cooldown,omitempty"`
  // The minimum time in seconds before the mob is able to snack again
  SnackingCooldownMin *float64 `json:"snacking_cooldown_min,omitempty"`
  // This is the chance that the mob will stop snacking, from 0 to 1
  SnackingStopChance *float64 `json:"snacking_stop_chance,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
