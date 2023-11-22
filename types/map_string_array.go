package types

import (
	jsoniter "github.com/json-iterator/go"
)

type MapStringArray []map[string]string

type IMapStringArray interface {
	Add(string, string)
	All() []map[string]string
	Clear()
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error
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
	var json = jsoniter.ConfigFastest
	var temp []map[string]string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}
