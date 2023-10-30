package condition

import (
	loot_table "github.com/respectZ/glowstone/bp/loot_table"
)

func NewKilledByPlayerOrPets() *loot_table.Condition {
	return &loot_table.Condition{
		Condition: "killed_by_player_or_pets",
	}
}
