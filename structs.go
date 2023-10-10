package glowstone

import (
	"log"

	animation "github.com/respectZ/glowstone/animation"
	animationController "github.com/respectZ/glowstone/animation_controller"
	attachable "github.com/respectZ/glowstone/attachable"
	entity "github.com/respectZ/glowstone/entity"
	item "github.com/respectZ/glowstone/item"
	recipe "github.com/respectZ/glowstone/recipe"
	sound "github.com/respectZ/glowstone/rp/sound"
	texture "github.com/respectZ/glowstone/rp/texture"

	rp "github.com/respectZ/glowstone/rp"
)

type glowstone struct {
	BPDir            string
	RPDir            string
	MinEngineVersion [3]int
	Logger           *logger

	// Data
	Lang        map[string]string
	Entities    map[string]*entity.Entity
	Items       map[string]*item.Item
	Attachables map[string]*attachable.Attachable

	// BP Specific
	BPAnimationController map[string]*animationController.BPAnimationController
	BPAnimation           map[string]*animation.BPAnimation
	Recipes               map[string]interface{}

	// RP Specific
	ItemTexture     *texture.ItemTexture
	TerrainTexture  *texture.TerrainTexture
	SoundDefinition *sound.SoundDefinition
	RPBlocks        *rp.Blocks

	// Settings
	IsUpfront bool
}

type logger struct {
	Warning *log.Logger
	Error   *log.Logger
}

type Glowstone interface {
	// GetBPDir returns the path to the blueprint directory
	GetBPDir() string
	// SetBPDir sets the path to the blueprint directory
	SetBPDir(string)
	// GetRPDir returns the path to the resourcepack directory
	GetRPDir() string
	// SetRPDir sets the path to the resourcepack directory
	SetRPDir(string)
	// GetMinEngineVersion returns the minimum engine version
	GetMinEngineVersion() [3]int
	// SetMinEngineVersion sets the minimum engine version
	SetMinEngineVersion([3]int)
	// GetLang returns the language data
	GetLang(string) string
	// SetLang sets the language data
	SetLang(map[string]string)
	// AddLang adds the language data
	AddLang(string, string)

	// Build builds the project
	Build() error
	// Initialize initializes the lang and item_texture.json
	Initialize() error

	// SetUpfront will cache the project data upfront.
	//
	// Default: false
	SetUpfront(bool)

	// Save the project data to the disk.
	Save()

	/******************* Entities *******************/

	// AddEntity adds the entity to the project
	//
	// Example:
	// 	glowstone.AddEntity(entity)
	AddEntity(...interface{})

	// GetEntities returns the entities
	GetEntities() map[string]*entity.Entity
	// GetEntity returns the entity
	//
	// Example:
	// 	entity, err := glowstone.GetEntity("minecraft:zombie")
	GetEntity(string) (*entity.Entity, error)

	// NewEntity creates a new entity
	//
	// Example:
	// 	entity := glowstone.NewEntity("minecraft", "zombie")
	NewEntity(string, string) *entity.Entity

	// PreloadEnitties Preloads the entities, so you can use GetEntity() by identifier.
	//
	// Example:
	// 	glowstone.PreloadEntities()
	PreloadEntities()

	/******************* Items *******************/

	// AddItem adds the item to the project
	//
	// Example:
	// 	glowstone.AddItem(item)
	AddItem(...interface{})

	// GetItems returns the items
	GetItems() map[string]*item.Item

	// GetItem returns the item
	//
	// Example:
	// 	item, err := glowstone.GetItem("minecraft:stick")
	GetItem(string) (*item.Item, error)

	// NewItem creates a new item
	//
	// Example:
	// 	item := glowstone.NewItem("minecraft", "stick")
	NewItem(string, string) *item.Item

	// PreloadItems Preloads the items, so you can use GetItem() by identifier.
	//
	// Example:
	// 	glowstone.PreloadItems()
	PreloadItems()

	/******************* Recipes *******************/

	// AddRecipe adds the recipe to the project
	//
	// Example:
	// 	glowstone.AddRecipe(recipe)
	AddRecipe(...interface{})

	// GetRecipes returns the recipes
	GetRecipes() map[string]interface{}

	// GetRecipe returns the recipe
	//
	// Example:
	// 	recipe, err := glowstone.GetRecipe("minecraft:stick")
	GetRecipe(string) (interface{}, error)

	// NewRecipeBrewingContainer creates a new recipeBrewingContainer
	//
	// Example:
	// 	recipe := glowstone.NewRecipeBrewingContainer("minecraft", "stick")
	NewRecipeBrewingContainer(string, string) *recipe.RecipeBrewingContainer

	// NewRecipeBrewingMix creates a new recipeBrewingMix.
	//
	// The recipeBrewingMix is used to mix a potion with an item.
	//
	// Example:
	// 	recipe := glowstone.NewRecipeBrewingMix("minecraft", "stick")
	NewRecipeBrewingMix(string, string) *recipe.RecipeBrewingMix

	// NewRecipeFurnace creates a new recipeFurnace
	//
	// Example:
	// 	recipe := glowstone.NewRecipeFurnace("minecraft", "stick")
	NewRecipeFurnace(string, string) *recipe.RecipeFurnace

	// NewRecipeShaped creates a new recipeShaped
	//
	// Example:
	// 	recipe := glowstone.NewRecipeShaped("minecraft", "stick")
	NewRecipeShaped(string, string) *recipe.RecipeShaped

	// NewRecipeShapeless creates a new recipeShapeless
	//
	// Example:
	// 	recipe := glowstone.NewRecipeShapeless("minecraft", "stick")
	NewRecipeShapeless(string, string) *recipe.RecipeShapeless

	// NewRecipeSmithingAddition creates a new recipeSmithingAddition
	//
	// Example:
	// 	recipe := glowstone.NewRecipeSmithingAddition("minecraft", "stick")
	NewRecipeSmithingTransform(string, string) *recipe.RecipeSmithingTransform

	/******************* Attachable *******************/

	// AddAttachable adds the attachable to the project
	//
	// Example:
	// 	glowstone.AddAttachable(attachable)
	AddAttachable(attachable.Attachable)

	// GetAttachables returns the attachables
	GetAttachables() map[string]*attachable.Attachable

	// GetAttachable returns the attachable
	//
	// Example:
	// 	attachable, err := glowstone.GetAttachable("minecraft:stick")
	GetAttachable(string) (*attachable.Attachable, error)

	// NewAttachable creates a new attachable
	//
	// Example:
	// 	attachable := glowstone.NewAttachable("minecraft", "stick")
	NewAttachable(string, string, ...string) *attachable.Attachable

	/******************* RPBlocks *******************/

	// Returns the blocks.json
	GetRPBlocks() *rp.Blocks

	/******************* Item Texture *******************/

	// GetItemTexture returns the item_texture.json
	GetItemTexture() *texture.ItemTexture

	/******************* Terrain Texture *******************/

	// GetTerrainTexture returns the terrain_texture.json
	GetTerrainTexture() *texture.TerrainTexture

	/******************* Sound Definition *******************/

	// GetSoundDefinition returns the sound_definitions.json
	GetSoundDefinition() *sound.SoundDefinition

	/******************* BPAnimationController *******************/

	// AddBPAnimationController adds the BPAnimationController to the project
	//
	// Example:
	// 	glowstone.AddBPAnimationController(bpAnimationController)
	AddBPAnimationController(...interface{})

	// NewBPAnimationController creates a new BPAnimationController
	//
	// Example:
	// 	bpAnimationController := glowstone.NewBPAnimationController("player.animation_controller.json")
	NewBPAnimationController(string) *animationController.BPAnimationController

	/******************* BPAnimation *******************/

	// AddBPAnimation adds the BPAnimation to the project
	//
	// Example:
	// 	glowstone.AddBPAnimation(bpAnimation)
	AddBPAnimation(...interface{})

	// NewBPAnimation creates a new BPAnimation
	//
	// Example:
	// 	bpAnimation := glowstone.NewBPAnimation("player.animation.json")
	NewBPAnimation(string) *animation.BPAnimation
}
