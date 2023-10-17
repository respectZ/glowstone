package function

// NewSetBookContents is the function used to set the book contents of the item.
//
// Example:
//
//  NewSetBookContents("Guidebook", "Mojang", []string{"This is a guidebook.", "It's a book that guides you."})
func NewSetBookContents(title string, author string, pages []string, conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "set_book_contents"
	r["title"] = title
	r["author"] = author
	r["pages"] = pages
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}
