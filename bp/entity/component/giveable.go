package component

// Code generated by scrapper/entityBp.ts;

// Defines sets of items that can be used to trigger events when used on this entity. The item will also be taken and placed in the entity's inventory.  
type Giveable struct {
  // An optional cool down in seconds to prevent spamming interactions.
  Cooldown float64 `json:"cooldown,omitempty"`
  // The list of items that can be given to the entity to place in their inventory.
  Items []interface{} `json:"items,omitempty"`
  // Event to fire when the correct item is given.
  OnGive string `json:"on_give,omitempty"`
}
