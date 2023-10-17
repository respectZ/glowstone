package function

// NewSetDamage is the function used to set the damage of the item.
//
// Example:
//
//  NewSetDamage(1, 1)
func NewSetDamage(min int, max int, conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "set_damage"
	r["damage"] = make(map[string]int)
	r["damage"].(map[string]int)["min"] = min
	r["damage"].(map[string]int)["max"] = max
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}
