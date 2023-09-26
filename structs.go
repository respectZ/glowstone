package glowstone

import (
	"log"

	entity "github.com/respectZ/glowstone/entity"
	item "github.com/respectZ/glowstone/item"
	texture "github.com/respectZ/glowstone/rp/texture"
)

type glowstone struct {
	ProjectName      string
	Namespace        string
	BPDir            string
	RPDir            string
	MinEngineVersion [3]int
	Logger           *logger

	// Data
	Lang        map[string]string
	Entities    map[string]*entity.Entity // TODO
	Items       map[string]*item.Item     // TODO
	Attachables map[string]interface{}    // TODO
	ItemTexture *texture.ItemTexture

	// Settings
	IsUpfront bool
}

type logger struct {
	Warning *log.Logger
	Error   *log.Logger
}

type Glowstone interface {
	// GetProjectName returns the name of the project
	GetProjectName() string
	// SetProjectName sets the name of the project
	SetProjectName(string)
	// GetNamespace returns the namespace of the project
	GetNamespace() string
	// SetNamespace sets the namespace of the project
	SetNamespace(string)
	// GetBPDir returns the path to the blueprint directory
	GetBPDir() string
	// SetBPDir sets the path to the blueprint directory
	SetBPDir(string)
	// GetRPDir returns the path to the resourcepack directory
	GetRPDir() string
	// SetRPDir sets the path to the resourcepack directory
	SetRPDir(string)
	// GetMinEngineVersion returns the minimum engine version
	GetMinEngineVersion() [3]int
	// SetMinEngineVersion sets the minimum engine version
	SetMinEngineVersion([3]int)
	// GetLang returns the language data
	GetLang(string) string
	// SetLang sets the language data
	SetLang(map[string]string)
	// AddLang adds the language data
	AddLang(string, string)

	// Build builds the project
	Build() error
	// Initialize
	Initialize() error

	// SetUpfront will cache the project data upfront.
	//
	// Default: false
	SetUpfront(bool)

	// Save the project data to the disk.
	Save()

	/******************* Entities *******************/

	// AddEntity adds the entity to the project
	//
	// Example:
	// 	glowstone.AddEntity(entity)
	AddEntity(...interface{})

	// GetEntities returns the entities
	GetEntities() map[string]*entity.Entity
	// GetEntity returns the entity
	//
	// Example:
	// 	entity, err := glowstone.GetEntity("minecraft:zombie")
	GetEntity(string) (*entity.Entity, error)

	// NewEntity creates a new entity
	//
	// Example:
	// 	entity := glowstone.NewEntity("minecraft", "zombie")
	NewEntity(string, string) *entity.Entity

	/******************* Items *******************/

	// AddItem adds the item to the project
	//
	// Example:
	// 	glowstone.AddItem(item)
	AddItem(...interface{})

	// GetItems returns the items
	GetItems() map[string]*item.Item

	// GetItem returns the item
	//
	// Example:
	// 	item, err := glowstone.GetItem("minecraft:stick")
	GetItem(string) (*item.Item, error)

	// NewItem creates a new item
	//
	// Example:
	// 	item := glowstone.NewItem("minecraft", "stick")
	NewItem(string, string) *item.Item

	/******************* Item Texture *******************/

	// GetItemTexture returns the item_texture.json
	GetItemTexture() *texture.ItemTexture
}
