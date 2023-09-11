package types

type Trigger struct {
	Event  string  `json:"event"`
	Target string  `json:"target,omitempty"`
	Filter *Filter `json:"filter,omitempty"`
}
