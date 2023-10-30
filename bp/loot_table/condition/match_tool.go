package condition

import (
	loot_table "github.com/respectZ/glowstone/bp/loot_table"
)

// NewMatchTool is a constructor for the `match_tool` condition
//
// Example:
//
//	NewMatchTool("minecraft:golden_pickaxe", []map[string]interface{}{
//		{
//			"enchantment": "fortune",
//			"level":      1,
//		},
func NewMatchTool(item string, enchantments ...[]map[string]interface{}) *loot_table.Condition {
	l := &loot_table.Condition{
		Condition: "match_tool",
		Item:      item,
	}
	if len(enchantments) > 0 {
		l.Enchantments = make([]map[string]interface{}, 0)
		l.Enchantments = append(l.Enchantments, enchantments[0]...)
	}
	return l
}
