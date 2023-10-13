package entity

func (e *clientEntityDescriptionScripts) GetParentSetup() string {
	return e.ParentSetup
}

func (e *clientEntityDescriptionScripts) SetParentSetup(parentSetup string) {
	e.ParentSetup = parentSetup
}

func (e *clientEntityDescriptionScripts) GetVariables() map[string]string {
	return e.Variables
}

func (e *clientEntityDescriptionScripts) SetVariables(variables map[string]string) {
	e.Variables = variables
}

func (e *clientEntityDescriptionScripts) AddVariable(name string, value string) {
	if e.Variables == nil {
		e.Variables = make(map[string]string)
	}
	e.Variables[name] = value
}

func (e *clientEntityDescriptionScripts) GetScaleX() string {
	return e.ScaleX
}

func (e *clientEntityDescriptionScripts) SetScaleX(scaleX string) {
	e.ScaleX = scaleX
}

func (e *clientEntityDescriptionScripts) GetScaleY() string {
	return e.ScaleY
}

func (e *clientEntityDescriptionScripts) SetScaleY(scaleY string) {
	e.ScaleY = scaleY
}

func (e *clientEntityDescriptionScripts) GetScaleZ() string {
	return e.ScaleZ
}

func (e *clientEntityDescriptionScripts) SetScaleZ(scaleZ string) {
	e.ScaleZ = scaleZ
}

func (e *clientEntityDescriptionScripts) GetScale() string {
	return e.Scale
}

func (e *clientEntityDescriptionScripts) SetScale(scale string) {
	e.Scale = scale
}

func (e *clientEntityDescriptionScripts) GetInitialize() []string {
	return e.Initialize
}

func (e *clientEntityDescriptionScripts) SetInitialize(initialize []string) {
	e.Initialize = initialize
}

func (e *clientEntityDescriptionScripts) AddInitialize(initialize string) {
	if e.Initialize == nil {
		e.Initialize = make([]string, 0)
	}
	e.Initialize = append(e.Initialize, initialize)
}

func (e *clientEntityDescriptionScripts) GetPreAnimation() []string {
	return e.PreAnimation
}

func (e *clientEntityDescriptionScripts) SetPreAnimation(preAnimation []string) {
	e.PreAnimation = preAnimation
}

func (e *clientEntityDescriptionScripts) AddPreAnimation(preAnimation string) {
	if e.PreAnimation == nil {
		e.PreAnimation = make([]string, 0)
	}
	e.PreAnimation = append(e.PreAnimation, preAnimation)
}

func (e *clientEntityDescriptionScripts) GetAnimate() []interface{} {
	return e.Animate
}

func (e *clientEntityDescriptionScripts) SetAnimate(animate []interface{}) {
	e.Animate = animate
}

func (e *clientEntityDescriptionScripts) AddAnimate(animate interface{}) {
	if e.Animate == nil {
		e.Animate = make([]interface{}, 0)
	}
	e.Animate = append(e.Animate, animate)
}
