package component

// Code generated by scrapper/entityBp.ts;

// Allows this entity to generate paths that include water.  
type Navigation_Swim struct {
  // Tells the pathfinder to avoid blocks that cause damage when finding a path
  AvoidDamageBlocks bool `json:"avoid_damage_blocks,omitempty"`
  // Tells the pathfinder to avoid portals (like nether portals) when finding a path
  AvoidPortals bool `json:"avoid_portals,omitempty"`
  // Whether or not the pathfinder should avoid tiles that are exposed to the sun when creating paths
  AvoidSun bool `json:"avoid_sun,omitempty"`
  // Tells the pathfinder to avoid water when creating a path
  AvoidWater bool `json:"avoid_water,omitempty"`
  // Tells the pathfinder which blocks to avoid when creating a path
  BlocksToAvoid []interface{} `json:"blocks_to_avoid,omitempty"`
  // Tells the pathfinder whether or not it can jump out of water (like a dolphin)
  CanBreach bool `json:"can_breach,omitempty"`
  // Tells the pathfinder that it can path through a closed door and break it
  CanBreakDoors bool `json:"can_break_doors,omitempty"`
  // Tells the pathfinder whether or not it can jump up blocks
  CanJump bool `json:"can_jump,omitempty"`
  // Tells the pathfinder that it can path through a closed door assuming the AI will open the door
  CanOpenDoors bool `json:"can_open_doors,omitempty"`
  // Tells the pathfinder that it can path through a closed iron door assuming the AI will open the door
  CanOpenIronDoors bool `json:"can_open_iron_doors,omitempty"`
  // Whether a path can be created through a door
  CanPassDoors bool `json:"can_pass_doors,omitempty"`
  // Tells the pathfinder that it can start pathing when in the air
  CanPathFromAir bool `json:"can_path_from_air,omitempty"`
  // Tells the pathfinder whether or not it can travel on the surface of the lava
  CanPathOverLava bool `json:"can_path_over_lava,omitempty"`
  // Tells the pathfinder whether or not it can travel on the surface of the water
  CanPathOverWater bool `json:"can_path_over_water,omitempty"`
  // Tells the pathfinder whether or not it will be pulled down by gravity while in water
  CanSink bool `json:"can_sink,omitempty"`
  // Tells the pathfinder whether or not it can path anywhere through water and plays swimming animation along that path
  CanSwim bool `json:"can_swim,omitempty"`
  // Tells the pathfinder whether or not it can walk on the ground outside water
  CanWalk bool `json:"can_walk,omitempty"`
  // Tells the pathfinder whether or not it can travel in lava like walking on ground
  CanWalkInLava bool `json:"can_walk_in_lava,omitempty"`
  // Tells the pathfinder whether or not it can walk on the ground underwater
  IsAmphibious bool `json:"is_amphibious,omitempty"`
}