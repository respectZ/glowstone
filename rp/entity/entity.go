package entity

import (
	"reflect"

	g_util "github.com/respectZ/glowstone/util"
)

func (e *Entity) Encode(minify ...bool) ([]byte, error) {
	if e.Entity.Description.Scripts != nil {
		if e.Entity.Description.Scripts.Animate.IsEmpty() {
			e.Entity.Description.Scripts.Animate = nil
		}
		if e.Entity.Description.Scripts.PreAnimation.IsEmpty() {
			e.Entity.Description.Scripts.PreAnimation = nil
		}
		if e.Entity.Description.Scripts.Initialize.IsEmpty() {
			e.Entity.Description.Scripts.Initialize = nil
		}
		if e.Entity.Description.Scripts.Variables.IsEmpty() {
			e.Entity.Description.Scripts.Variables = nil
		}
	}

	if e.Entity.Description.Animations.IsEmpty() {
		e.Entity.Description.Animations = nil
	}
	if e.Entity.Description.SoundEffects.IsEmpty() {
		e.Entity.Description.SoundEffects = nil
	}
	if e.Entity.Description.ParticleEffects.IsEmpty() {
		e.Entity.Description.ParticleEffects = nil
	}
	if e.Entity.Description.ParticleEmitters.IsEmpty() {
		e.Entity.Description.ParticleEmitters = nil
	}

	if reflect.ValueOf(e.Entity.Description.Scripts).Elem().IsZero() {
		e.Entity.Description.Scripts = nil
	}
	if reflect.ValueOf(e.Entity.Description.SpawnEgg).Elem().IsZero() {
		e.Entity.Description.SpawnEgg = nil
	}
	return g_util.MarshalJSON(e, minify...)
}

func (e *Entity) GetIdentifier() string {
	return e.Entity.Description.Identifier
}
