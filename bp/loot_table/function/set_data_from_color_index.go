package function

// NewSetDataFromColorIndex is the function used to set the data of the item from the color index.
//
// Example:
//
//  NewSetDataFromColorIndex()
func NewSetDataFromColorIndex(conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "set_data_from_color_index"
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}
