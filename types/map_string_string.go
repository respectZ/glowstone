package types

import (
	jsoniter "github.com/json-iterator/go"
)

type MapStringString map[string]string

type IMapStringString interface {
	Add(string, string)
	Get(string) (string, bool)
	Remove(string)
	Clear()
	All() map[string]string
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error
}

func (m *MapStringString) Add(key string, value string) {
	if *m == nil {
		*m = make(map[string]string)
	}
	(*m)[key] = value
}

func (m *MapStringString) Get(key string) (string, bool) {
	if *m == nil {
		*m = make(map[string]string)
	}
	value, ok := (*m)[key]
	return value, ok
}

func (m *MapStringString) Remove(key string) {
	if *m == nil {
		*m = make(map[string]string)
	}
	delete(*m, key)
}

func (m *MapStringString) Clear() {
	if *m == nil {
		*m = make(map[string]string)
	}
	*m = make(map[string]string)
}

func (m *MapStringString) All() map[string]string {
	if *m == nil {
		*m = make(map[string]string)
	}
	return *m
}

func (m *MapStringString) IsEmpty() bool {
	if *m == nil {
		*m = make(map[string]string)
	}
	return len(*m) == 0
}

func (m *MapStringString) Size() int {
	if *m == nil {
		*m = make(map[string]string)
	}
	return len(*m)
}

func (m *MapStringString) UnmarshalJSON(data []byte) error {
	if *m == nil {
		*m = make(map[string]string)
	}
	var json = jsoniter.ConfigFastest
	var temp map[string]string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*m = temp
	return nil
}
