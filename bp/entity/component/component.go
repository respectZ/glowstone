package component

import (
	f "glowstone/bp/types"
)

// Enables the mob to admire items that have been configured as admirable. Must be used in combination with the admire_item component
type Behavior_AdmireItem struct {
	// The sound event to play when admiring the item
	AdmireItemSound string `json:"admire_item_sound,omitempty"`

	// The range of time in seconds to randomly wait before playing the sound again.
	SoundInterval float64 `json:"sound_interval,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to avoid certain blocks.
type Behavior_AvoidBlock struct {
	// The sound event to play when the mob is avoiding a block.
	AvoidBlockSound string `json:"avoid_block_sound,omitempty"`

	// [Trigger] Escape trigger.
	OnEscape interface{} `json:"on_escape,omitempty"`

	// Maximum distance to look for a block in y.
	SearchHeight int `json:"search_height,omitempty"`

	// Maximum distance to look for a block in xz.
	SearchRange int `json:"search_range,omitempty"`

	// The range of time in seconds to randomly wait before playing the sound again.
	SoundInterval float64 `json:"sound_interval,omitempty"`

	// Modifier for sprint speed. 1.0 means keep the regular speed, while higher numbers make the sprint speed faster.
	SprintSpeedModifier float64 `json:"sprint_speed_modifier,omitempty"`

	// List of block types this mob avoids.
	TargetBlocks []interface{} `json:"target_blocks,omitempty"`

	// Block search method.
	TargetSelectionMethod string `json:"target_selection_method,omitempty"`

	// Should start tick interval.
	TickInterval int `json:"tick_interval,omitempty"`

	// Modifier for walking speed. 1.0 means keep the regular speed, while higher numbers make the walking speed faster.
	WalkSpeedModifier float64 `json:"walk_speed_modifier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to run away from other entities that meet the criteria specified.
type Behavior_AvoidMobType struct {
	// The sound event to play when the mob is avoiding another mob.
	AvoidMobSound string `json:"avoid_mob_sound,omitempty"`

	// The next target position the entity chooses to avoid another entity will be chosen within this XZ Distance.
	AvoidTargetXz int `json:"avoid_target_xz,omitempty"`

	// The next target position the entity chooses to avoid another entity will be chosen within this Y Distance.
	AvoidTargetY int `json:"avoid_target_y,omitempty"`

	// The list of conditions another entity must meet to be a valid target to avoid.
	EntityTypes *f.Filter `json:"entity_types,omitempty"`

	// Whether or not to ignore direct line of sight while this entity is running away from other specified entities.
	IgnoreVisibilty bool `json:"ignore_visibilty,omitempty"`

	// Maximum distance to look for an avoid target for the entity.
	MaxDist float64 `json:"max_dist,omitempty"`

	// How many blocks away from its avoid target the entity must be for it to stop fleeing from the avoid target.
	MaxFlee float64 `json:"max_flee,omitempty"`

	// [Trigger] Event that is triggered when escaping from a mob.
	OnEscapeEvent interface{} `json:"on_escape_event,omitempty"`

	// Percent chance this entity will stop avoiding another entity based on that entity's strength, where 1.0 = 100%.
	ProbabilityPerStrength float64 `json:"probability_per_strength,omitempty"`

	// Determine if we should remove target when fleeing or not.
	RemoveTarget bool `json:"remove_target,omitempty"`

	// The range of time in seconds to randomly wait before playing the sound again.
	SoundInterval float64 `json:"sound_interval,omitempty"`

	// How many blocks within range of its avoid target the entity must be for it to begin sprinting away from the avoid target.
	SprintDistance float64 `json:"sprint_distance,omitempty"`

	// Multiplier for sprint speed. 1.0 means keep the regular speed, while higher numbers make the sprint speed faster.
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Multiplier for walking speed. 1.0 means keep the regular speed, while higher numbers make the walking speed faster.
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Enables the mob to barter for items that have been configured as barter currency. Must be used in combination with the barter component

type Behavior_Barter struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this mob to look at and follow the player that holds food they like.
type Behavior_Beg struct {
	// List of items that this mob likes
	Items []interface{} `json:"items,omitempty"`

	// Distance in blocks the mob will beg from
	LookDistance float64 `json:"look_distance,omitempty"`

	// The range of time in seconds this mob will stare at the player holding a food they like, begging for it
	LookTime float64 `json:"look_time,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this mob to break doors.

type Behavior_BreakDoor struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this mob to breed with other mobs.
type Behavior_Breed struct {
	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to celebrate surviving a raid by making celebration sounds and jumping.
type Behavior_Celebrate struct {
	// The sound event to trigger during the celebration.
	CelebrationSound string `json:"celebration_sound,omitempty"`

	// The duration in seconds that the celebration lasts for.
	Duration float64 `json:"duration,omitempty"`

	// Minimum and maximum time between jumping (positive, in seconds).
	JumpInterval float64 `json:"jump_interval,omitempty"`

	// [Trigger] The event to trigger when the goal's duration expires.
	OnCelebrationEndEvent interface{} `json:"on_celebration_end_event,omitempty"`

	// Minimum and maximum time between sound events (positive, in seconds).
	SoundInterval float64 `json:"sound_interval,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to celebrate surviving a raid by shooting fireworks.
type Behavior_CelebrateSurvive struct {
	// The duration in seconds that the celebration lasts for.
	Duration float64 `json:"duration,omitempty"`

	// Minimum and maximum time between firework (positive, in seconds).
	FireworksInterval float64 `json:"fireworks_interval,omitempty"`

	// [Trigger] The event to trigger when the goal's duration expires.
	OnCelebrationEndEvent interface{} `json:"on_celebration_end_event,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to damage a target by using a running attack.
type Behavior_ChargeAttack struct {
	// A charge attack cannot start if the entity is farther than this distance to the target.
	MaxDistance float64 `json:"max_distance,omitempty"`

	// A charge attack cannot start if the entity is closer than this distance to the target.
	MinDistance float64 `json:"min_distance,omitempty"`

	// Modifies the entity's speed when charging toward the target.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Percent chance this entity will start a charge attack, if not already attacking (1.0 = 100%)
	SuccessRate float64 `json:"success_rate,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows an entity to charge and use their held item.
type Behavior_ChargeHeldItem struct {
	// The list of items that can be used to charge the held item. This list is required and must have at least one item in it.
	Items []interface{} `json:"items,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Causes an entity to circle around an anchor point placed near a point or target.
type Behavior_CircleAroundAnchor struct {
	// Number of degrees to change this entity's facing by, when the entity selects its next anchor point.
	AngleChange float64 `json:"angle_change,omitempty"`

	// Maximum distance from the anchor-point in which this entity considers itself to have reached the anchor point. This is to prevent the entity from bouncing back and forth trying to reach a specific spot.
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The number of blocks above the target that the next anchor point can be set. This value is used only when the entity is tracking a target.
	HeightAboveTargetRange float64 `json:"height_above_target_range,omitempty"`

	// Percent chance to determine how often to increase or decrease the current height around the anchor point. 1 = 100%. "height_change_chance" is deprecated and has been replaced with "height_adjustment_chance".
	HeightAdjustmentChance float64 `json:"height_adjustment_chance,omitempty"`

	// Vertical distance from the anchor point this entity must stay within, upon a successful height adjustment.
	HeightOffsetRange float64 `json:"height_offset_range,omitempty"`

	// Percent chance to determine how often to increase the size of the current movement radius around the anchor point. 1 = 100%. "radius_change_chance" is deprecated and has been replaced with "radius_adjustment_chance".
	RadiusAdjustmentChance float64 `json:"radius_adjustment_chance,omitempty"`

	// The number of blocks to increase the current movement radius by, upon successful "radius_adjustment_chance". If the current radius increases over the range maximum, the current radius will be set back to the range minimum and the entity will change between clockwise and counter-clockwise movement..
	RadiusChange float64 `json:"radius_change,omitempty"`

	// Horizontal distance from the anchor point this entity must stay within upon a successful radius adjustment.
	RadiusRange float64 `json:"radius_range,omitempty"`

	// Multiplies the speed at which this entity travels to its next desired position.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to be controlled by the player using an item in the item_controllable property (required). Also requires the minecraft:movement property, and the minecraft:rideable property. On every tick, the entity will attempt to rotate towards where the player is facing with the control item whilst simultaneously moving forward.
type Behavior_ControlledByPlayer struct {
	// The entity will attempt to rotate to face where the player is facing each tick. The entity will target this percentage of their difference in their current facing angles each tick (from 0.0 to 1.0 where 1.0 = 100%). This is limited by FractionalRotationLimit. A value of 0.0 will result in the entity no longer turning to where the player is facing.
	FractionalRotation float64 `json:"fractional_rotation,omitempty"`

	// Limits the total degrees the entity can rotate to face where the player is facing on each tick.
	FractionalRotationLimit float64 `json:"fractional_rotation_limit,omitempty"`

	// Speed multiplier of mount when controlled by player.
	MountSpeedMultiplier float64 `json:"mount_speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to croak at a random time interval with configurable conditions.
type Behavior_Croak struct {
	// Random range in seconds after which the croaking stops. Can also be a constant.
	Duration float64 `json:"duration,omitempty"`

	// Conditions for the behavior to start and keep running. The interval between runs only starts after passing the filters.
	Filters *f.Filter `json:"filters,omitempty"`

	// Random range in seconds between runs of this behavior. Can also be a constant.
	Interval float64 `json:"interval,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of entity types that this mob considers valid targets
type BehaviorDefendTrustedTargetEntityTypes struct {
	// The amount of time in seconds that the mob has to wait before selecting a target of the same type again
	Cooldown float64 `json:"cooldown,omitempty"`

	// Conditions that make this entry in the list valid
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum distance this mob can be away to be a valid choice
	MaxDist float64 `json:"max_dist,omitempty"`

	// If true, the mob has to be visible to be a valid choice
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// If true, the mob will stop being targeted if it stops meeting any conditions.
	ReevaluateDescription bool `json:"reevaluate_description,omitempty"`

	// Multiplier for the running speed. A value of 1.0 means the speed is unchanged
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the mob to target another mob that hurts an entity it trusts.
type Behavior_DefendTrustedTarget struct {
	// Sound to occasionally play while defending.
	AggroSound string `json:"aggro_sound,omitempty"`

	// Time in seconds between attacks
	AttackInterval int `json:"attack_interval,omitempty"`

	// Corresponding Type: BehaviorDefendTrustedTargetEntityTypes
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// If true, only entities in this mob's viewing range can be selected as targets
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// Distance in blocks that the target can be within to launch an attack
	WithinRadius float64 `json:"within_radius,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to stay in a village and defend the village from aggressors. If a player is in bad standing with the village this goal will cause the entity to attack the player regardless of filter conditions.
type Behavior_DefendVillageTarget struct {
	// The percentage chance that the entity has to attack aggressors of its village, where 1.0 = 100%.
	AttackChance float64 `json:"attack_chance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows an entity to attack, while also delaying the damage-dealt until a specific time in the attack animation.
type Behavior_DelayedAttack struct {
	// The entity's attack animation will play out over this duration (in seconds). Also controls attack cooldown.
	AttackDuration float64 `json:"attack_duration,omitempty"`

	// Allows the entity to use this attack behavior, only once EVER.
	AttackOnce bool `json:"attack_once,omitempty"`

	// Defines the entity types this entity will attack.
	AttackTypes string `json:"attack_types,omitempty"`

	// Cooldown time (in seconds) between attacks.
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// The percentage into the attack animation to apply the damage of the attack (1.0 = 100%).
	HitDelayPct float64 `json:"hit_delay_pct,omitempty"`

	// Time (in seconds) to add to attack path recalculation when the target is beyond the "path_inner_boundary".
	InnerBoundaryTimeIncrease float64 `json:"inner_boundary_time_increase,omitempty"`

	// Unused. No effect on "minecraft:behavior.melee_attack".
	MaxDist float64 `json:"max_dist,omitempty"`

	// Maximum base time (in seconds) to recalculate new attack path to target (before increases applied).
	MaxPathTime float64 `json:"max_path_time,omitempty"`

	// Field of view (in degrees) when using the sensing component to detect an attack target.
	MeleeFov float64 `json:"melee_fov,omitempty"`

	// Minimum base time (in seconds) to recalculate new attack path to target (before increases applied).
	MinPathTime float64 `json:"min_path_time,omitempty"`

	// [Trigger] Defines the event to trigger when this entity successfully attacks.
	OnAttack interface{} `json:"on_attack,omitempty"`

	// Time (in seconds) to add to attack path recalculation when the target is beyond the "path_outer_boundary".
	OuterBoundaryTimeIncrease float64 `json:"outer_boundary_time_increase,omitempty"`

	// Time (in seconds) to add to attack path recalculation when this entity cannot move along the current path.
	PathFailTimeIncrease float64 `json:"path_fail_time_increase,omitempty"`

	// Distance at which to increase attack path recalculation by "inner_boundary_tick_increase".
	PathInnerBoundary float64 `json:"path_inner_boundary,omitempty"`

	// Distance at which to increase attack path recalculation by "outer_boundary_tick_increase".
	PathOuterBoundary float64 `json:"path_outer_boundary,omitempty"`

	// This entity will have a 1 in N chance to stop it's current attack, where N = "random_stop_interval".
	RandomStopInterval int `json:"random_stop_interval,omitempty"`

	// Used with the base size of the entity to determine minimum target-distance before trying to deal attack damage.
	ReachMultiplier float64 `json:"reach_multiplier,omitempty"`

	// Toggles (on/off) the need to have a full path from the entity to the target when using this melee attack behavior.
	RequireCompletePath bool `json:"require_complete_path,omitempty"`

	// Allows the actor to be set to persist upon targeting a player
	SetPersistent bool `json:"set_persistent,omitempty"`

	// This multiplier modifies the attacking entity's speed when moving toward the target.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Unused. No effect on "minecraft:behavior.melee_attack".
	TargetDist float64 `json:"target_dist,omitempty"`

	// Allows the entity to track the attack target, even if the entity has no sensing.
	TrackTarget bool `json:"track_target,omitempty"`

	// Maximum rotation (in degrees), on the X-axis, this entity can rotate while trying to look at the target.
	XMaxRotation float64 `json:"x_max_rotation,omitempty"`

	// Maximum rotation (in degrees), on the Y-axis, this entity can rotate its head while trying to look at the target.
	YMaxHeadRotation float64 `json:"y_max_head_rotation,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to dig into the ground before despawning.
type Behavior_Dig struct {
	// If true, this behavior can run when this entity is named. Otherwise not.
	AllowDigWhenNamed bool `json:"allow_dig_when_named,omitempty"`

	// Indicates that the actor should start digging when it sees daylight
	DigsInDaylight bool `json:"digs_in_daylight,omitempty"`

	// Goal duration in seconds
	Duration float64 `json:"duration,omitempty"`

	// The minimum idle time in seconds between the last detected disturbance to the start of digging.
	IdleTime float64 `json:"idle_time,omitempty"`

	// [Trigger] Event(s) to run when the goal starts.
	OnStart interface{} `json:"on_start,omitempty"`

	// If true, finding new suspicious locations count as disturbances that may delay the start of this goal.
	SuspicionIsDisturbance bool `json:"suspicion_is_disturbance,omitempty"`

	// If true, vibrations count as disturbances that may delay the start of this goal.
	VibrationIsDisturbance bool `json:"vibration_is_disturbance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to open and close doors.

type Behavior_DoorInteract struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to attack a player by charging at them. The player is chosen by the "minecraft:behavior.dragonscanning". Can only be used by the Ender Dragon.
type Behavior_Dragonchargeplayer struct {
	// The speed this entity moves when this behavior has started or while it's active.
	ActiveSpeed float64 `json:"active_speed,omitempty"`

	// If the dragon is outside the "target_zone" for longer than "continue_charge_threshold_time" seconds, the charge is canceled.
	ContinueChargeThresholdTime float64 `json:"continue_charge_threshold_time,omitempty"`

	// The speed this entity moves while this behavior is not active.
	FlightSpeed float64 `json:"flight_speed,omitempty"`

	// Minimum and maximum distance, from the target, this entity can use this behavior.
	TargetZone float64 `json:"target_zone,omitempty"`

	// The speed at which this entity turns while using this behavior.
	TurnSpeed float64 `json:"turn_speed,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the dragon to go out with glory. This controls the Ender Dragon's death animation and can't be used by other mobs.

type Behavior_Dragondeath struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to use a flame-breath attack. Can only be used by the Ender Dragon.
type Behavior_Dragonflaming struct {
	// Time (in seconds), after roar, to breath flame.
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// Time (in seconds), after roar, to breath flame.
	FlameTime float64 `json:"flame_time,omitempty"`

	// Number of ground flame-breath attacks to use before flight-takeoff.
	GroundFlameCount int `json:"ground_flame_count,omitempty"`

	// Time (in seconds) to roar, before breathing flame.
	RoarTime float64 `json:"roar_time,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the Dragon to fly around in a circle around the center podium. Can only be used by the Ender Dragon.

type Behavior_Dragonholdingpattern struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the Dragon to stop flying and transition into perching mode. Can only be used by the Ender Dragon.

type Behavior_Dragonlanding struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the dragon to look around for a player to attack while in perch mode. Can only be used by the Ender Dragon.

type Behavior_Dragonscanning struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to fly around looking for a player to shoot fireballs at. Can only be used by the Ender Dragon.
type Behavior_Dragonstrafeplayer struct {
	// The speed this entity moves when this behavior has started or while it's active.
	ActiveSpeed float64 `json:"active_speed,omitempty"`

	// Maximum distance of this entity's fireball attack while strafing.
	FireballRange float64 `json:"fireball_range,omitempty"`

	// The speed this entity moves while this behavior is not active.
	FlightSpeed float64 `json:"flight_speed,omitempty"`

	// Percent chance to to switch this entity's strafe direction between clockwise and counterclockwise. Switch direction chance occurs each time a new target is chosen (1.0 = 100%).
	SwitchDirectionProbability float64 `json:"switch_direction_probability,omitempty"`

	// Time (in seconds) the target must be in fireball range, and in view [ie, no solid terrain in-between the target and this entity], before a fireball can be shot.
	TargetInRangeAndInViewTime float64 `json:"target_in_range_and_in_view_time,omitempty"`

	// Minimum and maximum distance, from the target, this entity can use this behavior.
	TargetZone float64 `json:"target_zone,omitempty"`

	// The speed at which this entity turns while using this behavior.
	TurnSpeed float64 `json:"turn_speed,omitempty"`

	// The target must be within "view_angle" degrees of the dragon's current rotation before a fireball can be shot.
	ViewAngle float64 `json:"view_angle,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the dragon to leave perch mode and go back to flying around. Can only be used by the Ender Dragon.

type Behavior_Dragontakeoff struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to drink milk based on specified environment conditions.
type Behavior_DrinkMilk struct {
	// Time (in seconds) that the goal is on cooldown before it can be used again.
	CooldownSeconds float64 `json:"cooldown_seconds,omitempty"`

	// Conditions that need to be met for the behavior to start.
	Filters *f.Filter `json:"filters,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] A list of potions that this entity can drink. Each potion entry has the following parameters:
type BehaviorDrinkPotionPotions struct {
	// The percent chance (from 0.0 to 1.0) of this potion being selected when searching for a potion to use.
	Chance float64 `json:"chance,omitempty"`

	// The filters to use when determining if this potion can be selected.
	Filters *f.Filter `json:"filters,omitempty"`

	// The registry ID of the potion to use
	Id int `json:"id,omitempty"`
}

// Allows the mob to drink potions based on specified environment conditions.
type Behavior_DrinkPotion struct {
	// A list of potions that this entity can drink. Each potion entry has the following parameters:
	Potions []BehaviorDrinkPotionPotions `json:"potions,omitempty"`

	// The movement speed modifier to apply to the entity while it is drinking a potion. A value of 0 represents no change in speed.
	SpeedModifier float64 `json:"speed_modifier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to move toward a target, and drop an item near the target. This goal requires a "minecraft:navigation" to execute.
type Behavior_DropItemFor struct {
	// Total time that the goal is on cooldown before it can be used again.
	Cooldown float64 `json:"cooldown,omitempty"`

	// The percent chance the entity will drop an item when using this goal.
	DropItemChance float64 `json:"drop_item_chance,omitempty"`

	// The list of conditions another entity must meet to be a valid target to drop an item for.
	EntityTypes *f.Filter `json:"entity_types,omitempty"`

	// Distance in blocks within the entity considers it has reached it's target position.
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The loot table that contains the possible loot the entity can drop with this goal.
	LootTable string `json:"loot_table,omitempty"`

	// The maximum height the entities head will look at when dropping the item. The entity will always be looking at its target.
	MaxHeadLookAtHeight float64 `json:"max_head_look_at_height,omitempty"`

	// If the target position is farther away than this distance on any tick, the entity will teleport to the target position.
	MinimumTeleportDistance float64 `json:"minimum_teleport_distance,omitempty"`

	// The preferred distance the entity tries to be from the target it is dropping an item for.
	OfferingDistance float64 `json:"offering_distance,omitempty"`

	// [Trigger] The event to trigger when the entity attempts to drop an item.
	OnDropAttempt interface{} `json:"on_drop_attempt,omitempty"`

	// The number of blocks each tick that the entity will check within its search range and height for a valid block to move to. A value of 0 will have the mob check every block within range in one tick.
	SearchCount int `json:"search_count,omitempty"`

	// The Height in blocks the entity will search within to find a valid target position.
	SearchHeight int `json:"search_height,omitempty"`

	// The distance in blocks the entity will search within to find a valid target position.
	SearchRange int `json:"search_range,omitempty"`

	// The numbers of seconds that will pass before the dropped entity can be picked up from the ground.
	SecondsBeforePickup float64 `json:"seconds_before_pickup,omitempty"`

	// Movement speed multiplier of the entity when using this Goal.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// [Vector3 [a,b,c]] The range in blocks within which the entity searches to find a target to drop an item for.
	TargetRange []float64 `json:"target_range,omitempty"`

	// [Vector3 [a,b,c]] When the entity teleports, offset the teleport position by this many blocks in the X, Y, and Z coordinate.
	TeleportOffset []float64 `json:"teleport_offset,omitempty"`

	// The valid times of day that this goal can be used. For reference: noon is 0.0, sunset is 0.25, midnight is 0.5, and sunrise is 0.75, and back to noon for 1.0.
	TimeOfDayRange float64 `json:"time_of_day_range,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to consume a block, replace the eaten block with another block, and trigger an event as a result.
type Behavior_EatBlock struct {
	// A collection of pairs of blocks; the first ("eat_block")is the block the entity should eat, the second ("replace_block") is the block that should replace the eaten block.
	EatAndReplaceBlockPairs []interface{} `json:"eat_and_replace_block_pairs,omitempty"`

	// [Trigger] The event to trigger when the block eating animation has completed.
	OnEat interface{} `json:"on_eat,omitempty"`

	// [Molang String] A molang expression defining the success chance the entity has to consume a block.
	SuccessChance string `json:"success_chance,omitempty"`

	// The amount of time (in seconds) it takes for the block to be eaten upon a successful eat attempt.
	TimeUntilEat float64 `json:"time_until_eat,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// If the mob is carrying a food item, the mob will eat it and the effects will be applied to the mob.
type Behavior_EatCarriedItem struct {
	// Time in seconds the mob should wait before eating the item.
	DelayBeforeEating float64 `json:"delay_before_eating,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to eat a specified Mob.
type Behavior_EatMob struct {
	// Sets the time in seconds the eat animation should play for.
	EatAnimationTime float64 `json:"eat_animation_time,omitempty"`

	// Sets the sound that should play when eating a mob.
	EatMobSound string `json:"eat_mob_sound,omitempty"`

	// The loot table for loot to be dropped when eating a mob.
	LootTable string `json:"loot_table,omitempty"`

	// Sets the force which the mob-to-be-eaten is pulled towards the eating mob.
	PullInForce float64 `json:"pull_in_force,omitempty"`

	// Sets the desired distance to be reached before eating the mob.
	ReachMobDistance float64 `json:"reach_mob_distance,omitempty"`

	// Sets the entity's speed when running toward the target.
	RunSpeed float64 `json:"run_speed,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to emerge from the ground
type Behavior_Emerge struct {
	// Time in seconds the mob has to wait before using the goal again
	CooldownTime int `json:"cooldown_time,omitempty"`

	// Goal duration in seconds
	Duration float64 `json:"duration,omitempty"`

	// [Trigger] Trigger to be executed when the goal execution is about to end
	OnDone interface{} `json:"on_done,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the enderman to drop a block they are carrying. Can only be used by Endermen.

type Behavior_EndermanLeaveBlock struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the enderman to take a block and carry it around. Can only be used by Endermen.

type Behavior_EndermanTakeBlock struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// The entity puts on the desired equipment.

type Behavior_EquipItem struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to first travel to a random point on the outskirts of the village, and then explore random points within a small distance. This goal requires "minecraft:dweller" and "minecraft:navigation" to execute.
type Behavior_ExploreOutskirts struct {
	// [Vector3 [a,b,c]] The distance from the boundary the villager must be within in to explore the outskirts.
	DistFromBoundary []float64 `json:"dist_from_boundary,omitempty"`

	// Total distance in blocks the the entity will explore beyond the village bounds when choosing its travel point.
	ExploreDist float64 `json:"explore_dist,omitempty"`

	// This is the maximum amount of time an entity will attempt to reach it's travel point on the outskirts of the village before the goal exits.
	MaxTravelTime float64 `json:"max_travel_time,omitempty"`

	// The wait time in seconds between choosing new explore points will be chosen on a random interval between this value and the minimum wait time. This value is also the total amount of time the entity will explore random points before the goal stops.
	MaxWaitTime float64 `json:"max_wait_time,omitempty"`

	// The entity must be within this distance for it to consider it has successfully reached its target.
	MinDistFromTarget float64 `json:"min_dist_from_target,omitempty"`

	// The minimum perimeter of the village required to run this goal.
	MinPerimeter float64 `json:"min_perimeter,omitempty"`

	// The wait time in seconds between choosing new explore points will be chosen on a random interval between this value and the maximum wait time.
	MinWaitTime float64 `json:"min_wait_time,omitempty"`

	// A new explore point will randomly be chosen within this XZ distance of the current target position when navigation has finished and the wait timer has elapsed.
	NextXz int `json:"next_xz,omitempty"`

	// A new explore point will randomly be chosen within this Y distance of the current target position when navigation has finished and the wait timer has elapsed.
	NextY int `json:"next_y,omitempty"`

	// The multiplier for speed while using this goal. 1.0 maintains the speed.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Each new explore point will be chosen on a random interval between the minimum and the maximum wait time, divided by this value. This does not apply to the first explore point chosen when the goal runs.
	TimerRatio float64 `json:"timer_ratio,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [EXPERIMENTAL] Allows this entity to feel happy for a specified duration
type Behavior_FeelingHappy struct {
	// Goal cooldown range in seconds
	CooldownRange float64 `json:"cooldown_range,omitempty"`

	// Goal duration range in seconds
	DurationRange float64 `json:"duration_range,omitempty"`

	// [Trigger] Event(s) to run when the goal end.
	OnEnd interface{} `json:"on_end,omitempty"`

	// [Trigger] Event(s) to run when the goal starts.
	OnStart interface{} `json:"on_start,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to search within an area for a growable crop block. If found, the mob will use any available fertilizer in their inventory on the crop. This goal will not execute if the mob does not have a fertilizer item in its inventory.
type Behavior_FertilizeFarmBlock struct {
	// Distance in blocks within the mob considers it has reached it's target position.
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The maximum number of times the mob will use fertilzer on the target block.
	MaxFertilizerUsage int `json:"max_fertilizer_usage,omitempty"`

	// The maximum amount of time in seconds that the goal can take before searching again. The time is chosen between 0 and this number.
	SearchCooldownMaxSeconds float64 `json:"search_cooldown_max_seconds,omitempty"`

	// The number of randomly selected blocks each tick that the mob will check within its search range and height for a valid block to move to. A value of 0 will have the mob check every block within range in one tick.
	SearchCount int `json:"search_count,omitempty"`

	// The Height in blocks the mob will search within to find a valid target position.
	SearchHeight int `json:"search_height,omitempty"`

	// The distance in blocks the mob will search within to find a valid target position.
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this Goal.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to seek shade.
type Behavior_FindCover struct {
	// Time in seconds the mob has to wait before using the goal again
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to look around for another mob to ride atop it.
type Behavior_FindMount struct {
	// If true, the mob will not go into water blocks when going towards a mount
	AvoidWater bool `json:"avoid_water,omitempty"`

	// This is the distance the mob needs to be, in blocks, from the desired mount to mount it. If the value is below 0, the mob will use its default attack distance
	MountDistance float64 `json:"mount_distance,omitempty"`

	// Time the mob will wait before starting to move towards the mount
	StartDelay int `json:"start_delay,omitempty"`

	// If true, the mob will only look for a mount if it has a target
	TargetNeeded bool `json:"target_needed,omitempty"`

	// Distance in blocks within which the mob will look for a mount
	WithinRadius float64 `json:"within_radius,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to move towards the nearest underwater ruin or shipwreck.
type Behavior_FindUnderwaterTreasure struct {
	// The range that the mob will search for a treasure chest within a ruin or shipwreck to move towards.
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The distance the mob will move before stopping.
	StopDistance float64 `json:"stop_distance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to run away from direct sunlight and seek shade.
type Behavior_FleeSun struct {
	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to stay afloat while swimming. Passengers will be kicked out the moment the mob's head goes underwater, which may not happen for tall mobs.
type Behavior_Float struct {
	// If true, the mob will keep sinking as long as it has passengers.
	SinkWithPassengers bool `json:"sink_with_passengers,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to float around like the Ghast.
type Behavior_FloatWander struct {
	// Range of time in seconds the mob will float around before landing and choosing to do something else
	FloatDuration float64 `json:"float_duration,omitempty"`

	// If true, the point has to be reachable to be a valid target
	MustReach bool `json:"must_reach,omitempty"`

	// If true, the mob will randomly pick a new point while moving to the previously selected one
	RandomReselect bool `json:"random_reselect,omitempty"`

	// Distance in blocks on ground that the mob will look for a new spot to move to. Must be at least 1
	XzDist int `json:"xz_dist,omitempty"`

	// Distance in blocks that the mob will look up or down for a new spot to move to. Must be at least 1
	YDist int `json:"y_dist,omitempty"`

	// Height in blocks to add to the selected target position
	YOffset float64 `json:"y_offset,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of entity types that this mob can follow in a caravan
type BehaviorFollowCaravanEntityTypes struct {
	// The amount of time in seconds that the mob has to wait before selecting a target of the same type again
	Cooldown float64 `json:"cooldown,omitempty"`

	// Conditions that make this entry in the list valid
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum distance this mob can be away to be a valid choice
	MaxDist float64 `json:"max_dist,omitempty"`

	// If true, the mob has to be visible to be a valid choice
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// If true, the mob will stop being targeted if it stops meeting any conditions.
	ReevaluateDescription bool `json:"reevaluate_description,omitempty"`

	// Multiplier for the running speed. A value of 1.0 means the speed is unchanged
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the mob to follow mobs that are in a caravan.
type Behavior_FollowCaravan struct {
	// Number of entities that can be in the caravan
	EntityCount int `json:"entity_count,omitempty"`

	// Corresponding Type: BehaviorFollowCaravanEntityTypes
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to follow other mobs.
type Behavior_FollowMob struct {
	// The distance in blocks it will look for a mob to follow
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The distance in blocks this mob stops from the mob it is following
	StopDistance float64 `json:"stop_distance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to follow the player that owns them.
type Behavior_FollowOwner struct {
	// Specify if the mob can teleport to the player if it is too far away
	CanTeleport bool `json:"can_teleport,omitempty"`

	// Specify if the mob will follow the owner if it has heard a vibration lately
	IgnoreVibration bool `json:"ignore_vibration,omitempty"`

	// The maximum distance in blocks this mob can be from its owner to start following, only used when canTeleport is false
	MaxDistance float64 `json:"max_distance,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The distance in blocks that the owner can be away from this mob before it starts following it
	StartDistance float64 `json:"start_distance,omitempty"`

	// The distance in blocks this mob will stop from its owner while following it
	StopDistance float64 `json:"stop_distance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to follow their parent around.
type Behavior_FollowParent struct {
	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows mob to move towards its current target captain.
type Behavior_FollowTargetCaptain struct {
	// Defines the distance in blocks the mob will stay from its target while following.
	FollowDistance float64 `json:"follow_distance,omitempty"`

	// Defines the maximum distance in blocks a mob can get from its target captain before giving up trying to follow it.
	WithinRadius float64 `json:"within_radius,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// The entity will attempt to toss the items from its inventory to a nearby recently played noteblock.
type Behavior_GoAndGiveItemsToNoteblock struct {
	// Sets the time an entity should continue delivering items to a noteblock after hearing it.
	ListenTime int `json:"listen_time,omitempty"`

	// [Trigger] Event(s) to run when this mob throws items.
	OnItemThrow interface{} `json:"on_item_throw,omitempty"`

	// Sets the desired distance to be reached before throwing the items towards the block.
	ReachBlockDistance float64 `json:"reach_block_distance,omitempty"`

	// Sets the entity's speed when running toward the block.
	RunSpeed float64 `json:"run_speed,omitempty"`

	// Sets the throw force.
	ThrowForce float64 `json:"throw_force,omitempty"`

	// Sound to play when this mob throws an item.
	ThrowSound string `json:"throw_sound,omitempty"`

	// Sets the vertical throw multiplier that is applied on top of the throw force in the vertical direction.
	VerticalThrowMul float64 `json:"vertical_throw_mul,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// The entity will attempt to toss the items from its inventory to its owner.
type Behavior_GoAndGiveItemsToOwner struct {
	// [Trigger] Event(s) to run when this mob throws items.
	OnItemThrow interface{} `json:"on_item_throw,omitempty"`

	// Sets the desired distance to be reached before giving items to owner.
	ReachMobDistance float64 `json:"reach_mob_distance,omitempty"`

	// Sets the entity's speed when running toward the owner.
	RunSpeed float64 `json:"run_speed,omitempty"`

	// Sets the throw force.
	ThrowForce float64 `json:"throw_force,omitempty"`

	// Sound to play when this mob throws an item.
	ThrowSound string `json:"throw_sound,omitempty"`

	// Sets the vertical throw multiplier that is applied on top of the throw force in the vertical direction.
	VerticalThrowMul float64 `json:"vertical_throw_mul,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to move back to the position they were spawned.
type Behavior_GoHome struct {
	// Distance in blocks that the mob is considered close enough to the end of the current path. A new path will then be calculated to continue toward home.
	CalculateNewPathRadius float64 `json:"calculate_new_path_radius,omitempty"`

	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot.
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// A random value to determine when to randomly move somewhere. This has a 1/interval chance to choose this goal.
	Interval int `json:"interval,omitempty"`

	// [Trigger] Event(s) to run when this goal fails.
	OnFailed interface{} `json:"on_failed,omitempty"`

	// [Trigger] Event(s) to run when this mob gets home.
	OnHome interface{} `json:"on_home,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to use a laser beam attack. Can only be used by Guardians and Elder Guardians.
type Behavior_GuardianAttack struct {
	// Amount of additional damage dealt from an elder guardian's magic attack.
	ElderExtraMagicDamage int `json:"elder_extra_magic_damage,omitempty"`

	// In hard difficulty, amount of additional damage dealt from a guardian's magic attack.
	HardModeExtraMagicDamage int `json:"hard_mode_extra_magic_damage,omitempty"`

	// Amount of damage dealt from a guardian's magic attack. Magic attack damage is added to the guardian's base attack damage.
	MagicDamage int `json:"magic_damage,omitempty"`

	// Guardian attack behavior stops if the target is closer than this distance (doesn't apply to elders).
	MinDistance float64 `json:"min_distance,omitempty"`

	// Time (in seconds) to wait after starting an attack before playing the guardian attack sound.
	SoundDelayTime float64 `json:"sound_delay_time,omitempty"`

	// Maximum rotation (in degrees), on the X-axis, this entity can rotate while trying to look at the target.
	XMaxRotation float64 `json:"x_max_rotation,omitempty"`

	// Maximum rotation (in degrees), on the Y-axis, this entity can rotate its head while trying to look at the target.
	YMaxHeadRotation float64 `json:"y_max_head_rotation,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to search within an area for farmland with air above it. If found, the entity will replace the air block by planting a seed item from its inventory on the farmland block. This goal requires "minecraft:inventory" and "minecraft:navigation" to execute. This goal will not execute if the entity does not have an item in its inventory.
type Behavior_HarvestFarmBlock struct {
	// Distance in blocks within the entity considers it has reached it's target position.
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The maximum amount of time in seconds that the goal can take before searching for the first harvest block. The time is chosen between 0 and this number.
	MaxSecondsBeforeSearch float64 `json:"max_seconds_before_search,omitempty"`

	// The maximum amount of time in seconds that the goal can take before searching again, after failing to find a a harvest block already. The time is chosen between 0 and this number.
	SearchCooldownMaxSeconds float64 `json:"search_cooldown_max_seconds,omitempty"`

	// The number of randomly selected blocks each tick that the entity will check within its search range and height for a valid block to move to. A value of 0 will have the mob check every block within range in one tick.
	SearchCount int `json:"search_count,omitempty"`

	// The Height in blocks the entity will search within to find a valid target position.
	SearchHeight int `json:"search_height,omitempty"`

	// The distance in blocks the entity will search within to find a valid target position.
	SearchRange int `json:"search_range,omitempty"`

	// The amount of time in seconds that the goal will cooldown after a successful reap/sow, before it can start again.
	SecondsUntilNewTask float64 `json:"seconds_until_new_task,omitempty"`

	// Movement speed multiplier of the entity when using this Goal.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows a mob with the hide component to attempt to move to - and hide at - an owned or nearby POI.
type Behavior_Hide struct {
	// Amount of time in seconds that the mob reacts.
	Duration float64 `json:"duration,omitempty"`

	// Defines what POI type to hide at.
	PoiType string `json:"poi_type,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The cooldown time in seconds before the goal can be reused after a internal failure or timeout condition.
	TimeoutCooldown float64 `json:"timeout_cooldown,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// The mob freezes and looks at the mob they are targeting.
type Behavior_HoldGround struct {
	// Whether to broadcast out the mob's target to other mobs of the same type.
	Broadcast bool `json:"broadcast,omitempty"`

	// Range in blocks for how far to broadcast.
	BroadcastRange float64 `json:"broadcast_range,omitempty"`

	// Minimum distance the target must be for the mob to run this goal.
	MinRadius float64 `json:"min_radius,omitempty"`

	// Event to run when target is within the radius. This event is broadcasted if broadcast is true.
	WithinRadiusEvent string `json:"within_radius_event,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of entity types that this mob can target when hurt by them
type BehaviorHurtByTargetEntityTypes struct {
	// The amount of time in seconds that the mob has to wait before selecting a target of the same type again
	Cooldown float64 `json:"cooldown,omitempty"`

	// Conditions that make this entry in the list valid
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum distance this mob can be away to be a valid choice
	MaxDist float64 `json:"max_dist,omitempty"`

	// If true, the mob has to be visible to be a valid choice
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// If true, the mob will stop being targeted if it stops meeting any conditions.
	ReevaluateDescription bool `json:"reevaluate_description,omitempty"`

	// Multiplier for the running speed. A value of 1.0 means the speed is unchanged
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the mob to target another mob that hurts them.
type Behavior_HurtByTarget struct {
	// If true, nearby mobs of the same type will be alerted about the damage
	AlertSameType bool `json:"alert_same_type,omitempty"`

	// Corresponding Type: BehaviorHurtByTargetEntityTypes
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// If true, the mob will hurt its owner and other mobs with the same owner as itself
	HurtOwner bool `json:"hurt_owner,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to inspect bookshelves.
type Behavior_InspectBookshelf struct {
	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The number of blocks each tick that the mob will check within its search range and height for a valid block to move to. A value of 0 will have the mob check every block within range in one tick
	SearchCount int `json:"search_count,omitempty"`

	// The height that the mob will search for bookshelves
	SearchHeight int `json:"search_height,omitempty"`

	// Distance in blocks the mob will look for books to inspect
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to move towards a "suspicious" position based on data gathered in minecraft:suspect_tracking
type Behavior_InvestigateSuspiciousLocation struct {
	// Distance in blocks within the entity considers it has reached it's target position.
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// Movement speed multiplier
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows an entity to jump to another random block.
type Behavior_JumpToBlock struct {
	// Minimum and maximum cooldown time-range (positive, in seconds) between each attempted jump.
	CooldownRange float64 `json:"cooldown_range,omitempty"`

	// Blocks that the mob can't jump to.
	ForbiddenBlocks []interface{} `json:"forbidden_blocks,omitempty"`

	// The maximum velocity with which the mob can jump.
	MaxVelocity float64 `json:"max_velocity,omitempty"`

	// The minimum distance (in blocks) from the mob to a block, in order to consider jumping to it.
	MinimumDistance int `json:"minimum_distance,omitempty"`

	// The minimum length (in blocks) of the mobs path to a block, in order to consider jumping to it.
	MinimumPathLength int `json:"minimum_path_length,omitempty"`

	// Blocks that the mob prefers jumping to.
	PreferredBlocks []interface{} `json:"preferred_blocks,omitempty"`

	// Chance (between 0.0 and 1.0) that the mob will jump to a preferred block, if in range. Only matters if preferred blocks are defined.
	PreferredBlocksChance float64 `json:"preferred_blocks_chance,omitempty"`

	// The scalefactor of the bounding box of the mob while it is jumping.
	ScaleFactor float64 `json:"scale_factor,omitempty"`

	// The height (in blocks, in range [2, 15]) of the search box, centered around the mob.
	SearchHeight int `json:"search_height,omitempty"`

	// The width (in blocks, in range [2, 15]) of the search box, centered around the mob.
	SearchWidth int `json:"search_width,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to perform a damaging knockback that affects all nearby entities.
type Behavior_KnockbackRoar struct {
	// The delay after which the knockback occurs (in seconds).
	AttackTime float64 `json:"attack_time,omitempty"`

	// Time (in seconds) the mob has to wait before using the goal again.
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// The list of conditions another entity must meet to be a valid target to apply damage to.
	DamageFilters *f.Filter `json:"damage_filters,omitempty"`

	// The max duration of the roar (in seconds).
	Duration float64 `json:"duration,omitempty"`

	// The damage dealt by the knockback roar.
	KnockbackDamage int `json:"knockback_damage,omitempty"`

	// The list of conditions another entity must meet to be a valid target to apply knockback to.
	KnockbackFilters *f.Filter `json:"knockback_filters,omitempty"`

	// The maximum height for vertical knockback.
	KnockbackHeightCap float64 `json:"knockback_height_cap,omitempty"`

	// The strength of the horizontal knockback.
	KnockbackHorizontalStrength int `json:"knockback_horizontal_strength,omitempty"`

	// The radius (in blocks) of the knockback effect.
	KnockbackRange int `json:"knockback_range,omitempty"`

	// The strength of the vertical knockback.
	KnockbackVerticalStrength int `json:"knockback_vertical_strength,omitempty"`

	// [Trigger] Event that is triggered when the roar ends.
	OnRoarEnd interface{} `json:"on_roar_end,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows mobs to lay down at times
type Behavior_LayDown struct {
	// A random value to determine at what intervals something can occur. This has a 1/interval chance to choose this goal
	Interval int `json:"interval,omitempty"`

	// a random value in which the goal can use to pull out of the behavior. This is a 1/interval chance to play the sound
	RandomStopInterval int `json:"random_stop_interval,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to lay an egg block on certain types of blocks if the mob is pregnant.
type Behavior_LayEgg struct {
	// Allows the mob to lay its eggs from below the target if it can't get there. This is useful if the target block is water with air above, since mobs may not be able to get to the air block above water.
	AllowLayingFromBelow bool `json:"allow_laying_from_below,omitempty"`

	// Block type for the egg to lay. If this is a turtle egg, the number of eggs in the block is randomly set.
	EggType string `json:"egg_type,omitempty"`

	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// Name of the sound event played when laying the egg. Defaults to lay_egg, which is used for Turtles.
	LayEggSound string `json:"lay_egg_sound,omitempty"`

	// Duration of the laying egg process in seconds.
	LaySeconds float64 `json:"lay_seconds,omitempty"`

	// [Trigger] Event to run when this mob lays the egg.
	OnLay interface{} `json:"on_lay,omitempty"`

	// Height in blocks the mob will look for a target block to move towards
	SearchHeight int `json:"search_height,omitempty"`

	// The distance in blocks it will look for a target block to move towards
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Blocks that the mob can lay its eggs on top of.
	TargetBlocks []interface{} `json:"target_blocks,omitempty"`

	// Types of materials that can exist above the target block. Valid types are Air, Water, and Lava.
	TargetMaterialsAboveBlock []interface{} `json:"target_materials_above_block,omitempty"`

	// Specifies if the default lay-egg animation should be played when the egg is placed or not.
	UseDefaultAnimation bool `json:"use_default_animation,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows monsters to jump at and attack their target. Can only be used by hostile mobs.
type Behavior_LeapAtTarget struct {
	// If true, the mob will only jump at its target if its on the ground. Setting it to false will allow it to jump even if its already in the air
	MustBeOnGround bool `json:"must_be_on_ground,omitempty"`

	// Allows the actor to be set to persist upon targeting a player
	SetPersistent bool `json:"set_persistent,omitempty"`

	// The height in blocks the mob jumps when leaping at its target
	Yd float64 `json:"yd,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to look at nearby entities.
type Behavior_LookAtEntity struct {
	// The angle in degrees that the mob can see in the Y-axis (up-down)
	AngleOfViewHorizontal int `json:"angle_of_view_horizontal,omitempty"`

	// The angle in degrees that the mob can see in the X-axis (left-right)
	AngleOfViewVertical int `json:"angle_of_view_vertical,omitempty"`

	// Filter to determine the conditions for this mob to look at the entity
	Filters *f.Filter `json:"filters,omitempty"`

	// The distance in blocks from which the entity will look at
	LookDistance float64 `json:"look_distance,omitempty"`

	// Time range to look at the entity
	LookTime float64 `json:"look_time,omitempty"`

	// The probability of looking at the target. A value of 1.00 is 100%
	Probability float64 `json:"probability,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to look at the player when the player is nearby.
type Behavior_LookAtPlayer struct {
	// The angle in degrees that the mob can see in the Y-axis (up-down)
	AngleOfViewHorizontal int `json:"angle_of_view_horizontal,omitempty"`

	// The angle in degrees that the mob can see in the X-axis (left-right)
	AngleOfViewVertical int `json:"angle_of_view_vertical,omitempty"`

	// The distance in blocks from which the entity will look at
	LookDistance float64 `json:"look_distance,omitempty"`

	// Time range to look at the entity
	LookTime float64 `json:"look_time,omitempty"`

	// The probability of looking at the target. A value of 1.00 is 100%
	Probability float64 `json:"probability,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to look at the entity they are targetting.
type Behavior_LookAtTarget struct {
	// The angle in degrees that the mob can see in the Y-axis (up-down)
	AngleOfViewHorizontal int `json:"angle_of_view_horizontal,omitempty"`

	// The angle in degrees that the mob can see in the X-axis (left-right)
	AngleOfViewVertical int `json:"angle_of_view_vertical,omitempty"`

	// The distance in blocks from which the entity will look at
	LookDistance float64 `json:"look_distance,omitempty"`

	// Time range to look at the entity
	LookTime float64 `json:"look_time,omitempty"`

	// The probability of looking at the target. A value of 1.00 is 100%
	Probability float64 `json:"probability,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to look at the player they are trading with.
type Behavior_LookAtTradingPlayer struct {
	// The angle in degrees that the mob can see in the Y-axis (up-down)
	AngleOfViewHorizontal int `json:"angle_of_view_horizontal,omitempty"`

	// The angle in degrees that the mob can see in the X-axis (left-right)
	AngleOfViewVertical int `json:"angle_of_view_vertical,omitempty"`

	// The distance in blocks from which the entity will look at
	LookDistance float64 `json:"look_distance,omitempty"`

	// Time range to look at the entity
	LookTime float64 `json:"look_time,omitempty"`

	// The probability of looking at the target. A value of 1.00 is 100%
	Probability float64 `json:"probability,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the villager to look for a mate to spawn other villagers with. Can only be used by Villagers.

type Behavior_MakeLove struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows an entity to deal damage through a melee attack.
type Behavior_MeleeAttack struct {
	// Allows the entity to use this attack behavior, only once EVER.
	AttackOnce bool `json:"attack_once,omitempty"`

	// Defines the entity types this entity will attack.
	AttackTypes string `json:"attack_types,omitempty"`

	// Cooldown time (in seconds) between attacks.
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// Time (in seconds) to add to attack path recalculation when the target is beyond the "path_inner_boundary".
	InnerBoundaryTimeIncrease float64 `json:"inner_boundary_time_increase,omitempty"`

	// Unused. No effect on "minecraft:behavior.melee_attack".
	MaxDist float64 `json:"max_dist,omitempty"`

	// Maximum base time (in seconds) to recalculate new attack path to target (before increases applied).
	MaxPathTime float64 `json:"max_path_time,omitempty"`

	// Field of view (in degrees) when using the sensing component to detect an attack target.
	MeleeFov float64 `json:"melee_fov,omitempty"`

	// Minimum base time (in seconds) to recalculate new attack path to target (before increases applied).
	MinPathTime float64 `json:"min_path_time,omitempty"`

	// [Trigger] Defines the event to trigger when this entity successfully attacks.
	OnAttack interface{} `json:"on_attack,omitempty"`

	// Time (in seconds) to add to attack path recalculation when the target is beyond the "path_outer_boundary".
	OuterBoundaryTimeIncrease float64 `json:"outer_boundary_time_increase,omitempty"`

	// Time (in seconds) to add to attack path recalculation when this entity cannot move along the current path.
	PathFailTimeIncrease float64 `json:"path_fail_time_increase,omitempty"`

	// Distance at which to increase attack path recalculation by "inner_boundary_tick_increase".
	PathInnerBoundary float64 `json:"path_inner_boundary,omitempty"`

	// Distance at which to increase attack path recalculation by "outer_boundary_tick_increase".
	PathOuterBoundary float64 `json:"path_outer_boundary,omitempty"`

	// This entity will have a 1 in N chance to stop it's current attack, where N = "random_stop_interval".
	RandomStopInterval int `json:"random_stop_interval,omitempty"`

	// Used with the base size of the entity to determine minimum target-distance before trying to deal attack damage.
	ReachMultiplier float64 `json:"reach_multiplier,omitempty"`

	// Toggles (on/off) the need to have a full path from the entity to the target when using this melee attack behavior.
	RequireCompletePath bool `json:"require_complete_path,omitempty"`

	// Allows the actor to be set to persist upon targeting a player
	SetPersistent bool `json:"set_persistent,omitempty"`

	// This multiplier modifies the attacking entity's speed when moving toward the target.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Unused. No effect on "minecraft:behavior.melee_attack".
	TargetDist float64 `json:"target_dist,omitempty"`

	// Allows the entity to track the attack target, even if the entity has no sensing.
	TrackTarget bool `json:"track_target,omitempty"`

	// Maximum rotation (in degrees), on the X-axis, this entity can rotate while trying to look at the target.
	XMaxRotation float64 `json:"x_max_rotation,omitempty"`

	// Maximum rotation (in degrees), on the Y-axis, this entity can rotate its head while trying to look at the target.
	YMaxHeadRotation float64 `json:"y_max_head_rotation,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows an entity to go to the village bell and mingle with other entities
type Behavior_Mingle struct {
	// Time in seconds the mob has to wait before using the goal again
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// Amount of time in seconds that the entity will chat with another entity
	Duration float64 `json:"duration,omitempty"`

	// The distance from its partner that this entity will mingle. If the entity type is not the same as the entity, this value needs to be identical on both entities.
	MingleDistance float64 `json:"mingle_distance,omitempty"`

	// The entity type that this entity is allowed to mingle with
	MinglePartnerType []interface{} `json:"mingle_partner_type,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to move around on its own while mounted seeking a target to attack.
type Behavior_MountPathing struct {
	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The distance at which this mob wants to be away from its target
	TargetDist float64 `json:"target_dist,omitempty"`

	// If true, this mob will chase after the target as long as it's a valid target
	TrackTarget bool `json:"track_target,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to move indoors.
type Behavior_MoveIndoors struct {
	// The movement speed modifier to apply to the entity while it is moving indoors.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The cooldown time in seconds before the goal can be reused after pathfinding fails
	TimeoutCooldown float64 `json:"timeout_cooldown,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to move outdoors.
type Behavior_MoveOutdoors struct {
	// The radius away from the target block to count as reaching the goal.
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The amount of times to try finding a random outdoors position before failing.
	SearchCount int `json:"search_count,omitempty"`

	// The y range to search for an outdoors position for.
	SearchHeight int `json:"search_height,omitempty"`

	// The x and z range to search for an outdoors position for.
	SearchRange int `json:"search_range,omitempty"`

	// The movement speed modifier to apply to the entity while it is moving outdoors.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The cooldown time in seconds before the goal can be reused after pathfinding fails
	TimeoutCooldown float64 `json:"timeout_cooldown,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Can only be used by Villagers. Allows the villagers to create paths around the village.
type Behavior_MoveThroughVillage struct {
	// If true, the mob will only move through the village during night time
	OnlyAtNight bool `json:"only_at_night,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows mob to move towards a block.
type Behavior_MoveToBlock struct {
	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// [Trigger] Event to run on block reached.
	OnReach interface{} `json:"on_reach,omitempty"`

	// [Trigger] Event to run on completing a stay of stay_duration at the block.
	OnStayCompleted interface{} `json:"on_stay_completed,omitempty"`

	// The height in blocks that the mob will look for the block.
	SearchHeight int `json:"search_height,omitempty"`

	// The distance in blocks that the mob will look for the block.
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Chance to start the behavior (applied after each random tick_interval).
	StartChance float64 `json:"start_chance,omitempty"`

	// Number of ticks needed to complete a stay at the block.
	StayDuration float64 `json:"stay_duration,omitempty"`

	// Block types to move to.
	TargetBlocks []interface{} `json:"target_blocks,omitempty"`

	// [Vector3 [a,b,c]] Offset to add to the selected target position.
	TargetOffset []float64 `json:"target_offset,omitempty"`

	// Kind of block to find fitting the specification. Valid values are "random" and "nearest".
	TargetSelectionMethod string `json:"target_selection_method,omitempty"`

	// Average interval in ticks to try to run this behavior.
	TickInterval int `json:"tick_interval,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to move back onto land when in water.
type Behavior_MoveToLand struct {
	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The number of blocks each tick that the mob will check within its search range and height for a valid block to move to. A value of 0 will have the mob check every block within range in one tick
	SearchCount int `json:"search_count,omitempty"`

	// Height in blocks the mob will look for land to move towards
	SearchHeight int `json:"search_height,omitempty"`

	// The distance in blocks it will look for land to move towards
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to move back into lava when on land.
type Behavior_MoveToLava struct {
	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The number of blocks each tick that the mob will check within its search range and height for a valid block to move to. A value of 0 will have the mob check every block within range in one tick
	SearchCount int `json:"search_count,omitempty"`

	// Height in blocks the mob will look for lava to move towards
	SearchHeight int `json:"search_height,omitempty"`

	// The distance in blocks it will look for lava to move towards
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to move into a liquid when on land.
type Behavior_MoveToLiquid struct {
	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The material type of the liquid block to find. Valid values are "Any", "Water", and "Lava".
	MaterialType string `json:"material_type,omitempty"`

	// The number of blocks each tick that the mob will check within its search range and height for a valid block to move to. A value of 0 will have the mob check every block within range in one tick
	SearchCount int `json:"search_count,omitempty"`

	// Height in blocks the mob will look for the liquid block to move towards
	SearchHeight int `json:"search_height,omitempty"`

	// The distance in blocks it will look for the liquid block to move towards
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to move to a POI if able to
type Behavior_MoveToPoi struct {
	// Tells the goal what POI type it should be looking for
	PoiType string `json:"poi_type,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows mob to move towards a random block.
type Behavior_MoveToRandomBlock struct {
	// Defines the distance from the mob, in blocks, that the block to move to will be chosen.
	BlockDistance float64 `json:"block_distance,omitempty"`

	// Defines the distance in blocks the mob has to be from the block for the movement to be finished.
	WithinRadius float64 `json:"within_radius,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to move into a random location within a village.
type Behavior_MoveToVillage struct {
	// Time in seconds the mob has to wait before using the goal again
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The distance in blocks to search for villages. If <= 0, find the closest village regardless of distance.
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to move back into water when on land.
type Behavior_MoveToWater struct {
	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The number of blocks each tick that the mob will check within its search range and height for a valid block to move to. A value of 0 will have the mob check every block within range in one tick
	SearchCount int `json:"search_count,omitempty"`

	// Height in blocks the mob will look for water to move towards
	SearchHeight int `json:"search_height,omitempty"`

	// The distance in blocks it will look for water to move towards
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows mobs with the dweller component to move toward their Village area that the mob should be restricted to.
type Behavior_MoveTowardsDwellingRestriction struct {
	// This multiplier modifies the entity's speed when moving towards it's restriction.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows mobs with the home component to move toward their pre-defined area that the mob should be restricted to.
type Behavior_MoveTowardsHomeRestriction struct {
	// This multiplier modifies the entity's speed when moving towards it's restriction.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows mob to move towards its current target.
type Behavior_MoveTowardsTarget struct {
	// Defines the radius in blocks that the mob tries to be from the target. A value of 0 means it tries to occupy the same block as the target
	WithinRadius float64 `json:"within_radius,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows mobs to occassionally stop and take a nap under certain conditions.
type Behavior_Nap struct {
	// Maximum time in seconds the mob has to wait before using the goal again
	CooldownMax float64 `json:"cooldown_max,omitempty"`

	// Minimum time in seconds the mob has to wait before using the goal again
	CooldownMin float64 `json:"cooldown_min,omitempty"`

	// The block distance in x and z that will be checked for mobs that this mob detects
	MobDetectDist float64 `json:"mob_detect_dist,omitempty"`

	// The block distance in y that will be checked for mobs that this mob detects
	MobDetectHeight float64 `json:"mob_detect_height,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] Filters which types of targets are valid for this entity.
type BehaviorNearestAttackableTargetEntityTypes struct {
	// Conditions that make this target a valid type.
	Filters *f.Filter `json:"filters,omitempty"`

	// To be a valid target choice, the target type cannot be farther away from this entity than "max_dist".
	MaxDist float64 `json:"max_dist,omitempty"`

	// Determines if target-validity requires this entity to be in range only, or both in range and in sight.
	MustSee bool `json:"must_see,omitempty"`

	// Time (in seconds) the target must not be seen by this entity to become invalid. Used only if "must_see" is true.
	MustSeeForgetDuration bool `json:"must_see_forget_duration,omitempty"`

	// If true, the mob will stop being targeted if it stops meeting any conditions.
	ReevaluateDescription bool `json:"reevaluate_description,omitempty"`
}

// Allows an entity to attack the closest target within a given subset of specific target types.
type Behavior_NearestAttackableTarget struct {
	// Time range (in seconds) between searching for an attack target, range is in (0, "attack_interval"]. Only used if "attack_interval" is greater than 0, otherwise "scan_interval" is used.
	AttackInterval int `json:"attack_interval,omitempty"`

	// Alias for "attack_interval"; provides the same functionality as "attack_interval".
	AttackIntervalMin int `json:"attack_interval_min,omitempty"`

	// If true, this entity can attack its owner.
	AttackOwner bool `json:"attack_owner,omitempty"`

	// Corresponding Type: BehaviorNearestAttackableTargetEntityTypes
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// If true, this entity requires a path to the target.
	MustReach bool `json:"must_reach,omitempty"`

	// Determines if target-validity requires this entity to be in range only, or both in range and in sight.
	MustSee bool `json:"must_see,omitempty"`

	// Time (in seconds) the target must not be seen by this entity to become invalid. Used only if "must_see" is true.
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// Time (in seconds) this entity can continue attacking the target after the target is no longer valid.
	PersistTime float64 `json:"persist_time,omitempty"`

	// Allows the attacking entity to update the nearest target, otherwise a target is only reselected after each "scan_interval" or "attack_interval".
	ReselectTargets bool `json:"reselect_targets,omitempty"`

	// If "attack_interval" is 0 or isn't declared, then between attacks: scanning for a new target occurs every amount of ticks equal to "scan_interval", minimum value is 1. Values under 10 can affect performance.
	ScanInterval int `json:"scan_interval,omitempty"`

	// Allows the actor to be set to persist upon targeting a player
	SetPersistent bool `json:"set_persistent,omitempty"`

	// Multiplied with the target's armor coverage percentage to modify "max_dist" when detecting an invisible target.
	TargetInvisibleMultiplier float64 `json:"target_invisible_multiplier,omitempty"`

	// Maximum vertical target-search distance, if it's greater than the target type's "max_dist". A negative value defaults to "entity_types" greatest "max_dist".
	TargetSearchHeight float64 `json:"target_search_height,omitempty"`

	// Multiplied with the target type's "max_dist" when trying to detect a sneaking target.
	TargetSneakVisibilityMultiplier float64 `json:"target_sneak_visibility_multiplier,omitempty"`

	// Maximum distance this entity can be from the target when following it, otherwise the target becomes invalid. This value is only used if the entity doesn't declare "minecraft:follow_range".
	WithinRadius float64 `json:"within_radius,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of entity types that this mob considers valid targets
type BehaviorNearestPrioritizedAttackableTargetEntityTypes struct {
	// The amount of time in seconds that the mob has to wait before selecting a target of the same type again
	Cooldown float64 `json:"cooldown,omitempty"`

	// Conditions that make this entry in the list valid
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum distance this mob can be away to be a valid choice
	MaxDist float64 `json:"max_dist,omitempty"`

	// If true, the mob has to be visible to be a valid choice
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// If true, the mob will stop being targeted if it stops meeting any conditions.
	ReevaluateDescription bool `json:"reevaluate_description,omitempty"`

	// Multiplier for the running speed. A value of 1.0 means the speed is unchanged
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the mob to check for and pursue the nearest valid target.
type Behavior_NearestPrioritizedAttackableTarget struct {
	// Time in seconds before selecting a target
	AttackInterval int `json:"attack_interval,omitempty"`

	// The amount of time in seconds that the mob has to wait before selecting a target of the same type again
	Cooldown float64 `json:"cooldown,omitempty"`

	// Corresponding Type: BehaviorNearestPrioritizedAttackableTargetEntityTypes
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// If true, only entities that this mob can path to can be selected as targets
	MustReach bool `json:"must_reach,omitempty"`

	// If true, only entities in this mob's viewing range can be selected as targets
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// Time in seconds for a valid target to stay targeted when it becomes and invalid target.
	PersistTime float64 `json:"persist_time,omitempty"`

	// Specifies the priority in which filtered enemy types should be attacked. Lower number means higher priority.
	Priority int `json:"priority,omitempty"`

	// If true, the target will change to the current closest entity whenever a different entity is closer
	ReselectTargets bool `json:"reselect_targets,omitempty"`

	// How many ticks to wait between scanning for a target.
	ScanInterval int `json:"scan_interval,omitempty"`

	// Allows the actor to be set to persist upon targeting a player
	SetPersistent bool `json:"set_persistent,omitempty"`

	// Height in blocks to search for a target mob. -1.0f means the height does not matter.
	TargetSearchHeight float64 `json:"target_search_height,omitempty"`

	// Distance in blocks that the target can be within to launch an attack
	WithinRadius float64 `json:"within_radius,omitempty"`
}

// Allows to mob to be able to sit in place like the ocelot.
type Behavior_OcelotSitOnBlock struct {
	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows an entity to attack by sneaking and pouncing.
type Behavior_Ocelotattack struct {
	// Time (in seconds) between attacks.
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// Max distance from the target, this entity will use this attack behavior.
	MaxDistance float64 `json:"max_distance,omitempty"`

	// Max distance from the target, this entity starts sneaking.
	MaxSneakRange float64 `json:"max_sneak_range,omitempty"`

	// Max distance from the target, this entity starts sprinting (sprinting takes priority over sneaking).
	MaxSprintRange float64 `json:"max_sprint_range,omitempty"`

	// Used with the base size of the entity to determine minimum target-distance before trying to deal attack damage.
	ReachMultiplier float64 `json:"reach_multiplier,omitempty"`

	// Modifies the attacking entity's movement speed while sneaking.
	SneakSpeedMultiplier float64 `json:"sneak_speed_multiplier,omitempty"`

	// Modifies the attacking entity's movement speed while sprinting.
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Modifies the attacking entity's movement speed when not sneaking or sprinting, but still within attack range.
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`

	// Maximum rotation (in degrees), on the X-axis, this entity can rotate while trying to look at the target.
	XMaxRotation float64 `json:"x_max_rotation,omitempty"`

	// Maximum rotation (in degrees), on the Y-axis, this entity can rotate its head while trying to look at the target.
	YMaxHeadRotation float64 `json:"y_max_head_rotation,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to offer a flower to another mob with the minecraft:take_flower behavior.
type Behavior_OfferFlower struct {
	// Percent chance that the mob will start this goal from 0.0 to 1.0 (where 1.0 = 100%).
	ChanceToStart float64 `json:"chance_to_start,omitempty"`

	// Conditions that need to be met for the behavior to start.
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum rotation (in degrees), on the Y-axis, this entity can rotate its head while trying to look at the target.
	MaxHeadRotationY float64 `json:"max_head_rotation_y,omitempty"`

	// The max amount of time (in seconds) that the mob will offer the flower for before exiting the Goal.
	MaxOfferFlowerDuration float64 `json:"max_offer_flower_duration,omitempty"`

	// Maximum rotation (in degrees), on the X-axis, this entity can rotate while trying to look at the target.
	MaxRotationX float64 `json:"max_rotation_x,omitempty"`

	// [Vector3 [a,b,c]] The dimensions of the AABB used to search for a potential mob to offer flower to.
	SearchArea []float64 `json:"search_area,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to open doors. Requires the mob to be able to path through doors, otherwise the mob won't even want to try opening them.
type Behavior_OpenDoor struct {
	// If true, the mob will close the door after opening it and going through it
	CloseDoorAfter bool `json:"close_door_after,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of entity types that this mob can target if they hurt their owner
type BehaviorOwnerHurtByTargetEntityTypes struct {
	// The amount of time in seconds that the mob has to wait before selecting a target of the same type again
	Cooldown float64 `json:"cooldown,omitempty"`

	// Conditions that make this entry in the list valid
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum distance this mob can be away to be a valid choice
	MaxDist float64 `json:"max_dist,omitempty"`

	// If true, the mob has to be visible to be a valid choice
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// If true, the mob will stop being targeted if it stops meeting any conditions.
	ReevaluateDescription bool `json:"reevaluate_description,omitempty"`

	// Multiplier for the running speed. A value of 1.0 means the speed is unchanged
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the mob to target another mob that hurts their owner.
type Behavior_OwnerHurtByTarget struct {
	// Corresponding Type: BehaviorOwnerHurtByTargetEntityTypes
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of entity types that this entity can target if the potential target is hurt by this mob's owner
type BehaviorOwnerHurtTargetEntityTypes struct {
	// The amount of time in seconds that the mob has to wait before selecting a target of the same type again
	Cooldown float64 `json:"cooldown,omitempty"`

	// Conditions that make this entry in the list valid
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum distance this mob can be away to be a valid choice
	MaxDist float64 `json:"max_dist,omitempty"`

	// If true, the mob has to be visible to be a valid choice
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// If true, the mob will stop being targeted if it stops meeting any conditions.
	ReevaluateDescription bool `json:"reevaluate_description,omitempty"`

	// Multiplier for the running speed. A value of 1.0 means the speed is unchanged
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the mob to target a mob that is hurt by their owner.
type Behavior_OwnerHurtTarget struct {
	// Corresponding Type: BehaviorOwnerHurtTargetEntityTypes
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to enter the panic state, which makes it run around and away from the damage source that made it enter this state.
type Behavior_Panic struct {
	// The list of Entity Damage Sources that will cause this mob to panic
	DamageSources []interface{} `json:"damage_sources,omitempty"`

	// If true, this mob will not stop panicking until it can't move anymore or the goal is removed from it
	Force bool `json:"force,omitempty"`

	// If true, the mob will not panic in response to damage from other mobs. This overrides the damage types in "damage_sources"
	IgnoreMobDamage bool `json:"ignore_mob_damage,omitempty"`

	// If true, the mob will prefer water over land
	PreferWater bool `json:"prefer_water,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to peek out. This is what the shulker uses to look out of its shell.

type Behavior_Peek struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the pet mob to move onto a bed with its owner while sleeping.
type Behavior_PetSleepWithOwner struct {
	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// Height in blocks from the owner the pet can be to sleep with owner.
	SearchHeight int `json:"search_height,omitempty"`

	// The distance in blocks from the owner the pet can be to sleep with owner.
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to pick up items on the ground.
type Behavior_PickupItems struct {
	// If true, the mob can pickup any item
	CanPickupAnyItem bool `json:"can_pickup_any_item,omitempty"`

	// If true, the mob can pickup items to its hand or armor slots
	CanPickupToHandOrEquipment bool `json:"can_pickup_to_hand_or_equipment,omitempty"`

	// List of items this mob will not pick up
	ExcludedItems []interface{} `json:"excluded_items,omitempty"`

	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// Maximum distance this mob will look for items to pick up
	MaxDist float64 `json:"max_dist,omitempty"`

	// If true, depending on the difficulty, there is a random chance that the mob may not be able to pickup items
	PickupBasedOnChance bool `json:"pickup_based_on_chance,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// If true, this mob will chase after the target as long as it's a valid target
	TrackTarget bool `json:"track_target,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to play with other mobs by chasing each other and moving around randomly.
type Behavior_Play struct {
	// Percent chance that the mob will start this goal, from 0 to 1.
	ChanceToStart float64 `json:"chance_to_start,omitempty"`

	// The distance (in blocks) that the mob tries to be in range of the friend it's following.
	FollowDistance int `json:"follow_distance,omitempty"`

	// [Vector3 [a,b,c]] The dimensions of the AABB used to search for a potential friend to play with.
	FriendSearchArea []float64 `json:"friend_search_area,omitempty"`

	// The entity type(s) to consider when searching for a potential friend to play with.
	FriendTypes []interface{} `json:"friend_types,omitempty"`

	// The max amount of seconds that the mob will play for before exiting the Goal.
	MaxPlayDurationSeconds float64 `json:"max_play_duration_seconds,omitempty"`

	// The height (in blocks) that the mob will search within to find a random position position to move to. Must be at least 1.
	RandomPosSearchHeight int `json:"random_pos_search_height,omitempty"`

	// The distance (in blocks) on ground that the mob will search within to find a random position to move to. Must be at least 1.
	RandomPosSearchRange int `json:"random_pos_search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to pretend to be dead to avoid being targeted by attackers.
type Behavior_PlayDead struct {
	// Whether the mob will receive the regeneration effect while playing dead.
	ApplyRegeneration bool `json:"apply_regeneration,omitempty"`

	// The list of Entity Damage Sources that will cause this mob to play dead.
	DamageSources []interface{} `json:"damage_sources,omitempty"`

	// The amount of time the mob will remain playing dead (in seconds).
	Duration float64 `json:"duration,omitempty"`

	// The list of other triggers that are required for the mob to activate play dead
	Filters *f.Filter `json:"filters,omitempty"`

	// The amount of health at which damage will cause the mob to play dead.
	ForceBelowHealth int `json:"force_below_health,omitempty"`

	// The range of damage that may cause the goal to start depending on randomness. Damage taken below the min will never cause the goal to start. Damage taken above the max will always cause the goal to start.
	RandomDamageRange float64 `json:"random_damage_range,omitempty"`

	// The likelihood of this goal starting upon taking damage.
	RandomStartChance float64 `json:"random_start_chance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to be ridden by the player after being tamed.

type Behavior_PlayerRideTamed struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to eat/raid crops out of farms until they are full.
type Behavior_RaidGarden struct {
	// Blocks that the mob is looking for to eat/raid
	Blocks []interface{} `json:"blocks,omitempty"`

	// Time in seconds between each time it eats/raids
	EatDelay int `json:"eat_delay,omitempty"`

	// Amount of time in seconds before this mob wants to eat/raid again after eating its maximum
	FullDelay int `json:"full_delay,omitempty"`

	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// Time in seconds before starting to eat/raid once it arrives at it
	InitialEatDelay int `json:"initial_eat_delay,omitempty"`

	// Maximum number of crops this entity wants to eat/raid. If set to zero or less then it doesn't have a maximum
	MaxToEat int `json:"max_to_eat,omitempty"`

	// Distance in blocks the mob will look for crops to eat
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to damage a target by using a running attack.
type Behavior_RamAttack struct {
	// The modifier to knockback that babies have.
	BabyKnockbackModifier float64 `json:"baby_knockback_modifier,omitempty"`

	// Minimum and maximum cooldown time-range (positive, in seconds) between each attempted ram attack.
	CooldownRange float64 `json:"cooldown_range,omitempty"`

	// The force of the knockback of the ram attack.
	KnockbackForce float64 `json:"knockback_force,omitempty"`

	// The height of the knockback of the ram attack.
	KnockbackHeight float64 `json:"knockback_height,omitempty"`

	// The minimum distance at which the mob can start a ram attack.
	MinRamDistance float64 `json:"min_ram_distance,omitempty"`

	// [Trigger] The event to trigger when attacking
	OnStart interface{} `json:"on_start,omitempty"`

	// The sound to play when an entity is about to perform a ram attack.
	PreRamSound string `json:"pre_ram_sound,omitempty"`

	// The distance at which the mob start to run with ram speed.
	RamDistance float64 `json:"ram_distance,omitempty"`

	// The sound to play when an entity is impacting on a ram attack.
	RamImpactSound string `json:"ram_impact_sound,omitempty"`

	// Sets the entity's speed when charging toward the target.
	RamSpeed float64 `json:"ram_speed,omitempty"`

	// Sets the entity's speed when running toward the target.
	RunSpeed float64 `json:"run_speed,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to randomly break surface of the water.
type Behavior_RandomBreach struct {
	// Time in seconds the mob has to wait before using the goal again
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// A random value to determine when to randomly move somewhere. This has a 1/interval chance to choose this goal
	Interval int `json:"interval,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Distance in blocks on ground that the mob will look for a new spot to move to. Must be at least 1
	XzDist int `json:"xz_dist,omitempty"`

	// Distance in blocks that the mob will look up or down for a new spot to move to. Must be at least 1
	YDist int `json:"y_dist,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows a mob to randomly fly around.
type Behavior_RandomFly struct {
	// If true, the mob can stop flying and land on a tree instead of the ground
	CanLandOnTrees bool `json:"can_land_on_trees,omitempty"`

	// Distance in blocks on ground that the mob will look for a new spot to move to. Must be at least 1
	XzDist int `json:"xz_dist,omitempty"`

	// Distance in blocks that the mob will look up or down for a new spot to move to. Must be at least 1
	YDist int `json:"y_dist,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to hover around randomly, close to the surface
type Behavior_RandomHover struct {
	// The height above the surface which the mob will try to maintain
	HoverHeight float64 `json:"hover_height,omitempty"`

	// A random value to determine when to randomly move somewhere. This has a 1/interval chance to choose this goal
	Interval int `json:"interval,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Distance in blocks on ground that the mob will look for a new spot to move to. Must be at least 1
	XzDist int `json:"xz_dist,omitempty"`

	// Distance in blocks that the mob will look up or down for a new spot to move to. Must be at least 1
	YDist int `json:"y_dist,omitempty"`

	// Height in blocks to add to the selected target position
	YOffset float64 `json:"y_offset,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to randomly look around.
type Behavior_RandomLookAround struct {
	// The range of time in seconds the mob will stay looking in a random direction before looking elsewhere
	LookTime float64 `json:"look_time,omitempty"`

	// The rightmost angle a mob can look at on the horizontal plane with respect to its initial facing direction.
	MaxAngleOfViewHorizontal int `json:"max_angle_of_view_horizontal,omitempty"`

	// The leftmost angle a mob can look at on the horizontal plane with respect to its initial facing direction.
	MinAngleOfViewHorizontal int `json:"min_angle_of_view_horizontal,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to randomly sit and look around for a duration. Note: Must have a sitting animation set up to use this.
type Behavior_RandomLookAroundAndSit struct {
	// If the goal should continue to be used as long as the mob is leashed.
	ContinueIfLeashed bool `json:"continue_if_leashed,omitempty"`

	// The rightmost angle a mob can look at on the horizontal plane with respect to its initial facing direction.
	MaxAngleOfViewHorizontal float64 `json:"max_angle_of_view_horizontal,omitempty"`

	// The max amount of unique looks a mob will have while looking around.
	MaxLookCount int `json:"max_look_count,omitempty"`

	// The max amount of time (in ticks) a mob will stay looking at a direction while looking around.
	MaxLookTime int `json:"max_look_time,omitempty"`

	// The leftmost angle a mob can look at on the horizontal plane with respect to its initial facing direction.
	MinAngleOfViewHorizontal float64 `json:"min_angle_of_view_horizontal,omitempty"`

	// The min amount of unique looks a mob will have while looking around.
	MinLookCount int `json:"min_look_count,omitempty"`

	// The min amount of time (in ticks) a mob will stay looking at a direction while looking around.
	MinLookTime int `json:"min_look_time,omitempty"`

	// The probability of randomly looking around/sitting.
	Probability float64 `json:"probability,omitempty"`

	// The cooldown in seconds before the goal can be used again.
	RandomLookAroundCooldown int `json:"random_look_around_cooldown,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [EXPERIMENTAL] Allows this entity to locate a random target block that it can path find to. Once found, the entity will move towards it and dig up an item. [Valid target block type: Dirt, Grass, Podzol, DirtWithRoots, MossBlock, Mud, MuddyMangroveRoots].
type Behavior_RandomSearchAndDig struct {
	// Goal cooldown range in seconds.
	CooldownRange float64 `json:"cooldown_range,omitempty"`

	// Digging duration in seconds.
	DiggingDurationRange float64 `json:"digging_duration_range,omitempty"`

	// Amount of retries to find a valid target position within search range.
	FindValidPositionRetries float64 `json:"find_valid_position_retries,omitempty"`

	// Distance in blocks within the entity to considers it has reached it's target position.
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// Weighted list of item(s) to spawn on successful dig
	Items []interface{} `json:"items,omitempty"`

	// [Trigger] Event to run when the goal ends searching has begins digging.
	OnDiggingStart interface{} `json:"on_digging_start,omitempty"`

	// [Trigger] Event to run when the goal failed while in digging state.
	OnFailDuringDigging interface{} `json:"on_fail_during_digging,omitempty"`

	// [Trigger] Event to run when the goal failed while in searching state.
	OnFailDuringSearching interface{} `json:"on_fail_during_searching,omitempty"`

	// [Trigger] Event to run when the goal find a item.
	OnItemFound interface{} `json:"on_item_found,omitempty"`

	// [Trigger] Event to run when the goal starts searching.
	OnSearchingStart interface{} `json:"on_searching_start,omitempty"`

	// [Trigger] Event to run when searching and digging has ended.
	OnSuccess interface{} `json:"on_success,omitempty"`

	// Width and length of the volume around the entity used to find a valid target position
	SearchRangeXz float64 `json:"search_range_xz,omitempty"`

	// Height of the volume around the entity used to find a valid target position
	SearchRangeY float64 `json:"search_range_y,omitempty"`

	// Digging duration before spawning item in seconds.
	SpawnItemAfterSeconds float64 `json:"spawn_item_after_seconds,omitempty"`

	// Distance to offset the item's spawn location in the direction the mob is facing.
	SpawnItemPosOffset float64 `json:"spawn_item_pos_offset,omitempty"`

	// Searching movement speed multiplier.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Dig target position offset from the feet position of the mob in their facing direction.
	TargetDigPositionOffset float64 `json:"target_dig_position_offset,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to randomly sit for a duration.
type Behavior_RandomSitting struct {
	// Time in seconds the mob has to wait before using the goal again
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// The minimum amount of time in seconds before the mob can stand back up
	MinSitTime float64 `json:"min_sit_time,omitempty"`

	// This is the chance that the mob will start this goal, from 0 to 1
	StartChance float64 `json:"start_chance,omitempty"`

	// This is the chance that the mob will stop this goal, from 0 to 1
	StopChance float64 `json:"stop_chance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows a mob to randomly stroll around.
type Behavior_RandomStroll struct {
	// A random value to determine when to randomly move somewhere. This has a 1/interval chance to choose this goal
	Interval int `json:"interval,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Distance in blocks on ground that the mob will look for a new spot to move to. Must be at least 1
	XzDist int `json:"xz_dist,omitempty"`

	// Distance in blocks that the mob will look up or down for a new spot to move to. Must be at least 1
	YDist int `json:"y_dist,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows an entity to randomly move through water
type Behavior_RandomSwim struct {
	// If true, the mob will avoid surface water blocks by swimming below them
	AvoidSurface bool `json:"avoid_surface,omitempty"`

	// A random value to determine when to randomly move somewhere. This has a 1/interval chance to choose this goal
	Interval int `json:"interval,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Distance in blocks on ground that the mob will look for a new spot to move to. Must be at least 1
	XzDist int `json:"xz_dist,omitempty"`

	// Distance in blocks that the mob will look up or down for a new spot to move to. Must be at least 1
	YDist int `json:"y_dist,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows an entity to attack by using ranged shots. "charge_shoot_trigger" must be greater than 0 to enable charged up burst-shot attacks. Requires minecraft:shooter to define projectile behaviour.
type Behavior_RangedAttack struct {
	// Alternative to "attack_interval_min" & "attack_interval_max". Consistent reload-time (in seconds), when not using a charged shot. Does not scale with target-distance.
	AttackInterval float64 `json:"attack_interval,omitempty"`

	// Maximum bound for reload-time range (in seconds), when not using a charged shot. Reload-time range scales with target-distance.
	AttackIntervalMax float64 `json:"attack_interval_max,omitempty"`

	// Minimum bound for reload-time range (in seconds), when not using a charged shot. Reload-time range scales with target-distance.
	AttackIntervalMin float64 `json:"attack_interval_min,omitempty"`

	// Minimum distance to target before this entity will attempt to shoot.
	AttackRadius float64 `json:"attack_radius,omitempty"`

	// Minimum distance the target can be for this mob to fire. If the target is closer, this mob will move first before firing
	AttackRadiusMin float64 `json:"attack_radius_min,omitempty"`

	// Time (in seconds) between each individual shot when firing a burst of shots from a charged up attack.
	BurstInterval float64 `json:"burst_interval,omitempty"`

	// Number of shots fired every time the attacking entity uses a charged up attack.
	BurstShots int `json:"burst_shots,omitempty"`

	// Time (in seconds, then add "charge_shoot_trigger"), before a charged up attack is done charging. Charge-time decays while target is not in sight.
	ChargeChargedTrigger float64 `json:"charge_charged_trigger,omitempty"`

	// Amount of time (in seconds, then doubled) a charged shot must be charging before reloading burst shots. Charge-time decays while target is not in sight.
	ChargeShootTrigger float64 `json:"charge_shoot_trigger,omitempty"`

	// Field of view (in degrees) when using sensing to detect a target for attack.
	RangedFov float64 `json:"ranged_fov,omitempty"`

	// Allows the actor to be set to persist upon targeting a player
	SetPersistent bool `json:"set_persistent,omitempty"`

	// During attack behavior, this multiplier modifies the entity's speed when moving toward the target.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// If a swing animation (using variable.attack_time) exists, this causes the actor to swing their arm(s) upon firing the ranged attack.
	Swing bool `json:"swing,omitempty"`

	// Minimum amount of time (in seconds) the attacking entity needs to see the target before moving toward it.
	TargetInSightTime float64 `json:"target_in_sight_time,omitempty"`

	// Maximum rotation (in degrees), on the X-axis, this entity can rotate while trying to look at the target.
	XMaxRotation float64 `json:"x_max_rotation,omitempty"`

	// Maximum rotation (in degrees), on the Y-axis, this entity can rotate its head while trying to look at the target.
	YMaxHeadRotation float64 `json:"y_max_head_rotation,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the villager to stop so another villager can breed with it. Can only be used by a Villager.

type Behavior_ReceiveLove struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to stay indoors during night time.

type Behavior_RestrictOpenDoor struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to automatically start avoiding the sun when its a clear day out.

type Behavior_RestrictSun struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to stay at a certain level when in liquid.
type Behavior_RiseToLiquidLevel struct {
	// Target distance down from the liquid surface. i.e. Positive values move the target Y down.
	LiquidYOffset float64 `json:"liquid_y_offset,omitempty"`

	// Movement up in Y per tick when below the liquid surface.
	RiseDelta float64 `json:"rise_delta,omitempty"`

	// Movement down in Y per tick when above the liquid surface.
	SinkDelta float64 `json:"sink_delta,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [EXPERIMENTAL] Allows this entity to rise for a specified duration
type Behavior_Rising struct {
	// Goal cooldown range in seconds
	CooldownRange float64 `json:"cooldown_range,omitempty"`

	// Goal duration range in seconds
	DurationRange float64 `json:"duration_range,omitempty"`

	// [Trigger] Event(s) to run when the goal end.
	OnEnd interface{} `json:"on_end,omitempty"`

	// [Trigger] Event(s) to run when the goal starts.
	OnStart interface{} `json:"on_start,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to roar at another entity based on data in minecraft:anger_level. Once the anger threshold specified in minecraft:anger_level has been reached, this entity will roar for the specified amount of time, look at the other entity, apply anger boost towards it, and finally target it.
type Behavior_Roar struct {
	// The amount of time to roar for.
	Duration float64 `json:"duration,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// This allows the mob to roll forward.
type Behavior_Roll struct {
	// The probability that the mob will use the goal.
	Probability float64 `json:"probability,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to run around aimlessly.
type Behavior_RunAroundLikeCrazy struct {
	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the a mob to become scared when the weather outside is thundering
type Behavior_Scared struct {
	// The interval in which a sound will play when active in a 1/delay chance to kick off
	SoundInterval int `json:"sound_interval,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [EXPERIMENTAL] Allows this entity to scent for a specified duration
type Behavior_Scenting struct {
	// Goal cooldown range in seconds
	CooldownRange float64 `json:"cooldown_range,omitempty"`

	// Goal duration range in seconds
	DurationRange float64 `json:"duration_range,omitempty"`

	// [Trigger] Event(s) to run when the goal end.
	OnEnd interface{} `json:"on_end,omitempty"`

	// [Trigger] Event(s) to run when the goal starts.
	OnStart interface{} `json:"on_start,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of events to send
type BehaviorSendEventSequence struct {
	// Amount of time in seconds before starting this step
	BaseDelay float64 `json:"base_delay,omitempty"`

	// The event to send to the entity
	Event string `json:"event,omitempty"`

	// The sound event to play when this step happens
	SoundEvent string `json:"sound_event,omitempty"`
}

// Allows the mob to send an event to another mob.
type Behavior_SendEvent struct {
	// Time in seconds for the entire event sending process
	CastDuration float64 `json:"cast_duration,omitempty"`

	// If true, the mob will face the entity it sends an event to
	LookAtTarget bool `json:"look_at_target,omitempty"`

	// List of events to send
	Sequence []BehaviorSendEventSequence `json:"sequence,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of entities this mob will share items with
type BehaviorShareItemsEntityTypes struct {
	// The amount of time in seconds that the mob has to wait before selecting a target of the same type again
	Cooldown float64 `json:"cooldown,omitempty"`

	// Conditions that make this entry in the list valid
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum distance this mob can be away to be a valid choice
	MaxDist float64 `json:"max_dist,omitempty"`

	// If true, the mob has to be visible to be a valid choice
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// If true, the mob will stop being targeted if it stops meeting any conditions.
	ReevaluateDescription bool `json:"reevaluate_description,omitempty"`

	// Multiplier for the running speed. A value of 1.0 means the speed is unchanged
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the mob to give items it has to others.
type Behavior_ShareItems struct {
	// Corresponding Type: BehaviorShareItemsEntityTypes
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// Maximum distance in blocks this mob will look for entities to share items with
	MaxDist float64 `json:"max_dist,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to go into stone blocks like Silverfish do. Currently it can only be used by Silverfish.

type Behavior_SilverfishMergeWithStone struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to alert mobs in nearby blocks to come out. Currently it can only be used by Silverfish.

type Behavior_SilverfishWakeUpFriends struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows Equine mobs to be Horse Traps and be triggered like them, spawning a lightning bolt and a bunch of horses when a player is nearby. Can only be used by Horses, Mules, Donkeys and Skeleton Horses.
type Behavior_SkeletonHorseTrap struct {
	// Amount of time in seconds the trap exists. After this amount of time is elapsed, the trap is removed from the world if it hasn't been activated
	Duration float64 `json:"duration,omitempty"`

	// Distance in blocks that the player has to be within to trigger the horse trap
	WithinRadius float64 `json:"within_radius,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows mobs that own a bed to in a village to move to and sleep in it.
type Behavior_Sleep struct {
	// If true, the mob will be able to use the sleep goal if riding something
	CanSleepWhileRiding bool `json:"can_sleep_while_riding,omitempty"`

	// Time in seconds the mob has to wait before using the goal again
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// The height of the mob's collider while sleeping
	SleepColliderHeight float64 `json:"sleep_collider_height,omitempty"`

	// The width of the mob's collider while sleeping
	SleepColliderWidth float64 `json:"sleep_collider_width,omitempty"`

	// The y offset of the mob's collider while sleeping
	SleepYOffset float64 `json:"sleep_y_offset,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The cooldown time in seconds before the goal can be reused after a internal failure or timeout condition
	TimeoutCooldown float64 `json:"timeout_cooldown,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Causes the entity to grow tired every once in a while, while attacking.
type Behavior_SlimeAttack struct {
	// Allows the actor to be set to persist upon targeting a player
	SetPersistent bool `json:"set_persistent,omitempty"`

	// During attack behavior, this multiplier modifies the entity's speed when moving toward the target.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Maximum rotation (in degrees), on the X-axis, this entity can rotate while trying to look at the target.
	XMaxRotation float64 `json:"x_max_rotation,omitempty"`

	// Maximum rotation (in degrees), on the Y-axis, this entity can rotate while trying to look at the target.
	YMaxRotation float64 `json:"y_max_rotation,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allow slimes to float in water / lava. Can only be used by Slime and Magma Cubes.
type Behavior_SlimeFloat struct {
	// Percent chance a slime or magma cube has to jump while in water / lava.
	JumpChancePercentage float64 `json:"jump_chance_percentage,omitempty"`

	// Determines the multiplier the entity's speed is modified by when moving through water / lava.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to continuously jump around like a slime.
type Behavior_SlimeKeepOnJumping struct {
	// Determines the multiplier this entity's speed is modified by when jumping around.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to move in random directions like a slime.
type Behavior_SlimeRandomDirection struct {
	// Additional time (in whole seconds), chosen randomly in the range of [0, "add_random_time_range"], to add to "min_change_direction_time".
	AddRandomTimeRange int `json:"add_random_time_range,omitempty"`

	// Constant minimum time (in seconds) to wait before choosing a new direction.
	MinChangeDirectionTime float64 `json:"min_change_direction_time,omitempty"`

	// Maximum rotation angle range (in degrees) when randomly choosing a new direction.
	TurnRange int `json:"turn_range,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to take a load off and snack on food that it found nearby.
type Behavior_Snacking struct {
	// Items that we are interested in snacking on
	Items []interface{} `json:"items,omitempty"`

	// The cooldown time in seconds before the mob is able to snack again
	SnackingCooldown float64 `json:"snacking_cooldown,omitempty"`

	// The minimum time in seconds before the mob is able to snack again
	SnackingCooldownMin float64 `json:"snacking_cooldown_min,omitempty"`

	// This is the chance that the mob will stop snacking, from 0 to 1
	SnackingStopChance float64 `json:"snacking_stop_chance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of entity types this mob will startle (cause to jump) when it sneezes.
type BehaviorSneezeEntityTypes struct {
	// The amount of time in seconds that the mob has to wait before selecting a target of the same type again
	Cooldown float64 `json:"cooldown,omitempty"`

	// Conditions that make this entry in the list valid
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum distance this mob can be away to be a valid choice
	MaxDist float64 `json:"max_dist,omitempty"`

	// If true, the mob has to be visible to be a valid choice
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// If true, the mob will stop being targeted if it stops meeting any conditions.
	ReevaluateDescription bool `json:"reevaluate_description,omitempty"`

	// Multiplier for the running speed. A value of 1.0 means the speed is unchanged
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the mob to stop and sneeze possibly startling nearby mobs and dropping an item.
type Behavior_Sneeze struct {
	// Time in seconds the mob has to wait before using the goal again
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// The probability that the mob will drop an item when it sneezes.
	DropItemChance float64 `json:"drop_item_chance,omitempty"`

	// Corresponding Type: BehaviorSneezeEntityTypes
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// Loot table to select dropped items from.
	LootTable string `json:"loot_table,omitempty"`

	// Sound to play when the sneeze is about to happen.
	PrepareSound string `json:"prepare_sound,omitempty"`

	// The time in seconds that the mob takes to prepare to sneeze (while the prepare_sound is playing).
	PrepareTime float64 `json:"prepare_time,omitempty"`

	// The probability of sneezing. A value of 1.00 is 100%
	Probability float64 `json:"probability,omitempty"`

	// Sound to play when the sneeze occurs.
	Sound string `json:"sound,omitempty"`

	// Distance in blocks that mobs will be startled.
	WithinRadius float64 `json:"within_radius,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this entity to detect the nearest player within "sniffing_radius" and update its "minecraft:suspect_tracking" component state
type Behavior_Sniff struct {
	// Cooldown range between sniffs in seconds
	CooldownRange float64 `json:"cooldown_range,omitempty"`

	// Sniffing duration in seconds
	Duration float64 `json:"duration,omitempty"`

	// Mob detection radius
	SniffingRadius float64 `json:"sniffing_radius,omitempty"`

	// Mob suspicion horizontal radius. When a player is within this radius horizontally, the anger level towards that player is increased
	SuspicionRadiusHorizontal float64 `json:"suspicion_radius_horizontal,omitempty"`

	// Mob suspicion vertical radius. When a player is within this radius vertically, the anger level towards that player is increased
	SuspicionRadiusVertical float64 `json:"suspicion_radius_vertical,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

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

// Allows the squid to dive down in water. Can only be used by the Squid.

type Behavior_SquidDive struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the squid to swim away. Can only be used by the Squid.

type Behavior_SquidFlee struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the squid to swim in place idly. Can only be used by the Squid.

type Behavior_SquidIdle struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the squid to move away from ground blocks and back to water. Can only be used by the Squid.

type Behavior_SquidMoveAwayFromGround struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the squid to stick to the ground when outside water. Can only be used by the Squid.

type Behavior_SquidOutOfWater struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows a mob to stalk a target, then once within range pounce onto a target, on success the target will be attacked dealing damage defined by the attack component. On failure, the mob will risk getting stuck
type Behavior_StalkAndPounceOnTarget struct {
	// The amount of time the mob will be interested before pouncing. This happens when the mob is within range of pouncing
	InterestTime float64 `json:"interest_time,omitempty"`

	// The distance in blocks the mob jumps in the direction of its target
	LeapDistance float64 `json:"leap_distance,omitempty"`

	// The height in blocks the mob jumps when leaping at its target
	LeapHeight float64 `json:"leap_height,omitempty"`

	// The maximum distance away a target can be before the mob gives up on stalking
	MaxStalkDist float64 `json:"max_stalk_dist,omitempty"`

	// The maximum distance away from the target in blocks to begin pouncing at the target
	PounceMaxDist float64 `json:"pounce_max_dist,omitempty"`

	// Allows the actor to be set to persist upon targeting a player
	SetPersistent bool `json:"set_persistent,omitempty"`

	// The movement speed in which you stalk your target
	StalkSpeed float64 `json:"stalk_speed,omitempty"`

	// The max distance away from the target when landing from the pounce that will still result in damaging the target
	StrikeDist float64 `json:"strike_dist,omitempty"`

	// The amount of time the mob will be stuck if they fail and land on a block they can be stuck on
	StuckTime float64 `json:"stuck_time,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// The entity will attempt to toss the items from its inventory to a nearby recently played noteblock.
type Behavior_StayNearNoteblock struct {
	// Sets the time an entity should stay near a noteblock after hearing it.
	ListenTime int `json:"listen_time,omitempty"`

	// Sets the entity's speed when moving toward the block.
	Speed float64 `json:"speed,omitempty"`

	// Sets the distance the entity needs to be away from the block to attempt to start the goal.
	StartDistance float64 `json:"start_distance,omitempty"`

	// Sets the distance from the block the entity will attempt to reach.
	StopDistance float64 `json:"stop_distance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to stay put while it is in a sitting state instead of doing something else.

type Behavior_StayWhileSitting struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows an entity to attack using stomp AoE damage behavior.
type Behavior_StompAttack struct {
	// Allows the entity to use this attack behavior, only once EVER.
	AttackOnce bool `json:"attack_once,omitempty"`

	// Defines the entity types this entity will attack.
	AttackTypes string `json:"attack_types,omitempty"`

	// Cooldown time (in seconds) between attacks.
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// Time (in seconds) to add to attack path recalculation when the target is beyond the "path_inner_boundary".
	InnerBoundaryTimeIncrease float64 `json:"inner_boundary_time_increase,omitempty"`

	// Unused. No effect on "minecraft:behavior.melee_attack".
	MaxDist float64 `json:"max_dist,omitempty"`

	// Maximum base time (in seconds) to recalculate new attack path to target (before increases applied).
	MaxPathTime float64 `json:"max_path_time,omitempty"`

	// Field of view (in degrees) when using the sensing component to detect an attack target.
	MeleeFov float64 `json:"melee_fov,omitempty"`

	// Minimum base time (in seconds) to recalculate new attack path to target (before increases applied).
	MinPathTime float64 `json:"min_path_time,omitempty"`

	// Multiplied with the final AoE damage range to determine a no damage range. The stomp attack will go on cooldown if target is in this no damage range.
	NoDamageRangeMultiplier float64 `json:"no_damage_range_multiplier,omitempty"`

	// [Trigger] Defines the event to trigger when this entity successfully attacks.
	OnAttack interface{} `json:"on_attack,omitempty"`

	// Time (in seconds) to add to attack path recalculation when the target is beyond the "path_outer_boundary".
	OuterBoundaryTimeIncrease float64 `json:"outer_boundary_time_increase,omitempty"`

	// Time (in seconds) to add to attack path recalculation when this entity cannot move along the current path.
	PathFailTimeIncrease float64 `json:"path_fail_time_increase,omitempty"`

	// Distance at which to increase attack path recalculation by "inner_boundary_tick_increase".
	PathInnerBoundary float64 `json:"path_inner_boundary,omitempty"`

	// Distance at which to increase attack path recalculation by "outer_boundary_tick_increase".
	PathOuterBoundary float64 `json:"path_outer_boundary,omitempty"`

	// This entity will have a 1 in N chance to stop it's current attack, where N = "random_stop_interval".
	RandomStopInterval int `json:"random_stop_interval,omitempty"`

	// Used with the base size of the entity to determine minimum target-distance before trying to deal attack damage.
	ReachMultiplier float64 `json:"reach_multiplier,omitempty"`

	// Toggles (on/off) the need to have a full path from the entity to the target when using this melee attack behavior.
	RequireCompletePath bool `json:"require_complete_path,omitempty"`

	// Allows the actor to be set to persist upon targeting a player
	SetPersistent bool `json:"set_persistent,omitempty"`

	// This multiplier modifies the attacking entity's speed when moving toward the target.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Multiplied with the base size of the entity to determine stomp AoE damage range.
	StompRangeMultiplier float64 `json:"stomp_range_multiplier,omitempty"`

	// Unused. No effect on "minecraft:behavior.melee_attack".
	TargetDist float64 `json:"target_dist,omitempty"`

	// Allows the entity to track the attack target, even if the entity has no sensing.
	TrackTarget bool `json:"track_target,omitempty"`

	// Maximum rotation (in degrees), on the X-axis, this entity can rotate while trying to look at the target.
	XMaxRotation float64 `json:"x_max_rotation,omitempty"`

	// Maximum rotation (in degrees), on the Y-axis, this entity can rotate its head while trying to look at the target.
	YMaxHeadRotation float64 `json:"y_max_head_rotation,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows this mob to stomp turtle eggs
type Behavior_StompTurtleEgg struct {
	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// A random value to determine when to randomly move somewhere. This has a 1/interval chance to choose this goal
	Interval int `json:"interval,omitempty"`

	// Height in blocks the mob will look for turtle eggs to move towards
	SearchHeight int `json:"search_height,omitempty"`

	// The distance in blocks it will look for turtle eggs to move towards
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to move into a random location within a village within the search range.
type Behavior_StrollTowardsVillage struct {
	// Time in seconds the mob has to wait before using the goal again
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// Distance in blocks within the mob considers it has reached the goal. This is the "wiggle room" to stop the AI from bouncing back and forth trying to reach a specific spot
	GoalRadius float64 `json:"goal_radius,omitempty"`

	// The distance in blocks to search for points inside villages. If <= 0, find the closest village regardless of distance.
	SearchRange int `json:"search_range,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// This is the chance that the mob will start this goal, from 0 to 1
	StartChance float64 `json:"start_chance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of steps for the spell. Each step has the following parameters:
type BehaviorSummonEntitySequence struct {
	// Amount of time in seconds to wait before this step starts
	BaseDelay float64 `json:"base_delay,omitempty"`

	// Amount of time in seconds before each entity is summoned in this step
	DelayPerSummon float64 `json:"delay_per_summon,omitempty"`

	// Amount of time in seconds that the spawned entity will be alive for. A value of -1.0 means it will remain alive for as long as it can
	EntityLifespan float64 `json:"entity_lifespan,omitempty"`

	// The entity type of the entities we will spawn in this step
	EntityType string `json:"entity_type,omitempty"`

	// Number of entities that will be spawned in this step
	NumEntitiesSpawned int `json:"num_entities_spawned,omitempty"`

	// The base shape of this step. Valid values are circle and line
	Shape string `json:"shape,omitempty"`

	// The base size of the entity
	Size float64 `json:"size,omitempty"`

	// The sound event to play for this step
	SoundEvent string `json:"sound_event,omitempty"`

	// Maximum number of summoned entities at any given time
	SummonCap int `json:"summon_cap,omitempty"`

	//     SummonCapRadius float64 `json:"summon_cap_radius,omitempty"`

	// The target of the spell. This is where the spell will start (line will start here, circle will be centered here)
	Target string `json:"target,omitempty"`
}

// [Not a component] List of spells for the mob to use to summon entities. Each spell has the following parameters:
type BehaviorSummonEntitySummonChoices struct {
	// Time in seconds the spell casting will take
	CastDuration float64 `json:"cast_duration,omitempty"`

	// Time in seconds the mob has to wait before using the spell again
	CooldownTime float64 `json:"cooldown_time,omitempty"`

	// If true, the mob will do the casting animations and render spell particles
	DoCasting bool `json:"do_casting,omitempty"`

	//     Filters *f.Filter `json:"filters,omitempty"`

	// Upper bound of the activation distance in blocks for this spell, must not be negative.
	MaxActivationRange float64 `json:"max_activation_range,omitempty"`

	// Lower bound of the activation distance in blocks for this spell, must not be negative.
	MinActivationRange float64 `json:"min_activation_range,omitempty"`

	// The color of the particles for this spell
	ParticleColor int `json:"particle_color,omitempty"`

	// List of steps for the spell. Each step has the following parameters:
	Sequence []BehaviorSummonEntitySequence `json:"sequence,omitempty"`

	// The sound event to play when using this spell
	StartSoundEvent string `json:"start_sound_event,omitempty"`

	// The weight of this spell. Controls how likely the mob is to choose this spell when casting one
	Weight float64 `json:"weight,omitempty"`
}

// Allows the mob to attack the player by summoning other entities.
type Behavior_SummonEntity struct {
	// List of spells for the mob to use to summon entities. Each spell has the following parameters:
	SummonChoices []BehaviorSummonEntitySummonChoices `json:"summon_choices,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the creeper to swell up when a player is nearby. It can only be used by Creepers.
type Behavior_Swell struct {
	// This mob starts swelling when a target is at least this many blocks away
	StartDistance float64 `json:"start_distance,omitempty"`

	// This mob stops swelling when a target has moved away at least this many blocks
	StopDistance float64 `json:"stop_distance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity go idle, if swimming. Entity must be in water.
type Behavior_SwimIdle struct {
	// Amount of time (in seconds) to stay idle.
	IdleTime float64 `json:"idle_time,omitempty"`

	// Percent chance this entity will go idle, 1.0 = 100%.
	SuccessRate float64 `json:"success_rate,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity to wander around while swimming, when not path-finding.
type Behavior_SwimWander struct {
	// Percent chance to start wandering, when not path-finding. 1 = 100%
	Interval float64 `json:"interval,omitempty"`

	// Distance to look ahead for obstacle avoidance, while wandering.
	LookAhead float64 `json:"look_ahead,omitempty"`

	// This multiplier modifies the entity's speed when wandering.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Amount of time (in seconds) to wander after wandering behavior was successfully started.
	WanderTime float64 `json:"wander_time,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the entity follow another entity. Both entities must be swimming [ie, in water].
type Behavior_SwimWithEntity struct {
	// The multiplier this entity's speed is modified by when matching another entity's direction.
	CatchUpMultiplier float64 `json:"catch_up_multiplier,omitempty"`

	// Distance, from the entity being followed, at which this entity will speed up to reach that entity.
	CatchUpThreshold float64 `json:"catch_up_threshold,omitempty"`

	// Percent chance to stop following the current entity, if they're riding another entity or they're not swimming. 1.0 = 100%
	ChanceToStop float64 `json:"chance_to_stop,omitempty"`

	// Filters which types of entities are valid to follow.
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// Distance, from the entity being followed, at which this entity will try to match that entity's direction
	MatchDirectionThreshold float64 `json:"match_direction_threshold,omitempty"`

	// Radius around this entity to search for another entity to follow.
	SearchRange float64 `json:"search_range,omitempty"`

	// The multiplier this entity's speed is modified by when trying to catch up to the entity being followed.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Time (in seconds) between checks to determine if this entity should catch up to the entity being followed or match the direction of the entity being followed.
	StateCheckInterval float64 `json:"state_check_interval,omitempty"`

	// Distance, from the entity being followed, at which this entity will stop following that entity.
	StopDistance float64 `json:"stop_distance,omitempty"`

	// Percent chance to start following another entity, if not already doing so. 1.0 = 100%
	SuccessRate float64 `json:"success_rate,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows an entity to attack using swoop attack behavior; Ideal for use with flying mobs. The behavior ends if the entity has a horizontal collision or gets hit.
type Behavior_SwoopAttack struct {
	// Added to the base size of the entity, to determine the target's maximum allowable distance, when trying to deal attack damage.
	DamageReach float64 `json:"damage_reach,omitempty"`

	// Minimum and maximum cooldown time-range (in seconds) between each attempted swoop attack.
	DelayRange float64 `json:"delay_range,omitempty"`

	// During swoop attack behavior, this determines the multiplier the entity's speed is modified by when moving toward the target.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to accept flowers from another mob with the minecraft:offer_flower behavior.
type Behavior_TakeFlower struct {
	// Conditions that need to be met for the behavior to start.
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum rotation (in degrees), on the Y-axis, this entity can rotate its head while trying to look at the target.
	MaxHeadRotationY float64 `json:"max_head_rotation_y,omitempty"`

	// Maximum rotation (in degrees), on the X-axis, this entity can rotate while trying to look at the target.
	MaxRotationX float64 `json:"max_rotation_x,omitempty"`

	// The maximum amount of time (in seconds) for the mob to randomly wait for before taking the flower.
	MaxWaitTime float64 `json:"max_wait_time,omitempty"`

	// Minimum distance (in blocks) for the entity to be considered having reached its target.
	MinDistanceToTarget float64 `json:"min_distance_to_target,omitempty"`

	// The minimum amount of time (in seconds) for the mob to randomly wait for before taking the flower.
	MinWaitTime float64 `json:"min_wait_time,omitempty"`

	// [Vector3 [a,b,c]] The dimensions of the AABB used to search for a potential mob to take a flower from.
	SearchArea []float64 `json:"search_area,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to be tempted by food they like.
type Behavior_Tempt struct {
	// If true, the mob can stop being tempted if the player moves too fast while close to this mob
	CanGetScared bool `json:"can_get_scared,omitempty"`

	// If true, vertical distance to the player will be considered when tempting.
	CanTemptVertically bool `json:"can_tempt_vertically,omitempty"`

	// If true, the mob can be tempted even if it has a passenger (i.e. if being ridden).
	CanTemptWhileRidden bool `json:"can_tempt_while_ridden,omitempty"`

	// List of items this mob is tempted by
	Items []interface{} `json:"items,omitempty"`

	// Range of random ticks to wait between tempt sounds.
	SoundInterval float64 `json:"sound_interval,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// Sound to play while the mob is being tempted.
	TemptSound string `json:"tempt_sound,omitempty"`

	// Distance in blocks this mob can get tempted by a player holding an item they like
	WithinRadius float64 `json:"within_radius,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to look at a player that is holding a tradable item.
type Behavior_TradeInterest struct {
	// The max time in seconds that the trader will hold an item before attempting to switch for a different item that takes the same trade
	CarriedItemSwitchTime float64 `json:"carried_item_switch_time,omitempty"`

	// The time in seconds before the trader can use this goal again
	Cooldown float64 `json:"cooldown,omitempty"`

	// The max time in seconds that the trader will be interested with showing its trade items
	InterestTime float64 `json:"interest_time,omitempty"`

	// The max time in seconds that the trader will wait when you no longer have items to trade
	RemoveItemTime float64 `json:"remove_item_time,omitempty"`

	// Distance in blocks this mob can be interested by a player holding an item they like
	WithinRadius float64 `json:"within_radius,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the player to trade with this mob. When the goal starts, it will stop the mob's navigation.
type Behavior_TradeWithPlayer struct {
	// Conditions that need to be met for the behavior to start.
	Filters *f.Filter `json:"filters,omitempty"`

	// The max distance that the mob can be from the player before exiting the goal.
	MaxDistanceFromPlayer float64 `json:"max_distance_from_player,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of entities this mob can copy the owner from
type BehaviorVexCopyOwnerTargetEntityTypes struct {
	// The amount of time in seconds that the mob has to wait before selecting a target of the same type again
	Cooldown float64 `json:"cooldown,omitempty"`

	// Conditions that make this entry in the list valid
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum distance this mob can be away to be a valid choice
	MaxDist float64 `json:"max_dist,omitempty"`

	// If true, the mob has to be visible to be a valid choice
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// If true, the mob will stop being targeted if it stops meeting any conditions.
	ReevaluateDescription bool `json:"reevaluate_description,omitempty"`

	// Multiplier for the running speed. A value of 1.0 means the speed is unchanged
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the mob to target the same entity its owner is targeting.
type Behavior_VexCopyOwnerTarget struct {
	// Corresponding Type: BehaviorVexCopyOwnerTargetEntityTypes
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the mob to move around randomly like the Vex.

type Behavior_VexRandomMove struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the wither to launch random attacks. Can only be used by the Wither Boss.

type Behavior_WitherRandomAttackPosGoal struct {
	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// [Not a component] List of entity types the wither takes into account to find who dealt the most damage to it
type BehaviorWitherTargetHighestDamageEntityTypes struct {
	// The amount of time in seconds that the mob has to wait before selecting a target of the same type again
	Cooldown float64 `json:"cooldown,omitempty"`

	// Conditions that make this entry in the list valid
	Filters *f.Filter `json:"filters,omitempty"`

	// Maximum distance this mob can be away to be a valid choice
	MaxDist float64 `json:"max_dist,omitempty"`

	// If true, the mob has to be visible to be a valid choice
	MustSee bool `json:"must_see,omitempty"`

	// Determines the amount of time in seconds that this mob will look for a target before forgetting about it and looking for a new one when the target isn't visible any more
	MustSeeForgetDuration float64 `json:"must_see_forget_duration,omitempty"`

	// If true, the mob will stop being targeted if it stops meeting any conditions.
	ReevaluateDescription bool `json:"reevaluate_description,omitempty"`

	// Multiplier for the running speed. A value of 1.0 means the speed is unchanged
	SprintSpeedMultiplier float64 `json:"sprint_speed_multiplier,omitempty"`

	// Multiplier for the walking speed. A value of 1.0 means the speed is unchanged
	WalkSpeedMultiplier float64 `json:"walk_speed_multiplier,omitempty"`
}

// Allows the wither to focus its attacks on whichever mob has dealt the most damage to it.
type Behavior_WitherTargetHighestDamage struct {
	// Corresponding Type: BehaviorWitherTargetHighestDamageEntityTypes
	EntityTypes interface{} `json:"entity_types,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the NPC to use the POI
type Behavior_Work struct {
	// The amount of ticks the NPC will stay in their the work location
	ActiveTime int `json:"active_time,omitempty"`

	// If true, this entity can work when their jobsite POI is being rained on.
	CanWorkInRain bool `json:"can_work_in_rain,omitempty"`

	// The amount of ticks the goal will be on cooldown before it can be used again
	GoalCooldown int `json:"goal_cooldown,omitempty"`

	// [Trigger] Event to run when the mob reaches their jobsite.
	OnArrival interface{} `json:"on_arrival,omitempty"`

	// The max interval in which a sound will play.
	SoundDelayMax int `json:"sound_delay_max,omitempty"`

	// The min interval in which a sound will play.
	SoundDelayMin int `json:"sound_delay_min,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// If "can_work_in_rain" is false, this is the maximum number of ticks left in the goal where rain will not interrupt the goal
	WorkInRainTolerance int `json:"work_in_rain_tolerance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Allows the NPC to use the composter POI to convert excess seeds into bone meal.
type Behavior_WorkComposter struct {
	// The amount of ticks the NPC will stay in their the work location
	ActiveTime int `json:"active_time,omitempty"`

	// The maximum number of times the mob will interact with the composter.
	BlockInteractionMax int `json:"block_interaction_max,omitempty"`

	// Determines whether the mob can empty a full composter.
	CanEmptyComposter bool `json:"can_empty_composter,omitempty"`

	// Determines whether the mob can add items to a composter given that it is not full.
	CanFillComposter bool `json:"can_fill_composter,omitempty"`

	// If true, this entity can work when their jobsite POI is being rained on.
	CanWorkInRain bool `json:"can_work_in_rain,omitempty"`

	// The amount of ticks the goal will be on cooldown before it can be used again
	GoalCooldown int `json:"goal_cooldown,omitempty"`

	// The maximum number of items which can be added to the composter per block interaction.
	ItemsPerUseMax int `json:"items_per_use_max,omitempty"`

	// Limits the amount of each compostable item the mob can use. Any amount held over this number will be composted if possible
	MinItemCount int `json:"min_item_count,omitempty"`

	// [Trigger] Event to run when the mob reaches their jobsite.
	OnArrival interface{} `json:"on_arrival,omitempty"`

	// Unused.
	SoundDelayMax int `json:"sound_delay_max,omitempty"`

	// Unused.
	SoundDelayMin int `json:"sound_delay_min,omitempty"`

	// Movement speed multiplier of the mob when using this AI Goal
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`

	// The maximum interval in which the mob will interact with the composter.
	UseBlockMax int `json:"use_block_max,omitempty"`

	// The minimum interval in which the mob will interact with the composter.
	UseBlockMin int `json:"use_block_min,omitempty"`

	// If "can_work_in_rain" is false, this is the maximum number of ticks left in the goal where rain will not interrupt the goal
	WorkInRainTolerance int `json:"work_in_rain_tolerance,omitempty"`

	// The priority of this behavior. Lower values are higher priority.
	Priority int `json:"priority,omitempty"`
}

// Defines an entity's melee attack and any additional effects on it.
type Attack struct {
	// Range of the random amount of damage the melee attack deals. A negative value can heal the entity instead of hurting it
	Damage float64 `json:"damage,omitempty"`

	// Duration in seconds of the status ailment applied to the damaged entity
	EffectDuration float64 `json:"effect_duration,omitempty"`

	// Identifier of the status ailment to apply to an entity attacked by this entity's melee attack
	EffectName string `json:"effect_name,omitempty"`
}

// [Not a component] List of effects to add to this entity after adding this component
type SpellEffectsAddEffects struct {
	// Effect to add to this entity. Includes 'duration' in seconds, 'amplifier' level, 'ambient' if it is to be considered an ambient effect, and 'visible' if the effect should be visible
	Effect string `json:"effect,omitempty"`
}

// Defines what mob effects to add and remove to the entity when adding this component.
type SpellEffects struct {
	// List of effects to add to this entity after adding this component
	AddEffects []SpellEffectsAddEffects `json:"add_effects,omitempty"`

	// List of identifiers of effects to be removed from this entity after adding this component
	RemoveEffects string `json:"remove_effects,omitempty"`
}

// Defines the entity's strength to carry items.
type Strength struct {
	// The maximum strength of this entity
	Max int `json:"max,omitempty"`

	// The initial value of the strength
	Value int `json:"value,omitempty"`
}

// Adds a rider to the entity. Requires minecraft:rideable.
type Addrider struct {
	// The entity type that will be riding this entity.
	EntityType string `json:"entity_type,omitempty"`

	// The spawn event that will be used when the riding entity is created.
	SpawnEvent string `json:"spawn_event,omitempty"`
}

// Causes the mob to ignore attackable targets for a given duration.
type AdmireItem struct {
	// Duration, in seconds, for which mob won't admire items if it was hurt
	CooldownAfterBeingAttacked int `json:"cooldown_after_being_attacked,omitempty"`

	// Duration, in seconds, that the mob is pacified.
	Duration int `json:"duration,omitempty"`
}

// Adds a timer for the entity to grow up. It can be accelerated by giving the entity the items it likes as defined by feedItems.
type Ageable struct {
	// List of items that the entity drops when it grows up.
	DropItems []interface{} `json:"drop_items,omitempty"`

	// Amount of time before the entity grows up, -1 for always a baby.
	Duration float64 `json:"duration,omitempty"`

	// List of items that can be fed to the entity. Includes 'item' for the item name and 'growth' to define how much time it grows up by
	FeedItems []interface{} `json:"feed_items,omitempty"`

	// Event to run when this entity grows up.
	GrowUp string `json:"grow_up,omitempty"`

	// The feed item used will transform to this item upon successful interaction. Format: itemName:auxValue
	TransformToItem string `json:"transform_to_item,omitempty"`
}

// Allows this entity to track anger towards a set of nuisances
type AngerLevel struct {
	// Anger level will decay over time. Defines how often anger towards all nuisances will be decreased by one
	AngerDecrementInterval float64 `json:"anger_decrement_interval,omitempty"`

	// Anger boost applied to angry threshold when mob gets angry
	AngryBoost uint `json:"angry_boost,omitempty"`

	// Threshold that define when the mob is considered angry at a nuisance
	AngryThreshold uint `json:"angry_threshold,omitempty"`

	// The default amount of annoyingness for any given nuisance. Specifies how much to raise anger level on each provocation
	DefaultAnnoyingness string `json:"default_annoyingness,omitempty"`

	// The maximum anger level that can be reached. Applies to any nuisance
	MaxAnger uint `json:"max_anger,omitempty"`

	// Filter that is applied to determine if a mob can be a nuisance
	NuisanceFilter *f.Filter `json:"nuisance_filter,omitempty"`

	// Sounds to play when the entity is getting provoked. Evaluated in order. First matching condition wins
	OnIncreaseSounds []interface{} `json:"on_increase_sounds,omitempty"`

	// Defines if the mob should remove target if it falls below 'angry' threshold
	RemoveTargetsBelowAngryThreshold bool `json:"remove_targets_below_angry_threshold,omitempty"`
}

// Defines the entity's 'angry' state using a timer.
type Angry struct {
	// The sound event to play when the mob is angry
	AngrySound string `json:"angry_sound,omitempty"`

	// If true, other entities of the same entity definition within the broadcastRange will also become angry
	BroadcastAnger bool `json:"broadcast_anger,omitempty"`

	// If true, other entities of the same entity definition within the broadcastRange will also become angry whenever this mob attacks
	BroadcastAngerOnAttack bool `json:"broadcast_anger_on_attack,omitempty"`

	// If true, other entities of the same entity definition within the broadcastRange will also become angry whenever this mob is attacked
	BroadcastAngerOnBeingAttacked bool `json:"broadcast_anger_on_being_attacked,omitempty"`

	// Conditions that make this entry in the list valid
	BroadcastFilters *f.Filter `json:"broadcast_filters,omitempty"`

	// Distance in blocks within which other entities of the same entity definition will become angry
	BroadcastRange int `json:"broadcast_range,omitempty"`

	// A list of entity families to broadcast anger to
	BroadcastTargets []interface{} `json:"broadcast_targets,omitempty"`

	// Event to run after the number of seconds specified in duration expires (when the entity stops being 'angry')
	CalmEvent string `json:"calm_event,omitempty"`

	// The amount of time in seconds that the entity will be angry
	Duration int `json:"duration,omitempty"`

	// Variance in seconds added to the duration [-delta, delta]
	DurationDelta int `json:"duration_delta,omitempty"`

	// Filter out mob types that it should not attack while angry (other Piglins)
	Filters *f.Filter `json:"filters,omitempty"`

	// The range of time in seconds to randomly wait before playing the sound again
	SoundInterval float64 `json:"sound_interval,omitempty"`
}

// Allows the actor to break doors assuming that that flags set up for the component to use in navigation
type Annotation_BreakDoor struct {
	// The time in seconds required to break through doors.
	BreakTime float64 `json:"break_time,omitempty"`

	// The minimum difficulty that the world must be on for this entity to break doors.
	MinDifficulty string `json:"min_difficulty,omitempty"`
}

// Allows the actor to open doors assuming that that flags set up for the component to use in navigation

type Annotation_OpenDoor struct {
}

// A component that does damage to entities that get within range.
type AreaAttack struct {
	// The type of damage that is applied to entities that enter the damage range.
	Cause string `json:"cause,omitempty"`

	// Attack cooldown (in seconds) for how often this entity can attack a target.
	DamageCooldown float64 `json:"damage_cooldown,omitempty"`

	// How much damage per tick is applied to entities that enter the damage range.
	DamagePerTick int `json:"damage_per_tick,omitempty"`

	// How close a hostile entity must be to have the damage applied.
	DamageRange float64 `json:"damage_range,omitempty"`

	// The set of entities that are valid to apply the damage to when within range.
	EntityFilter *f.Filter `json:"entity_filter,omitempty"`

	// If the entity should play their attack sound when attacking a target.
	PlayAttackSound bool `json:"play_attack_sound,omitempty"`
}

// Adds a cooldown to a mob. The intention of this cooldown is to be used to prevent the mob from attempting to aquire new attack targets.
type AttackCooldown struct {
	// [Trigger] Event to be runned when the cooldown is complete.
	AttackCooldownCompleteEvent interface{} `json:"attack_cooldown_complete_event,omitempty"`

	// Amount of time in seconds for the cooldown. Can be specified as a number or a pair of numbers (min and max).
	AttackCooldownTime float64 `json:"attack_cooldown_time,omitempty"`
}

// Enables the component to drop an item as a barter exchange.
type Barter struct {
	// Loot table that's used to drop a random item.
	BarterTable string `json:"barter_table,omitempty"`

	// Duration, in seconds, for which mob won't barter items if it was hurt
	CooldownAfterBeingAttacked int `json:"cooldown_after_being_attacked,omitempty"`
}

// Allows the player to detect and maneuver on the scaffolding block.

type BlockClimber struct {
}

// Fires off a specified event when a block in the block list is broken within the sensor range.
type BlockSensor struct {
	// List of blocks to watch for being broken to fire off a specified event. If a block is in multiple lists, multiple events will fire.
	OnBreak []interface{} `json:"on_break,omitempty"`

	// The maximum radial distance in which a specified block can be detected. The biggest radius is 32.0.
	SensorRadius float64 `json:"sensor_radius,omitempty"`

	// List of sources that break the block to listen for. If none are specified, all block breaks will be detected.
	Sources []interface{} `json:"sources,omitempty"`
}

// [Not a component] List of items that can be used to boost while riding this entity. Each item has the following properties:
type BoostableBoostItems struct {
	// This is the damage that the item will take each time it is used.
	Damage int `json:"damage,omitempty"`

	// Name of the item that can be used to boost.
	Item string `json:"item,omitempty"`

	// The item used to boost will become this item once it is used up.
	ReplaceItem string `json:"replace_item,omitempty"`
}

// Defines the conditions and behavior of a rideable entity's boost.
type Boostable struct {
	// List of items that can be used to boost while riding this entity. Each item has the following properties:
	BoostItems []BoostableBoostItems `json:"boost_items,omitempty"`

	// Time in seconds for the boost.
	Duration float64 `json:"duration,omitempty"`

	// Factor by which the entity's normal speed increases. E.g. 2.0 means go twice as fast. Requires "format_version" of 1.20 or more, otherwise the value 1.35 will be used.
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
}

// The current state of the boss for updating the boss HUD.
type Boss struct {
	// The max distance from the boss at which the boss's health bar is present on the players screen.
	HudRange int `json:"hud_range,omitempty"`

	// The name that will be displayed above the boss's health bar.
	Name string `json:"name,omitempty"`

	// Whether the sky should darken in the presence of the boss.
	ShouldDarkenSky bool `json:"should_darken_sky,omitempty"`
}

// Specifies the blocks that this entity can break as it moves around.
type BreakBlocks struct {
	// A list of the blocks that can be broken as this entity moves around
	BreakableBlocks []interface{} `json:"breakable_blocks,omitempty"`
}

// Defines what blocks this entity can breathe in and gives them the ability to suffocate.
type Breathable struct {
	// List of blocks this entity can breathe in, in addition to the other "breathes" parameters.
	BreatheBlocks []interface{} `json:"breathe_blocks,omitempty"`

	// If true, this entity can breathe in air.
	BreathesAir bool `json:"breathes_air,omitempty"`

	// If true, this entity can breathe in lava.
	BreathesLava bool `json:"breathes_lava,omitempty"`

	// If true, this entity can breathe in solid blocks.
	BreathesSolids bool `json:"breathes_solids,omitempty"`

	// If true, this entity can breathe in water.
	BreathesWater bool `json:"breathes_water,omitempty"`

	// If true, this entity will have visible bubbles while in water.
	GeneratesBubbles bool `json:"generates_bubbles,omitempty"`

	// Time in seconds to recover breath to maximum.
	InhaleTime float64 `json:"inhale_time,omitempty"`

	// List of blocks this entity can't breathe in, in addition to the other "breathes" parameters.
	NonBreatheBlocks []interface{} `json:"non_breathe_blocks,omitempty"`

	// Time in seconds between suffocation damage.
	SuffocateTime int `json:"suffocate_time,omitempty"`

	// Time in seconds the entity can hold its breath.
	TotalSupply int `json:"total_supply,omitempty"`
}

// [Not a component] The list of entity definitions that this entity can breed with.
type BreedableBreedsWith struct {
	// The entity definition of this entity's babies.
	BabyType string `json:"baby_type,omitempty"`

	// Event to run when this entity breeds.
	BreedEvent string `json:"breed_event,omitempty"`

	// The entity definition of this entity's mate.
	MateType string `json:"mate_type,omitempty"`
}

// [Not a component] Determines how likely the baby of parents with the same variant will deny that variant and take a random variant within the given range instead.
type BreedableDenyParentsVariant struct {
	// The percentage chance of denying the parents' variant.
	Chance float64 `json:"chance,omitempty"`

	// The inclusive maximum of the variant range.
	MaxVariant int `json:"max_variant,omitempty"`

	// The inclusive minimum of the variant range.
	MinVariant int `json:"min_variant,omitempty"`
}

// [Not a component] Determines how likely the babies are to NOT inherit one of their parent's variances. Values are between 0.0 and 1.0, with a higher number meaning more likely to mutate.
type BreedableMutationFactor struct {
	// The percentage chance of a mutation on the entity's color.
	Color float64 `json:"color,omitempty"`

	// The percentage chance of a mutation on the entity's extra variant type.
	ExtraVariant float64 `json:"extra_variant,omitempty"`

	// The percentage chance of a mutation on the entity's variant type.
	Variant float64 `json:"variant,omitempty"`
}

// Defines the way an entity can get into the 'love' state.
type Breedable struct {
	// If true, entities can breed while sitting
	AllowSitting bool `json:"allow_sitting,omitempty"`

	// If true, the entities will blend their attributes in the offspring after they breed.
	BlendAttributes bool `json:"blend_attributes,omitempty"`

	// Time in seconds before the Entity can breed again.
	BreedCooldown float64 `json:"breed_cooldown,omitempty"`

	// The list of items that can be used to get the entity into the 'love' state
	BreedItems []interface{} `json:"breed_items,omitempty"`

	// The list of entity definitions that this entity can breed with.
	BreedsWith []BreedableBreedsWith `json:"breeds_with,omitempty"`

	// If true, the entity will become pregnant instead of spawning a baby.
	CausesPregnancy bool `json:"causes_pregnancy,omitempty"`

	// Determines how likely the baby of parents with the same variant will deny that variant and take a random variant within the given range instead.
	DenyParentsVariant BreedableDenyParentsVariant `json:"deny_parents_variant,omitempty"`

	// The list of nearby block requirements to get the entity into the 'love' state.
	EnvironmentRequirements []interface{} `json:"environment_requirements,omitempty"`

	// Chance that up to 16 babies will spawn between 0.0 and 1.0, where 1.0 is 100%.
	ExtraBabyChance float64 `json:"extra_baby_chance,omitempty"`

	// If true, the babies will be automatically tamed if its parents are
	InheritTamed bool `json:"inherit_tamed,omitempty"`

	// The filters to run when attempting to fall in love.
	LoveFilters *f.Filter `json:"love_filters,omitempty"`

	// Determines how likely the babies are to NOT inherit one of their parent's variances. Values are between 0.0 and 1.0, with a higher number meaning more likely to mutate.
	MutationFactor BreedableMutationFactor `json:"mutation_factor,omitempty"`

	// Strategy used for mutating variants and extra variants for offspring. Current valid alternatives are 'random' and 'none'.
	MutationStrategy string `json:"mutation_strategy,omitempty"`

	// [EXPERIMENTAL] List of attributes that should benefit from parent centric attribute blending. For example, horses blend their health, movement, and jump_strength in their offspring.
	ParentCentricAttributeBlending []interface{} `json:"parent_centric_attribute_blending,omitempty"`

	// Range used to determine random extra variant.
	RandomExtraVariantMutationInterval float64 `json:"random_extra_variant_mutation_interval,omitempty"`

	// Range used to determine random variant.
	RandomVariantMutationInterval float64 `json:"random_variant_mutation_interval,omitempty"`

	// If true, the entity needs to be at full health before it can breed.
	RequireFullHealth bool `json:"require_full_health,omitempty"`

	// If true, the entities need to be tamed first before they can breed.
	RequireTame bool `json:"require_tame,omitempty"`

	// The breed item used will transform to this item upon successful interaction. Format: itemName:auxValue
	TransformToItem string `json:"transform_to_item,omitempty"`
}

// Defines the way an entity can get into the 'bribed' state.
type Bribeable struct {
	// Time in seconds before the Entity can be bribed again.
	BribeCooldown float64 `json:"bribe_cooldown,omitempty"`

	// The list of items that can be used to bribe the entity.
	BribeItems []interface{} `json:"bribe_items,omitempty"`
}

// Enables an entity to float on the specified liquid blocks.
type Buoyant struct {
	// Applies gravity each tick. Causes more of a wave simulation, but will cause more gravity to be applied outside liquids.
	ApplyGravity bool `json:"apply_gravity,omitempty"`

	// Base buoyancy used to calculate how much will a mob float.
	BaseBuoyancy float64 `json:"base_buoyancy,omitempty"`

	// Probability for a big wave hitting the entity. Only used if `simulate_waves` is true.
	BigWaveProbability float64 `json:"big_wave_probability,omitempty"`

	// Multiplier for the speed to make a big wave. Triggered depending on 'big_wave_probability'.
	BigWaveSpeed float64 `json:"big_wave_speed,omitempty"`

	// How much an actor will be dragged down when the Buoyancy Component is removed.
	DragDownOnBuoyancyRemoved float64 `json:"drag_down_on_buoyancy_removed,omitempty"`

	// List of blocks this entity can float on. Must be a liquid block.
	LiquidBlocks []interface{} `json:"liquid_blocks,omitempty"`

	// Should the movement simulate waves going through.
	SimulateWaves bool `json:"simulate_waves,omitempty"`
}

// Specifies if/how a mob burns in daylight.

type BurnsInDaylight struct {
}

// Specifies hunt celebration behaviour.
type CelebrateHunt struct {
	// If true, celebration will be broadcasted to other entities in the radius.
	Broadcast bool `json:"broadcast,omitempty"`

	// The list of conditions that target of hunt must satisfy to initiate celebration.
	CeleberationTargets *f.Filter `json:"celeberation_targets,omitempty"`

	// The sound event to play when the mob is celebrating
	CelebrateSound string `json:"celebrate_sound,omitempty"`

	// Duration, in seconds, of celebration
	Duration int `json:"duration,omitempty"`

	// If broadcast is enabled, specifies the radius in which it will notify other entities for celebration.
	Radius float64 `json:"radius,omitempty"`

	// The range of time in seconds to randomly wait before playing the sound again
	SoundInterval float64 `json:"sound_interval,omitempty"`
}

// Sets the width and height of the Entity's collision box.
type CollisionBox struct {
	// Height of the collision box in blocks. A negative value will be assumed to be 0.
	Height float64 `json:"height,omitempty"`

	// Width of the collision box in blocks. A negative value will be assumed to be 0.
	Width float64 `json:"width,omitempty"`
}

// Gives Regeneration I and removes Mining Fatigue from the mob that kills the Actor's attack target.
type CombatRegeneration struct {
	// Determines if the mob will grant mobs of the same type combat buffs if they kill the target.
	ApplyToFamily bool `json:"apply_to_family,omitempty"`

	// Determines if the mob will grant itself the combat buffs if it kills the target.
	ApplyToSelf bool `json:"apply_to_self,omitempty"`

	// The duration in seconds of Regeneration I added to the mob.
	RegenerationDuration int `json:"regeneration_duration,omitempty"`
}

// Defines the Conditional Spatial Update Bandwidth Optimizations of this entity.
type ConditionalBandwidthOptimization struct {
	// The object containing the conditional bandwidth optimization values.
	ConditionalValues []interface{} `json:"conditional_values,omitempty"`

	// The object containing the default bandwidth optimization values.
	DefaultValues interface{} `json:"default_values,omitempty"`
}

// List of hitboxes for melee and ranged hits against the entity.
type CustomHitTest struct {
	// Comma seperated list of hitboxes.
	Hitboxes []interface{} `json:"hitboxes,omitempty"`
}

// Applies defined amount of damage to the entity at specified intervals.
type DamageOverTime struct {
	// Amount of damage caused each hurt.
	DamagePerHurt int `json:"damage_per_hurt,omitempty"`

	// Time in seconds between damage.
	TimeBetweenHurt float64 `json:"time_between_hurt,omitempty"`
}

// [Not a component] List of triggers with the events to call when taking specific kinds of damage.
type DamageSensorTriggers struct {
	// Type of damage that triggers the events.
	Cause string `json:"cause,omitempty"`

	// A modifier that adds to/removes from the base damage from the damage cause. It does not reduce damage to less than 0.
	DamageModifier float64 `json:"damage_modifier,omitempty"`

	// A multiplier that modifies the base damage from the damage cause. If deals_damage is true the multiplier can only reduce the damage the entity will take to a minimum of 1.
	DamageMultiplier float64 `json:"damage_multiplier,omitempty"`

	// If true, the damage dealt to the entity will take away health from it, set to false to make the entity ignore that damage.
	DealsDamage bool `json:"deals_damage,omitempty"`

	// Specifies filters for entity definitions and events.
	OnDamage interface{} `json:"on_damage,omitempty"`

	// Defines what sound to play, if any, when the on_damage filters are met.
	OnDamageSoundEvent string `json:"on_damage_sound_event,omitempty"`
}

// Defines what events to call when this entity is damaged by specific entities or items.
type DamageSensor struct {
	// List of triggers with the events to call when taking specific kinds of damage.
	Triggers []DamageSensorTriggers `json:"triggers,omitempty"`
}

// [Not a component] Specifies if the "min_distance" and "max_distance" are used in the standard despawn rules.
type DespawnDespawnFromDistance struct {
	// maximum distance for standard despawn rules to instantly despawn the mob.
	MaxDistance int `json:"max_distance,omitempty"`

	// minimum distance for standard despawn rules to try to despawn the mob.
	MinDistance int `json:"min_distance,omitempty"`
}

// Despawns the Actor when the despawn rules or optional filters evaluate to true.
type Despawn struct {
	// Determines if "min_range_random_chance" is used in the standard despawn rules
	DespawnFromChance bool `json:"despawn_from_chance,omitempty"`

	// Specifies if the "min_distance" and "max_distance" are used in the standard despawn rules.
	DespawnFromDistance DespawnDespawnFromDistance `json:"despawn_from_distance,omitempty"`

	// Determines if the "min_range_inactivity_timer" is used in the standard despawn rules.
	DespawnFromInactivity bool `json:"despawn_from_inactivity,omitempty"`

	// Determines if the mob is instantly despawned at the edge of simulation distance in the standard despawn rules.
	DespawnFromSimulationEdge bool `json:"despawn_from_simulation_edge,omitempty"`

	// The list of conditions that must be satisfied before the Actor is despawned. If a filter is defined then standard despawn rules are ignored.
	Filters *f.Filter `json:"filters,omitempty"`

	// The amount of time in seconds that the mob must be inactive.
	MinRangeInactivityTimer int `json:"min_range_inactivity_timer,omitempty"`

	// A random chance between 1 and the given value.
	MinRangeRandomChance int `json:"min_range_random_chance,omitempty"`

	// If true, all entities linked to this entity in a child relationship (eg. leashed) will also be despawned.
	RemoveChildEntities bool `json:"remove_child_entities,omitempty"`
}

// Adds a timer for drying out that will count down and fire 'dried_out_event' or will stop as soon as the entity will get under rain or water and fire 'stopped_drying_out_event'
type DryingOutTimer struct {
	// Event to fire when the drying out time runs out.
	DriedOutEvent string `json:"dried_out_event,omitempty"`

	// Event to fire when entity was already dried out but received increase in water supply.
	RecoverAfterDriedOutEvent string `json:"recover_after_dried_out_event,omitempty"`

	// Event to fire when entity stopped drying out, for example got into water or under rain.
	StoppedDryingOutEvent string `json:"stopped_drying_out_event,omitempty"`

	// Amount of time in seconds to dry out fully.
	TotalTime float64 `json:"total_time,omitempty"`

	// Optional amount of additional time in seconds given by using splash water bottle on entity.
	WaterBottleRefillTime float64 `json:"water_bottle_refill_time,omitempty"`
}

// Defines this entity's ability to trade with players.
type EconomyTradeTable struct {
	// Determines when the mob transforms, if the trades should be converted when the new mob has a economy_trade_table. When the trades are converted, the mob will generate a new trade list with their new trade table, but then it will try to convert any of the same trades over to have the same enchantments and user data. For example, if the original has a Emerald to Enchanted Iron Sword (Sharpness 1), and the new trade also has an Emerald for Enchanted Iron Sword, then the enchantment will be Sharpness 1.
	ConvertTradesEconomy bool `json:"convert_trades_economy,omitempty"`

	// How much should the discount be modified by when the player has cured the Zombie Villager. Can be specified as a pair of numbers (low-tier trade discount and high-tier trade discount)
	CuredDiscount float64 `json:"cured_discount,omitempty"`

	// Name to be displayed while trading with this entity
	DisplayName string `json:"display_name,omitempty"`

	// Used in legacy prices to determine how much should Demand be modified by when the player has the Hero of the Village mob effect
	HeroDemandDiscount int `json:"hero_demand_discount,omitempty"`

	// The max the discount can be modified by when the player has cured the Zombie Villager. Can be specified as a pair of numbers (low-tier trade discount and high-tier trade discount)
	MaxCuredDiscount float64 `json:"max_cured_discount,omitempty"`

	// The max the discount can be modified by when the player has cured a nearby Zombie Villager
	MaxNearbyCuredDiscount int `json:"max_nearby_cured_discount,omitempty"`

	// How much should the discount be modified by when the player has cured a nearby Zombie Villager
	NearbyCuredDiscount int `json:"nearby_cured_discount,omitempty"`

	// Used to determine if trading with entity opens the new trade screen
	NewScreen bool `json:"new_screen,omitempty"`

	// Determines if the trades should persist when the mob transforms. This makes it so that the next time the mob is transformed to something with a trade_table or economy_trade_table, then it keeps their trades.
	PersistTrades bool `json:"persist_trades,omitempty"`

	// Show an in game trade screen when interacting with the mob.
	ShowTradeScreen bool `json:"show_trade_screen,omitempty"`

	// File path relative to the resource pack root for this entity's trades
	Table string `json:"table,omitempty"`

	// Determines whether the legacy formula is used to determines the trade prices.
	UseLegacyPriceFormula bool `json:"use_legacy_price_formula,omitempty"`
}

// A component that fires an event when a set of conditions are met by other entities within the defined range.
type EntitySensor struct {
	// Event to fire when the conditions are met.
	Event string `json:"event,omitempty"`

	// The set of conditions that must be satisfied to fire the event.
	EventFilters *f.Filter `json:"event_filters,omitempty"`

	// The maximum number of entities that must pass the filter conditions for the event to send.
	MaximumCount int `json:"maximum_count,omitempty"`

	// The minimum number of entities that must pass the filter conditions for the event to send.
	MinimumCount int `json:"minimum_count,omitempty"`

	// If true the sensor range is additive on top of the entity's size.
	RelativeRange bool `json:"relative_range,omitempty"`

	// If true requires all nearby entities to pass the filter conditions for the event to send.
	RequireAll bool `json:"require_all,omitempty"`

	// The maximum distance another entity can be from this and have the filters checked against it.
	SensorRange float64 `json:"sensor_range,omitempty"`
}

// Creates a trigger based on environment conditions.
type EnvironmentSensor struct {
	// The list of triggers that fire when the environment conditions match the given filter criteria.
	Triggers []interface{} `json:"triggers,omitempty"`
}

// The entity puts on the desired equipment.

type EquipItem struct {
}

// [Not a component] List of slots and the item that can be equipped.
type EquippableSlots struct {
	// The list of items that can go in this slot.
	AcceptedItems []interface{} `json:"accepted_items,omitempty"`

	// Text to be displayed when the entity can be equipped with this item when playing with Touch-screen controls.
	InteractText string `json:"interact_text,omitempty"`

	// Identifier of the item that can be equipped for this slot.
	Item string `json:"item,omitempty"`

	// Event to trigger when this entity is equipped with this item.
	OnEquip string `json:"on_equip,omitempty"`

	// Event to trigger when this item is removed from this entity.
	OnUnequip string `json:"on_unequip,omitempty"`

	// The slot number of this slot.
	Slot int `json:"slot,omitempty"`
}

// Defines an entity's behavior for having items equipped to it.
type Equippable struct {
	// List of slots and the item that can be equipped.
	Slots []EquippableSlots `json:"slots,omitempty"`
}

// Defines how much exhaustion each player action should take.
type ExhaustionValues struct {
	// Amount of exhaustion applied when attacking.
	Attack float64 `json:"attack,omitempty"`

	// Amount of exhaustion applied when taking damage.
	Damage float64 `json:"damage,omitempty"`

	// Amount of exhaustion applied when healed through food regeneration.
	Heal float64 `json:"heal,omitempty"`

	// Amount of exhaustion applied when jumping.
	Jump float64 `json:"jump,omitempty"`

	// Amount of exhaustion applied when mining.
	Mine float64 `json:"mine,omitempty"`

	// Amount of exhaustion applied when sprinting.
	Sprint float64 `json:"sprint,omitempty"`

	// Amount of exhaustion applied when sprint jumping.
	SprintJump float64 `json:"sprint_jump,omitempty"`

	// Amount of exhaustion applied when swimming.
	Swim float64 `json:"swim,omitempty"`

	// Amount of exhaustion applied when walking.
	Walk float64 `json:"walk,omitempty"`
}

// .
type ExperienceReward struct {
	// [Molang String] A Molang expression defining the amount of experience rewarded when this entity is successfully bred. An array of expressions adds each expression's result together for a final total.
	OnBred string `json:"on_bred,omitempty"`

	// [Molang String] A Molang expression defining the amount of experience rewarded when this entity dies. An array of expressions adds each expression's result together for a final total.
	OnDeath string `json:"on_death,omitempty"`
}

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

// Allows entities to flock in groups in water or not.
type Flocking struct {
	// The amount of blocks away the entity will look at to push away from.
	BlockDistance float64 `json:"block_distance,omitempty"`

	// The weight of the push back away from blocks.
	BlockWeight float64 `json:"block_weight,omitempty"`

	// The amount of push back given to a flocker that breaches out of the water.
	BreachInfluence float64 `json:"breach_influence,omitempty"`

	// The threshold in which to start applying cohesion.
	CohesionThreshold float64 `json:"cohesion_threshold,omitempty"`

	// The weight applied for the cohesion steering of the flock.
	CohesionWeight float64 `json:"cohesion_weight,omitempty"`

	// The weight on which to apply on the goal output.
	GoalWeight float64 `json:"goal_weight,omitempty"`

	// Determines the high bound amount of entities that can be allowed in the flock.
	HighFlockLimit int `json:"high_flock_limit,omitempty"`

	// Tells the Flocking Component if the entity exists in water.
	InWater bool `json:"in_water,omitempty"`

	// The area around the entity that allows others to be added to the flock.
	InfluenceRadius float64 `json:"influence_radius,omitempty"`

	// The distance in which the flocker will stop applying cohesion.
	InnnerCohesionThreshold float64 `json:"innner_cohesion_threshold,omitempty"`

	// The percentage chance between 0-1 that a fish will spawn and not want to join flocks. Invalid values will be capped at the end points.
	LonerChance float64 `json:"loner_chance,omitempty"`

	// Determines the low bound amount of entities that can be allowed in the flock.
	LowFlockLimit int `json:"low_flock_limit,omitempty"`

	// Tells the flockers that they can only match similar entities that also match the variant, mark variants, and color data of the other potential flockers.
	MatchVariants bool `json:"match_variants,omitempty"`

	// The max height allowable in the air or water.
	MaxHeight float64 `json:"max_height,omitempty"`

	// The min height allowable in the air or water.
	MinHeight float64 `json:"min_height,omitempty"`

	// The distance that is determined to be to close to another flocking and to start applying separation.
	SeparationThreshold float64 `json:"separation_threshold,omitempty"`

	// The weight applied to the separation of the flock.
	SeparationWeight float64 `json:"separation_weight,omitempty"`

	// Tells the flockers that they will follow flocks based on the center of mass.
	UseCenterOfMass bool `json:"use_center_of_mass,omitempty"`
}

// Allows an entity to emit `entityMove`, `swim` and `flap` game events, depending on the block the entity is moving through. It is added by default to every mob. Add it again to override its behavior.
type GameEventMovementTracking struct {
	// If true, the `flap` game event will be emitted when the entity moves through air.
	EmitFlap bool `json:"emit_flap,omitempty"`

	// If true, the `entityMove` game event will be emitted when the entity moves on ground or through a solid.
	EmitMove bool `json:"emit_move,omitempty"`

	// If true, the `swim` game event will be emitted when the entity moves through a liquid.
	EmitSwim bool `json:"emit_swim,omitempty"`
}

// [Not a component] The range of positive integer allele values for this gene. Spawned mobs will have a random number in this range assigned to them.
type GeneticsAlleleRange struct {
	// Upper bound of the allele values for this gene.
	RangeMax int `json:"range_max,omitempty"`

	// Lower bound of the allele values for this gene.
	RangeMin int `json:"range_min,omitempty"`
}

// [Not a component] The list of genetic variants for this gene. These check for particular allele combinations and fire events when all of them are satisfied.
type GeneticsGeneticVariants struct {
	// Event to run when this mob is created and matches the allele conditions.
	BirthEvent string `json:"birth_event,omitempty"`

	// If this value is non-negative, compare both the mob's main and hidden alleles with this value for a match with both. Can also be a range of integers.
	BothAllele int `json:"both_allele,omitempty"`

	// If this value is non-negative, compare both the mob's main and hidden alleles with this value for a match with either. Can also be a range of integers.
	EitherAllele int `json:"either_allele,omitempty"`

	// If this value is non-negative, compare the mob's hidden allele with this value for a match. Can also be a range of integers.
	HiddenAllele int `json:"hidden_allele,omitempty"`

	// If this value is non-negative, compare the mob's main allele with this value for a match. Can also be a range of integers.
	MainAllele int `json:"main_allele,omitempty"`
}

// [Not a component] The list of genes that this entity has and will cross with a partner during breeding.
type GeneticsGenes struct {
	// The range of positive integer allele values for this gene. Spawned mobs will have a random number in this range assigned to them.
	AlleleRange GeneticsAlleleRange `json:"allele_range,omitempty"`

	// The list of genetic variants for this gene. These check for particular allele combinations and fire events when all of them are satisfied.
	GeneticVariants []GeneticsGeneticVariants `json:"genetic_variants,omitempty"`

	// If this value is non-negative, overrides the chance for this gene that an allele will be replaced with a random one instead of the parent's allele during birth. Non-negative values greater than `1` will be the same as the value `1`.
	MutationRate float64 `json:"mutation_rate,omitempty"`

	// The name of the gene.
	Name string `json:"name,omitempty"`
}

// Defines the way a mob's genes and alleles are passed on to its offspring, and how those traits manifest in the child. Compatible parent genes are crossed together, the alleles are handed down from the parents to the child, and any matching genetic variants fire off JSON events to modify the child and express the traits.
type Genetics struct {
	// The list of genes that this entity has and will cross with a partner during breeding.
	Genes []GeneticsGenes `json:"genes,omitempty"`

	// Chance that an allele will be replaced with a random one instead of the parent's allele during birth.
	MutationRate float64 `json:"mutation_rate,omitempty"`
}

// Defines sets of items that can be used to trigger events when used on this entity. The item will also be taken and placed in the entity's inventory.
type Giveable struct {
	// An optional cool down in seconds to prevent spamming interactions.
	Cooldown float64 `json:"cooldown,omitempty"`

	// The list of items that can be given to the entity to place in their inventory.
	Items []interface{} `json:"items,omitempty"`

	// Event to fire when the correct item is given.
	OnGive string `json:"on_give,omitempty"`
}

// Keeps track of entity group size in the given radius.
type GroupSize struct {
	// The list of conditions that must be satisfied for other entities to be counted towards group size.
	Filters *f.Filter `json:"filters,omitempty"`

	// Radius from center of entity.
	Radius float64 `json:"radius,omitempty"`
}

// Could increase crop growth when entity walks over crop
type GrowsCrop struct {
	// Value between 0-1. Chance of success per tick.
	Chance float64 `json:"chance,omitempty"`

	// Number of charges
	Charges int `json:"charges,omitempty"`
}

// [Not a component] The array of items that can be used to heal this entity.
type HealableItems struct {
	// The amount of health this entity gains when fed this item.
	HealAmount int `json:"heal_amount,omitempty"`

	// Item identifier that can be used to heal this entity.
	Item string `json:"item,omitempty"`
}

// Defines the interactions with this entity for healing it.
type Healable struct {
	// The filter group that defines the conditions for using this item to heal the entity.
	Filters *f.Filter `json:"filters,omitempty"`

	// Determines if item can be used regardless of entity being at full health.
	ForceUse bool `json:"force_use,omitempty"`

	// The array of items that can be used to heal this entity.
	Items []HealableItems `json:"items,omitempty"`
}

// Defines the entity's heartbeat.
type Heartbeat struct {
	// [Molang String] A Molang expression defining the inter-beat interval in seconds. A value of zero or less means no heartbeat.
	Interval string `json:"interval,omitempty"`

	// Level sound event to be played as the heartbeat sound.
	SoundEvent string `json:"sound_event,omitempty"`
}

// Saves a home pos for when the the entity is spawned.
type Home struct {
	// Optional block list that the home position will be associated with. If any of the blocks no longer exist at that position, the home restriction is removed. Example syntax: minecraft:sand. Not supported: minecraft:sand:1
	HomeBlockList []interface{} `json:"home_block_list,omitempty"`

	// The radius that the entity will be restricted to in relation to its home
	RestrictionRadius int `json:"restriction_radius,omitempty"`
}

// Defines a set of conditions under which an entity should take damage.
type HurtOnCondition struct {
	// List of damage conditions that when met can cause damage to the entity.
	DamageConditions []interface{} `json:"damage_conditions,omitempty"`
}

// Verifies whether the entity is inside any of the listed blocks.
type InsideBlockNotifier struct {
	// List of blocks, with certain block states, that we are monitoring to see if the entity is inside.
	BlockList []interface{} `json:"block_list,omitempty"`
}

// Adds a timer since last rested to see if phantoms should spawn.
type Insomnia struct {
	// Number of days the mob has to stay up until the insomnia effect begins.
	DaysUntilInsomnia float64 `json:"days_until_insomnia,omitempty"`
}

// Despawns the Actor immediately.
type InstantDespawn struct {
	// If true, all entities linked to this entity in a child relationship (eg. leashed) will also be despawned.
	RemoveChildEntities bool `json:"remove_child_entities,omitempty"`
}

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

	// The entity's equipment slot to equip the item to, if any, upon successful interaction.
	EquipItemSlot int `json:"equip_item_slot,omitempty"`

	// The amount of health this entity will recover or hurt when interacting with this item. Negative values will harm the entity.
	HealthAmount int `json:"health_amount,omitempty"`

	// The amount of damage the item will take when used to interact with this entity. A value of 0 means the item won't lose durability.
	HurtItem int `json:"hurt_item,omitempty"`

	// Text to show when the player is able to interact in this way with this entity when playing with Touch-screen controls.
	InteractText string `json:"interact_text,omitempty"`

	// Event to fire when the interaction occurs.
	OnInteract string `json:"on_interact,omitempty"`

	// Particle effect that will be triggered at the start of the interaction.
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

	// Vibration to emit when the interaction occurs. Admitted values are entity_interact (used by default), shear, and none (no vibration emitted).
	Vibration string `json:"vibration,omitempty"`
}

// Defines this entity's inventory properties.
type Inventory struct {
	// Number of slots that this entity can gain per extra strength
	AdditionalSlotsPerStrength int `json:"additional_slots_per_strength,omitempty"`

	// If true, the contents of this inventory can be removed by a hopper
	CanBeSiphonedFrom bool `json:"can_be_siphoned_from,omitempty"`

	// Type of container this entity has. Can be horse, minecart_chest, chest_boat, minecart_hopper, inventory, container or hopper
	ContainerType string `json:"container_type,omitempty"`

	// Number of slots the container has
	InventorySize int `json:"inventory_size,omitempty"`

	// If true, the entity will not drop its inventory on death
	Private bool `json:"private,omitempty"`

	// If true, the entity's inventory can only be accessed by its owner or itself
	RestrictToOwner bool `json:"restrict_to_owner,omitempty"`
}

// Determines that this entity is an item hopper.

type ItemHopper struct {
}

// Defines a dynamic type jump control that will change jump properties based on the speed modifier of the mob.

type Jump_Dynamic struct {
}

// Gives the entity the ability to jump.
type Jump_Static struct {
	// The initial vertical velocity for the jump
	JumpPower float64 `json:"jump_power,omitempty"`
}

// Allows this entity to be leashed and defines the conditions and events for this entity when is leashed.
type Leashable struct {
	// If true, players can leash this entity even if it is already leashed to another mob.
	CanBeStolen bool `json:"can_be_stolen,omitempty"`

	// Distance in blocks at which the leash stiffens, restricting movement.
	HardDistance float64 `json:"hard_distance,omitempty"`

	// Distance in blocks at which the leash breaks.
	MaxDistance float64 `json:"max_distance,omitempty"`

	// Event to call when this entity is leashed.
	OnLeash string `json:"on_leash,omitempty"`

	// Event to call when this entity is unleashed.
	OnUnleash string `json:"on_unleash,omitempty"`

	// Distance in blocks at which the 'spring' effect starts acting to keep this entity close to the entity that leashed it.
	SoftDistance float64 `json:"soft_distance,omitempty"`
}

// Defines the behavior when another entity looks at this entity.
type Lookat struct {
	// If true, invulnerable entities (e.g. Players in creative mode) are considered valid targets.
	AllowInvulnerable bool `json:"allow_invulnerable,omitempty"`

	// Defines the entities that can trigger this component.
	Filters *f.Filter `json:"filters,omitempty"`

	// The range for the random amount of time during which the entity is 'cooling down' and won't get angered or look for a target.
	LookCooldown float64 `json:"look_cooldown,omitempty"`

	// The event identifier to run when the entities specified in filters look at this entity.
	LookEvent string `json:"look_event,omitempty"`

	// Maximum distance this entity will look for another entity looking at it.
	SearchRadius float64 `json:"search_radius,omitempty"`

	// If true, this entity will set the attack target as the entity that looked at it.
	SetTarget bool `json:"set_target,omitempty"`
}

// This component is used to implement part of the Wandering Trader behavior

type ManagedWanderingTrader struct {
}

// A component that applies a mob effect to entities that get within range.
type MobEffect struct {
	// Time in seconds to wait between each application of the effect.
	CooldownTime int `json:"cooldown_time,omitempty"`

	// How close a hostile entity must be to have the mob effect applied.
	EffectRange float64 `json:"effect_range,omitempty"`

	// How long the applied mob effect lasts in seconds.
	EffectTime int `json:"effect_time,omitempty"`

	// The set of entities that are valid to apply the mob effect to.
	EntityFilter *f.Filter `json:"entity_filter,omitempty"`

	// The mob effect that is applied to entities that enter this entities effect range.
	MobEffect string `json:"mob_effect,omitempty"`
}

// This move control allows the mob to swim in water and walk on land.
type Movement_Amphibious struct {
	// The maximum number in degrees the mob can turn per tick.
	MaxTurn float64 `json:"max_turn,omitempty"`
}

// This component accents the movement of an entity.
type Movement_Basic struct {
	// The maximum number in degrees the mob can turn per tick.
	MaxTurn float64 `json:"max_turn,omitempty"`
}

// This move control causes the mob to fly.
type Movement_Fly struct {
	// The maximum number in degrees the mob can turn per tick.
	MaxTurn float64 `json:"max_turn,omitempty"`
}

// This move control allows a mob to fly, swim, climb, etc.
type Movement_Generic struct {
	// The maximum number in degrees the mob can turn per tick.
	MaxTurn float64 `json:"max_turn,omitempty"`
}

// This move control causes the mob to hover.
type Movement_Hover struct {
	// The maximum number in degrees the mob can turn per tick.
	MaxTurn float64 `json:"max_turn,omitempty"`
}

// Move control that causes the mob to jump as it moves with a specified delay between jumps.
type Movement_Jump struct {
	// Delay after landing when using the slime move control.
	JumpDelay float64 `json:"jump_delay,omitempty"`

	// The maximum number in degrees the mob can turn per tick.
	MaxTurn float64 `json:"max_turn,omitempty"`
}

// This move control causes the mob to hop as it moves.
type Movement_Skip struct {
	// The maximum number in degrees the mob can turn per tick.
	MaxTurn float64 `json:"max_turn,omitempty"`
}

// This move control causes the mob to sway side to side giving the impression it is swimming.
type Movement_Sway struct {
	// The maximum number in degrees the mob can turn per tick.
	MaxTurn float64 `json:"max_turn,omitempty"`

	// Strength of the sway movement.
	SwayAmplitude float64 `json:"sway_amplitude,omitempty"`

	// Multiplier for the frequency of the sway movement.
	SwayFrequency float64 `json:"sway_frequency,omitempty"`
}

// [Not a component] Describes the special names for this entity and the events to call when the entity acquires those names
type NameableNameActions struct {
	// List of special names that will cause the events defined in 'on_named' to fire
	NameFilter string `json:"name_filter,omitempty"`

	// Event to be called when this entity acquires the name specified in 'name_filter'
	OnNamed string `json:"on_named,omitempty"`
}

// Allows this entity to be named (e.g. using a name tag).
type Nameable struct {
	// If true, this entity can be renamed with name tags
	AllowNameTagRenaming bool `json:"allow_name_tag_renaming,omitempty"`

	// If true, the name will always be shown
	AlwaysShow bool `json:"always_show,omitempty"`

	// Trigger to run when the entity gets named
	DefaultTrigger string `json:"default_trigger,omitempty"`

	// Describes the special names for this entity and the events to call when the entity acquires those names
	NameActions NameableNameActions `json:"name_actions,omitempty"`
}

// Allows this entity to generate paths that include vertical walls like the vanilla Spiders do.
type Navigation_Climb struct {
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

// Allows this entity to generate paths by flying around the air like the regular Ghast.
type Navigation_Float struct {
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

// Allows this entity to generate paths in the air like the vanilla Parrots do.
type Navigation_Fly struct {
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

// Allows this entity to generate paths by walking, swimming, flying and/or climbing around and jumping up and down a block.
type Navigation_Generic struct {
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

// Allows this entity to generate paths in the air like the vanilla Bees do. Keeps them from falling out of the skies and doing predictive movement.
type Navigation_Hover struct {
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

// Allows this entity to generate paths by walking around and jumping up and down a block like regular mobs.
type Navigation_Walk struct {
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

// Defines the entity's 'out of control' state.

type OutOfControl struct {
}

// Defines the entity's 'peek' behavior, defining the events that should be called during it.
type Peek struct {
	// Event to call when the entity is done peeking.
	OnClose string `json:"on_close,omitempty"`

	// Event to call when the entity starts peeking.
	OnOpen string `json:"on_open,omitempty"`

	// Event to call when the entity's target entity starts peeking.
	OnTargetOpen string `json:"on_target_open,omitempty"`
}

// Defines whether an entity should be persistent in the game world.

type Persistent struct {
}

// Defines physics properties of an actor, including if it is affected by gravity or if it collides with objects.
type Physics struct {
	// Whether or not the object collides with things.
	HasCollision bool `json:"has_collision,omitempty"`

	// Whether or not the entity is affected by gravity.
	HasGravity bool `json:"has_gravity,omitempty"`

	// Whether or not the entity should be pushed towards the nearest open area when stuck inside a block.
	PushTowardsClosestSpace bool `json:"push_towards_closest_space,omitempty"`
}

// Specifies costing information for mobs that prefer to walk on preferred paths.
type PreferredPath struct {
	// Cost for non-preferred blocks
	DefaultBlockCost float64 `json:"default_block_cost,omitempty"`

	// Added cost for jumping up a node
	JumpCost int `json:"jump_cost,omitempty"`

	// Distance mob can fall without taking damage
	MaxFallBlocks int `json:"max_fall_blocks,omitempty"`

	// A list of blocks with their associated cost
	PreferredPathBlocks []interface{} `json:"preferred_path_blocks,omitempty"`
}

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
	Gravity float64 `json:"gravity,omitempty"`

	// The sound that plays when the projectile hits something
	HitSound string `json:"hit_sound,omitempty"`

	// If true, the projectile homes in to the nearest entity
	Homing bool `json:"homing,omitempty"`

	// The fraction of the projectile's speed maintained every frame while traveling in air
	Inertia float64 `json:"inertia,omitempty"`

	// If true, the projectile will be treated as dangerous to the players
	IsDangerous bool `json:"is_dangerous,omitempty"`

	// If true, the projectile will knock back the entity it hits
	Knockback bool `json:"knockback,omitempty"`

	// If true, the entity hit will be struck by lightning
	Lightning bool `json:"lightning,omitempty"`

	// The fraction of the projectile's speed maintained every frame while traveling in water
	LiquidInertia float64 `json:"liquid_inertia,omitempty"`

	// If true, the projectile can hit multiple entities per flight
	MultipleTargets bool `json:"multiple_targets,omitempty"`

	// [Vector3 [a,b,c]] The offset from the entity's anchor where the projectile will spawn
	Offset []float64 `json:"offset,omitempty"`

	// Time in seconds that the entity hit will be on fire for
	OnFireTime float64 `json:"on_fire_time,omitempty"`

	// Particle to use upon collision
	Particle string `json:"particle,omitempty"`

	// Defines the effect the arrow will apply to the entity it hits
	PotionEffect int `json:"potion_effect,omitempty"`

	// Determines the velocity of the projectile
	Power float64 `json:"power,omitempty"`

	// If true, this entity will be reflected back when hit
	ReflectOnHurt bool `json:"reflect_on_hurt,omitempty"`

	// If true, damage will be randomized based on damage and speed
	SemiRandomDiffDamage bool `json:"semi_random_diff_damage,omitempty"`

	// The sound that plays when the projectile is shot
	ShootSound string `json:"shoot_sound,omitempty"`

	// If true, the projectile will be shot towards the target of the entity firing it
	ShootTarget bool `json:"shoot_target,omitempty"`

	// If true, the projectile will bounce upon hit
	ShouldBounce bool `json:"should_bounce,omitempty"`

	// If true, the projectile will be treated like a splash potion
	SplashPotion bool `json:"splash_potion,omitempty"`

	// Radius in blocks of the 'splash' effect
	SplashRange float64 `json:"splash_range,omitempty"`

	// The base accuracy. Accuracy is determined by the formula uncertaintyBase - difficultyLevel * uncertaintyMultiplier
	UncertaintyBase float64 `json:"uncertainty_base,omitempty"`

	// Determines how much difficulty affects accuracy. Accuracy is determined by the formula uncertaintyBase - difficultyLevel * uncertaintyMultiplier
	UncertaintyMultiplier float64 `json:"uncertainty_multiplier,omitempty"`
}

// Defines what can push an entity between other entities and pistons.
type Pushable struct {
	// Whether the entity can be pushed by other entities.
	IsPushable bool `json:"is_pushable,omitempty"`

	// Whether the entity can be pushed by pistons safely.
	IsPushableByPiston bool `json:"is_pushable_by_piston,omitempty"`
}

// Attempts to trigger a raid at the entity's location.
type RaidTrigger struct {
	// Event to run when a raid is triggered on the village.
	TriggeredEvent string `json:"triggered_event,omitempty"`
}

// Defines the entity's movement on the rails. An entity with this component is only allowed to move on the rail.
type RailMovement struct {
	// Maximum speed that this entity will move at when on the rail.
	MaxSpeed float64 `json:"max_speed,omitempty"`
}

type RailSensor struct {
	// If true, on tick this entity will trigger its on_deactivate behavior
	CheckBlockTypes bool `json:"check_block_types,omitempty"`

	// If true, this entity will eject all of its riders when it passes over an activated rail
	EjectOnActivate bool `json:"eject_on_activate,omitempty"`

	// If true, this entity will eject all of its riders when it passes over a deactivated rail
	EjectOnDeactivate bool `json:"eject_on_deactivate,omitempty"`

	// Event to call when the rail is activated
	OnActivate string `json:"on_activate,omitempty"`

	// Event to call when the rail is deactivated
	OnDeactivate string `json:"on_deactivate,omitempty"`

	// If true, command blocks will start ticking when passing over an activated rail
	TickCommandBlockOnActivate bool `json:"tick_command_block_on_activate,omitempty"`

	// If false, command blocks will stop ticking when passing over a deactivated rail
	TickCommandBlockOnDeactivate bool `json:"tick_command_block_on_deactivate,omitempty"`
}

// Defines the ravager's response to their melee attack being blocked.
type RavagerBlocked struct {
	// The strength with which blocking entities should be knocked back
	KnockbackStrength float64 `json:"knockback_strength,omitempty"`

	// A list of weighted responses to the melee attack being blocked
	ReactionChoices []interface{} `json:"reaction_choices,omitempty"`
}

// [Not a component] The list of positions and number of riders for each position for entities riding this entity
type RideableSeats struct {
	// Angle in degrees that a rider is allowed to rotate while riding this entity. Omit this property for no limit
	LockRiderRotation float64 `json:"lock_rider_rotation,omitempty"`

	// Defines the maximum number of riders that can be riding this entity for this seat to be valid
	MaxRiderCount int `json:"max_rider_count,omitempty"`

	// Defines the minimum number of riders that need to be riding this entity before this seat can be used
	MinRiderCount int `json:"min_rider_count,omitempty"`

	// [Vector3 [a,b,c]] Position of this seat relative to this entity's position
	Position []float64 `json:"position,omitempty"`

	// [Molang String] Offset to rotate riders by
	RotateRiderBy string `json:"rotate_rider_by,omitempty"`
}

// Determines whether this entity can be ridden. Allows specifying the different seat positions and quantity.
type Rideable struct {
	// The seat that designates the driver of the entity. This is only observed by the horse/boat styles of riding; minecarts/entities with "minecraft:controlled_by_player" give control to any player in any seat.
	ControllingSeat int `json:"controlling_seat,omitempty"`

	// If true, this entity can't be interacted with if the entity interacting with it is crouching
	CrouchingSkipInteract bool `json:"crouching_skip_interact,omitempty"`

	// List of entities that can ride this entity
	FamilyTypes []interface{} `json:"family_types,omitempty"`

	// The text to display when the player can interact with the entity when playing with Touch-screen controls
	InteractText string `json:"interact_text,omitempty"`

	// The max width a mob can be to be a passenger. A value of 0 ignores this parameter.
	PassengerMaxWidth float64 `json:"passenger_max_width,omitempty"`

	// This field may exist in old data but isn't used by minecraft:rideable.
	Priority int `json:"priority,omitempty"`

	// If true, this entity will pull in entities that are in the correct family_types into any available seats
	PullInEntities bool `json:"pull_in_entities,omitempty"`

	// If true, this entity will be picked when looked at by the rider
	RiderCanInteract bool `json:"rider_can_interact,omitempty"`

	// The number of entities that can ride this entity at the same time
	SeatCount int `json:"seat_count,omitempty"`

	// The list of positions and number of riders for each position for entities riding this entity
	Seats []RideableSeats `json:"seats,omitempty"`
}

// Defines the entity's size interpolation based on the entity's age.
type ScaleByAge struct {
	// Ending scale of the entity when it's fully grown.
	EndScale float64 `json:"end_scale,omitempty"`

	// Initial scale of the newborn entity.
	StartScale float64 `json:"start_scale,omitempty"`
}

// Fires off scheduled mob events at time of day events.
type Scheduler struct {
	// The list of triggers that fire when the conditions match the given filter criteria. If any filter criteria overlap the first defined event will be picked.
	ScheduledEvents []interface{} `json:"scheduled_events,omitempty"`
}

// Defines a list of items the mob wants to share or pick up. Each item must have the following parameters:
type Shareables struct {
	// A bucket for all other items in the game. Note this category is always least priority items.
	AllItems bool `json:"all_items,omitempty"`

	// Maximum number of this item the mob will hold.
	AllItemsMaxAmount int `json:"all_items_max_amount,omitempty"`

	// Number of this item considered extra that the entity wants to share.
	AllItemsSurplusAmount int `json:"all_items_surplus_amount,omitempty"`

	// Number of this item this entity wants to share.
	AllItemsWantAmount int `json:"all_items_want_amount,omitempty"`

	// List of items that the entity wants to share.
	Items []interface{} `json:"items,omitempty"`

	// Controls if the mob is able to pick up more of the same item if it is already holding that item
	SingularPickup bool `json:"singular_pickup,omitempty"`
}

// Defines the entity's ranged attack behavior. The "minecraft:behavior.ranged_attack" goal uses this component to determine which projectiles to shoot.
type Shooter struct {
	// ID of the Potion effect for the default projectile to be applied on hit.
	AuxVal int `json:"aux_val,omitempty"`

	// Actor definition to use as the default projectile for the ranged attack. The actor definition must have the projectile component to be able to be shot as a projectile.
	Def string `json:"def,omitempty"`

	// Sets whether the projectiles being used are flagged as magic. If set, the ranged attack goal will not be used at the same time as other magic goals, such as minecraft:behavior.drink_potion
	Magic bool `json:"magic,omitempty"`

	// Velocity in which the projectiles will be shot at. A power of 0 will be overwritten by the default projectile throw power.
	Power float64 `json:"power,omitempty"`

	// List of projectiles that can be used by the shooter. Projectiles are evaluated in the order of the list; After a projectile is chosen, the rest of the list is ignored.
	Projectiles []interface{} `json:"projectiles,omitempty"`

	// Sound that is played when the shooter shoots a projectile.
	Sound string `json:"sound,omitempty"`
}

// Defines the entity's 'sit' state.
type Sittable struct {
	// Event to run when the entity enters the 'sit' state
	SitEvent string `json:"sit_event,omitempty"`

	// Event to run when the entity exits the 'sit' state
	StandEvent string `json:"stand_event,omitempty"`
}

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

// Allows this entity to remember suspicious locations

type SuspectTracking struct {
}

// Defines the rules for a mob to be tamed by the player.
type Tameable struct {
	// The chance of taming the entity with each item use between 0.0 and 1.0, where 1.0 is 100%
	Probability float64 `json:"probability,omitempty"`

	// Event to run when this entity becomes tamed
	TameEvent string `json:"tame_event,omitempty"`

	// The list of items that can be used to tame this entity
	TameItems []interface{} `json:"tame_items,omitempty"`
}

// [Not a component] The list of items that, if carried while interacting with the entity, will anger it.
type TamemountAutoRejectItems struct {
	// Name of the item this entity dislikes and will cause it to get angry if used while untamed.
	Item string `json:"item,omitempty"`
}

// [Not a component] The list of items that can be used to increase the entity's temper and speed up the taming process.
type TamemountFeedItems struct {
	// Name of the item this entity likes and can be used to increase this entity's temper.
	Item string `json:"item,omitempty"`

	// The amount of temper this entity gains when fed this item.
	TemperMod int `json:"temper_mod,omitempty"`
}

// Allows the Entity to be tamed by mounting it.
type Tamemount struct {
	// The amount the entity's temper will increase when mounted.
	AttemptTemperMod int `json:"attempt_temper_mod,omitempty"`

	// The list of items that, if carried while interacting with the entity, will anger it.
	AutoRejectItems TamemountAutoRejectItems `json:"autoRejectItems,omitempty"`

	// The list of items that can be used to increase the entity's temper and speed up the taming process.
	FeedItems TamemountFeedItems `json:"feed_items,omitempty"`

	// The text that shows in the feeding interact button.
	FeedText string `json:"feed_text,omitempty"`

	// The maximum value for the entity's random starting temper.
	MaxTemper int `json:"max_temper,omitempty"`

	// The minimum value for the entity's random starting temper.
	MinTemper int `json:"min_temper,omitempty"`

	// The text that shows in the riding interact button.
	RideText string `json:"ride_text,omitempty"`

	// Event that triggers when the entity becomes tamed.
	TameEvent string `json:"tame_event,omitempty"`
}

// Defines the entity's range within which it can see or sense other entities to target them.
type TargetNearbySensor struct {
	// Maximum distance in blocks that another entity will be considered in the 'inside' range
	InsideRange float64 `json:"inside_range,omitempty"`

	// Whether the other entity needs to be visible to trigger 'inside' events
	MustSee bool `json:"must_see,omitempty"`

	// Event to call when an entity gets in the inside range. Can specify 'event' for the name of the event and 'target' for the target of the event
	OnInsideRange string `json:"on_inside_range,omitempty"`

	// Event to call when an entity gets in the outside range. Can specify 'event' for the name of the event and 'target' for the target of the event
	OnOutsideRange string `json:"on_outside_range,omitempty"`

	// Event to call when an entity exits visual range. Can specify 'event' for the name of the event and 'target' for the target of the event
	OnVisionLostInsideRange string `json:"on_vision_lost_inside_range,omitempty"`

	// Maximum distance in blocks that another entity will be considered in the 'outside' range
	OutsideRange float64 `json:"outside_range,omitempty"`
}

// Defines an entity's teleporting behavior.
type Teleport struct {
	// Modifies the chance that the entity will teleport if the entity is in darkness
	DarkTeleportChance float64 `json:"dark_teleport_chance,omitempty"`

	// Modifies the chance that the entity will teleport if the entity is in daylight
	LightTeleportChance float64 `json:"light_teleport_chance,omitempty"`

	// Maximum amount of time in seconds between random teleports
	MaxRandomTeleportTime float64 `json:"max_random_teleport_time,omitempty"`

	// Minimum amount of time in seconds between random teleports
	MinRandomTeleportTime float64 `json:"min_random_teleport_time,omitempty"`

	// [Vector3 [a,b,c]] Entity will teleport to a random position within the area defined by this cube
	RandomTeleportCube []float64 `json:"random_teleport_cube,omitempty"`

	// If true, the entity will teleport randomly
	RandomTeleports bool `json:"random_teleports,omitempty"`

	// Maximum distance the entity will teleport when chasing a target
	TargetDistance float64 `json:"target_distance,omitempty"`

	// The chance that the entity will teleport between 0.0 and 1.0. 1.0 means 100%
	TargetTeleportChance float64 `json:"target_teleport_chance,omitempty"`
}

// Defines if the entity ticks the world and the radius around it to tick.
type TickWorld struct {
	// The distance at which the closest player has to be before this entity despawns. This option will be ignored if never_despawn is true. Min: 128 blocks.
	DistanceToPlayers float64 `json:"distance_to_players,omitempty"`

	// If true, this entity will not despawn even if players are far away. If false, distance_to_players will be used to determine when to despawn.
	NeverDespawn bool `json:"never_despawn,omitempty"`

	// The area around the entity to tick. Default: 2. Allowed range: 2-6.
	Radius uint `json:"radius,omitempty"`
}

// Adds a timer after which an event will fire.
type Timer struct {
	// If true, the timer will restart every time after it fires.
	Looping bool `json:"looping,omitempty"`

	// If true, the amount of time on the timer will be random between the min and max values specified in time.
	RandomInterval bool `json:"randomInterval,omitempty"`

	// This is a list of objects, representing one value in seconds that can be picked before firing the event and an optional weight. Incompatible with time.
	RandomTimeChoices []interface{} `json:"random_time_choices,omitempty"`

	// Amount of time in seconds for the timer. Can be specified as a number or a pair of numbers (min and max). Incompatible with random_time_choices.
	Time float64 `json:"time,omitempty"`

	// Event to fire when the time on the timer runs out.
	TimeDownEvent string `json:"time_down_event,omitempty"`
}

// Defines this entity's ability to trade with players.
type TradeTable struct {
	// Determines when the mob transforms, if the trades should be converted when the new mob has a economy_trade_table. When the trades are converted, the mob will generate a new trade list with their new trade table, but then it will try to convert any of the same trades over to have the same enchantments and user data. For example, if the original has a Emerald to Enchanted Iron Sword (Sharpness 1), and the new trade also has an Emerald for Enchanted Iron Sword, then the enchantment will be Sharpness 1.
	ConvertTradesEconomy bool `json:"convert_trades_economy,omitempty"`

	// Name to be displayed while trading with this entity.
	DisplayName string `json:"display_name,omitempty"`

	// Used to determine if trading with entity opens the new trade screen.
	NewScreen bool `json:"new_screen,omitempty"`

	// Determines if the trades should persist when the mob transforms. This makes it so that the next time the mob is transformed to something with a trade_table or economy_trade_table, then it keeps their trades.
	PersistTrades bool `json:"persist_trades,omitempty"`

	// File path relative to the behavior pack root for this entity's trades.
	Table string `json:"table,omitempty"`
}

// Causes an entity to leave a trail of blocks as it moves about the world.
type Trail struct {
	// The type of block you wish to be spawned by the entity as it move about the world. Solid blocks may not be spawned at an offset of (0,0,0).
	BlockType string `json:"block_type,omitempty"`

	// One or more conditions that must be met in order to cause the chosen block type to spawn.
	SpawnFilter *f.Filter `json:"spawn_filter,omitempty"`

	// [Vector3 [a,b,c]] The distance from the entities current position to spawn the block. Capped at up to 16 blocks away. The X value is left/right(-/+), the Z value is backward/forward(-/+), the Y value is below/above(-/+).
	SpawnOffset []float64 `json:"spawn_offset,omitempty"`
}

// [Not a component] List of components to add to the entity after the transformation
type TransformationAdd struct {
	// Names of component groups to add
	ComponentGroups []interface{} `json:"component_groups,omitempty"`
}

// [Not a component] Defines the properties of the delay for the transformation
type TransformationDelay struct {
	// Chance that the entity will look for nearby blocks that can speed up the transformation. Value must be between 0.0 and 1.0
	BlockAssistChance float64 `json:"block_assist_chance,omitempty"`

	// Chance that, once a block is found, will help speed up the transformation
	BlockChance float64 `json:"block_chance,omitempty"`

	// Maximum number of blocks the entity will look for to aid in the transformation. If not defined or set to 0, it will be set to the block radius
	BlockMax int `json:"block_max,omitempty"`

	// Distance in Blocks that the entity will search for blocks that can help the transformation
	BlockRadius int `json:"block_radius,omitempty"`

	// List of blocks that can help the transformation of this entity
	BlockTypes []interface{} `json:"block_types,omitempty"`

	// Time in seconds before the entity transforms
	Value float64 `json:"value,omitempty"`
}

// Defines an entity's transformation from the current definition into another
type Transformation struct {
	// List of components to add to the entity after the transformation
	Add TransformationAdd `json:"add,omitempty"`

	// Sound to play when the transformation starts
	BeginTransformSound string `json:"begin_transform_sound,omitempty"`

	// Defines the properties of the delay for the transformation
	Delay TransformationDelay `json:"delay,omitempty"`

	// Cause the entity to drop all equipment upon transformation
	DropEquipment bool `json:"drop_equipment,omitempty"`

	// Cause the entity to drop all items in inventory upon transformation
	DropInventory bool `json:"drop_inventory,omitempty"`

	// Entity Definition that this entity will transform into
	Into string `json:"into,omitempty"`

	// If this entity has trades and has leveled up, it should maintain that level after transformation.
	KeepLevel bool `json:"keep_level,omitempty"`

	// If this entity is owned by another entity, it should remain owned after transformation.
	KeepOwner bool `json:"keep_owner,omitempty"`

	// Cause the entity to keep equipment after going through transformation
	PreserveEquipment bool `json:"preserve_equipment,omitempty"`

	// Sound to play when the entity is done transforming
	TransformationSound string `json:"transformation_sound,omitempty"`
}

// Defines the rules for a mob to trust players.
type Trusting struct {
	// The chance of the entity trusting with each item use between 0.0 and 1.0, where 1.0 is 100%.
	Probability float64 `json:"probability,omitempty"`

	// Event to run when this entity becomes trusting.
	TrustEvent string `json:"trust_event,omitempty"`

	// The list of items that can be used to get the entity to trust players.
	TrustItems []interface{} `json:"trust_items,omitempty"`
}

// Entities with this component will have a maximum auto step height that is different depending on whether they are on a block that prevents jumping. Incompatible with "runtime_identifier": "minecraft:horse".
type VariableMaxAutoStep struct {
	// The maximum auto step height when on any other block.
	BaseValue float64 `json:"base_value,omitempty"`

	// The maximum auto step height when on any other block and controlled by the player.
	ControlledValue float64 `json:"controlled_value,omitempty"`

	// The maximum auto step height when on a block that prevents jumping.
	JumpPreventedValue float64 `json:"jump_prevented_value,omitempty"`
}

// Vibrations emitted by this entity will be ignored.

type VibrationDamper struct {
}

type WaterMovement struct {
	// Drag factor to determine movement speed when in water.
	DragFactor float64 `json:"drag_factor,omitempty"`
}

// Sets the entity's delay between playing its ambient sound.
type AmbientSoundInterval struct {
	// Level sound event to be played as the ambient sound.
	EventName string `json:"event_name,omitempty"`

	// List of dynamic level sound events, with conditions for choosing between them. Evaluated in order, first one wins. If none evaluate to true, 'event_name' will take precedence.
	EventNames []interface{} `json:"event_names,omitempty"`

	// Maximum time in seconds to randomly add to the ambient sound delay time.
	Range float64 `json:"range,omitempty"`

	// Minimum time in seconds before the entity plays its ambient sound again.
	Value float64 `json:"value,omitempty"`
}

// Allows this entity to climb up ladders.

type CanClimb struct {
}

// Marks the entity as being able to fly, the pathfinder won't be restricted to paths where a solid block is required underneath it.

type CanFly struct {
}

// Allows the entity to power jump like the horse does in vanilla.

type CanPowerJump struct {
}

// Defines the entity's color. Only works on vanilla entities that have predefined color values (sheep, llama, shulker).
type Color struct {
	// The Palette Color value of the entity.
	Value int `json:"value,omitempty"`
}

// Defines the entity's second texture color. Only works on vanilla entities that have a second predefined color values (tropical fish).
type Color2 struct {
	// The second Palette Color value of the entity.
	Value int `json:"value,omitempty"`
}

// Sets this entity's default head rotation angle.
type DefaultLookAngle struct {
	// Angle in degrees.
	Value float64 `json:"value,omitempty"`
}

// Sets the Equipment table to use for this Entity.
type Equipment struct {
	// A list of slots with the chance to drop an equipped item from that slot.
	SlotDropChance []interface{} `json:"slot_drop_chance,omitempty"`

	// The file path to the equipment table, relative to the behavior pack's root.
	Table string `json:"table,omitempty"`
}

// Sets that this entity doesn't take damage from fire.

type FireImmune struct {
}

// Sets that this entity can float in liquid blocks.

type FloatsInLiquid struct {
}

// Speed in Blocks that this entity flies at.
type FlyingSpeed struct {
	// Flying speed in blocks per tick.
	Value float64 `json:"value,omitempty"`
}

// Defines how much friction affects this entity.
type FrictionModifier struct {
	// The higher the number, the more the friction affects this entity. A value of 1.0 means regular friction, while 2.0 means twice as much.
	Value float64 `json:"value,omitempty"`
}

// Sets the offset from the ground that the entity is actually at.
type GroundOffset struct {
	// The value of the entity's offset from the terrain, in blocks.
	Value float64 `json:"value,omitempty"`
}

// When configured as a rideable entity, the entity will be controlled using WASD controls. Beginning with 1.19.50 the default auto step height for rideable entities is half a block. Consider adding the `variable_max_auto_step` component to increase it.

type InputGroundControlled struct {
}

// Sets that this entity is a baby.

type IsBaby struct {
}

// Sets that this entity is charged.

type IsCharged struct {
}

// Sets that this entity is currently carrying a chest.

type IsChested struct {
}

// Allows dyes to be used on this entity to change its color.
type IsDyeable struct {
	// The text that will display when interacting with this entity with a dye when playing with Touch-screen controls.
	InteractText string `json:"interact_text,omitempty"`
}

// Sets that this entity can hide from hostile mobs while invisible.

type IsHiddenWhenInvisible struct {
}

// Sets that this entity is currently on fire.

type IsIgnited struct {
}

// Sets that this entity is an illager captain.

type IsIllagerCaptain struct {
}

// Sets that this entity is currently pregnant.

type IsPregnant struct {
}

// Sets that this entity is currently saddled.

type IsSaddled struct {
}

// Sets that this entity is currently shaking.

type IsShaking struct {
}

// Sets that this entity is currently sheared.

type IsSheared struct {
}

// Sets that this entity can be stacked.

type IsStackable struct {
}

// Sets that this entity is currently stunned.

type IsStunned struct {
}

// Sets that this entity is currently tamed.

type IsTamed struct {
}

// Defines what items can be used to control this entity while ridden.
type ItemControllable struct {
	// List of items that can be used to control this entity.
	ControlItems []interface{} `json:"control_items,omitempty"`
}

// Sets the loot table for what items this entity drops upon death.
type Loot struct {
	// The path to the loot table, relative to the Behavior Pack's root.
	Table string `json:"table,omitempty"`
}

// Additional variant value. Can be used to further differentiate variants.
type MarkVariant struct {
	// The ID of the variant. By convention, 0 is the ID of the base entity.
	Value int `json:"value,omitempty"`
}

// Sets the offset used to determine the next step distance for playing a movement sound.
type MovementSoundDistanceOffset struct {
	// The higher the number, the less often the movement sound will be played.
	Value float64 `json:"value,omitempty"`
}

// Sets the distance through which the entity can push through.
type PushThrough struct {
	// The value of the entity's push-through, in blocks.
	Value float64 `json:"value,omitempty"`
}

// Sets the entity's visual size.
type Scale struct {
	// The value of the scale. 1.0 means the entity will appear at the scale they are defined in their model. Higher numbers make the entity bigger.
	Value float64 `json:"value,omitempty"`
}

// Skin ID value. Can be used to differentiate skins, such as base skins for villagers.
type SkinId struct {
	// The ID of the skin. By convention, 0 is the ID of the base skin.
	Value int `json:"value,omitempty"`
}

// Sets the entity's base volume for sound effects.
type SoundVolume struct {
	// The value of the volume the entity uses for sound effects.
	Value float64 `json:"value,omitempty"`
}

// Defines the families this entity belongs to.
type TypeFamily struct {
	// List of family names.
	Family []interface{} `json:"family,omitempty"`
}

// Used to differentiate the component group of a variant of an entity from others. (e.g. ocelot, villager)
type Variant struct {
	// The ID of the variant. By convention, 0 is the ID of the base entity.
	Value int `json:"value,omitempty"`
}

// Sets the speed multiplier for this entity's walk animation speed.
type WalkAnimationSpeed struct {
	// The higher the number, the faster the animation for walking plays. A value of 1.0 means normal speed, while 2.0 means twice as fast.
	Value float64 `json:"value,omitempty"`
}

// Sets that this entity wants to become a jockey.

type WantsJockey struct {
}

// Only usable by the Ender Dragon. Adds a trigger to call on this entity's death.
type OnDeath struct {
	// The event to run when the conditions for this trigger are met.
	Event string `json:"event,omitempty"`

	// The list of conditions for this trigger to execute.
	Filters *f.Filter `json:"filters,omitempty"`

	// The target of the event.
	Target string `json:"target,omitempty"`
}

// Adds a trigger that will run when a nearby entity of the same type as this entity becomes Angry.
type OnFriendlyAnger struct {
	// The event to run when the conditions for this trigger are met.
	Event string `json:"event,omitempty"`

	// The list of conditions for this trigger to execute.
	Filters *f.Filter `json:"filters,omitempty"`

	// The target of the event.
	Target string `json:"target,omitempty"`
}

// Adds a trigger to call when this entity takes damage.
type OnHurt struct {
	// The event to run when the conditions for this trigger are met.
	Event string `json:"event,omitempty"`

	// The list of conditions for this trigger to execute.
	Filters *f.Filter `json:"filters,omitempty"`

	// The target of the event.
	Target string `json:"target,omitempty"`
}

// Adds a trigger to call when this entity is attacked by the player.
type OnHurtByPlayer struct {
	// The event to run when the conditions for this trigger are met.
	Event string `json:"event,omitempty"`

	// The list of conditions for this trigger to execute.
	Filters *f.Filter `json:"filters,omitempty"`

	// The target of the event.
	Target string `json:"target,omitempty"`
}

// Adds a trigger to call when this entity is set on fire.
type OnIgnite struct {
	// The event to run when the conditions for this trigger are met.
	Event string `json:"event,omitempty"`

	// The list of conditions for this trigger to execute.
	Filters *f.Filter `json:"filters,omitempty"`

	// The target of the event.
	Target string `json:"target,omitempty"`
}

// Only usable by the Ender Dragon. Adds a trigger to call when this entity lands.
type OnStartLanding struct {
	// The event to run when the conditions for this trigger are met.
	Event string `json:"event,omitempty"`

	// The list of conditions for this trigger to execute.
	Filters *f.Filter `json:"filters,omitempty"`

	// The target of the event.
	Target string `json:"target,omitempty"`
}

// Only usable by the Ender Dragon. Adds a trigger to call when this entity starts flying.
type OnStartTakeoff struct {
	// The event to run when the conditions for this trigger are met.
	Event string `json:"event,omitempty"`

	// The list of conditions for this trigger to execute.
	Filters *f.Filter `json:"filters,omitempty"`

	// The target of the event.
	Target string `json:"target,omitempty"`
}

// Adds a trigger to call when this entity finds a target.
type OnTargetAcquired struct {
	// The event to run when the conditions for this trigger are met.
	Event string `json:"event,omitempty"`

	// The list of conditions for this trigger to execute.
	Filters *f.Filter `json:"filters,omitempty"`

	// The target of the event.
	Target string `json:"target,omitempty"`
}

// Adds a trigger to call when this entity loses the target it currently has.
type OnTargetEscape struct {
	// The event to run when the conditions for this trigger are met.
	Event string `json:"event,omitempty"`

	// The list of conditions for this trigger to execute.
	Filters *f.Filter `json:"filters,omitempty"`

	// The target of the event.
	Target string `json:"target,omitempty"`
}

// Adds a trigger to call when this pet's owner awakes after sleeping with the pet.
type OnWakeWithOwner struct {
	// The event to run when the conditions for this trigger are met.
	Event string `json:"event,omitempty"`

	// The list of conditions for this trigger to execute.
	Filters *f.Filter `json:"filters,omitempty"`

	// The target of the event.
	Target string `json:"target,omitempty"`
}

// Sounds to play when the entity is getting provoked. Evaluated in order. First matching condition wins
type AngerLevelOn_IncreaseSounds struct {
	// A Molang expression describing under which conditions to play this sound, given that the entity was provoked
	Condition string `json:"condition,omitempty"`

	// The sound to play
	Sound string `json:"sound,omitempty"`
}

// List of dynamic level sound events, with conditions for choosing between them. Evaluated in order, first one wins. If none evaluate to true, 'event_name' will take precedence.
type AmbientSoundInterval_EventNames struct {
	// List of dynamic level sound events, with conditions for choosing between them. Evaluated in order, first one wins. If none evaluate to true, 'event_name' will take precedence.
	Condition string `json:"condition,omitempty"`

	// Level sound event to be played as the ambient sound
	EventName string `json:"event_name,omitempty"`
}

// List of damage conditions that when met can cause damage to the entity.
type HurtOnCondition_DamageConditions struct {
	// The kind of damage that is caused to the entity. Various armors and spells use this to determine if the entity is immune.
	Cause string `json:"cause,omitempty"`

	// The amount of damage done each tick that the conditions are met.
	DamagePerTick int `json:"damage_per_tick,omitempty"`

	// The set of conditions that must be satisfied before the entity takes the defined damage.
	Filters *f.Filter `json:"filters,omitempty"`
}

// Defines the Conditional Spatial Update Bandwidth Optimizations of this entity.
type ConditionalBandwithOptimizationConditionalValues struct {
	// Conditions that must be met for these optimization values to be used.
	Conditional_values *f.Filter `json:"Conditional_values,omitempty"`

	// In relation to the optimization value, determines the maximum ticks spatial update packets can be not sent.
	Max_dropped_ticks int `json:"Max_dropped_ticks,omitempty"`

	// The maximum distance considered during bandwidth optimizations. Any value below the max is interpolated to find optimization, and any value greater than or equal to this max results in max optimization.
	Max_optimized_distance float64 `json:"Max_optimized_distance,omitempty"`

	// When set to true, smaller motion packets will be sent during drop packet intervals, resulting in the same amount of packets being sent as without optimizations but with much less data being sent. This should be used when actors are travelling very quickly or teleporting to prevent visual oddities.
	Use_motion_prediction_hints bool `json:"Use_motion_prediction_hints,omitempty"`
}

// Defines the Conditional Spatial Update Bandwidth Optimizations of this entity.
type ConditionalBandwithOptimizationDefaultValues struct {
	// In relation to the optimization value, determines the maximum ticks spatial update packets can be not sent.
	Max_dropped_ticks int `json:"Max_dropped_ticks,omitempty"`

	// The maximum distance considered during bandwidth optimizations. Any value below the max is interpolated to find optimization, and any value greater than or equal to this max results in max optimization.
	Max_optimized_distance float64 `json:"Max_optimized_distance,omitempty"`

	// When set to true, smaller motion packets will be sent during drop packet intervals, resulting in the same amount of packets being sent as without optimizations but with much less data being sent. This should be used when actors are travelling very quickly or teleporting to prevent visual oddities.
	Use_motion_prediction_hints bool `json:"Use_motion_prediction_hints,omitempty"`
}

// Defines the health of this entity.
type Health struct {
	// The amount of health the entity has.
	Value int `json:"Value,omitempty"`

	// The maximum amount of health the entity can have.
	Max int `json:"Max,omitempty"`

	// The minimum amount of health the entity can have.
	Min int `json:"Min,omitempty"`
}
