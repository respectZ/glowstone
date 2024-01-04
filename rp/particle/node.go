package particle

import (
	"fmt"

	g_util "github.com/respectZ/glowstone/util"
)

// type Nodes map[string]*Node
type Nodes map[string]interface{}

type Node struct {
	Value float64 `json:"value"`
	Slope float64 `json:"slope"`
}

type INodes interface {
	UnmarshalJSON([]byte) error

	Add(float64, interface{})
	Get(float64) interface{}
	Clear()
	Remove(float64)
	IsEmpty() bool
	Size() int
	All() map[string]interface{}

	New(float64) interface{}
}

func (n *Nodes) UnmarshalJSON(data []byte) error {
	var temp map[string]interface{}
	fmt.Printf("%s\n", data)
	fmt.Println("b")
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	fmt.Println("a")
	*n = temp
	return nil
}

func (n *Nodes) Add(s string, v interface{}) {
	(*n)[s] = v
}

func (n *Nodes) Get(s string) interface{} {
	return (*n)[s]
}

func (n *Nodes) Clear() {
	*n = make(Nodes, 0)
}

func (n *Nodes) Remove(s string) {
	delete(*n, s)
}

func (n *Nodes) IsEmpty() bool {
	return len(*n) == 0
}

func (n *Nodes) Size() int {
	return len(*n)
}

func (n *Nodes) All() map[string]interface{} {
	return *n
}

func (n *Nodes) New(s float64) interface{} {
	node := &Node{}
	(*n)[fmt.Sprintf("%f", s)] = node
	return node
}
