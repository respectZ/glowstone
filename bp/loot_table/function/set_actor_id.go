package function

// NewSetActorId is the function used to set the actor id of the item.
//
// Example:
//
//  NewSetActorId("minecraft:zombie")
func NewSetActorId(Id string, conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "set_actor_id"
	r["id"] = Id
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}
