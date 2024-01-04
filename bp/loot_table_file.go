package bp

import (
	bp "github.com/respectZ/glowstone/bp/loot_table"
)

type LootTableFile struct {
	Data bp.LootTable

	// Other stuff
}

// Create a new animation controller file
//
// Example:
//
//	loot_table := bp.NewLootTable("vanilla/player.animation.controller.json")
func NewLootTable(dest string) *LootTableFile {
	return &LootTableFile{
		Data: bp.New(),
	}
}

func LoadLootTable(src string) (bp.LootTable, error) {
	animation, err := bp.Load(src)
	if err != nil {
		return nil, err
	}
	return animation, nil
}

func (e *LootTableFile) Encode() ([]byte, error) {
	animation, err := e.Data.Encode()
	if err != nil {
		return nil, err
	}
	return animation, nil
}
