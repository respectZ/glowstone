package function

// NewSetLore is the function used to set the data of the item from the color index.
//
// Example:
//
//  NewSetLore([]string{"Hello", "World"})
func NewSetLore(lore []string, conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "set_lore"
	r["lore"] = lore
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}
