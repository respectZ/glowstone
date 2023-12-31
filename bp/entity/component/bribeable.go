package component

// Code generated by scrapper/entityBp.ts;

// Defines the way an entity can get into the 'bribed' state.  
type Bribeable struct {
  // Time in seconds before the Entity can be bribed again.
  BribeCooldown *float64 `json:"bribe_cooldown,omitempty"`
  // The list of items that can be used to bribe the entity.
  BribeItems []interface{} `json:"bribe_items,omitempty"`
}
