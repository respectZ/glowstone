package types

import (
	"fmt"

	g_util "github.com/respectZ/glowstone/util"
)

type StringArray []string

type IStringArray interface {
	// Set the value of the array to v.
	Set([]string)

	// Add a new element to the array.
	Add(...string)

	// Get all elements of the array.
	All() []string

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

func (a *StringArray) Set(v []string) {
	*a = v
}

func (a *StringArray) Add(s ...string) {
	if *a == nil {
		*a = []string{}
	}
	*a = append(*a, s...)
}

func (a *StringArray) All() []string {
	if *a == nil {
		*a = []string{}
	}
	return *a
}

func (a *StringArray) Remove(index int) error {
	if *a == nil {
		*a = []string{}
		return nil
	}
	if index < 0 || index >= len(*a) {
		return fmt.Errorf("index out of range")
	}
	*a = append((*a)[:index], (*a)[index+1:]...)
	return nil
}

func (a *StringArray) Clear() {
	*a = []string{}
}

func (a *StringArray) IsEmpty() bool {
	if *a == nil {
		*a = []string{}
	}
	return len(*a) == 0
}

func (a *StringArray) Size() int {
	if *a == nil {
		*a = []string{}
	}
	return len(*a)
}

func (a *StringArray) UnmarshalJSON(data []byte) error {
	var temp []string
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}
