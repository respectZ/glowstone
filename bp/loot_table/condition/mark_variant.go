package condition

import (
	loot_table "github.com/respectZ/glowstone/bp/loot_table"
)

func NewHasMarkVariant(value int) *loot_table.Condition {
	return &loot_table.Condition{
		Condition: "has_mark_variant",
		Value:     value,
	}
}
