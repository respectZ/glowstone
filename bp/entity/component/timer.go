package component

// Code generated by scrapper/entityBp.ts;

// Adds a timer after which an event will fire.  
type Timer struct {
  // If true, the timer will restart every time after it fires.
  Looping *bool `json:"looping,omitempty"`
  // If true, the amount of time on the timer will be random between the min and max values specified in time.
  RandomInterval *bool `json:"randomInterval,omitempty"`
  // This is a list of objects, representing one value in seconds that can be picked before firing the event and an optional weight. Incompatible with time.
  RandomTimeChoices []interface{} `json:"random_time_choices,omitempty"`
  // Amount of time in seconds for the timer. Can be specified as a number or a pair of numbers (min and max). Incompatible with random_time_choices.
  Time *float64 `json:"time,omitempty"`
  // Event to fire when the time on the timer runs out.
  TimeDownEvent string `json:"time_down_event,omitempty"`
}
