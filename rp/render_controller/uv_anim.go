package rendercontroller

import (
	g_util "github.com/respectZ/glowstone/util"
)

type UvAnim struct {
	Offset IArrayMisc `json:"offset,omitempty"`
	Scale  IArrayMisc `json:"scale,omitempty"`
}

type ArrayMisc []interface{}

type IArrayMisc interface {
	UnmarshalJSON([]byte) error
	Add(interface{})
	Get(int) (interface{}, bool)
	Remove(int)
	Size() int
	All() []interface{}
	IsEmpty() bool
	Set(int, interface{})
	Clear()
}

func (u *UvAnim) UnmarshalJSON(b []byte) error {
	var temp struct {
		Offset ArrayMisc `json:"offset,omitempty"`
		Scale  ArrayMisc `json:"scale,omitempty"`
	}
	if err := g_util.UnmarshalJSON(b, &temp); err != nil {
		return err
	}
	u.Offset = &temp.Offset
	u.Scale = &temp.Scale
	return nil
}

func (a *ArrayMisc) UnmarshalJSON(b []byte) error {
	var temp []interface{}
	if err := g_util.UnmarshalJSON(b, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}

func (a *ArrayMisc) Add(value interface{}) {
	*a = append(*a, value)
}

func (a *ArrayMisc) Get(index int) (interface{}, bool) {
	if index < 0 || index >= len(*a) {
		return nil, false
	}
	return (*a)[index], true
}

func (a *ArrayMisc) Remove(index int) {
	if index < 0 || index >= len(*a) {
		return
	}
	*a = append((*a)[:index], (*a)[index+1:]...)
}

func (a *ArrayMisc) Size() int {
	return len(*a)
}

func (a *ArrayMisc) All() []interface{} {
	return *a
}

func (a *ArrayMisc) IsEmpty() bool {
	return len(*a) == 0
}

func (a *ArrayMisc) Set(index int, value interface{}) {
	if index < 0 || index >= len(*a) {
		return
	}
	(*a)[index] = value
}

func (a *ArrayMisc) Clear() {
	*a = make([]interface{}, 0)
}
