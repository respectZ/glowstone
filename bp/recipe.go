package glowstone

import (
	"fmt"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/recipe"
	"github.com/respectZ/glowstone/recipe"
	g_util "github.com/respectZ/glowstone/util"
)

type RecipeBP map[string]*recipe.GenericRecipe

type IRecipeBP interface {
	Add(string, *recipe.GenericRecipe)
	Get(string) (*recipe.GenericRecipe, bool)
	Remove(string)
	Clear()
	All() map[string]*recipe.GenericRecipe
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

func (e *RecipeBP) Add(key string, value *recipe.GenericRecipe) {
	if *e == nil {
		*e = make(map[string]*recipe.GenericRecipe)
	}
	(*e)[key] = value
}

func (e *RecipeBP) Get(key string) (*recipe.GenericRecipe, bool) {
	if *e == nil {
		*e = make(map[string]*recipe.GenericRecipe)
	}
	value, ok := (*e)[key]
	return value, ok
}

func (e *RecipeBP) Remove(key string) {
	if *e == nil {
		*e = make(map[string]*recipe.GenericRecipe)
	}
	delete(*e, key)
}

func (e *RecipeBP) Clear() {
	if *e == nil {
		*e = make(map[string]*recipe.GenericRecipe)
	}
	*e = make(map[string]*recipe.GenericRecipe)
}

func (e *RecipeBP) All() map[string]*recipe.GenericRecipe {
	if *e == nil {
		*e = make(map[string]*recipe.GenericRecipe)
	}
	return *e
}

func (e *RecipeBP) IsEmpty() bool {
	if *e == nil {
		*e = make(map[string]*recipe.GenericRecipe)
	}
	return len(*e) == 0
}

func (e *RecipeBP) Size() int {
	if *e == nil {
		*e = make(map[string]*recipe.GenericRecipe)
	}
	return len(*e)
}

func (e *RecipeBP) UnmarshalJSON(data []byte) error {
	if *e == nil {
		*e = make(map[string]*recipe.GenericRecipe)
	}
	return g_util.UnmarshalJSON(data, e)
}

func (e *RecipeBP) NewBrewingContainer(identifier string, option ...struct{ Subdir string }) *bp.RecipeBrewingContainer {
	r := bp.NewBrewingContainer(identifier)
	generic := &recipe.GenericRecipe{
		Data: r,
	}

	if len(option) > 0 {
		generic.Subdir = option[0].Subdir
	}

	e.Add(identifier, generic)

	return &r
}

func (e *RecipeBP) NewBrewingMix(identifier string, option ...struct{ Subdir string }) *bp.RecipeBrewingMix {
	r := bp.NewBrewingMix(identifier)
	generic := &recipe.GenericRecipe{
		Data: r,
	}

	if len(option) > 0 {
		generic.Subdir = option[0].Subdir
	}

	e.Add(identifier, generic)

	return &r
}

func (e *RecipeBP) NewFurnace(identifier string, option ...struct{ Subdir string }) *bp.RecipeFurnace {
	r := bp.NewFurnace(identifier)
	generic := &recipe.GenericRecipe{
		Data: r,
	}

	if len(option) > 0 {
		generic.Subdir = option[0].Subdir
	}

	e.Add(identifier, generic)

	return &r
}

func (e *RecipeBP) NewShaped(identifier string, option ...struct{ Subdir string }) *bp.RecipeShaped {
	r := bp.NewShaped(identifier)
	generic := &recipe.GenericRecipe{
		Data: r,
	}

	if len(option) > 0 {
		generic.Subdir = option[0].Subdir
	}

	e.Add(identifier, generic)

	return &r
}

func (e *RecipeBP) NewShapeless(identifier string, option ...struct{ Subdir string }) *bp.RecipeShapeless {
	r := bp.NewShapeless(identifier)
	generic := &recipe.GenericRecipe{
		Data: r,
	}

	if len(option) > 0 {
		generic.Subdir = option[0].Subdir
	}

	e.Add(identifier, generic)

	return &r
}

func (e *RecipeBP) NewSmithingTransform(identifier string, option ...struct{ Subdir string }) *bp.RecipeSmithingTransform {
	r := bp.NewSmithingTransform(identifier)
	generic := &recipe.GenericRecipe{
		Data: r,
	}

	if len(option) > 0 {
		generic.Subdir = option[0].Subdir
	}

	e.Add(identifier, generic)

	return &r
}

func (e *RecipeBP) Save(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]*recipe.GenericRecipe)
	}
	for _, v := range *e {
		b, err := v.Encode()
		if err != nil {
			return err
		}
		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.Recipe, v.GetSubdir(), fmt.Sprintf("%s.json", v.GetIdentifier())), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *RecipeBP) LoadAll(pathToBP string) error {
	if *e == nil {
		*e = make(map[string]*recipe.GenericRecipe)
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
		file = strings.TrimPrefix(file, pathToBP+string(filepath.Separator)+destDirectory.Recipe+string(filepath.Separator))
		// Get all the directories
		subdir := filepath.Dir(file)
		// Set subdir
		r.Subdir = subdir
		e.Add(r.Data.GetIdentifier(), r)
	}
	return nil
}
