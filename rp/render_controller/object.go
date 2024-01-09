package rendercontroller

import g_util "github.com/respectZ/glowstone/util"

type Object map[string]interface{}

type Objects []Object

type IObjects interface {
	// Add adds a new object to the list.
	//
	//  Example:
	//
	//      objects.Add("*", true)
	Add(string, interface{})
	Get(int) (interface{}, bool)
	Remove(int)
	Clear()
	All() []Object
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error
}

func (o *Object) UnmarshalJSON(data []byte) error {
	var temp map[string]interface{}
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*o = temp
	return nil
}

func (o *Objects) Add(key string, value interface{}) {
	*o = append(*o, Object{key: value})
}

func (o *Objects) Get(index int) (interface{}, bool) {
	if index < 0 || index >= len(*o) {
		return nil, false
	}
	return (*o)[index], true
}

func (o *Objects) Remove(index int) {
	if index < 0 || index >= len(*o) {
		return
	}
	*o = append((*o)[:index], (*o)[index+1:]...)
}

func (o *Objects) Clear() {
	*o = make([]Object, 0)
}

func (o *Objects) All() []Object {
	return *o
}

func (o *Objects) IsEmpty() bool {
	return len(*o) == 0
}

func (o *Objects) Size() int {
	return len(*o)
}

func (o *Objects) UnmarshalJSON(b []byte) error {
	var temp []Object
	if err := g_util.UnmarshalJSON(b, &temp); err != nil {
		return err
	}
	*o = temp
	return nil
}
