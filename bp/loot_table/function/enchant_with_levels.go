package function

func NewEnchantWithLevels(treasure bool, maxLevel int, minLevel int, conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "enchant_with_levels"
	r["treasure"] = treasure
	r["level"] = make(map[string]int)
	r["level"].(map[string]int)["max"] = maxLevel
	r["level"].(map[string]int)["min"] = minLevel
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}
