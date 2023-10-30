package loottable

import (
	loot_table "github.com/respectZ/glowstone/bp/loot_table"
)

type LootTable struct {
	Data loot_table.LootTable

	// Other stuff
	Dest string
}

// New creates a new LootTable. Dest doesn't need to start with "loot_tables/"
//
// Example:
//
//	a := New("chests/simple_dungeon.json")
//	b := New("chests/abandoned_mineshaft.json")
func New(dest string) *LootTable {
	return &LootTable{
		Data: loot_table.New(),
		Dest: dest,
	}
}

// Load loads the LootTable from the given path. This is not setting the Dest.
//
// Example:
//
//	a, err := Load("chests/simple_dungeon.json")
func Load(src string) (*LootTable, error) {
	a, err := loot_table.Load(src)
	if err != nil {
		return nil, err
	}
	return &LootTable{
		Data: a,
	}, nil
}

// Encode returns []byte of the LootTable.
//
// Example:
//
//	lootTable, err := e.Encode()
func (a *LootTable) Encode() ([]byte, error) {
	return a.Data.Encode()
}
