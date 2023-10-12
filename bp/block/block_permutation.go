package block

func (b *blockPermutation) GetCondition() string {
	return b.Condition
}

func (b *blockPermutation) SetCondition(condition string) {
	b.Condition = condition
}

func (b *blockPermutation) SetComponent(component string, value interface{}) {
	b.Components[component] = value
}
