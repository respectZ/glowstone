package entity

import g_util "github.com/respectZ/glowstone/util"

type EntityAliasData map[string]interface{}

type IEntityAliasData interface {
	UnmarshalJSON([]byte) error
	Get(string) (interface{}, bool)
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]interface{}

	// Add a property to the alias.
	//
	// name: The name of the property.
	// value: The value of the property.
	//
	// Example:
	//
	//     alias.Add("property:variant", 1)
	Add(string, interface{})
}

func (d *EntityAliasData) UnmarshalJSON(data []byte) error {
	var temp map[string]interface{}
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}

	*d = temp
	return nil
}

func (d *EntityAliasData) Get(key string) (interface{}, bool) {
	v, ok := (*d)[key]
	return v, ok
}

func (d *EntityAliasData) Clear() {
	*d = make(map[string]interface{})
}

func (d *EntityAliasData) Remove(key string) {
	delete(*d, key)
}

func (d *EntityAliasData) IsEmpty() bool {
	return len(*d) == 0
}

func (d *EntityAliasData) Size() int {
	return len(*d)
}

func (d *EntityAliasData) All() map[string]interface{} {
	return *d
}

func (d *EntityAliasData) Add(key string, value interface{}) {
	(*d)[key] = value
}

type EntityAlias map[string]*EntityAliasData

type IEntityAlias interface {
	UnmarshalJSON([]byte) error
	Add(string, *EntityAliasData)
	Get(string) (*EntityAliasData, bool)
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]*EntityAliasData

	// Create a new alias.
	//
	// name: The name of the alias.
	//
	// return: The new alias.
	//
	// Example:
	//
	//     alias := entity.Aliases.New("minecraft:entity_spawned")
	//     alias.Add("property:variant", 1)
	New(string) *EntityAliasData
}

func (a *EntityAlias) UnmarshalJSON(data []byte) error {
	var temp map[string]*EntityAliasData
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}

	*a = temp
	return nil
}

func (a *EntityAlias) Add(key string, value *EntityAliasData) {
	(*a)[key] = value
}

func (a *EntityAlias) Get(key string) (*EntityAliasData, bool) {
	v, ok := (*a)[key]
	return v, ok
}

func (a *EntityAlias) Clear() {
	*a = make(map[string]*EntityAliasData)
}

func (a *EntityAlias) Remove(key string) {
	delete(*a, key)
}

func (a *EntityAlias) IsEmpty() bool {
	return len(*a) == 0
}

func (a *EntityAlias) Size() int {
	return len(*a)
}

func (a *EntityAlias) All() map[string]*EntityAliasData {
	return *a
}

func (a *EntityAlias) New(key string) *EntityAliasData {
	e := &EntityAliasData{}
	(*a)[key] = e
	return e
}
