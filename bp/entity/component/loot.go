package component

// Code generated by scrapper/entityBp.ts;

// Sets the loot table for what items this entity drops upon death.  
type Loot struct {
  // The path to the loot table, relative to the Behavior Pack's root.
  Table string `json:"table,omitempty"`
}
