package glowstone

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	bp "github.com/respectZ/glowstone/bp"
	rp "github.com/respectZ/glowstone/rp"
)

/*
TODO
- Move to BP
- Write lang and some other options
- Middlewares
*/

type Project struct {
	Logger *logger
	BP     *bp.GlowstoneBP
	RP     *rp.GlowstoneRP
}

func newLogger() *logger {
	return &logger{
		Warning: log.New(os.Stdout, "\033[33m[WARNING]\033[0m ", log.Ltime),
		Error:   log.New(os.Stderr, "\033[31m[ERROR]\033[0m ", log.Ltime),
	}
}

func NewProjectV2() *Project {
	return &Project{
		Logger: newLogger(),
		BP:     bp.New(filepath.Join("packs", "BP")),
		RP:     rp.New(filepath.Join("packs", "RP")),
	}
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
}

func (g *Project) Save() {
	// TODO: Middlewares
	err := g.BP.Save()
	if err != nil {
		g.Logger.Error.Println(err)
	}
	// TODO: Save RP
	err = g.RP.Save()
	if err != nil {
		g.Logger.Error.Println(err)
	}
}
