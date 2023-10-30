package condition

import (
	loot_table "github.com/respectZ/glowstone/bp/loot_table"
)

// RandomRegionalDifficultyChance is a constructor for the `random_regional_difficulty_chance` condition
//
// Example:
//
//	RandomRegionalDifficultyChance(0.5)
func RandomRegionalDifficultyChance(max_chance float64) *loot_table.Condition {
	return &loot_table.Condition{
		Condition: "random_regional_difficulty_chance",
		MaxChance: max_chance,
	}
}
