package component

// Code generated by scrapper/entityBp.ts;

// Allows an entity to emit `entityMove`, `swim` and `flap` game events, depending on the block the entity is moving through. It is added by default to every mob. Add it again to override its behavior.  
type GameEventMovementTracking struct {
  // If true, the `flap` game event will be emitted when the entity moves through air.
  EmitFlap bool `json:"emit_flap,omitempty"`
  // If true, the `entityMove` game event will be emitted when the entity moves on ground or through a solid.
  EmitMove bool `json:"emit_move,omitempty"`
  // If true, the `swim` game event will be emitted when the entity moves through a liquid.
  EmitSwim bool `json:"emit_swim,omitempty"`
}
