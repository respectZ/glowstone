package glowstone

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/respectZ/glowstone/entity"
	item "github.com/respectZ/glowstone/item"
	g_util "github.com/respectZ/glowstone/util"

	entityBP "github.com/respectZ/glowstone/bp/entity"
	itemBP "github.com/respectZ/glowstone/bp/item"
	entityRP "github.com/respectZ/glowstone/rp/entity"
)

var MIN_ENGINE_VERSION = [3]int{1, 20, 0}

func NewProject(ProjectName string, Namespace string) Glowstone {
	return &glowstone{
		ProjectName:      ProjectName,
		Namespace:        Namespace,
		BPDir:            "./packs/BP/",
		RPDir:            "./packs/RP/",
		MinEngineVersion: MIN_ENGINE_VERSION,

		Logger: &logger{
			Warning: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
			Error:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		},
		Lang:      make(map[string]string),
		IsUpfront: false,
	}
}

func (g *glowstone) GetProjectName() string {
	return g.ProjectName
}

func (g *glowstone) SetProjectName(name string) {
	g.ProjectName = name
}

func (g *glowstone) GetNamespace() string {
	return g.Namespace
}

func (g *glowstone) SetNamespace(namespace string) {
	g.Namespace = namespace
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
	return g.Lang[key]
}

func (g *glowstone) SetLang(data map[string]string) {
	g.Lang = data
}

func (g *glowstone) Build() error {
	g_util.Writelang(g.RPDir+"texts/en_US.lang", g.Lang)
	return nil
}

func (g *glowstone) Initialize() error {
	// Read Lang
	data, err := g_util.Loadlang(g.RPDir + "texts/en_US.lang")
	if err != nil {
		g.Logger.Warning.Println("Failed to read lang file, creting new lang file.")
		g.SetLang(make(map[string]string))
		g_util.Writefile(g.RPDir+"texts/en_US.lang", []byte{})
	} else {
		g.SetLang(data)
	}
	// TODO: Upfront here.
	return nil
}

func (g *glowstone) SetUpfront(upfront bool) {
	g.IsUpfront = upfront
}

func (g *glowstone) Save() {
	// Entity
	for _, e := range g.Entities {
		bp, rp, err := e.Encode()
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		g_util.Writefile(path.Join(g.BPDir, "entities", e.Subdir, fmt.Sprintf("%s.json", e.GetIdentifier())), bp)
		g_util.Writefile(path.Join(g.RPDir, "entity", e.Subdir, fmt.Sprintf("%s.json", e.GetIdentifier())), rp)
		// Lang
		full := strings.Split(e.GetNamespaceIdentifier(), ":")
		namespace := full[0]
		identifier := full[1]

		if e.Lang == "" {
			lang := g_util.TitleCase(strings.ReplaceAll(identifier, "_", " "))
			e.Lang = lang
		}
		g.Lang[fmt.Sprintf("item.%s:%s.name", namespace, identifier)] = e.Lang
	}

	// Item
	for _, i := range g.Items {
		bp, err := i.Encode()
		if err != nil {
			g.Logger.Error.Println(err)
			continue
		}
		g_util.Writefile(path.Join(g.BPDir, "items", i.Subdir, fmt.Sprintf("%s.json", i.GetIdentifier())), bp)
		// Lang
		full := strings.Split(i.GetNamespaceIdentifier(), ":")
		namespace := full[0]
		identifier := full[1]

		if i.Lang == "" {
			lang := g_util.TitleCase(strings.ReplaceAll(identifier, "_", " "))
			i.Lang = lang
		}
		g.Lang[fmt.Sprintf("item.%s:%s.name", namespace, identifier)] = i.Lang
	}
}

/******************* Upfronts *******************/

func (g *glowstone) cacheEntities() {
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
			g.Logger.Error.Printf("entity %s not found in BP", e.GetIdentifier())
			continue
		}
		g.Entities[e.GetIdentifier()].RP = e
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
	return i
}
