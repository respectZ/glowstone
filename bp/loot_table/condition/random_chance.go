package condition

import (
	loot_table "github.com/respectZ/glowstone/bp/loot_table"
)

// NewRandomChance is a constructor for the `random_chance` condition
//
// Example:
//
//	NewRandomChance(0.5)
func NewRandomChance(chance float64) *loot_table.Condition {
	return &loot_table.Condition{
		Condition: "random_chance",
		Chance:    chance,
	}
}
