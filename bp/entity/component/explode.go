package component

// Code generated by scrapper/entityBp.ts;

// Defines how the entity explodes.  
type Explode struct {
  // If true, the explosion will destroy blocks in the explosion radius.
  BreaksBlocks bool `json:"breaks_blocks,omitempty"`
  // If true, blocks in the explosion radius will be set on fire.
  CausesFire bool `json:"causes_fire,omitempty"`
  // If true, whether the explosion breaks blocks is affected by the mob griefing game rule.
  DestroyAffectedByGriefing bool `json:"destroy_affected_by_griefing,omitempty"`
  // If true, whether the explosion causes fire is affected by the mob griefing game rule.
  FireAffectedByGriefing bool `json:"fire_affected_by_griefing,omitempty"`
  // The range for the random amount of time the fuse will be lit before exploding, a negative value means the explosion will be immediate.
  FuseLength float64 `json:"fuse_length,omitempty"`
  // If true, the fuse is already lit when this component is added to the entity.
  FuseLit bool `json:"fuse_lit,omitempty"`
  // A blocks explosion resistance will be capped at this value when an explosion occurs.
  MaxResistance float64 `json:"max_resistance,omitempty"`
  // The radius of the explosion in blocks and the amount of damage the explosion deals.
  Power float64 `json:"power,omitempty"`
}
