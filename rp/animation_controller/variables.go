package animation_controller

import (
	"github.com/respectZ/glowstone/types"
	g_util "github.com/respectZ/glowstone/util"
)

type Variables map[string]*variables

type IVariables interface {
	UnmarshalJSON([]byte) error
	Add(string, *variables)
	Get(string) *variables
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]*variables

	// Create a new variable.
	//
	// name: The name of the variable.
	//
	// return: The new variable.
	//
	// Example:
	//
	//  v := variable.New("v.test")
	New(string) *variables
}

func (a *Variables) UnmarshalJSON(data []byte) error {
	var temp map[string]*variables
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*a = temp
	return nil
}

func (a *Variables) Add(name string, animationController *variables) {
	(*a)[name] = animationController
}

func (a *Variables) Get(name string) *variables {
	return (*a)[name]
}

func (a *Variables) Clear() {
	*a = make(Variables)
}

func (a *Variables) Remove(name string) {
	delete(*a, name)
}

func (a *Variables) IsEmpty() bool {
	return len(*a) == 0
}

func (a *Variables) Size() int {
	return len(*a)
}

func (a *Variables) All() map[string]*variables {
	return *a
}

func (a *Variables) New(name string) *variables {
	v := &variables{
		RemapCurve: &types.MapStringFloat{},
	}
	(*a)[name] = v
	return v
}
