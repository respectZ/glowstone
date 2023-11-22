package animation_controller

func (a *animationControllerState) AddParticleEffect(effect string, locator string, preEffectScript string, bindToActor bool) {
	if a.ParticleEffects == nil {
		a.ParticleEffects = make([]*particleEffect, 0)
	}
	a.ParticleEffects = append(a.ParticleEffects, &particleEffect{
		Effect:          effect,
		Locator:         locator,
		PreEffectScript: preEffectScript,
		BindToActor:     &bindToActor,
	})
}

func (a *animationControllerState) AddSoundEffect(effect string) {
	if a.SoundEffects == nil {
		a.SoundEffects = make([]*soundEffect, 0)
	}
	a.SoundEffects = append(a.SoundEffects, &soundEffect{
		Effect: effect,
	})
}

func (a *animationControllerState) AddTransition(to string, trigger string) {
	a.Transitions = append(a.Transitions, map[string]string{
		to: trigger,
	})
}

func (a *animationControllerState) AddOnEntry(trigger string) {
	if a.OnEntry == nil {
		a.OnEntry = make([]string, 0)
	}
	a.OnEntry = append(a.OnEntry, trigger)
}

func (a *animationControllerState) AddOnExit(trigger string) {
	if a.OnExit == nil {
		a.OnExit = make([]string, 0)
	}
	a.OnExit = append(a.OnExit, trigger)
}

func (a *animationControllerState) AddAnimation(animation string, condition ...string) {
	if a.Animations == nil {
		a.Animations = make([]interface{}, 0)
	}
	if len(condition) > 0 {
		a.Animations = append(a.Animations, map[string]interface{}{
			animation: condition[0],
		})
	} else {
		a.Animations = append(a.Animations, animation)
	}
}

func (a *animationControllerState) SetBlendTransition(transition float64, viaShortestPath bool) {
	a.BlendTransition = transition
	a.BlendViaShortestPath = &viaShortestPath
}

func (a *animationControllerState) AddVariable(name string, input string, remapCurve map[string]float64) {
	if a.Variables == nil {
		a.Variables = make(map[string]*variables)
	}
	a.Variables[name] = &variables{
		Input:      input,
		RemapCurve: remapCurve,
	}
}
