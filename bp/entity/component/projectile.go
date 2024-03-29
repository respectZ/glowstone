package component

// Code generated by scrapper/entityBp.ts;

// Allows the entity to be a thrown entity.  
type Projectile struct {
  // Determines the angle at which the projectile is thrown
  AngleOffset float64 `json:"angle_offset,omitempty"`
  // If true, the entity hit will be set on fire
  CatchFire bool `json:"catch_fire,omitempty"`
  // If true, the projectile will produce additional particles when a critical hit happens
  CritParticleOnHurt bool `json:"crit_particle_on_hurt,omitempty"`
  // If true, this entity will be destroyed when hit
  DestroyOnHurt bool `json:"destroy_on_hurt,omitempty"`
  // Entity Definitions defined here can't be hurt by the projectile
  Filter string `json:"filter,omitempty"`
  // If true, whether the projectile causes fire is affected by the mob griefing game rule
  FireAffectedByGriefing bool `json:"fire_affected_by_griefing,omitempty"`
  // The gravity applied to this entity when thrown. The higher the value, the faster the entity falls
  Gravity *float64 `json:"gravity,omitempty"`
  // The sound that plays when the projectile hits something
  HitSound string `json:"hit_sound,omitempty"`
  // If true, the projectile homes in to the nearest entity
  Homing bool `json:"homing,omitempty"`
  // The fraction of the projectile's speed maintained every frame while traveling in air
  Inertia *float64 `json:"inertia,omitempty"`
  // If true, the projectile will be treated as dangerous to the players
  IsDangerous bool `json:"is_dangerous,omitempty"`
  // If true, the projectile will knock back the entity it hits
  Knockback *bool `json:"knockback,omitempty"`
  // If true, the entity hit will be struck by lightning
  Lightning bool `json:"lightning,omitempty"`
  // The fraction of the projectile's speed maintained every frame while traveling in water
  LiquidInertia *float64 `json:"liquid_inertia,omitempty"`
  // If true, the projectile can hit multiple entities per flight
  MultipleTargets *bool `json:"multiple_targets,omitempty"`
  // [Vector3 [a,b,c]] The offset from the entity's anchor where the projectile will spawn
  Offset []float64 `json:"offset,omitempty"`
  // Time in seconds that the entity hit will be on fire for
  OnFireTime *float64 `json:"on_fire_time,omitempty"`
  // Particle to use upon collision
  Particle *string `json:"particle,omitempty"`
  // Defines the effect the arrow will apply to the entity it hits
  PotionEffect *int `json:"potion_effect,omitempty"`
  // Determines the velocity of the projectile
  Power *float64 `json:"power,omitempty"`
  // If true, this entity will be reflected back when hit
  ReflectOnHurt bool `json:"reflect_on_hurt,omitempty"`
  // If true, damage will be randomized based on damage and speed
  SemiRandomDiffDamage bool `json:"semi_random_diff_damage,omitempty"`
  // The sound that plays when the projectile is shot
  ShootSound string `json:"shoot_sound,omitempty"`
  // If true, the projectile will be shot towards the target of the entity firing it
  ShootTarget *bool `json:"shoot_target,omitempty"`
  // If true, the projectile will bounce upon hit
  ShouldBounce bool `json:"should_bounce,omitempty"`
  // If true, the projectile will be treated like a splash potion
  SplashPotion bool `json:"splash_potion,omitempty"`
  // Radius in blocks of the 'splash' effect
  SplashRange *float64 `json:"splash_range,omitempty"`
  // The base accuracy. Accuracy is determined by the formula uncertaintyBase - difficultyLevel * uncertaintyMultiplier
  UncertaintyBase float64 `json:"uncertainty_base,omitempty"`
  // Determines how much difficulty affects accuracy. Accuracy is determined by the formula uncertaintyBase - difficultyLevel * uncertaintyMultiplier
  UncertaintyMultiplier float64 `json:"uncertainty_multiplier,omitempty"`
}
