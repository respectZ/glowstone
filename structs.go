package glowstone

import (
	"log"

	animation "github.com/respectZ/glowstone/animation"
	animationController "github.com/respectZ/glowstone/animation_controller"
	attachable "github.com/respectZ/glowstone/attachable"
	block "github.com/respectZ/glowstone/block"
	entity "github.com/respectZ/glowstone/entity"
	item "github.com/respectZ/glowstone/item"
	loot_table "github.com/respectZ/glowstone/loot_table"
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
	Blocks                map[string]*block.Block
	LootTables            map[string]*loot_table.LootTable

	// RP Specific
	RPAnimationController map[string]*animationController.RPAnimationController
	ItemTexture           *texture.ItemTexture
	TerrainTexture        *texture.TerrainTexture
	SoundDefinition       *sound.SoundDefinition
	RPBlocks              *rp.Blocks

	// Settings
	MinifyJSON bool
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
	// 	entity := glowstone.NewEntity("glowstone:zombie", "mob")
	NewEntity(string) *entity.Entity

	// PreloadEnitties Preloads the entities, so you can use GetEntity() by identifier.
	//
	// Example:
	// 	glowstone.PreloadEntities()
	PreloadEntities()

	/******************* Blocks *******************/

	// AddBlock adds the block to the project
	//
	// Example:
	// 	glowstone.AddBlock(block)
	AddBlock(...interface{})

	// GetBlocks returns the blocks
	GetBlocks() map[string]*block.Block

	// GetBlock returns the block
	//
	// Example:
	// 	block, err := glowstone.GetBlock("minecraft:stone")
	GetBlock(string) (*block.Block, error)

	// NewBlock creates a new block
	//
	// Example:
	// 	block := glowstone.NewBlock("glowstone:stone", "stones")
	NewBlock(string, ...string) *block.Block

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
	NewItem(string) *item.Item

	// PreloadItems Preloads the items, so you can use GetItem() by identifier.
	//
	// Example:
	// 	glowstone.PreloadItems()
	PreloadItems()

	/******************* Loot Tables *******************/

	// AddLootTable adds the loot table to the project
	//
	// Example:
	// 	glowstone.AddLootTable("entities/zombie.json", lootTable)
	AddLootTable(name string, lootTable *loot_table.LootTable)

	// GetLootTables returns the loot tables
	GetLootTables() map[string]*loot_table.LootTable

	// GetLootTable returns the loot table
	//
	// Example:
	// 	lootTable, err := glowstone.GetLootTable("entities/zombie.json")
	GetLootTable(string) (*loot_table.LootTable, error)

	// NewLootTable creates a new loot table
	//
	// Example:
	// 	lootTable := glowstone.NewLootTable("entities/zombie.json")
	NewLootTable(string) *loot_table.LootTable

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

	/******************* RPAnimationController *******************/

	// AddRPAnimationController adds the RPAnimationController to the project
	//
	// Example:
	// 	glowstone.AddRPAnimationController(bpAnimationController)
	AddRPAnimationController(...interface{})

	// NewRPAnimationController creates a new RPAnimationController
	//
	// Example:
	// 	bpAnimationController := glowstone.NewRPAnimationController("player.animation_controller.json")
	NewRPAnimationController(string) *animationController.RPAnimationController

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
