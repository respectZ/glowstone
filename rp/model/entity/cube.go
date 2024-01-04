package entity

import (
	g_util "github.com/respectZ/glowstone/util"
)

type Cube struct {
	Origin      []float64   `json:"origin"`
	Size        []float64   `json:"size"`
	Uv          interface{} `json:"uv"`
	Inflate     float64     `json:"inflate,omitempty"`
	NeverRender *bool       `json:"never_render,omitempty"`
}

type Cubes []*Cube

type ICubes interface {
	Add(...*Cube)
	All() []*Cube
	Clear()
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	New() *Cube
}

func (c *Cubes) Add(s ...*Cube) {
	if *c == nil {
		*c = make(Cubes, 0)
	}
	*c = append(*c, s...)
}

func (c *Cubes) All() []*Cube {
	if *c == nil {
		*c = make(Cubes, 0)
	}
	return *c
}

func (c *Cubes) Clear() {
	*c = make(Cubes, 0)
}

func (c *Cubes) IsEmpty() bool {
	if *c == nil {
		*c = make(Cubes, 0)
	}
	return len(*c) == 0
}

func (c *Cubes) Size() int {
	if *c == nil {
		*c = make(Cubes, 0)
	}
	return len(*c)
}

func (c *Cubes) UnmarshalJSON(data []byte) error {
	var temp []*Cube
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*c = temp
	return nil
}

func (c *Cubes) New() *Cube {
	cube := &Cube{}
	*c = append(*c, cube)
	return cube
}
