package types

type StringArrayConditional []interface{}

type IStringArrayConditional interface {
	AddString(...string)
	AddStringConditional(string, string)
	Clear()
	IsEmpty() bool
	Size() int
}

func (a *StringArrayConditional) AddString(s ...string) {
	*a = append(*a, s)
}

func (a *StringArrayConditional) AddStringConditional(s, c string) {
	*a = append(*a, map[string]string{s: c})
}

func (a *StringArrayConditional) Clear() {
	*a = []interface{}{}
}

func (a *StringArrayConditional) IsEmpty() bool {
	return len(*a) == 0
}

func (a *StringArrayConditional) Size() int {
	return len(*a)
}
