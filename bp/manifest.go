package bp

import (
	"path/filepath"

	"github.com/google/uuid"
	g_util "github.com/respectZ/glowstone/util"
)

type Manifest struct {
	FormatVersion int                  `json:"format_version"`
	Header        *ManifestHeader      `json:"header,omitempty"`
	Modules       ManifestModules      `json:"modules,omitempty"`
	Dependencies  ManifestDependencies `json:"dependencies,omitempty"`
}

type ManifestHeader struct {
	Name             string `json:"name,omitempty"`
	Uuid             string `json:"uuid,omitempty"`
	Description      string `json:"description,omitempty"`
	Version          []int  `json:"version,omitempty"`
	MinEngineVersion []int  `json:"min_engine_version,omitempty"`
}

type ManifestModule struct {
	Type        string      `json:"type,omitempty"`
	Uuid        string      `json:"uuid,omitempty"`
	Description string      `json:"description,omitempty"`
	Version     interface{} `json:"version,omitempty"`
	Entry       string      `json:"entry,omitempty"`
	Language    string      `json:"language,omitempty"`
}

type ManifestDependency struct {
	Uuid       string      `json:"uuid,omitempty"`
	ModuleName string      `json:"module_name,omitempty"`
	Version    interface{} `json:"version,omitempty"`
}

type (
	ManifestModules      []*ManifestModule
	ManifestDependencies []*ManifestDependency
)

type IManifestModules interface {
	Add(...*ManifestModule)
	All() []*ManifestModule
	Clear()
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// NewBP creates a new Module for Behavior Pack.
	//
	// Example:
	//
	//	a := NewBP("My Behavior Pack", []int{1, 0, 0})
	NewBP(string, []int) *ManifestModule

	// NewScript creates a new Module for Script.
	//
	// Example:
	//
	//	a := NewScript("scripts/main.js", []int{1, 0, 0})
	NewScript(string, []int) *ManifestModule

	GetBPModules() []*ManifestModule
	GetScriptModules() []*ManifestModule

	Save(string) error
}

type IManifestDependencies interface {
	Add(...*ManifestDependency)
	All() []*ManifestDependency
	Clear()
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	New(string) *ManifestDependency
}

func (m *ManifestModules) Add(values ...*ManifestModule) {
	if *m == nil {
		*m = make([]*ManifestModule, 0)
	}
	*m = append(*m, values...)
}

func (m *ManifestModules) All() []*ManifestModule {
	if *m == nil {
		*m = make([]*ManifestModule, 0)
	}
	return *m
}

func (m *ManifestModules) Clear() {
	if *m == nil {
		*m = make([]*ManifestModule, 0)
	}
	*m = make([]*ManifestModule, 0)
}

func (m *ManifestModules) IsEmpty() bool {
	if *m == nil {
		*m = make([]*ManifestModule, 0)
	}
	return len(*m) == 0
}

func (m *ManifestModules) Size() int {
	if *m == nil {
		*m = make([]*ManifestModule, 0)
	}
	return len(*m)
}

func (m *ManifestModules) UnmarshalJSON(data []byte) error {
	var temp []*ManifestModule
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*m = temp
	return nil
}

func (m *ManifestModules) NewBP(description string, version []int) *ManifestModule {
	return &ManifestModule{
		Uuid:        uuid.NewString(),
		Type:        "data",
		Description: description,
		Version:     version,
	}
}

func (m *ManifestModules) NewScript(entry string, version []int) *ManifestModule {
	return &ManifestModule{
		Uuid:     uuid.NewString(),
		Type:     "script",
		Language: "javascript",
		Version:  version,
		Entry:    entry,
	}
}

func (m *ManifestModules) GetBPModules() []*ManifestModule {
	if *m == nil {
		*m = make([]*ManifestModule, 0)
	}
	var modules []*ManifestModule
	for _, module := range *m {
		if module.Type == "data" {
			modules = append(modules, module)
		}
	}
	return modules
}

func (m *ManifestModules) GetScriptModules() []*ManifestModule {
	if *m == nil {
		*m = make([]*ManifestModule, 0)
	}
	var modules []*ManifestModule
	for _, module := range *m {
		if module.Type == "script" {
			modules = append(modules, module)
		}
	}
	return modules
}

func (m *ManifestDependencies) Add(values ...*ManifestDependency) {
	if *m == nil {
		*m = make([]*ManifestDependency, 0)
	}
	*m = append(*m, values...)
}

func (m *ManifestDependencies) All() []*ManifestDependency {
	if *m == nil {
		*m = make([]*ManifestDependency, 0)
	}
	return *m
}

func (m *ManifestDependencies) Clear() {
	if *m == nil {
		*m = make([]*ManifestDependency, 0)
	}
	*m = make([]*ManifestDependency, 0)
}

func (m *ManifestDependencies) IsEmpty() bool {
	if *m == nil {
		*m = make([]*ManifestDependency, 0)
	}
	return len(*m) == 0
}

func (m *ManifestDependencies) Size() int {
	if *m == nil {
		*m = make([]*ManifestDependency, 0)
	}
	return len(*m)
}

func (m *ManifestDependencies) UnmarshalJSON(data []byte) error {
	var temp []*ManifestDependency
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*m = temp
	return nil
}

// Create a new dependency module
//
// name: The name of the module.
//
// return: The new dependency
//
// Example:
//
//	dependency := New("@minecraft/server", "1.7.0")
func (m *ManifestDependencies) NewModule(name string, version string) *ManifestDependency {
	return &ManifestDependency{
		Uuid:       uuid.NewString(),
		ModuleName: name,
		Version:    version,
	}
}

func NewManifest() *Manifest {
	return &Manifest{
		Header:       &ManifestHeader{},
		Dependencies: []*ManifestDependency{},
		Modules:      []*ManifestModule{},
	}
}

func LoadManifest(src string) (*Manifest, error) {
	manifest := NewManifest()
	err := g_util.LoadJSON(src, &manifest)
	return manifest, err
}

func (e *Manifest) Encode() ([]byte, error) {
	if e.Dependencies.IsEmpty() {
		e.Dependencies = nil
	}
	if e.Modules.IsEmpty() {
		e.Modules = nil
	}
	return g_util.MarshalJSON(e)
}

func (e *Manifest) Save(pathToBP string) error {
	data, err := e.Encode()
	if err != nil {
		return err
	}
	err = g_util.Writefile(filepath.Join(pathToBP, "manifest.json"), data)
	if err != nil {
		return err
	}
	return nil
}

func (e *Manifest) Load(src string) error {
	if filepath.Base(src) != "manifest.json" {
		src = filepath.Join(src, "manifest.json")
	}
	return g_util.LoadJSON(src, e)
}

func (e *Manifest) IsEmpty() bool {
	return e.FormatVersion == 0 && e.Header == nil && e.Dependencies.IsEmpty() && e.Modules.IsEmpty()
}
