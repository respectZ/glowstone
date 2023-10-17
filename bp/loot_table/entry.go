package loottable

type Entry interface {
	GetType() string
	GetName() string
	GetWeight() int
	SetWeight(int)
	SetType(string)
	SetName(string)
	AddFunction(*map[string]interface{})
}

func (e *entry) GetType() string {
	return e.Type
}

func (e *entry) GetName() string {
	return e.Name
}

func (e *entry) GetWeight() int {
	return e.Weight
}

func (e *entry) SetWeight(weight int) {
	e.Weight = weight
}

func (e *entry) SetType(t string) {
	e.Type = t
}

func (e *entry) SetName(name string) {
	e.Name = name
}

func (e *entry) AddFunction(function *map[string]interface{}) {
	if e.Functions == nil {
		e.Functions = make([]*map[string]interface{}, 0)
	}
	e.Functions = append(e.Functions, function)
}
