package component

// Code generated by scrapper/item/main.ts;

// When an item has a food component, it becomes edible to the player. Must have the minecraft:use_modifiers component in order to function properly.
type Food struct {
  // If true you can always eat this item (even when not hungry).
  CanAlwaysEat bool `json:"can_always_eat,omitempty"`

  // List of Events to fire off when consumed
  Effects []interface{} `json:"effects,omitempty"`

  // The value that is added to the actor's nutrition when the item is used.
  Nutrition int `json:"nutrition,omitempty"`

  // Saturation Modifier is used in this formula: (nutrition * saturation_modifier * 2) when applying the saturation buff.
  SaturationModifier float64 `json:"saturation_modifier,omitempty"`

  // When used, converts to the item specified by the string in this field.
  UsingConvertsTo interface{} `json:"using_converts_to,omitempty"`

}


