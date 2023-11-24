package entity

import (
	g_util "github.com/respectZ/glowstone/util"
)

type SetProperty map[string]interface{}

type ISetProperty interface {
	UnmarshalJSON([]byte) error
	SetBool(string, bool)
	SetInt(string, int)
	SetString(string, string)
	IsEmpty() bool
}

func (s *SetProperty) UnmarshalJSON(data []byte) error {
	var temp map[string]interface{}
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*s = temp
	return nil
}

func (s *SetProperty) SetBool(name string, value bool) {
	(*s)[name] = value
}

func (s *SetProperty) SetInt(name string, value int) {
	(*s)[name] = value
}

func (s *SetProperty) SetString(name string, value string) {
	(*s)[name] = value
}

func (s *SetProperty) IsEmpty() bool {
	return len(*s) == 0
}
