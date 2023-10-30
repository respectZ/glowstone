package condition

import (
	loot_table "github.com/respectZ/glowstone/bp/loot_table"
)

func NewEntityProperties() *loot_table.Condition {
	return &loot_table.Condition{
		Condition: "entity_properties",
		Entity:    "this",
		Value: map[string]bool{
			"on_fire": true,
		},
	}
}
