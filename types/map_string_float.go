package types

import (
	g_util "github.com/respectZ/glowstone/util"
)

type MapStringFloat map[string]float64

type IMapStringFloat interface {
	// Set the value of the map to v.
	Set(map[string]float64)

	// Add a new element to the map.
	Add(string, float64)

	// Get the value of the specified key from map.
	Get(string) (float64, bool)

	// Remove the element of the specified key from map.
	Remove(string)

	// Clear the map.
	Clear()

	// Get all elements of the map.
	All() map[string]float64

	// Check if the map is empty.
	IsEmpty() bool

	// Get the size of the map.
	Size() int

	// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
	UnmarshalJSON([]byte) error
}

func (m *MapStringFloat) Set(v map[string]float64) {
	*m = v
}

func (m *MapStringFloat) Add(key string, value float64) {
	if *m == nil {
		*m = make(map[string]float64)
	}
	(*m)[key] = value
}

func (m *MapStringFloat) Get(key string) (float64, bool) {
	if *m == nil {
		*m = make(map[string]float64)
	}
	value, ok := (*m)[key]
	return value, ok
}

func (m *MapStringFloat) Remove(key string) {
	if *m == nil {
		*m = make(map[string]float64)
	}
	delete(*m, key)
}

func (m *MapStringFloat) Clear() {
	if *m == nil {
		*m = make(map[string]float64)
	}
	*m = make(map[string]float64)
}

func (m *MapStringFloat) All() map[string]float64 {
	if *m == nil {
		*m = make(map[string]float64)
	}
	return *m
}

func (m *MapStringFloat) IsEmpty() bool {
	if *m == nil {
		*m = make(map[string]float64)
	}
	return len(*m) == 0
}

func (m *MapStringFloat) Size() int {
	if *m == nil {
		*m = make(map[string]float64)
	}
	return len(*m)
}

func (m *MapStringFloat) UnmarshalJSON(data []byte) error {
	if *m == nil {
		*m = make(map[string]float64)
	}
	var temp map[string]float64
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*m = temp
	return nil
}
