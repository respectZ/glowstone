package component

// Code generated by scrapper/entityBp.ts;

import (
  f "github.com/respectZ/glowstone/bp/types"
)

// Adds a timer after which this entity will spawn another entity or item (similar to vanilla's chicken's egg-laying behavior).  
type SpawnEntity struct {
  // If present, the specified entity will only spawn if the filter evaluates to true.
  Filters *f.Filter `json:"filters,omitempty"`
  // Maximum amount of time to randomly wait in seconds before another entity is spawned.
  MaxWaitTime int `json:"max_wait_time,omitempty"`
  // Minimum amount of time to randomly wait in seconds before another entity is spawned.
  MinWaitTime int `json:"min_wait_time,omitempty"`
  // The number of entities of this type to spawn each time that this triggers.
  NumToSpawn int `json:"num_to_spawn,omitempty"`
  // If true, this the spawned entity will be leashed to the parent.
  ShouldLeash bool `json:"should_leash,omitempty"`
  // If true, this component will only ever spawn the specified entity once.
  SingleUse bool `json:"single_use,omitempty"`
  // Identifier of the entity to spawn, leave empty to spawn the item defined by "spawn_item" instead.
  SpawnEntity string `json:"spawn_entity,omitempty"`
  // Event to call on the spawned entity when it spawns.
  SpawnEvent string `json:"spawn_event,omitempty"`
  // Item identifier of the item to spawn.
  SpawnItem string `json:"spawn_item,omitempty"`
  // [Trigger] Event to call on this entity when the item is spawned.
  SpawnItemEvent interface{} `json:"spawn_item_event,omitempty"`
  // Method to use to spawn the entity.
  SpawnMethod string `json:"spawn_method,omitempty"`
  // Identifier of the sound effect to play when the entity is spawned.
  SpawnSound string `json:"spawn_sound,omitempty"`
}
