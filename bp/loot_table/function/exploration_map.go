package function

// NewExplorationMap is the constructor for a new ExplorationMap Function.
//
// Available destinations:
//
//  - "ancientcity"
//  - "buriedtreasure"
//  - "endcity"
//  - "fortress"
//  - "mansion"
//  - "mineshaft"
//  - "monument"
//  - "pillageroutpost"
//  - "ruins"
//  - "shipwreck"
//  - "stronghold"
//  - "temple"
//  - "village"
func NewExplorationMap(destination string, conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "exploration_map"
	r["destination"] = destination
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}

const (
	ExplorationMapAncientCity     string = "ancientcity"
	ExplorationMapBuriedTreasure  string = "buriedtreasure"
	ExplorationMapEndCity         string = "endcity"
	ExplorationMapFortress        string = "fortress"
	ExplorationMapMansion         string = "mansion"
	ExplorationMapMineshaft       string = "mineshaft"
	ExplorationMapMonument        string = "monument"
	ExplorationMapPillagerOutpost string = "pillageroutpost"
	ExplorationMapRuins           string = "ruins"
	ExplorationMapShipwreck       string = "shipwreck"
	ExplorationMapStronghold      string = "stronghold"
	ExplorationMapTemple          string = "temple"
	ExplorationMapVillage         string = "village"
)
