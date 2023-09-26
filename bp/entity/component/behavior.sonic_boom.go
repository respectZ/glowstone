package component

// Code generated by scrapper/entityBp.ts;

// Allows this entity to perform a 'sonic boom' ranged attack  
type Behavior_SonicBoom struct {
  // Cooldown in seconds required after using this attack until the entity can use sonic boom again.
  AttackCooldown float64 `json:"attack_cooldown,omitempty"`
  // Attack damage of the sonic boom.
  AttackDamage float64 `json:"attack_damage,omitempty"`
  // Horizontal range (in blocks) at which the sonic boom can damage the target.
  AttackRangeHorizontal float64 `json:"attack_range_horizontal,omitempty"`
  // Vertical range (in blocks) at which the sonic boom can damage the target.
  AttackRangeVertical float64 `json:"attack_range_vertical,omitempty"`
  // Sound event for the attack.
  AttackSound string `json:"attack_sound,omitempty"`
  // Sound event for the charge up.
  ChargeSound string `json:"charge_sound,omitempty"`
  // Goal duration in seconds
  Duration float64 `json:"duration,omitempty"`
  // Duration in seconds until the attack sound is played.
  DurationUntilAttackSound float64 `json:"duration_until_attack_sound,omitempty"`
  // Height cap of the attack knockback's vertical delta.
  KnockbackHeightCap float64 `json:"knockback_height_cap,omitempty"`
  // Horizontal strength of the attack's knockback applied to the attack target.
  KnockbackHorizontalStrength float64 `json:"knockback_horizontal_strength,omitempty"`
  // Vertical strength of the attack's knockback applied to the attack target.
  KnockbackVerticalStrength float64 `json:"knockback_vertical_strength,omitempty"`
  // This multiplier modifies the attacking entity's speed when moving toward the target.
  SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
  // The priority of this behavior. Lower values are higher priority.
  Priority int `json:"priority,omitempty"`
}
