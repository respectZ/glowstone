package glowstone

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	bp "github.com/respectZ/glowstone/bp"
	rp "github.com/respectZ/glowstone/rp"
)

type Project struct {
	// Logger
	Logger *logger

	// Contains Behavior Pack data.
	BP *bp.GlowstoneBP

	// Contains Resource Pack data.
	RP *rp.GlowstoneRP

	// Middlewares
	//
	// Contains functions that will be executed before saving the project.
	//
	// Example:
	//
	//	g := glowstone.NewProjectV2()
	//	func entity_WriteLang(g *Project) error {
	// 		// code
	//		return nil
	//	}
	//	g.Middlewares.Add(entity_WriteLang)
	Middlewares IMiddlewares
}

type values struct {
	Bool  bool
	Int   int
	Float float64
}

var ZeroValue = values{false, 0, 0.0}

func newLogger() *logger {
	return &logger{
		Warning: log.New(os.Stdout, "\033[33m[WARNING]\033[0m ", log.Ltime),
		Error:   log.New(os.Stderr, "\033[31m[ERROR]\033[0m ", log.Ltime),
	}
}

func NewProject() *Project {
	project := &Project{
		Logger: newLogger(),
		BP:     bp.New(filepath.Join("packs", "BP")),
		RP:     rp.New(filepath.Join("packs", "RP")),
		Middlewares: &Middlewares{
			entity_WriteLang,
			item_WriteLang,
		},
	}
	return project
}

// Preload loads all files from the project.
func (g *Project) Preload() {
	err := g.BP.Preload()
	if err != nil {
		g.Logger.Error.Println(fmt.Errorf("BP: %w", err))
	}
	err = g.RP.Preload()
	if err != nil {
		g.Logger.Error.Println(fmt.Errorf("RP: %w", err))
	}
	// TODO: post preload functions to load entity lang.
}

func (g *Project) Save() {
	// Middlewares
	err := g.Middlewares.Apply(g)
	if err != nil {
		g.Logger.Error.Println(fmt.Errorf("Middlewares: %w", err))
	}

	err = g.BP.Save()
	if err != nil {
		g.Logger.Error.Println(err)
	}
	err = g.RP.Save()
	if err != nil {
		g.Logger.Error.Println(err)
	}
}

// Create a new entity (BP and RP).
//
// Returns BPFile and RPFile of the entity, which contain additional fields along with the entity.
//
// Example:
//
//	bp, rp := g.NewEntity("glostone:diamond_sword")
func (g *Project) NewEntity(identifier string) (*bp.EntityFile, *rp.EntityFile) {
	g.BP.Entity.New(identifier)
	g.RP.Entity.New(identifier)

	bp := g.BP.Entity.All()[identifier]
	rp := g.RP.Entity.All()[identifier]

	return bp, rp
}
