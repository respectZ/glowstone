package types

import (
	g_util "github.com/respectZ/glowstone/util"
)

type StringArray []string

type IStringArray interface {
	Set([]string)
	Add(...string)
	All() []string
	Clear()
	IsEmpty() bool
	Size() int
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
