package animation_controller

import g_util "github.com/respectZ/glowstone/util"

type SoundEffects []*soundEffect

type ISoundEffects interface {
	Add(...*soundEffect)
	All() []*soundEffect
	Clear()
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error
}

func (p *SoundEffects) Add(s ...*soundEffect) {
	if *p == nil {
		*p = make(SoundEffects, 0)
	}
	*p = append(*p, s...)
}

func (p *SoundEffects) All() []*soundEffect {
	if *p == nil {
		*p = make(SoundEffects, 0)
	}
	return *p
}

func (p *SoundEffects) Clear() {
	*p = make(SoundEffects, 0)
}

func (p *SoundEffects) IsEmpty() bool {
	if *p == nil {
		*p = make(SoundEffects, 0)
	}
	return len(*p) == 0
}

func (p *SoundEffects) Size() int {
	if *p == nil {
		*p = make(SoundEffects, 0)
	}
	return len(*p)
}

func (p *SoundEffects) UnmarshalJSON(data []byte) error {
	var temp []*soundEffect
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*p = temp
	return nil
}
