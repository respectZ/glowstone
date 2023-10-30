package function

// NewSetData is the function used to set the data of the item.
//
// Example:
//
//  NewSetData(1, 1000)
func NewSetData(min int, max int, conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "set_data"
	r["data"] = make(map[string]int)
	r["data"].(map[string]int)["min"] = min
	r["data"].(map[string]int)["max"] = max
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}
