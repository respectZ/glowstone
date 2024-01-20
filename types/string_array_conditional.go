package types

import (
	"fmt"

	g_util "github.com/respectZ/glowstone/util"
)

type StringArrayConditional []interface{}

type IStringArrayConditional interface {
	// Set the value of the array to v.
	Set([]interface{})

	// Add a new element to the array.
	AddString(...string)

	// Add a new element to the array.
	AddStringConditional(string, string)

	// Get all elements of the array.
	All() []interface{}

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

func (a *StringArrayConditional) Set(v []interface{}) {
	*a = v
}

func (a *StringArrayConditional) AddString(s ...string) {
	if *a == nil {
		*a = []interface{}{}
	}
	for _, v := range s {
		*a = append(*a, v)
	}
}

func (a *StringArrayConditional) AddStringConditional(s, c string) {
	if *a == nil {
		*a = []interface{}{}
	}
	*a = append(*a, map[string]string{s: c})
}

func (a *StringArrayConditional) All() []interface{} {
	if *a == nil {
		*a = []interface{}{}
	}
	return *a
}

func (a *StringArrayConditional) Remove(index int) error {
	if *a == nil {
		*a = []interface{}{}
		return nil
	}
	if index < 0 || index >= len(*a) {
		return fmt.Errorf("index out of range")
	}
	*a = append((*a)[:index], (*a)[index+1:]...)
	return nil
}

func (a *StringArrayConditional) Clear() {
	*a = []interface{}{}
}

func (a *StringArrayConditional) IsEmpty() bool {
	if *a == nil {
		*a = []interface{}{}
	}
	return len(*a) == 0
}

func (a *StringArrayConditional) Size() int {
	if *a == nil {
		*a = []interface{}{}
	}
	return len(*a)
}

func (a *StringArrayConditional) UnmarshalJSON(data []byte) error {
	var temp []interface{}
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}
