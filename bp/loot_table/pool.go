package loottable

type Pool interface {
	GetConditions() []*Condition
	SetConditions([]*Condition)
	AddCondition(Condition)
	ResetConditions()

	GetEntries() []Entry
	SetEntries([]*entry)
	AddEntry(*entry)
	NewEntryItem(string) Entry
	NewEntryLootTable(string) Entry
	NewEntryEmpty() Entry
	ResetEntries()

	GetRolls() int
	SetRolls(int)

	GetTiers() *tier
	SetTiers(*tier)
}

func (p *pool) GetConditions() []*Condition {
	return p.Conditions
}

func (p *pool) SetConditions(conditions []*Condition) {
	p.Conditions = conditions
}

func (p *pool) AddCondition(condition Condition) {
	p.Conditions = append(p.Conditions, &condition)
}

func (p *pool) ResetConditions() {
	p.Conditions = []*Condition{}
}

func (p *pool) GetEntries() []Entry {
	r := []Entry{}
	for _, e := range p.Entries {
		r = append(r, e)
	}
	return r
}

func (p *pool) SetEntries(entries []*entry) {
	p.Entries = entries
}

func (p *pool) AddEntry(entry *entry) {
	p.Entries = append(p.Entries, entry)
}

func (p *pool) NewEntryItem(name string) Entry {
	e := &entry{
		Type: "item",
		Name: name,
	}
	if p.Entries == nil {
		p.Entries = make([]*entry, 0)
	}
	p.Entries = append(p.Entries, e)
	return e
}

func (p *pool) NewEntryLootTable(name string) Entry {
	e := &entry{
		Type: "loot_table",
		Name: name,
	}
	if p.Entries == nil {
		p.Entries = make([]*entry, 0)
	}
	p.Entries = append(p.Entries, e)
	return e
}

func (p *pool) NewEntryEmpty() Entry {
	e := &entry{
		Type: "empty",
	}
	if p.Entries == nil {
		p.Entries = make([]*entry, 0)
	}
	p.Entries = append(p.Entries, e)
	return e
}

func (p *pool) ResetEntries() {
	p.Entries = []*entry{}
}

func (p *pool) GetRolls() int {
	return p.Rolls
}

func (p *pool) SetRolls(rolls int) {
	p.Rolls = rolls
}

func (p *pool) GetTiers() *tier {
	return p.Tiers
}

func (p *pool) SetTiers(tiers *tier) {
	p.Tiers = tiers
}
