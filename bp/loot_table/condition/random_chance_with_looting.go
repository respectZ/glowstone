package condition

import (
	loot_table "github.com/respectZ/glowstone/bp/loot_table"
)

// NewRandomChanceWithLooting is a constructor for the `random_chance_with_looting` condition
//
// Example:
//
//	NewRandomChanceWithLooting(0.5, 0.01)
func NewRandomChanceWithLooting(chance float64, looting_multiplier float64) *loot_table.Condition {
	return &loot_table.Condition{
		Condition:         "random_chance_with_looting",
		Chance:            chance,
		LootingMultiplier: looting_multiplier,
	}
}
