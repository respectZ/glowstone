package glowstone

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/recipe"
	"github.com/respectZ/glowstone/recipe"
	g_util "github.com/respectZ/glowstone/util"
)

type RecipeBP map[string]interface{}

type IRecipeBP interface {
	Add(string, interface{})
	Get(string) (interface{}, bool)
	Remove(string)
	Clear()
	All() map[string]interface{}
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	Save(string) error

	// NewBrewingContainer creates a new recipe for a brewing container.
	//
	// Example:
	//
	//	a := New("glowstone:chair")
	//	b := New("glowstone:table")
	NewBrewingContainer(string, ...struct{ Subdir string }) *bp.RecipeBrewingContainer

	// NewBrewingMix creates a new recipe for a brewing container with a mix.
	//
	// Example:
	//
	//	a := New("glowstone:chair")
	//	b := New("glowstone:table")
	NewBrewingMix(string, ...struct{ Subdir string }) *bp.RecipeBrewingMix

	// NewFurnace creates a new recipe to smelt items in a furnace.
	//
	// Example:
	//
	//	a := New("glowstone:chair")
	//	b := New("glowstone:table")
	NewFurnace(string, ...struct{ Subdir string }) *bp.RecipeFurnace

	// NewShaped creates a new crafting recipe with a shape.
	//
	// Example:
	//
	//	a := New("glowstone:chair")
	//	b := New("glowstone:table")
	NewShaped(string, ...struct{ Subdir string }) *bp.RecipeShaped

	// NewShapeless creates a new crafting recipe without a shape.
	//
	// Example:
	//
	//	a := New("glowstone:chair")
	//	b := New("glowstone:table")
	NewShapeless(string, ...struct{ Subdir string }) *bp.RecipeShapeless

	// NewSmithingTransform creates a new smithing recipe.
	//
	// Example:
	//
	//	a := New("glowstone:chair")
	//	b := New("glowstone:table")
	NewSmithingTransform(string, ...struct{ Subdir string }) *bp.RecipeSmithingTransform

	LoadAll(string) error
}

func (e *RecipeBP) Add(key string, value interface{}) {
	if *e == nil {
		*e = make(map[string]interface{})
	}
	(*e)[key] = value
}

func (e *RecipeBP) Get(key string) (interface{}, bool) {
	if *e == nil {
		*e = make(map[string]interface{})
	}
	value, ok := (*e)[key]
	return value, ok
}

func (e *RecipeBP) Remove(key string) {
	if *e == nil {
		*e = make(map[string]interface{})
	}
	delete(*e, key)
}

func (e *RecipeBP) Clear() {
	if *e == nil {
		*e = make(map[string]interface{})
	}
	*e = make(map[string]interface{})
}

func (e *RecipeBP) All() map[string]interface{} {
	if *e == nil {
		*e = make(map[string]interface{})
	}
	return *e
}

func (e *RecipeBP) IsEmpty() bool {
	if *e == nil {
		*e = make(map[string]interface{})
	}
	return len(*e) == 0
}

func (e *RecipeBP) Size() int {
	if *e == nil {
		*e = make(map[string]interface{})
	}
	return len(*e)
}

func (e *RecipeBP) UnmarshalJSON(data []byte) error {
	if *e == nil {
		*e = make(map[string]interface{})
	}
	return g_util.UnmarshalJSON(data, e)
}

func (e *RecipeBP) NewBrewingContainer(identifier string, option ...struct{ Subdir string }) *bp.RecipeBrewingContainer {
	f := strings.Split(identifier, ":")
	namespace := f[0]
	identifier = f[1]

	p := recipe.NewBrewingContainer(namespace, identifier)
	e.Add(fmt.Sprintf("%s:%s", namespace, identifier), &p)

	if len(option) > 0 {
		p.Subdir = option[0].Subdir
	}

	return &p.Data
}

func (e *RecipeBP) NewBrewingMix(identifier string, option ...struct{ Subdir string }) *bp.RecipeBrewingMix {
	f := strings.Split(identifier, ":")
	namespace := f[0]
	identifier = f[1]

	p := recipe.NewBrewingMix(namespace, identifier)
	e.Add(fmt.Sprintf("%s:%s", namespace, identifier), &p)

	if len(option) > 0 {
		p.Subdir = option[0].Subdir
	}

	return &p.Data
}

func (e *RecipeBP) NewFurnace(identifier string, option ...struct{ Subdir string }) *bp.RecipeFurnace {
	f := strings.Split(identifier, ":")
	namespace := f[0]
	identifier = f[1]

	p := recipe.NewFurnace(namespace, identifier)
	e.Add(fmt.Sprintf("%s:%s", namespace, identifier), &p)

	if len(option) > 0 {
		p.Subdir = option[0].Subdir
	}

	return &p.Data
}

func (e *RecipeBP) NewShaped(identifier string, option ...struct{ Subdir string }) *bp.RecipeShaped {
	f := strings.Split(identifier, ":")
	namespace := f[0]
	identifier = f[1]

	p := recipe.NewShaped(namespace, identifier)
	e.Add(fmt.Sprintf("%s:%s", namespace, identifier), &p)

	if len(option) > 0 {
		p.Subdir = option[0].Subdir
	}

	return &p.Data
}

func (e *RecipeBP) NewShapeless(identifier string, option ...struct{ Subdir string }) *bp.RecipeShapeless {
	f := strings.Split(identifier, ":")
	namespace := f[0]
	identifier = f[1]

	p := recipe.NewShapeless(namespace, identifier)
	e.Add(fmt.Sprintf("%s:%s", namespace, identifier), &p)

	if len(option) > 0 {
		p.Subdir = option[0].Subdir
	}

	return &p.Data
}

func (e *RecipeBP) NewSmithingTransform(identifier string, option ...struct{ Subdir string }) *bp.RecipeSmithingTransform {
	f := strings.Split(identifier, ":")
	namespace := f[0]
	identifier = f[1]

	p := recipe.NewSmithingTransform(namespace, identifier)
	e.Add(fmt.Sprintf("%s:%s", namespace, identifier), &p)

	if len(option) > 0 {
		p.Subdir = option[0].Subdir
	}

	return &p.Data
}

func (e *RecipeBP) Save(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]interface{})
	}
	for _, v := range *e {
		data, ok := v.(recipe.RecipeInterface)
		if !ok {
			return fmt.Errorf("recipe: %v is not a recipe", v)
		}
		b, err := data.Encode()
		if err != nil {
			return err
		}
		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.Recipe, data.GetSubdir(), fmt.Sprintf("%s.json", data.GetIdentifier())), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *RecipeBP) LoadAll(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]interface{})
	}
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.Recipe))
	if err != nil {
		return err
	}
	for _, file := range files {
		r, err := recipe.Load(file)
		if err != nil {
			return err
		}
		// Strip the pathToBP from the file path
		file = strings.TrimPrefix(file, pathToBP+string(filepath.Separator)+destDirectory.Entity+string(filepath.Separator))
		// Get all the directories
		subdir := path.Dir(file)
		// Set subdir
		r.Subdir = subdir
		e.Add(r.Data.GetIdentifier(), r.Data)
	}
	return nil
}
