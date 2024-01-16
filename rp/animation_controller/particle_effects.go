package animation_controller

import g_util "github.com/respectZ/glowstone/util"

type ParticleEffects []*particleEffect

type IParticleEffects interface {
	Add(string) *particleEffect
	All() []*particleEffect
	Clear()
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error
}

func (p *ParticleEffects) Add(effect string) *particleEffect {
	if *p == nil {
		*p = make(ParticleEffects, 0)
	}
	e := &particleEffect{
		Effect: effect,
	}
	*p = append(*p, e)
	return e
}

func (p *ParticleEffects) All() []*particleEffect {
	if *p == nil {
		*p = make(ParticleEffects, 0)
	}
	return *p
}

func (p *ParticleEffects) Clear() {
	*p = make(ParticleEffects, 0)
}

func (p *ParticleEffects) IsEmpty() bool {
	if *p == nil {
		*p = make(ParticleEffects, 0)
	}
	return len(*p) == 0
}

func (p *ParticleEffects) Size() int {
	if *p == nil {
		*p = make(ParticleEffects, 0)
	}
	return len(*p)
}

func (p *ParticleEffects) UnmarshalJSON(data []byte) error {
	var temp []*particleEffect
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*p = temp
	return nil
}
