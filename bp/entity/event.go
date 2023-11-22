package entity

import (
	jsoniter "github.com/json-iterator/go"
	f "github.com/respectZ/glowstone/bp/types"
	"github.com/respectZ/glowstone/types"
)

type EntityEvent struct {
	Add         *ComponentGroup `json:"add,omitempty"`
	Remove      *ComponentGroup `json:"remove,omitempty"`
	SetProperty ISetProperty    `json:"set_property,omitempty"`
	Trigger     interface{}     `json:"trigger,omitempty"` // TODO: implement trigger
	Randomize   []*EntityEvent  `json:"randomize,omitempty"`
	Sequence    []*EntityEvent  `json:"sequence,omitempty"`
	Filters     *f.Filter       `json:"filters,omitempty"`
}

func (e *EntityEvent) UnmarshalJSON(data []byte) error {
	type Alias EntityEvent

	aux := &struct {
		Add         *ComponentGroup `json:"add,omitempty"`
		Remove      *ComponentGroup `json:"remove,omitempty"`
		SetProperty ISetProperty    `json:"set_property,omitempty"`
		Trigger     interface{}     `json:"trigger,omitempty"`
		Randomize   []*EntityEvent  `json:"randomize,omitempty"`
		Sequence    []*EntityEvent  `json:"sequence,omitempty"`
		Filters     *f.Filter       `json:"filters,omitempty"`
	}{
		Add:         &ComponentGroup{ComponentGroups: &types.StringArray{}},
		Remove:      &ComponentGroup{ComponentGroups: &types.StringArray{}},
		SetProperty: &SetProperty{},
	}

	var json = jsoniter.ConfigFastest

	if err := json.Unmarshal(data, (*Alias)(aux)); err != nil {
		return err
	}

	*e = EntityEvent(*aux)

	return nil
}
