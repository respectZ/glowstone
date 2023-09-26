package component

// Code generated by scrapper/entityBp.ts;

// Sets the Equipment table to use for this Entity.  
type Equipment struct {
  // A list of slots with the chance to drop an equipped item from that slot.
  SlotDropChance []interface{} `json:"slot_drop_chance,omitempty"`
  // The file path to the equipment table, relative to the behavior pack's root.
  Table string `json:"table,omitempty"`
}
