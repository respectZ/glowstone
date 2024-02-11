package types

import (
	"fmt"

	g_util "github.com/respectZ/glowstone/util"
)

type MapStringArray []map[string]string

type IMapStringArray interface {
	// Set the value of the array to v.
	Set([]map[string]string)

	// Add a new element to the array.
	Add(string, string)

	// Get all elements of the array.
	All() []map[string]string

	// Remove the element of the specified index from array.
	Remove(int) error

	// Clear the array.
	Clear()

	// Check if the array is empty.
	IsEmpty() bool

	// Get the size of the array.
	Size() int

	// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
	UnmarshalJSON([]byte) error
}

func (a *MapStringArray) Set(v []map[string]string) {
	*a = v
}

func (a *MapStringArray) Add(s, c string) {
	if *a == nil {
		*a = []map[string]string{}
	}
	*a = append(*a, map[string]string{s: c})
}

func (a *MapStringArray) All() []map[string]string {
	if *a == nil {
		*a = []map[string]string{}
	}
	return *a
}

func (a *MapStringArray) Remove(index int) error {
	if *a == nil {
		*a = []map[string]string{}
		return nil
	}
	if index < 0 || index >= len(*a) {
		return fmt.Errorf("index %d out of bound", index)
	}
	*a = append((*a)[:index], (*a)[index+1:]...)
	return nil
}

func (a *MapStringArray) Clear() {
	if *a == nil {
		*a = []map[string]string{}
	}
	*a = []map[string]string{}
}

func (a *MapStringArray) IsEmpty() bool {
	if *a == nil {
		*a = []map[string]string{}
	}
	return len(*a) == 0
}

func (a *MapStringArray) Size() int {
	if *a == nil {
		*a = []map[string]string{}
	}
	return len(*a)
}

func (a *MapStringArray) UnmarshalJSON(data []byte) error {
	var temp []map[string]string
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}
