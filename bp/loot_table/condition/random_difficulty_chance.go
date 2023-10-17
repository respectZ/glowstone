package condition

import (
	loot_table "github.com/respectZ/glowstone/bp/loot_table"
)

// RandomDifficultyChance is a constructor for the `random_difficulty_chance` condition
//
// Example:
//
//	RandomDifficultyChance(0.5, 0.2, 1)
func RandomDifficultyChance(chance float64, hard float64, peaceful float64) *loot_table.Condition {
	return &loot_table.Condition{
		Condition: "random_difficulty_chance",
		Chance:    chance,
		Hard:      hard,
		Peaceful:  peaceful,
	}
}
