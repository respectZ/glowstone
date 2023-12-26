package glowstone

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	animation "github.com/respectZ/glowstone/animation"
	animationController "github.com/respectZ/glowstone/animation_controller"
	attachable "github.com/respectZ/glowstone/attachable"
	block "github.com/respectZ/glowstone/block"
	"github.com/respectZ/glowstone/entity"
	item "github.com/respectZ/glowstone/item"
	recipe "github.com/respectZ/glowstone/recipe"
	"github.com/respectZ/glowstone/rp"
	g_util "github.com/respectZ/glowstone/util"

	blockBP "github.com/respectZ/glowstone/bp/block"
	entityBP "github.com/respectZ/glowstone/bp/entity"
	itemBP "github.com/respectZ/glowstone/bp/item"
	recipeBP "github.com/respectZ/glowstone/bp/recipe"
	entityRP "github.com/respectZ/glowstone/rp/entity"
	sound "github.com/respectZ/glowstone/rp/sound"
	texture "github.com/respectZ/glowstone/rp/texture"

	blockBPCompnent "github.com/respectZ/glowstone/bp/block/component"
	entityBPComponent "github.com/respectZ/glowstone/bp/entity/component"
	itemBPComponent "github.com/respectZ/glowstone/bp/item/component"
)

var MIN_ENGINE_VERSION = [3]int{1, 20, 40}

func NewProject() Glowstone {
	return &glowstone{
		BPDir:            "./packs/BP/",
		RPDir:            "./packs/RP/",
		MinEngineVersion: MIN_ENGINE_VERSION,

		Logger: &logger{
			Warning: log.New(os.Stdout, "\033[33m[WARNING]\033[0m ", log.Ltime),
			Error:   log.New(os.Stderr, "\033[31m[ERROR]\033[0m ", log.Ltime),
		},
		Lang: make(map[string]string),

		Entities:              make(map[string]*entity.Entity),
		Items:                 make(map[string]*item.Item),
		BPAnimationController: make(map[string]*animationController.BPAnimationController),
		BPAnimation:           make(map[string]*animation.BPAnimation),
		Recipes:               make(map[string]interface{}),
		Attachables:           make(map[string]*attachable.Attachable),

		RPAnimationController: make(map[string]*animationController.RPAnimationController),

		IsUpfront: false,
	}
}

func (g *glowstone) GetBPDir() string {
	return g.BPDir
}

func (g *glowstone) SetBPDir(bpdir string) {
	g.BPDir = bpdir
}

func (g *glowstone) GetRPDir() string {
	return g.RPDir
}

func (g *glowstone) SetRPDir(rpdir string) {
	g.RPDir = rpdir
}

func (g *glowstone) GetMinEngineVersion() [3]int {
	return g.MinEngineVersion
}

func (g *glowstone) SetMinEngineVersion(version [3]int) {
	g.MinEngineVersion = version
}

func (g *glowstone) GetLang(key string) string {
	if v, ok := g.Lang[key]; ok {
		return v
	}
	return ""
}

func (g *glowstone) SetLang(data map[string]string) {
	g.Lang = data
}

func (g *glowstone) AddLang(key string, value string) {
	g.Lang[key] = value
}

func (g *glowstone) Build() error {
	g_util.Writelang(g.RPDir+"texts/en_US.lang", g.Lang)
	return nil
}

func (g *glowstone) Initialize() error {
	// Read Lang
	data, err := g_util.Loadlang(g.RPDir + "texts/en_US.lang")
	if err != nil {
		g.Logger.Warning.Println("Failed to read lang file, creating new lang file.")
		g.SetLang(make(map[string]string))
		g_util.Writefile(g.RPDir+"texts/en_US.lang", []byte{})
	} else {
		g.SetLang(data)
	}

	// Read blocks.json
	g.RPBlocks, _ = rp.LoadBlocks(g.RPDir + "blocks.json")

	// Read ItemTexture
	g.ItemTexture, err = texture.LoadItemTexture(g.RPDir + "textures/item_texture.json")
	if err != nil {
		g.Logger.Warning.Println("Failed to read item_texture.json file, creating new item_texture.json file.")
		g.ItemTexture = texture.NewItemTexture()
	}

	// Read TerrainTexture
	g.TerrainTexture, _ = texture.LoadTerrainTexture(g.RPDir + "textures/terrain_texture.json")

	return nil
}

func (g *glowstone) SetUpfront(upfront bool) {
	g.IsUpfront = upfront
}

func (g *glowstone) Save() {
	// RPBlocks
	// Do some workaround to read format_version from blocks.json
	blocksFormatVersion := [3]int{1, 1, 0}
	blocksSrc := path.Join(g.RPDir, "blocks.json")
	var blocks map[string][3]int
	err := g_util.LoadJSON(blocksSrc, blocks)
	if err == nil {
		blocksFormatVersion = blocks["format_version"]
	}

	if g.RPBlocks != nil {
		data, err := g.RPBlocks.Encode()
		if err == nil {
			// Reload to write format_version
			var temp map[string]interface{}

			g_util.UnmarshalJSON(data, &temp)
			temp["format_version"] = blocksFormatVersion
			data, err = g_util.MarshalJSON(temp)
			if err != nil {
				g.Logger.Error.Println(err)
			}

			// Encode
			g_util.Writefile(path.Join(g.RPDir, "blocks.json"), data)
		}
	}

	// ItemTexture
	data, err := g.ItemTexture.Encode()
	if err != nil {
		g.Logger.Error.Println(err)
	} else {
		g_util.Writefile(path.Join(g.RPDir, "textures", "item_texture.json"), data)
	}

	// TerrainTexture
	if g.TerrainTexture != nil {
		data, err = g.TerrainTexture.Encode()
		if err == nil {
			g_util.Writefile(path.Join(g.RPDir, "textures", "terrain_texture.json"), data)
		}
	}

	// Sound Definition
	if g.SoundDefinition != nil {
		data, err := g.SoundDefinition.Encode()
		if err != nil {
			g.Logger.Error.Println(err)
		} else {
			g_util.Writefile(path.Join(g.RPDir, "sounds", "sound_definitions.json"), data)
		}
	}
	// Lang
	defer g_util.Writelang(path.Join(g.RPDir, "texts", "en_US.lang"), g.Lang)

	// Entity
	for _, e := range g.Entities {
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		if e.BP != nil {
			if e.RideHint == "" {
				// Check if rideable
				var ridebable entityBPComponent.Rideable
				_, err := e.BP.Entity.Components.Get(&ridebable)
				if err == nil {
					// Add to lang
					e.SetRideHint(fmt.Sprintf("Tap jump to exit the %s", e.Lang))
				}
			}
			bp, err := e.BP.Encode()
			if err != nil {
				g.Logger.Error.Println(err)
				continue
			}
			g_util.Writefile(path.Join(g.BPDir, "entities", e.Subdir, fmt.Sprintf("%s.json", e.GetIdentifier())), bp)
		}
		if e.RP != nil {
			// AutoSpawnEggTexture
			if e.AutoSpawnEggTexture {
				// Get texture
				rp_textures := e.RP.Entity.Description.Textures
				// Get first texture
				var texture string
				for _, v := range rp_textures.All() {
					texture = v
					break
				}
				// Add rpdir
				texture = path.Join(g.RPDir, texture)
				// Add .png
				texture = texture + ".png"
				// Check if texture exists
				if _, err := os.Stat(texture); err == nil {
					primary, secondary, err := g_util.GetEggColor(texture)
					if err == nil {
						// Set color
						e.RP.Entity.Description.SpawnEgg.BaseColor = primary
						e.RP.Entity.Description.SpawnEgg.OverlayColor = secondary
						// e.RP.GetSpawnEgg().SetBaseColor(primary)
						// e.RP.GetSpawnEgg().SetOverlayColor(secondary)
					} else {
						g.Logger.Warning.Printf("failed to get egg color for %s: %s", e.GetIdentifier(), err)
					}
				} else {
					g.Logger.Warning.Printf("missing texture auto for %s: %s", e.GetIdentifier(), err)
				}
			}
			rp, err := e.RP.Encode()
			if err != nil {
				g.Logger.Error.Println(err)
				continue
			}
			g_util.Writefile(path.Join(g.RPDir, "entity", e.Subdir, fmt.Sprintf("%s.json", e.GetIdentifier())), rp)
		}
	}

	// Loot Table
	if g.LootTables != nil {
		for _, lootTable := range g.LootTables {
			data, err := lootTable.Encode()
			if err != nil {
				g.Logger.Error.Println(err)
				continue
			}
			g_util.Writefile(path.Join(g.BPDir, "loot_tables", lootTable.Dest), data)
		}
	}

	// Block
	if g.Blocks != nil {
		for _, b := range g.Blocks {
			// Display name
			var displayName blockBPCompnent.DisplayName
			_, err = b.BP.GetComponent(&displayName)
			if err != nil {
				// Add component
				displayName = blockBPCompnent.DisplayName(fmt.Sprintf("block.%s.name", b.GetNamespaceIdentifier()))
				b.BP.AddComponent(&displayName)
			}
			bp, err := b.Encode()
			if err != nil {
				g.Logger.Error.Println(err)
				continue
			}
			g_util.Writefile(path.Join(g.BPDir, "blocks", b.Subdir, fmt.Sprintf("%s.json", b.GetIdentifier())), bp)
		}
	}

	// Item
	for _, i := range g.Items {
		// Display name
		var displayName itemBPComponent.DisplayName
		_, err := i.BP.GetComponent(&displayName)
		if err != nil {
			g.Logger.Error.Println(err)
			// Add component
			displayName.Value = fmt.Sprintf("item.%s.name", i.GetNamespaceIdentifier())
			i.BP.AddComponent(&displayName)
		}
		bp, err := i.Encode()
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		g_util.Writefile(path.Join(g.BPDir, "items", i.Subdir, fmt.Sprintf("%s.json", i.GetIdentifier())), bp)
	}

	// BPAnimationController
	for _, a := range g.BPAnimationController {
		data, err := a.Encode()
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		g_util.Writefile(path.Join(g.BPDir, "animation_controllers", a.Dest), data)
	}

	// RPAnimationController
	for _, a := range g.RPAnimationController {
		data, err := a.Encode()
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		g_util.Writefile(path.Join(g.RPDir, "animation_controllers", a.Dest), data)
	}

	// BPAnimation
	for _, a := range g.BPAnimation {
		data, err := a.Encode()
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		g_util.Writefile(path.Join(g.BPDir, "animations", a.Dest), data)
	}

	// Recipes
	for _, r := range g.Recipes {
		switch r := r.(type) {
		case recipe.RecipeInterface:
			data, err := r.Encode()
			if err != nil {
				g.Logger.Error.Println(err)
				continue
			}
			g_util.Writefile(path.Join(g.BPDir, "recipes", r.GetSubdir(), r.GetIdentifier()+".json"), data)
		case recipeBP.RecipeInterface:
			data, err := r.Encode()
			if err != nil {
				g.Logger.Error.Println(err)
				continue
			}
			name := strings.ReplaceAll(r.GetIdentifier(), ":", "_")
			g_util.Writefile(path.Join(g.BPDir, "recipes", name+".json"), data)
		default:
			g.Logger.Error.Printf("invalid type %T", r)
		}
	}

	// Attachable
	for _, a := range g.Attachables {
		data, err := a.Encode()
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		g_util.Writefile(path.Join(g.RPDir, "attachables", a.Subdir, fmt.Sprintf("%s.json", a.GetIdentifier())), data)
	}
}

/******************* Preloads *******************/

func (g *glowstone) PreloadEntities() {
	if g.Entities == nil {
		g.Entities = make(map[string]*entity.Entity)
	}

	// BP
	files, err := g_util.Walk(path.Join(g.BPDir, "entities"))
	if err != nil {
		g.Logger.Error.Println(err)
		return
	}
	for _, file := range files {
		e, err := entity.LoadBP(file)
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		if g.Entities[e.GetIdentifier()] == nil {
			en := &entity.Entity{}
			g.Entities[e.GetIdentifier()] = en
			// Get basepath
			basepath := filepath.Dir(file)
			subdir := strings.ReplaceAll(basepath, filepath.Join(g.BPDir, "entities"), "")
			en.Subdir = subdir

			// Get lang
			full := strings.Split(e.GetIdentifier(), ":")
			namespace := full[0]
			identifier := full[1]

			lang := g.GetLang(fmt.Sprintf("entity.%s:%s.name", namespace, identifier))
			en.SetLang(lang)

			spawnLang := g.GetLang(fmt.Sprintf("item.spawn_egg.entity.%s:%s.name", namespace, identifier))
			en.SpawnLang = spawnLang
		}
		g.Entities[e.GetIdentifier()].BP = e
	}
	// RP
	files, err = g_util.Walk(path.Join(g.RPDir, "entity"))
	if err != nil {
		g.Logger.Error.Println(err)
		return
	}
	for _, file := range files {
		e, err := entity.LoadRP(file)
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		// Check if identifier exists
		if _, ok := g.Entities[e.GetIdentifier()]; !ok {
			g.Logger.Warning.Printf("entity %s not found in BP", e.GetIdentifier())
			continue
		}
		if g.Entities[e.GetIdentifier()] == nil {
			g.Entities[e.GetIdentifier()] = &entity.Entity{}
			// Get basepath
			basepath := filepath.Dir(file)
			subdir := strings.ReplaceAll(basepath, filepath.Join(g.BPDir, "entities"), "")
			g.Entities[e.GetIdentifier()].Subdir = subdir
		}
		g.Entities[e.GetIdentifier()].RP = e
	}
}

func (g *glowstone) PreloadItems() {
	if g.Items == nil {
		g.Items = make(map[string]*item.Item)
	}

	files, err := g_util.Walk(path.Join(g.BPDir, "items"))
	if err != nil {
		g.Logger.Error.Println(err)
		return
	}
	for _, file := range files {
		i, err := item.LoadBP(file)
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		if g.Items[i.GetIdentifier()] == nil {
			it := &item.Item{}
			g.Items[i.GetIdentifier()] = it

			// Get basepath
			basepath := filepath.Dir(file)
			subdir := strings.ReplaceAll(basepath, filepath.Join(g.BPDir, "items"), "")
			it.Subdir = subdir

			// Get lang
			full := strings.Split(i.GetIdentifier(), ":")
			namespace := full[0]
			identifier := full[1]

			lang := g.GetLang(fmt.Sprintf("item.%s:%s.name", namespace, identifier))
			it.SetLang(lang)

		}
		g.Items[i.GetIdentifier()].BP = i
	}
}

/******************* Entities *******************/

func (g *glowstone) AddEntity(entities ...interface{}) {
	for _, e := range entities {
		switch e := e.(type) {
		case *entity.Entity:
			g.Entities[e.GetNamespaceIdentifier()] = e
		case entityBP.Entity:
			p := &entity.Entity{
				BP: &e,
			}
			old, ok := g.Entities[p.GetNamespaceIdentifier()]
			if ok {
				old.BP = p.BP
			} else {
				g.Entities[p.GetNamespaceIdentifier()] = p
			}
		case *entityBP.Entity:
			p := &entity.Entity{
				BP: e,
			}
			old, ok := g.Entities[p.GetNamespaceIdentifier()]
			if ok {
				old.BP = p.BP
			} else {
				g.Entities[p.GetNamespaceIdentifier()] = p
			}
		case entityRP.Entity:
			p := &entity.Entity{
				RP: &e,
			}
			old, ok := g.Entities[p.GetNamespaceIdentifier()]
			if ok {
				old.RP = p.RP
			} else {
				g.Entities[p.GetNamespaceIdentifier()] = p
			}
		case *entityRP.Entity:
			p := &entity.Entity{
				RP: e,
			}
			old, ok := g.Entities[p.GetNamespaceIdentifier()]
			if ok {
				old.RP = p.RP
			} else {
				g.Entities[p.GetNamespaceIdentifier()] = p
			}
		default:
			g.Logger.Error.Printf("invalid type %T", e)
		}
	}
}

func (g *glowstone) GetEntities() map[string]*entity.Entity {
	return g.Entities
}

func (g *glowstone) GetEntity(identifier string) (*entity.Entity, error) {
	if e, ok := g.Entities[identifier]; ok {
		return e, nil
	}
	return nil, fmt.Errorf("entity %s not found", identifier)
}

func (g *glowstone) NewEntity(namespace string, identifier string) *entity.Entity {
	e := entity.New(namespace, identifier)
	g.Entities[fmt.Sprintf("%s:%s", namespace, identifier)] = e
	// Set lang
	lang := g_util.TitleCase(strings.ReplaceAll(identifier, "_", " "))
	e.Lang = lang

	entityLang := fmt.Sprintf("entity.%s:%s.name", namespace, identifier)
	if g.GetLang(entityLang) == "" {
		g.AddLang(entityLang, lang)
	}
	spawnEggLang := fmt.Sprintf("item.spawn_egg.entity.%s:%s.name", namespace, identifier)
	if g.GetLang(spawnEggLang) == "" {
		g.AddLang(spawnEggLang, fmt.Sprintf("Spawn %s", lang))
	}
	return e
}

/******************* Block *******************/

func (g *glowstone) AddBlock(blocks ...interface{}) {
	if g.Blocks == nil {
		g.Blocks = make(map[string]*block.Block)
	}
	for _, b := range blocks {
		switch b := b.(type) {
		case *block.Block:
			g.Blocks[b.GetNamespaceIdentifier()] = b
		case blockBP.Block:
			p := &block.Block{
				BP: b,
			}
			g.Blocks[p.GetNamespaceIdentifier()] = p
		default:
			g.Logger.Error.Printf("invalid type %T", b)
		}
	}
}

func (g *glowstone) GetBlocks() map[string]*block.Block {
	return g.Blocks
}

func (g *glowstone) GetBlock(identifier string) (*block.Block, error) {
	if b, ok := g.Blocks[identifier]; ok {
		return b, nil
	}
	return nil, fmt.Errorf("block %s not found", identifier)
}

func (g *glowstone) NewBlock(namespace string, identifier string, subdir ...string) *block.Block {
	b := block.New(namespace, identifier)
	if g.Blocks == nil {
		g.Blocks = make(map[string]*block.Block)
	}
	g.Blocks[fmt.Sprintf("%s:%s", namespace, identifier)] = b
	// Set lang
	lang := g_util.TitleCase(strings.ReplaceAll(identifier, "_", " "))
	b.Lang = lang
	g.AddLang(fmt.Sprintf("block.%s:%s.name", namespace, identifier), b.Lang)
	if len(subdir) > 0 {
		b.Subdir = subdir[0]
	}
	return b
}

/******************* Items *******************/

func (g *glowstone) AddItem(items ...interface{}) {
	for _, i := range items {
		switch i := i.(type) {
		case *item.Item:
			g.Items[i.GetNamespaceIdentifier()] = i
		case itemBP.Item:
			p := &item.Item{
				BP: i,
			}
			g.Items[p.GetNamespaceIdentifier()] = p
		default:
			g.Logger.Error.Printf("invalid type %T", i)
		}
	}
}

func (g *glowstone) GetItems() map[string]*item.Item {
	return g.Items
}

func (g *glowstone) GetItem(identifier string) (*item.Item, error) {
	if i, ok := g.Items[identifier]; ok {
		return i, nil
	}
	return nil, fmt.Errorf("item %s not found", identifier)
}

func (g *glowstone) NewItem(namespace string, identifier string) *item.Item {
	i := item.New(namespace, identifier)
	g.Items[fmt.Sprintf("%s:%s", namespace, identifier)] = i
	// Set lang
	lang := g_util.TitleCase(strings.ReplaceAll(identifier, "_", " "))
	i.Lang = lang
	if g.GetLang(fmt.Sprintf("item.%s:%s.name", namespace, identifier)) == "" {
		g.AddLang(fmt.Sprintf("item.%s:%s.name", namespace, identifier), i.Lang)
	}
	return i
}

/******************* Recipes *******************/

func (g *glowstone) AddRecipe(recipes ...interface{}) {
	for _, r := range recipes {
		switch r := r.(type) {
		case recipe.RecipeInterface:
			g.Recipes[r.GetNamespaceIdentifier()] = r
		case recipeBP.RecipeInterface:
			g.Recipes[r.GetIdentifier()] = r
		default:
			g.Logger.Error.Printf("invalid type %T", r)
		}
	}
}

func (g *glowstone) GetRecipes() map[string]interface{} {
	return g.Recipes
}

func (g *glowstone) GetRecipe(identifier string) (interface{}, error) {
	if r, ok := g.Recipes[identifier]; ok {
		return r, nil
	}
	return nil, fmt.Errorf("recipe %s not found", identifier)
}

func (g *glowstone) NewRecipeBrewingContainer(namespace string, identifier string) *recipe.RecipeBrewingContainer {
	r := recipe.NewBrewingContainer(namespace, identifier)
	g.Recipes[r.GetNamespaceIdentifier()] = r
	return r
}

func (g *glowstone) NewRecipeBrewingMix(namespace string, identifier string) *recipe.RecipeBrewingMix {
	r := recipe.NewBrewingMix(namespace, identifier)
	g.Recipes[r.GetNamespaceIdentifier()] = r
	return r
}

func (g *glowstone) NewRecipeFurnace(namespace string, identifier string) *recipe.RecipeFurnace {
	r := recipe.NewFurnace(namespace, identifier)
	g.Recipes[r.GetNamespaceIdentifier()] = r
	return r
}

func (g *glowstone) NewRecipeShaped(namespace string, identifier string) *recipe.RecipeShaped {
	r := recipe.NewShaped(namespace, identifier)
	g.Recipes[r.GetNamespaceIdentifier()] = r
	return r
}

func (g *glowstone) NewRecipeShapeless(namespace string, identifier string) *recipe.RecipeShapeless {
	r := recipe.NewShapeless(namespace, identifier)
	g.Recipes[r.GetNamespaceIdentifier()] = r
	return r
}

func (g *glowstone) NewRecipeSmithingTransform(namespace string, identifier string) *recipe.RecipeSmithingTransform {
	r := recipe.NewSmithingTransform(namespace, identifier)
	g.Recipes[r.GetNamespaceIdentifier()] = r
	return r
}

/******************* Attachable *******************/

func (g *glowstone) AddAttachable(a attachable.Attachable) {
	if g.Attachables == nil {
		g.Attachables = make(map[string]*attachable.Attachable)
	}
	g.Attachables[a.GetNamespaceIdentifier()] = &a
}

func (g *glowstone) GetAttachables() map[string]*attachable.Attachable {
	return g.Attachables
}

func (g *glowstone) GetAttachable(identifier string) (*attachable.Attachable, error) {
	if a, ok := g.Attachables[identifier]; ok {
		return a, nil
	}
	return nil, fmt.Errorf("attachable %s not found", identifier)
}

func (g *glowstone) NewAttachable(namespace string, identifier string, subdir ...string) *attachable.Attachable {
	a := attachable.New(namespace, identifier)
	if len(subdir) > 0 {
		a.Subdir = subdir[0]
	}
	g.Attachables[fmt.Sprintf("%s:%s", namespace, identifier)] = a
	return a
}

/******************* RPBlocks *******************/

func (g *glowstone) GetRPBlocks() *rp.Blocks {
	if g.RPBlocks == nil {
		g.Logger.Warning.Println("Trying to access nil blocks.json.")
		g.Logger.Warning.Println("Creating new blocks.json file.")
		g.RPBlocks = rp.NewBlocks()
	}
	return g.RPBlocks
}

/******************* ItemTexture *******************/

func (g *glowstone) GetItemTexture() *texture.ItemTexture {
	return g.ItemTexture
}

/******************* TerrainTexture *******************/

func (g *glowstone) GetTerrainTexture() *texture.TerrainTexture {
	if g.TerrainTexture == nil {
		g.Logger.Warning.Println("Trying to access nil terrain_texture.json.")
		g.Logger.Warning.Println("Creating new terrain_texture.json file.")
		g.TerrainTexture = texture.NewTerrainTexture()
	}
	return g.TerrainTexture
}

/******************* Sound Definition *******************/

func (g *glowstone) GetSoundDefinition() *sound.SoundDefinition {
	if g.SoundDefinition == nil {
		r, err := sound.Load(path.Join(g.RPDir, "sounds", "sound_definitions.json"))
		if err != nil {
			g.Logger.Warning.Println("Failed to read sound_definitions.json file, creating new sound_definitions.json file.")
			g.SoundDefinition = sound.New()
		} else {
			g.SoundDefinition = r
		}
	}
	return g.SoundDefinition
}

/******************* BPAnimationController *******************/

func (g *glowstone) AddBPAnimationController(bpAnimationControllers ...interface{}) {
	for _, a := range bpAnimationControllers {
		switch a := a.(type) {
		case *animationController.BPAnimationController:
			g.BPAnimationController[a.Dest] = a
		case animationController.BPAnimationController:
			g.BPAnimationController[a.Dest] = &a
		default:
			g.Logger.Error.Printf("invalid type %T", a)
		}
	}
}

func (g *glowstone) NewBPAnimationController(dest string) *animationController.BPAnimationController {
	a := animationController.NewBP(dest)
	g.BPAnimationController[dest] = a
	return a
}

/******************* RPAnimationController *******************/

func (g *glowstone) AddRPAnimationController(bpAnimationControllers ...interface{}) {
	for _, a := range bpAnimationControllers {
		switch a := a.(type) {
		case *animationController.RPAnimationController:
			g.RPAnimationController[a.Dest] = a
		case animationController.RPAnimationController:
			g.RPAnimationController[a.Dest] = &a
		default:
			g.Logger.Error.Printf("invalid type %T", a)
		}
	}
}

func (g *glowstone) NewRPAnimationController(dest string) *animationController.RPAnimationController {
	a := animationController.NewRP(dest)
	g.RPAnimationController[dest] = a
	return a
}

/******************* BPAnimation *******************/

func (g *glowstone) AddBPAnimation(bpAnimations ...interface{}) {
	for _, a := range bpAnimations {
		switch a := a.(type) {
		case *animation.BPAnimation:
			g.BPAnimation[a.Dest] = a
		case animation.BPAnimation:
			g.BPAnimation[a.Dest] = &a
		default:
			g.Logger.Error.Printf("invalid type %T", a)
		}
	}
}

func (g *glowstone) NewBPAnimation(dest string) *animation.BPAnimation {
	a := animation.NewBP(dest)
	g.BPAnimation[dest] = a
	return a
}
