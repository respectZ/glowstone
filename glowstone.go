package glowstone

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	animation "github.com/respectZ/glowstone/animation"
	"github.com/respectZ/glowstone/entity"
	item "github.com/respectZ/glowstone/item"
	recipe "github.com/respectZ/glowstone/recipe"
	g_util "github.com/respectZ/glowstone/util"

	entityBP "github.com/respectZ/glowstone/bp/entity"
	itemBP "github.com/respectZ/glowstone/bp/item"
	recipeBP "github.com/respectZ/glowstone/bp/recipe"
	entityRP "github.com/respectZ/glowstone/rp/entity"
	sound "github.com/respectZ/glowstone/rp/sound"
	texture "github.com/respectZ/glowstone/rp/texture"

	entityBPComponent "github.com/respectZ/glowstone/bp/entity/component"
)

var MIN_ENGINE_VERSION = [3]int{1, 20, 0}

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

		Entities:    make(map[string]*entity.Entity),
		Items:       make(map[string]*item.Item),
		BPAnimation: make(map[string]*animation.BPAnimation),
		Recipes:     make(map[string]interface{}),

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
	// Read ItemTexture
	g.ItemTexture, err = texture.Load(g.RPDir + "textures/item_texture.json")
	if err != nil {
		g.Logger.Warning.Println("Failed to read item_texture.json file, creating new item_texture.json file.")
		g.ItemTexture = texture.New()
	}
	// TODO: Upfront here.
	return nil
}

func (g *glowstone) SetUpfront(upfront bool) {
	g.IsUpfront = upfront
}

func (g *glowstone) Save() {
	// ItemTexture
	data, err := g.ItemTexture.Encode()
	if err != nil {
		g.Logger.Error.Println(err)
	} else {
		g_util.Writefile(path.Join(g.RPDir, "textures", "item_texture.json"), data)
	}
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
		bp, rp, err := e.Encode()
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		if bp != nil {
			g_util.Writefile(path.Join(g.BPDir, "entities", e.Subdir, fmt.Sprintf("%s.json", e.GetIdentifier())), bp)
			if e.RideHint == "" {
				// Check if rideable
				var ridebable entityBPComponent.Rideable
				_, err := e.BP.GetComponent(&ridebable)
				if err == nil {
					// Add to lang
					e.SetRideHint(fmt.Sprintf("Tap jump to exit the %s", e.Lang))
				}
			}
		}
		if rp != nil {
			g_util.Writefile(path.Join(g.RPDir, "entity", e.Subdir, fmt.Sprintf("%s.json", e.GetIdentifier())), rp)
		}
	}

	// Item
	for _, i := range g.Items {
		bp, err := i.Encode()
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		g_util.Writefile(path.Join(g.BPDir, "items", i.Subdir, fmt.Sprintf("%s.json", i.GetIdentifier())), bp)
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
	g.AddLang(fmt.Sprintf("entity.%s:%s.name", namespace, identifier), e.Lang)
	g.AddLang(fmt.Sprintf("item.spawn_egg.entity.%s:%s.name", namespace, identifier), fmt.Sprintf("Spawn %s", lang))
	return e
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
	g.AddLang(fmt.Sprintf("item.%s:%s.name", namespace, identifier), i.Lang)
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

/******************* ItemTexture *******************/

func (g *glowstone) GetItemTexture() *texture.ItemTexture {
	return g.ItemTexture
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
