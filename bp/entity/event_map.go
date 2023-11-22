package entity

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/respectZ/glowstone/types"
)

type EntityEventMap map[string]*EntityEvent

type IEntityEventMap interface {
	UnmarshalJSON([]byte) error
	Add(string, *EntityEvent)
	Get(string) (*EntityEvent, bool)
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]*EntityEvent

	// Create a new event.
	//
	// name: The name of the event.
	//
	// return: The new event.
	//
	// Example:
	//
	//     event := entity.Events.New("minecraft:entity_spawned")
	New(string) *EntityEvent
}

func (m *EntityEventMap) UnmarshalJSON(data []byte) error {
	var json = jsoniter.ConfigFastest
	var temp map[string]*EntityEvent
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*m = temp
	return nil
}

func (m *EntityEventMap) Add(name string, event *EntityEvent) {
	if *m == nil {
		*m = make(EntityEventMap)
	}
	(*m)[name] = event
}

func (m *EntityEventMap) Get(name string) (*EntityEvent, bool) {
	if *m == nil {
		*m = make(EntityEventMap)
	}
	event, ok := (*m)[name]
	return event, ok
}

func (m *EntityEventMap) Clear() {
	if *m == nil {
		*m = make(EntityEventMap)
	}
	*m = make(EntityEventMap)
}

func (m *EntityEventMap) Remove(name string) {
	if *m == nil {
		*m = make(EntityEventMap)
	}
	delete(*m, name)
}

func (m *EntityEventMap) IsEmpty() bool {
	if *m == nil {
		*m = make(EntityEventMap)
	}
	return len(*m) == 0
}

func (m *EntityEventMap) Size() int {
	if *m == nil {
		*m = make(EntityEventMap)
	}
	return len(*m)
}

func (m *EntityEventMap) All() map[string]*EntityEvent {
	if *m == nil {
		*m = make(EntityEventMap)
	}
	return *m
}

func (m *EntityEventMap) New(name string) *EntityEvent {
	if *m == nil {
		*m = make(EntityEventMap)
	}
	event := &EntityEvent{
		Add:         &ComponentGroup{ComponentGroups: &types.StringArray{}},
		Remove:      &ComponentGroup{ComponentGroups: &types.StringArray{}},
		SetProperty: &SetProperty{},
	}
	(*m)[name] = event
	return event
}
