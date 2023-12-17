package types

import (
	g_util "github.com/respectZ/glowstone/util"
)

type StringArrayConditional []interface{}

type IStringArrayConditional interface {
	AddString(...string)
	AddStringConditional(string, string)
	All() []interface{}
	Clear()
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error
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
