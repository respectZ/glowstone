package entity

import (
	g_util "github.com/respectZ/glowstone/util"
)

type Geometry struct {
	FormatVersion string             `json:"format_version"`
	Geometry      IMinecraftGeometry `json:"minecraft:geometry"`
}

type Description struct {
	Identifier          string    `json:"identifier"`
	TextureWidth        int       `json:"texture_width,omitempty"`
	TextureHeight       int       `json:"texture_height,omitempty"`
	VisibleBoundsWidth  float64   `json:"visible_bounds_width,omitempty"`
	VisibleBoundsHeight float64   `json:"visible_bounds_height,omitempty"`
	VisibleBoundsOffset []float64 `json:"visible_bounds_offset,omitempty"`
}

type TextureMesh struct {
	Position   []float64 `json:"position"`
	Rotation   []float64 `json:"rotation"`
	Texture    string    `json:"texture"`
	LocalPivot []float64 `json:"local_pivot,omitempty"`
}

func New(identifier string) *Geometry {
	p := &Geometry{
		FormatVersion: "1.16.0",
		Geometry:      &MinecraftGeometries{},
	}
	p.Geometry.New(identifier)
	return p
}

func Load(src string) (*Geometry, error) {
	p := New("")
	err := g_util.LoadJSON(src, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Geometry) Encode() ([]byte, error) {
	return g_util.MarshalJSON(p)
}

func (p *Geometry) GetIdentifier() string {
	if p.Geometry == nil {
		return ""
	}
	if p.Geometry.IsEmpty() {
		return ""
	}
	return p.Geometry.All()[0].Description.Identifier
}

// Todo: older versions of geometry
