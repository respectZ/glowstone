package component

// Code generated by scrapper/item/main.ts;

// Determines what enchantments can be applied to the item. Not all enchantments will have an effect on all item components.
type Enchantable struct {
  // What enchantments can be applied (ex. Using bow would allow this item to be enchanted as if it were a bow).
  Slot string `json:"slot,omitempty"`

  // The value of the enchantment (minimum of 0).
  Value interface{} `json:"value,omitempty"`

}


