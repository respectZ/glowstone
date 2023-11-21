package types

type StringArray []string

type IStringArray interface {
	Add(...string)
	Clear()
	IsEmpty() bool
	Size() int
}

func (a *StringArray) Add(s ...string) {
	*a = append(*a, s...)
}

func (a *StringArray) Clear() {
	*a = []string{}
}

func (a *StringArray) IsEmpty() bool {
	return len(*a) == 0
}

func (a *StringArray) Size() int {
	return len(*a)
}
