package loottable

import g_util "github.com/respectZ/glowstone/util"

type LootTable interface {
	Encode(...bool) ([]byte, error)
	GetPools() []Pool

	// Create a new pool
	//
	// Example:
	//
	// 	pool := lootTable.NewPool()
	NewPool() Pool
}

func (l *lootTable) Encode(minify ...bool) ([]byte, error) {
	return g_util.MarshalJSON(l, minify...)
}

func (l *lootTable) GetPools() []Pool {
	p := make([]Pool, 0)
	for _, pool := range l.Pools {
		p = append(p, pool)
	}
	return p
}

func (l *lootTable) NewPool() Pool {
	pool := &pool{
		Entries: []*entry{},
		Rolls:   1,
	}
	l.Pools = append(l.Pools, pool)
	return pool
}

func Load(src string) (LootTable, error) {
	a := New()
	err := g_util.LoadJSON(src, a)
	return a, err
}
