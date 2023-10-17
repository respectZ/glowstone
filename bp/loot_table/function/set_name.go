package function

// NewName is the function used to set the name of the item.
//
// Example:
//
//  NewName("Hello World")
func NewName(name string, conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "set_name"
	r["name"] = name
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}
