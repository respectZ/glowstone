package component

// Code generated by scrapper/entityBp.ts;

// Allows an entity to charge and use their held item.  
type Behavior_ChargeHeldItem struct {
  // The list of items that can be used to charge the held item. This list is required and must have at least one item in it.
  Items []interface{} `json:"items,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}