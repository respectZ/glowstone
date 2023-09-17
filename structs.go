package glowstone

import (
	"log"

	entity "github.com/respectZ/glowstone/entity"
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
	Items       map[string]interface{}    // TODO
	Attachables map[string]interface{}    // TODO

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

	// Entities

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
}
