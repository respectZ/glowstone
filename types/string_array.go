package types

import (
	jsoniter "github.com/json-iterator/go"
)

type StringArray []string

type IStringArray interface {
	Add(...string)
	All() []string
	Clear()
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error
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
	var json = jsoniter.ConfigFastest
	var temp []string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}
