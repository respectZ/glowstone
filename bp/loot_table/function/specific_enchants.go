package function

// NewSpecificEnchants is the function used to set the name of the item.
//
// Example:
//
//  NewSpecificEnchants([]SpecificEnchants{
//  	{
//  		Id:    SpecificEnchantsSharpness,
//  		Level: [2]int{1, 5},
// 		},
// 	})
func NewSpecificEnchants(enchants []SpecificEnchants, conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "specific_enchants"
	r["enchants"] = enchants
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}

type SpecificEnchants struct {
	Id    string `json:"id,omitempty"`
	Level [2]int `json:"level,omitempty"`
}

const (
	// Enchantments
	SpecificEnchantsAquaAffinity         string = "aqua_affinity"
	SpecificEnchantsBaneOfArthropods     string = "bane_of_arthropods"
	SpecificEnchantsBindingCurse         string = "binding"
	SpecificEnchantsBlastProtection      string = "blast_protection"
	SpecificEnchantsChanneling           string = "channeling"
	SpecificEnchantsDepthStrider         string = "depth_strider"
	SpecificEnchantsEfficiency           string = "efficiency"
	SpecificEnchantsFeatherFalling       string = "feather_falling"
	SpecificEnchantsFireAspect           string = "fire_aspect"
	SpecificEnchantsFireProtection       string = "fire_protection"
	SpecificEnchantsFlame                string = "flame"
	SpecificEnchantsFortune              string = "fortune"
	SpecificEnchantsFrostWalker          string = "frost_walker"
	SpecificEnchantsImpaling             string = "impaling"
	SpecificEnchantsInfinity             string = "infinity"
	SpecificEnchantsKnockback            string = "knockback"
	SpecificEnchantsLooting              string = "looting"
	SpecificEnchantsLoyalty              string = "loyalty"
	SpecificEnchantsLuckOfTheSea         string = "luck_of_the_sea"
	SpecificEnchantsLure                 string = "lure"
	SpecificEnchantsMending              string = "mending"
	SpecificEnchantsMultishot            string = "multishot"
	SpecificEnchantsPiercing             string = "piercing"
	SpecificEnchantsPower                string = "power"
	SpecificEnchantsProjectileProtection string = "projectile_protection"
	SpecificEnchantsProtection           string = "protection"
	SpecificEnchantsPunch                string = "punch"
	SpecificEnchantsQuickCharge          string = "quick_charge"
	SpecificEnchantsRespiration          string = "respiration"
	SpecificEnchantsRiptide              string = "riptide"
	SpecificEnchantsSharpness            string = "sharpness"
	SpecificEnchantsSilkTouch            string = "silk_touch"
	SpecificEnchantsSmite                string = "smite"
	SpecificEnchantsSoulSpeed            string = "soul_speed"
	SpecificEnchantsSwiftSneak           string = "swift_sneak"
	SpecificEnchantsThorns               string = "thorns"
	SpecificEnchantsUnbreaking           string = "unbreaking"
	SpecificEnchantsVanishingCurse       string = "vanishing"
)
