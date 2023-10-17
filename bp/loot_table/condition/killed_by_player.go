package condition

import (
	loot_table "github.com/respectZ/glowstone/bp/loot_table"
)

func NewKilledByPlayer() *loot_table.Condition {
	return &loot_table.Condition{
		Condition: "killed_by_player",
	}
}
