package item

type Category string

/*
NONE
COMMANDS
CONSTRUCTION
EQUIPMENT
ITEMS
NATURE
*/
const (
	CATEGORY_NONE         Category = "none"
	CATEGORY_COMMANDS     Category = "commands"
	CATEGORY_CONSTRUCTION Category = "construction"
	CATEGORY_EQUIPMENT    Category = "equipment"
	CATEGORY_ITEMS        Category = "items"
	CATEGORY_NATURE       Category = "nature"
)

type Group string

/*
itemGroup.name.anvil
itemGroup.name.arrow
itemGroup.name.axe
itemGroup.name.banner
itemGroup.name.banner_pattern
itemGroup.name.bed
itemGroup.name.boat
itemGroup.name.boots
itemGroup.name.buttons
itemGroup.name.candles
itemGroup.name.chalkboard
itemGroup.name.chest
itemGroup.name.chestboat
itemGroup.name.chestplate
itemGroup.name.concrete
itemGroup.name.concretePowder
itemGroup.name.cookedFood
itemGroup.name.copper
itemGroup.name.coral
itemGroup.name.coral_decorations
itemGroup.name.crop
itemGroup.name.door
itemGroup.name.dye
itemGroup.name.enchantedBook
itemGroup.name.fence
itemGroup.name.fenceGate
itemGroup.name.firework
itemGroup.name.fireworkStars
itemGroup.name.flower
itemGroup.name.glass
itemGroup.name.glassPane
itemGroup.name.glazedTerracotta
itemGroup.name.goatHorn
itemGroup.name.grass
itemGroup.name.hanging_sign
itemGroup.name.helmet
itemGroup.name.hoe
itemGroup.name.horseArmor
itemGroup.name.leaves
itemGroup.name.leggings
itemGroup.name.lingeringPotion
itemGroup.name.log
itemGroup.name.minecart
itemGroup.name.miscFood
itemGroup.name.mobEgg
itemGroup.name.monsterStoneEgg
itemGroup.name.mushroom
itemGroup.name.netherWartBlock
itemGroup.name.ore
itemGroup.name.permission
itemGroup.name.pickaxe
itemGroup.name.planks
itemGroup.name.potion
itemGroup.name.potterySherds
itemGroup.name.pressurePlate
itemGroup.name.rail
itemGroup.name.rawFood
itemGroup.name.record
itemGroup.name.sandstone
itemGroup.name.sapling
itemGroup.name.sculk
itemGroup.name.seed
itemGroup.name.shovel
itemGroup.name.shulkerBox
itemGroup.name.sign
itemGroup.name.skull
itemGroup.name.slab
itemGroup.name.smithing_templates
itemGroup.name.splashPotion
itemGroup.name.stainedClay
itemGroup.name.stairs
itemGroup.name.stone
itemGroup.name.stoneBrick
itemGroup.name.sword
itemGroup.name.trapdoor
itemGroup.name.walls
itemGroup.name.wood
itemGroup.name.wool
itemGroup.name.woolCarpet
*/
const (
	GROUP_ANVIL          Group = "itemGroup.name.anvil"
	GROUP_ARROW          Group = "itemGroup.name.arrow"
	GROUP_AXE            Group = "itemGroup.name.axe"
	GROUP_BANNER         Group = "itemGroup.name.banner"
	GROUP_BANNER_PATTERN Group = "itemGroup.name.banner_pattern"
	GROUP_BED            Group = "itemGroup.name.bed"
	GROUP_BOAT           Group = "itemGroup.name.boat"
	GROUP_BOOTS          Group = "itemGroup.name.boots"
	GROUP_BUTTONS        Group = "itemGroup.name.buttons"
	GROUP_CANDLES        Group = "itemGroup.name.candles"
	GROUP_CHALKBOARD     Group = "itemGroup.name.chalkboard"
	GROUP_CHEMISTRYTABLE Group = "itemGroup.name.chemistry_table"
	GROUP_CHEST          Group = "itemGroup.name.chest"
	GROUP_CHESTBOAT      Group = "itemGroup.name.chest_boat"
	GROUP_CHESTPLATE     Group = "itemGroup.name.chestplate"
	GROUP_CONCRETE       Group = "itemGroup.name.concrete"
	GROUP_CONCRETEPOWDER Group = "itemGroup.name.concretePowder"
	GROUP_COOKEDFOOD     Group = "itemGroup.name.cookedFood"
	GROUP_COPPER         Group = "itemGroup.name.copper"
	GROUP_CORAL          Group = "itemGroup.name.coral"
	GROUP_CORALDECOR     Group = "itemGroup.name.coral_decorations"
	GROUP_CROP           Group = "itemGroup.name.crop"
	GROUP_DOOR           Group = "itemGroup.name.door"
	GROUP_DYE            Group = "itemGroup.name.dye"
	GROUP_ENCHANTEDBOOK  Group = "itemGroup.name.enchantedBook"
	GROUP_FENCE          Group = "itemGroup.name.fence"
	GROUP_FENCEGATE      Group = "itemGroup.name.fenceGate"
	GROUP_FIREWORK       Group = "itemGroup.name.firework"
	GROUP_FIREWORKSTARS  Group = "itemGroup.name.fireworkStars"
	GROUP_FLOWER         Group = "itemGroup.name.flower"
	GROUP_GLASS          Group = "itemGroup.name.glass"
	GROUP_GLASSPANE      Group = "itemGroup.name.glassPane"
	GROUP_GLAZEDTERRA    Group = "itemGroup.name.glazedTerracotta"
	GROUP_GOATHORN       Group = "itemGroup.name.goatHorn"
	GROUP_GRASS          Group = "itemGroup.name.grass"
	GROUP_HANGINGSIGN    Group = "itemGroup.name.hanging_sign"
	GROUP_HELMET         Group = "itemGroup.name.helmet"
	GROUP_HOE            Group = "itemGroup.name.hoe"
	GROUP_HORSEARMOR     Group = "itemGroup.name.horseArmor"
	GROUP_LEAVES         Group = "itemGroup.name.leaves"
	GROUP_LEGGINGS       Group = "itemGroup.name.leggings"
	GROUP_LINGERINGPOT   Group = "itemGroup.name.lingeringPotion"
	GROUP_LOG            Group = "itemGroup.name.log"
	GROUP_MINECART       Group = "itemGroup.name.minecart"
	GROUP_MISCFOOD       Group = "itemGroup.name.miscFood"
	GROUP_MOBEGG         Group = "itemGroup.name.mobEgg"
	GROUP_MONSTEREGG     Group = "itemGroup.name.monsterStoneEgg"
	GROUP_MUSHROOM       Group = "itemGroup.name.mushroom"
	GROUP_NETHERWART     Group = "itemGroup.name.netherWartBlock"
	GROUP_ORE            Group = "itemGroup.name.ore"
	GROUP_PERMISSION     Group = "itemGroup.name.permission"
	GROUP_PICKAXE        Group = "itemGroup.name.pickaxe"
	GROUP_PLANKS         Group = "itemGroup.name.planks"
	GROUP_POTION         Group = "itemGroup.name.potion"
	GROUP_POTTERYSHERDS  Group = "itemGroup.name.potterySherds"
	GROUP_PRESSUREPLATE  Group = "itemGroup.name.pressurePlate"
	GROUP_RAIL           Group = "itemGroup.name.rail"
	GROUP_RAWFOOD        Group = "itemGroup.name.rawFood"
	GROUP_RECORD         Group = "itemGroup.name.record"
	GROUP_SANDSTONE      Group = "itemGroup.name.sandstone"
	GROUP_SAPLING        Group = "itemGroup.name.sapling"
	GROUP_SCULK          Group = "itemGroup.name.sculk"
	GROUP_SEED           Group = "itemGroup.name.seed"
	GROUP_SHOVEL         Group = "itemGroup.name.shovel"
	GROUP_SHULKERBOX     Group = "itemGroup.name.shulkerBox"
	GROUP_SIGN           Group = "itemGroup.name.sign"
	GROUP_SKULL          Group = "itemGroup.name.skull"
	GROUP_SLAB           Group = "itemGroup.name.slab"
	GROUP_SMITHING       Group = "itemGroup.name.smithing_templates"
	GROUP_SPLASHPOTION   Group = "itemGroup.name.splashPotion"
	GROUP_STAINEDCLAY    Group = "itemGroup.name.stainedClay"
	GROUP_STAIRS         Group = "itemGroup.name.stairs"
	GROUP_STONE          Group = "itemGroup.name.stone"
	GROUP_STONEBRICK     Group = "itemGroup.name.stoneBrick"
	GROUP_SWORD          Group = "itemGroup.name.sword"
	GROUP_TRAPDOOR       Group = "itemGroup.name.trapdoor"
	GROUP_WALLS          Group = "itemGroup.name.walls"
	GROUP_WOOD           Group = "itemGroup.name.wood"
	GROUP_WOOL           Group = "itemGroup.name.wool"
	GROUP_WOOLCARPET     Group = "itemGroup.name.woolCarpet"
)
