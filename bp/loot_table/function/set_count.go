package function

// NewSetCount is the function used to set the count of the item.
//
// Example:
//
//  NewSetCount(1, 1)
func NewSetCount(min int, max int, conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "set_count"
	r["count"] = make(map[string]int)
	r["count"].(map[string]int)["min"] = min
	r["count"].(map[string]int)["max"] = max
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}
