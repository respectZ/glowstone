package entity

import g_util "github.com/respectZ/glowstone/util"

type Locators map[string]*Vector3

type ILocators interface {
	UnmarshalJSON([]byte) error

	Add(string, *Vector3)
	Get(string) *Vector3
	Clear()
	Remove(string)
	IsEmpty() bool
	Size() int
	All() map[string]*Vector3

	New(string, float64, float64, float64)
}

func (l *Locators) UnmarshalJSON(data []byte) error {
	var temp map[string]*Vector3
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*l = temp
	return nil
}

func (l *Locators) Add(s string, v *Vector3) {
	(*l)[s] = v
}

func (l *Locators) Get(s string) *Vector3 {
	return (*l)[s]
}

func (l *Locators) Clear() {
	*l = make(Locators, 0)
}

func (l *Locators) Remove(s string) {
	delete(*l, s)
}

func (l *Locators) IsEmpty() bool {
	return len(*l) == 0
}

func (l *Locators) Size() int {
	return len(*l)
}

func (l *Locators) All() map[string]*Vector3 {
	return *l
}

func (l *Locators) New(s string, x, y, z float64) {
	(*l)[s] = &Vector3{x, y, z}
}
