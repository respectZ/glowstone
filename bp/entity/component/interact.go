package component

// Code generated by scrapper/entityBp.ts;

// [Not a component] Loot table with items to add to the player's inventory upon successful interaction.
type InteractAddItems struct {
  // File path, relative to the Behavior Pack's path, to the loot table file.
  Table string `json:"table,omitempty"`
}

// [Not a component] Loot table with items to drop on the ground upon successful interaction.
type InteractSpawnItems struct {
  // File path, relative to the Behavior Pack's path, to the loot table file.
  Table string `json:"table,omitempty"`
}

// Defines interactions with this entity.  
type Interact struct {
  // Loot table with items to add to the player's inventory upon successful interaction.
  AddItems InteractAddItems `json:"add_items,omitempty"`
  // Time in seconds before this entity can be interacted with again.
  Cooldown float64 `json:"cooldown,omitempty"`
  // Time in seconds before this entity can be interacted with after being attacked.
  CooldownAfterBeingAttacked float64 `json:"cooldown_after_being_attacked,omitempty"`
  // The entity's equipment slot to remove and drop the item from, if any, upon successful interaction.
  DropItemSlot *int `json:"drop_item_slot,omitempty"`
  // The entity's equipment slot to equip the item to, if any, upon successful interaction.
  EquipItemSlot *int `json:"equip_item_slot,omitempty"`
  // The amount of health this entity will recover or hurt when interacting with this item. Negative values will harm the entity.
  HealthAmount int `json:"health_amount,omitempty"`
  // The amount of damage the item will take when used to interact with this entity. A value of 0 means the item won't lose durability.
  HurtItem int `json:"hurt_item,omitempty"`
  // Text to show when the player is able to interact in this way with this entity when playing with Touch-screen controls.
  InteractText string `json:"interact_text,omitempty"`
  // Event to fire when the interaction occurs.
  OnInteract string `json:"on_interact,omitempty"`
  // Particle effect that will be triggered at the start of the interaction.particle_offset_towards_interactor#
  ParticleOnStart interface{} `json:"particle_on_start,omitempty"`
  // List of sounds to play when the interaction occurs.
  PlaySounds string `json:"play_sounds,omitempty"`
  // List of entities to spawn when the interaction occurs.
  SpawnEntities string `json:"spawn_entities,omitempty"`
  // Loot table with items to drop on the ground upon successful interaction.
  SpawnItems InteractSpawnItems `json:"spawn_items,omitempty"`
  // If true, the player will do the 'swing' animation when interacting with this entity.
  Swing bool `json:"swing,omitempty"`
  // The item used will transform to this item upon successful interaction. Format: itemName:auxValue
  TransformToItem string `json:"transform_to_item,omitempty"`
  // If true, the interaction will use an item.
  UseItem bool `json:"use_item,omitempty"`
  // Vibration to emit when the interaction occurs. Admitted values are none (no vibration emitted), shear, entity_act, entity_interact.
  Vibration *string `json:"vibration,omitempty"`
}
