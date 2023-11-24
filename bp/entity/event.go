package entity

import (
	f "github.com/respectZ/glowstone/bp/types"
	"github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

type EntityEvent struct {
	Add         *ComponentGroup `json:"add,omitempty"`
	Remove      *ComponentGroup `json:"remove,omitempty"`
	SetProperty ISetProperty    `json:"set_property,omitempty"`
	Trigger     *f.Trigger      `json:"trigger,omitempty"` // TODO: implement trigger
	Randomize   IEntityEvent    `json:"randomize,omitempty"`
	Sequence    IEntityEvent    `json:"sequence,omitempty"`
	Filters     *f.Filter       `json:"filters,omitempty"`
}

func (e *EntityEvent) UnmarshalJSON(data []byte) error {
	type Alias EntityEvent

	aux := &struct {
		Add         *ComponentGroup `json:"add,omitempty"`
		Remove      *ComponentGroup `json:"remove,omitempty"`
		SetProperty ISetProperty    `json:"set_property,omitempty"`
		Trigger     *f.Trigger      `json:"trigger,omitempty"`
		Randomize   IEntityEvent    `json:"randomize,omitempty"`
		Sequence    IEntityEvent    `json:"sequence,omitempty"`
		Filters     *f.Filter       `json:"filters,omitempty"`
	}{
		Add:         &ComponentGroup{ComponentGroups: &types.StringArray{}},
		Remove:      &ComponentGroup{ComponentGroups: &types.StringArray{}},
		SetProperty: &SetProperty{},
		Randomize:   &EntityEventArray{},
		Sequence:    &EntityEventArray{},
	}

	if err := g_util.UnmarshalJSON(data, (*Alias)(aux)); err != nil {
		return err
	}

	*e = EntityEvent(*aux)

	return nil
}

type EntityEventArray []*EntityEvent

type IEntityEvent interface {
	UnmarshalJSON([]byte) error
	Add(...*EntityEvent)
	All() []*EntityEvent
	New() *EntityEvent
	Clear()
	IsEmpty() bool
	Size() int
}

func (a *EntityEventArray) UnmarshalJSON(data []byte) error {
	var temp []*EntityEvent
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}

func (a *EntityEventArray) Add(s ...*EntityEvent) {
	if *a == nil {
		*a = []*EntityEvent{}
	}
	*a = append(*a, s...)
}

func (a *EntityEventArray) All() []*EntityEvent {
	if *a == nil {
		*a = []*EntityEvent{}
	}
	return *a
}

func (a *EntityEventArray) New() *EntityEvent {
	event := &EntityEvent{
		Add:         &ComponentGroup{ComponentGroups: &types.StringArray{}},
		Remove:      &ComponentGroup{ComponentGroups: &types.StringArray{}},
		SetProperty: &SetProperty{},
		Randomize:   &EntityEventArray{},
		Sequence:    &EntityEventArray{},
	}
	// Add event to array
	a.Add(event)
	return event
}

func (a *EntityEventArray) Clear() {
	*a = []*EntityEvent{}
}

func (a *EntityEventArray) IsEmpty() bool {
	if *a == nil {
		*a = []*EntityEvent{}
	}
	return len(*a) == 0
}

func (a *EntityEventArray) Size() int {
	if *a == nil {
		*a = []*EntityEvent{}
	}
	return len(*a)
}
