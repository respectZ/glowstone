package loottable

type lootTable struct {
	Pools []*pool `json:"pools,omitempty"`
}

type pool struct {
	Conditions []*Condition `json:"conditions,omitempty"`
	Entries    []*entry     `json:"entries,omitempty"`
	Rolls      int          `json:"rolls,omitempty"`
	Tiers      *tier        `json:"tiers,omitempty"`
}

type entry struct {
	Type      string                    `json:"type,omitempty"`
	Weight    int                       `json:"weight,omitempty"`
	Name      string                    `json:"name,omitempty"`
	Pools     []*pool                   `json:"pools,omitempty"`
	Functions []*map[string]interface{} `json:"functions,omitempty"` // TODO
}

type Condition struct {
	Condition         string                   `json:"condition,omitempty"`
	Value             interface{}              `json:"value,omitempty"`
	Entity            string                   `json:"entity,omitempty"`
	Properties        map[string]interface{}   `json:"properties,omitempty"`
	Item              string                   `json:"item,omitempty"`
	Enchantments      []map[string]interface{} `json:"enchantments,omitempty"`
	Chance            float64                  `json:"chance,omitempty"`
	LootingMultiplier float64                  `json:"looting_multiplier,omitempty"`
	Hard              float64                  `json:"hard,omitempty"`
	Peaceful          float64                  `json:"peaceful,omitempty"`
	MaxChance         float64                  `json:"max_chance,omitempty"`
}

type tier struct {
	BonusChance  float64 `json:"bonus_chance,omitempty"`
	BonusRolls   int     `json:"bonus_rolls,omitempty"`
	InitialRange int     `json:"initial_range,omitempty"`
}

func New() LootTable {
	return &lootTable{
		Pools: make([]*pool, 0),
	}
}
