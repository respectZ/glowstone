package bp

import (
	"os"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/recipe"
	g_util "github.com/respectZ/glowstone/util"
)

type RecipeBP map[string]*RecipeFile

type IRecipeBP interface {
	Add(string, bp.RecipeInterface)
	Get(string) (bp.RecipeInterface, bool)
	Remove(string)
	Clear()
	All() map[string]*RecipeFile
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

func (e *RecipeBP) Add(key string, value bp.RecipeInterface) {
	p, ok := (*e)[key]
	if ok {
		p.Data = value
	} else {
		(*e)[key] = &RecipeFile{
			Data: value,
		}
	}
}

func (e *RecipeBP) Get(key string) (bp.RecipeInterface, bool) {
	value, ok := (*e)[key]
	if !ok {
		return nil, false
	}
	return value.Data, true
}

func (e *RecipeBP) Remove(key string) {
	delete(*e, key)
}

func (e *RecipeBP) Clear() {
	*e = make(map[string]*RecipeFile)
}

func (e *RecipeBP) All() map[string]*RecipeFile {
	return *e
}

func (e *RecipeBP) IsEmpty() bool {
	return len(*e) == 0
}

func (e *RecipeBP) Size() int {
	return len(*e)
}

func (e *RecipeBP) UnmarshalJSON(data []byte) error {
	return g_util.UnmarshalJSON(data, e)
}

func (e *RecipeBP) NewBrewingContainer(identifier string, option ...struct{ Subdir string }) *bp.RecipeBrewingContainer {
	r := bp.NewBrewingContainer(identifier)
	generic := &RecipeFile{
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
	generic := &RecipeFile{
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
	generic := &RecipeFile{
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
	generic := &RecipeFile{
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
	generic := &RecipeFile{
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
	generic := &RecipeFile{
		Data: r,
	}

	if len(option) > 0 {
		generic.Subdir = option[0].Subdir
	}

	e.Add(identifier, generic)

	return &r
}

func (e *RecipeBP) Save(pathToBP string) error {
	for _, v := range *e {
		b, err := v.Encode()
		if err != nil {
			return err
		}
		err = g_util.Writefile(filepath.Join(pathToBP, destDirectory.Recipe, v.Subdir, v.GetFilename()), b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *RecipeBP) LoadAll(pathToBP string) error {
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.Recipe))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, file := range files {
		r, err := bp.Load(file)
		if err != nil {
			return err
		}
		e.Add(r.GetIdentifier(), r)
		// Strip the pathToBP from the file path
		file = strings.TrimPrefix(file, pathToBP+string(filepath.Separator)+destDirectory.Recipe+string(filepath.Separator))
		// Get all the directories
		subdir := filepath.Dir(file)
		// Filename
		filename := filepath.Base(file)

		(*e)[r.GetIdentifier()].Filename = filename
		(*e)[r.GetIdentifier()].Subdir = subdir
	}
	return nil
}
