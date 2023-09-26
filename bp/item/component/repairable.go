package component

// Code generated by scrapper/item/main.ts;

// Repairable item component: Determines which items can be used to repair a defined item, as well as the amount of durability specified items will repair. In format versions prior to 1.20.10, this component requires the 'Holiday Creator Features' experimental toggle.
type Repairable struct {
  // List of repair item entries.
  RepairItems []interface{} `json:"repair_items,omitempty"`

}


