package particle

import (
	g_util "github.com/respectZ/glowstone/util"
)

type Particle struct {
	FormatVersion  string          `json:"format_version"`
	ParticleEffect *ParticleEffect `json:"particle_effect"`
}

type ParticleEffect struct {
	Description *Description           `json:"description"`
	Curves      *Curves                `json:"curves,omitempty"`
	Events      interface{}            `json:"events,omitempty"`     // TODO: parse events
	Components  map[string]interface{} `json:"components,omitempty"` // TODO: add components
}

type Description struct {
	Identifier            string                 `json:"identifier"`
	BasicRenderParameters *BasicRenderParameters `json:"basic_render_parameters"`
}

type BasicRenderParameters struct {
	Material string `json:"material"`
	Texture  string `json:"texture"`
}

func New(identifier string) *Particle {
	return &Particle{
		FormatVersion: "1.20.0",
		ParticleEffect: &ParticleEffect{
			Description: &Description{
				Identifier:            identifier,
				BasicRenderParameters: &BasicRenderParameters{},
			},
			Curves: &Curves{},
		},
	}
}

func Load(src string) (*Particle, error) {
	a := New("")
	err := g_util.LoadJSON(src, a)
	return a, err
}

func (a *Particle) Encode(minify ...bool) ([]byte, error) {
	if a.ParticleEffect.Curves.IsEmpty() {
		a.ParticleEffect.Curves = nil
	}
	return g_util.MarshalJSON(a, minify...)
}

func (a *Particle) GetIdentifier() string {
	return a.ParticleEffect.Description.Identifier
}
