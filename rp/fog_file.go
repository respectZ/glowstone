package rp

import (
	"strings"

	rp "github.com/respectZ/glowstone/rp/fog"
)

type FogFile struct {
	Data *rp.Fog

	Subdir   string
	Filename string
}

func LoadFog(src string) (*rp.Fog, error) {
	data, err := rp.Load(src)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f *FogFile) Encode() ([]byte, error) {
	fog, err := f.Data.Encode()
	if err != nil {
		return nil, err
	}
	return fog, nil
}

// GetIdentifier returns the identifier of the fog without namespace
//
// Example:
//
//	identifier := e.GetIdentifier()
func (f *FogFile) GetIdentifier() string {
	return strings.Split(f.Data.FogSettings.Description.Identifier, ":")[1]
}

// GetNamespaceIdentifier returns the identifier of the fog with namespace
//
// Example:
//
//	identifier := e.GetNamespaceIdentifier()
func (f *FogFile) GetNamespaceIdentifier() string {
	return f.Data.FogSettings.Description.Identifier
}

func (f *FogFile) GetFilename() string {
	if f.Filename == "" {
		return f.GetIdentifier() + ".json"
	}
	return f.Filename
}
