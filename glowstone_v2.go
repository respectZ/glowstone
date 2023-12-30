package glowstone

import (
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

func newBPV2() *bp.GlowstoneBP {
	data := &bp.GlowstoneBP{
		Path:                filepath.Join("packs", "BP"),
		AnimationController: &bp.AnimationControllerBP{},
		Animation:           &bp.AnimationBP{},
		Block:               &bp.BlockBP{},
		Entity:              &bp.EntityBP{},
		Item:                &bp.ItemBP{},
		LootTable:           &bp.LootTable{},
		Recipe:              &bp.RecipeBP{},
		Manifest:            &bp.Manifest{},
	}
	return data
}

func newRPV2() *rp.GlowstoneRP {
	data := &rp.GlowstoneRP{
		Path: filepath.Join("packs", "RP"),
	}
	return data
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
		BP:     newBPV2(),
		RP:     newRPV2(),
	}
}

func (g *Project) Initialize() {
	// Animatinon Controller
	err := g.BP.AnimationController.LoadAll(g.BP.Path)
	if err != nil {
		g.Logger.Error.Println(err)
	}
	// Animation
	err = g.BP.Animation.LoadAll(g.BP.Path)
	if err != nil {
		g.Logger.Error.Println(err)
	}
	// Block
	err = g.BP.Block.LoadAll(g.BP.Path)
	if err != nil {
		g.Logger.Error.Println(err)
	}
	// Entity
	err = g.BP.Entity.LoadAll(g.BP.Path)
	if err != nil {
		g.Logger.Error.Println(err)
	}
	// Item
	err = g.BP.Item.LoadAll(g.BP.Path)
	if err != nil {
		g.Logger.Error.Println(err)
	}
	// Manifest
	err = g.BP.Manifest.Load(filepath.Join(g.BP.Path, "manifest.json"))
	if err != nil {
		g.Logger.Error.Println(err)
	}
}

func (g *Project) Save() {
	err := g.BP.Save()
	if err != nil {
		g.Logger.Error.Println(err)
	}
	// TODO: Save RP
	// TODO: Middlewares
}
