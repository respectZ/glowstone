package types

type MapStringArray []map[string]string

type IMapStringArray interface {
	Add(string, string)
	Clear()
	IsEmpty() bool
	Size() int
}

func (a *MapStringArray) Add(s, c string) {
	*a = append(*a, map[string]string{s: c})
}

func (a *MapStringArray) Clear() {
	*a = []map[string]string{}
}

func (a *MapStringArray) IsEmpty() bool {
	return len(*a) == 0
}

func (a *MapStringArray) Size() int {
	return len(*a)
}
