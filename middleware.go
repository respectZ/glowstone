package glowstone

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	entityBPComponent "github.com/respectZ/glowstone/bp/entity/component"
	g_util "github.com/respectZ/glowstone/util"
)

type Middleware func(*Project) error

type Middlewares []Middleware

type IMiddlewares interface {
	Apply(*Project) error
	All() []Middleware
	Add(Middleware)
}

func (m Middlewares) Apply(g *Project) error {
	for _, v := range m {
		err := v(g)
		if err != nil {
			return fmt.Errorf("%s: %w", getFuncName(v), err)
		}
	}
	return nil
}

func (m Middlewares) All() []Middleware {
	return m
}

func (m *Middlewares) Add(f Middleware) {
	*m = append(*m, f)
}

func getFuncName(f Middleware) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func entity_WriteLang(g *Project) error {
	if g.RP.Lang == nil {
		return nil
	}

	rp := g.RP.Entity.All()
	bp := g.BP.Entity.All()
	for k, v := range rp {
		// Check to prevent overwriting
		k1 := fmt.Sprintf("entity.%s.name", v.GetNamespaceIdentifier())
		if v.Lang == "" {
			if !g.RP.Lang.Has(k1) {
				v.Lang = g_util.TitleCase(strings.ReplaceAll(v.GetIdentifier(), "_", " "))
				g.RP.Lang.Add(k1, v.Lang)
			}
		}
		k1 = fmt.Sprintf("item.spawn_egg.entity.%s.name", v.GetNamespaceIdentifier())
		if v.SpawnLang == "" {
			if !g.RP.Lang.Has(k1) {
				v.SpawnLang = "Spawn " + v.Lang
				g.RP.Lang.Add(k1, v.SpawnLang)
			}
		}
		k1 = fmt.Sprintf("action.hint.exit.%s", v.GetNamespaceIdentifier())
		// Rideable
		if v.RideHintLang == "" {
			var ridebable entityBPComponent.Rideable
			_, err := bp[k].Data.Entity.Components.Get(&ridebable)
			if err == nil {
				if !g.RP.Lang.Has(k1) {
					v.RideHintLang = "Tap jump to exit"
					g.RP.Lang.Add(k1, v.RideHintLang)
				}
			}
		} else {
			g.RP.Lang.Add(k1, v.RideHintLang)
		}
	}
	return nil
}

func item_WriteLang(g *Project) error {
	bp := g.BP.Item.All()
	for _, v := range bp {
		// Check to prevent overwriting
		k := fmt.Sprintf("item.%s.name", v.GetNamespaceIdentifier())
		if !g.RP.Lang.Has(k) {
			p := g_util.TitleCase(strings.ReplaceAll(v.GetIdentifier(), "_", " "))
			g.RP.Lang.Add(k, p)
		}
	}
	return nil
}
