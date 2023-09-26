package component

// Code generated by scrapper/entityBp.ts;

// [Not a component] The list of items that, if carried while interacting with the entity, will anger it.
type TamemountAutoRejectItems struct {
  // Name of the item this entity dislikes and will cause it to get angry if used while untamed.
  Item string `json:"item,omitempty"`
}

// [Not a component] The list of items that can be used to increase the entity's temper and speed up the taming process.
type TamemountFeedItems struct {
  // Name of the item this entity likes and can be used to increase this entity's temper.
  Item string `json:"item,omitempty"`
  // The amount of temper this entity gains when fed this item.
  TemperMod int `json:"temper_mod,omitempty"`
}

// Allows the Entity to be tamed by mounting it.  
type Tamemount struct {
  // The amount the entity's temper will increase when mounted.
  AttemptTemperMod int `json:"attempt_temper_mod,omitempty"`
  // The list of items that, if carried while interacting with the entity, will anger it.
  AutoRejectItems TamemountAutoRejectItems `json:"autoRejectItems,omitempty"`
  // The list of items that can be used to increase the entity's temper and speed up the taming process.
  FeedItems TamemountFeedItems `json:"feed_items,omitempty"`
  // The text that shows in the feeding interact button.
  FeedText string `json:"feed_text,omitempty"`
  // The maximum value for the entity's random starting temper.
  MaxTemper int `json:"max_temper,omitempty"`
  // The minimum value for the entity's random starting temper.
  MinTemper int `json:"min_temper,omitempty"`
  // The text that shows in the riding interact button.
  RideText string `json:"ride_text,omitempty"`
  // Event that triggers when the entity becomes tamed.
  TameEvent string `json:"tame_event,omitempty"`
}