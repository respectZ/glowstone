package glowstone

import (
	"fmt"

	loot_table "github.com/respectZ/glowstone/loot_table"
)

func (g *glowstone) AddLootTable(name string, lootTable *loot_table.LootTable) {
}

func (g *glowstone) GetLootTable(name string) (*loot_table.LootTable, error) {
	if g.LootTables == nil {
		g.LootTables = make(map[string]*loot_table.LootTable)
	}
	lootTable, ok := g.LootTables[name]
	if !ok {
		return nil, fmt.Errorf("loot table %s not found", name)
	}
	return lootTable, nil
}

func (g *glowstone) GetLootTables() map[string]*loot_table.LootTable {
	if g.LootTables == nil {
		g.LootTables = make(map[string]*loot_table.LootTable)
	}
	return g.LootTables
}

func (g *glowstone) NewLootTable(name string) *loot_table.LootTable {
	if g.LootTables == nil {
		g.LootTables = make(map[string]*loot_table.LootTable)
	}
	lootTable := loot_table.New(name)
	g.LootTables[name] = lootTable
	return lootTable
}
