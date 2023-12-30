package glowstone

import (
	"fmt"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/loot_table"
	g_util "github.com/respectZ/glowstone/util"
)

type LootTable map[string]bp.LootTable

type ILootTable interface {
	Add(string, bp.LootTable)
	Get(string) (bp.LootTable, bool)
	Remove(string)
	Clear()
	All() map[string]bp.LootTable
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// Create a new loot table
	//
	// Example:
	//
	// 	lootTable := New("entities/pig.json")
	New(string) LootTable

	Save(string) error

	LoadAll(string) error
}

func (l *LootTable) Add(key string, value bp.LootTable) {
	if *l == nil {
		*l = make(map[string]bp.LootTable)
	}
	(*l)[key] = value
}

func (l *LootTable) Get(key string) (bp.LootTable, bool) {
	if *l == nil {
		*l = make(map[string]bp.LootTable)
	}
	value, ok := (*l)[key]
	return value, ok
}

func (l *LootTable) Remove(key string) {
	if *l == nil {
		*l = make(map[string]bp.LootTable)
	}
	delete(*l, key)
}

func (l *LootTable) Clear() {
	*l = make(map[string]bp.LootTable)
}

func (l *LootTable) All() map[string]bp.LootTable {
	if *l == nil {
		*l = make(map[string]bp.LootTable)
	}
	return *l
}

func (l *LootTable) IsEmpty() bool {
	if *l == nil {
		*l = make(map[string]bp.LootTable)
	}
	return len(*l) == 0
}

func (l *LootTable) Size() int {
	if *l == nil {
		*l = make(map[string]bp.LootTable)
	}
	return len(*l)
}

func (l *LootTable) UnmarshalJSON(data []byte) error {
	if *l == nil {
		*l = make(map[string]bp.LootTable)
	}
	return g_util.UnmarshalJSON(data, l)
}

func (l *LootTable) New(dest string) LootTable {
	a := bp.New()
	l.Add(dest, a)
	return *l
}

func (l *LootTable) Save(pathToBP string) error {
	if *l == nil {
		*l = make(map[string]bp.LootTable)
	}
	for dest, v := range l.All() {
		data, err := v.Encode()
		if err != nil {
			return err
		}
		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.LootTable, dest), data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *LootTable) LoadAll(pathToBP string) error {
	if *l == nil {
		*l = make(map[string]bp.LootTable)
	}
	flies, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.LootTable))
	if err != nil {
		return err
	}
	for _, file := range flies {
		lootTable, err := bp.Load(file)
		if err != nil {
			return err
		}
		// Remove pathToBP and destDirectory.LootTable from file
		file := strings.TrimPrefix(file, fmt.Sprintf("%s%s%s%s", pathToBP, string(filepath.Separator), destDirectory.LootTable, string(filepath.Separator)))
		l.Add(file, lootTable)
	}
	return nil
}
