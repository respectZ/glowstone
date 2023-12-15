package animation

import (
	"fmt"

	g_util "github.com/respectZ/glowstone/util"
)

type Timeline map[string][]string

type ITimeline interface {
	UnmarshalJSON([]byte) error

	// Add a molang or command to specific timeline.
	//
	// - Example:
	//
	//    entity.Entity.Timeline.Add(0.0, "/say hi")
	Add(float64, ...string)
	Get(float64) ([]string, bool)
	Clear()
	Remove(float64)
	IsEmpty() bool
	Size() int
	All() map[string][]string
}

func (m *Timeline) UnmarshalJSON(data []byte) error {
	var temp map[string][]string
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*m = temp
	return nil
}

func (m *Timeline) Add(name float64, timeline ...string) {
	if *m == nil {
		*m = make(Timeline)
	}
	(*m)[fmt.Sprintf("%f", name)] = timeline
}

func (m *Timeline) Get(name float64) ([]string, bool) {
	if *m == nil {
		*m = make(Timeline)
	}
	timeline, ok := (*m)[fmt.Sprintf("%f", name)]
	return timeline, ok
}

func (m *Timeline) Clear() {
	if *m == nil {
		*m = make(Timeline)
	}
	*m = make(Timeline)
}

func (m *Timeline) Remove(name float64) {
	if *m == nil {
		*m = make(Timeline)
	}
	delete(*m, fmt.Sprintf("%f", name))
}

func (m *Timeline) IsEmpty() bool {
	if *m == nil {
		*m = make(Timeline)
	}
	return len(*m) == 0
}

func (m *Timeline) Size() int {
	if *m == nil {
		*m = make(Timeline)
	}
	return len(*m)
}

func (m *Timeline) All() map[string][]string {
	if *m == nil {
		*m = make(Timeline)
	}
	return *m
}
