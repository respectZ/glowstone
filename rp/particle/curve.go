package particle

import (
	g_util "github.com/respectZ/glowstone/util"
)

type Curves map[string]*Curve

// TODO: since the node can either be a map[string]Node or []float64, we need to use interface{} here

type Curve struct {
	Type            string      `json:"type"`
	Input           interface{} `json:"input,omitempty"`
	HorizontalRange interface{} `json:"horizontal_range,omitempty"`

	// Nodes           *Nodes      `json:"nodes,omitempty"`

	Node interface{} `json:"nodes,omitempty"`
}

type ICurves interface {
	UnmarshalJSON([]byte) error
	Add(string, *Curve)
	Get(string) *Curve
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]*Curve

	New(string) *Curve
}

func (c *Curves) UnmarshalJSON(data []byte) error {
	var temp map[string]*Curve
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*c = temp
	return nil
}

func (c *Curves) Add(s string, v *Curve) {
	(*c)[s] = v
}

func (c *Curves) Get(s string) *Curve {
	return (*c)[s]
}

func (c *Curves) Clear() {
	*c = make(Curves, 0)
}

func (c *Curves) Remove(s string) {
	delete(*c, s)
}

func (c *Curves) IsEmpty() bool {
	return len(*c) == 0
}

func (c *Curves) Size() int {
	return len(*c)
}

func (c *Curves) All() map[string]*Curve {
	return *c
}

func (c *Curves) New(s string) *Curve {
	curve := &Curve{}
	(*c)[s] = curve
	return curve
}
