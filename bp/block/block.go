package block

import (
	"fmt"
	"reflect"

	g_util "github.com/respectZ/glowstone/util"
	comp "github.com/respectZ/glowstone/util/component"
)

const (
	FORMAT_VERSION string = "1.20.30"
)

func New(namespace string, identifier string) Block {
	return &block{
		FormatVersion: FORMAT_VERSION,
		MinecraftBlock: &minecraftBlock{
			Description: &blockDescription{
				Identifier: fmt.Sprintf("%s:%s", namespace, identifier),
			},
			Component: make(map[string]interface{}),
		},
	}
}

func Load(src string) (Block, error) {
	b := &block{}
	err := g_util.LoadJSON(src, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (b *block) Encode() ([]byte, error) {
	return g_util.MarshalJSON(b)
}

func (b *block) GetIdentifier() string {
	return b.MinecraftBlock.Description.Identifier
}

func (b *block) SetIdentifier(identifier string) {
	b.MinecraftBlock.Description.Identifier = identifier
}

func (b *block) GetDescription() BlockDescription {
	return b.MinecraftBlock.Description
}

func (b *block) GetComponent(component interface{}) error {
	componentName := comp.GetComponentName(component)
	if c, ok := b.MinecraftBlock.Component[componentName]; ok {
		// If the type is match, return it
		if reflect.TypeOf(c) == reflect.TypeOf(component) {
			component = c
			return nil
		}
		// Convert map to struct
		err := comp.ConvertMapToStruct(c.(map[string]interface{}), component)
		if err != nil {
			return err
		}
		// Assign component to the struct
		b.MinecraftBlock.Component[componentName] = component
	}
	return fmt.Errorf("component %s not found", componentName)
}

func (b *block) GetPermutations() []BlockPermutation {
	if b.MinecraftBlock.Permutations == nil {
		return nil
	}
	if b.MinecraftBlock.permutations == nil {
		b.MinecraftBlock.permutations = make([]BlockPermutation, 0)
		for _, permutation := range b.MinecraftBlock.Permutations {
			b.MinecraftBlock.permutations = append(b.MinecraftBlock.permutations, permutation)
		}
	}
	return b.MinecraftBlock.permutations
}

func (b *block) AddComponent(components ...interface{}) {
	for _, component := range components {
		componentName := comp.GetComponentName(component)
		// If component is a string, add it, so the result will "component": {}
		c := reflect.TypeOf(component)
		if c.Kind() == reflect.String && c.Name() == "string" {
			b.MinecraftBlock.Component[componentName] = struct{}{}
		} else {
			b.MinecraftBlock.Component[componentName] = component
		}
	}
}

func (b *block) AddPermutation(condition string, components ...interface{}) {
	if b.MinecraftBlock.Permutations == nil {
		b.MinecraftBlock.Permutations = make([]*blockPermutation, 0)
	}
	permutation := &blockPermutation{
		Condition:  condition,
		Components: make(map[string]interface{}),
	}
	for _, component := range components {
		componentName := comp.GetComponentName(component)
		permutation.Components[componentName] = component
	}
	// Add to struct
	b.MinecraftBlock.Permutations = append(b.MinecraftBlock.Permutations, permutation)
	// Add to interface
	if b.MinecraftBlock.permutations == nil {
		b.MinecraftBlock.permutations = make([]BlockPermutation, 0)
	}
	b.MinecraftBlock.permutations = append(b.MinecraftBlock.permutations, permutation)
}
