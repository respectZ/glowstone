package function

// This function was only designed to be used on books, rolling for a single enchantment across all possible non-curse enchantments, including treasure enchantments. The function doesn't adapt to the current item. If used on a book, an enchantment will always successfully be applied; if used on something else enchantable, it's possible the item won't be successfully enchanted.
func NewEnchantBookTrading(base_cost int, base_random_cost int, per_level_cost int, per_level_random_cost int) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "enchant_book_for_trading"
	r["base_cost"] = base_cost
	r["base_random_cost"] = base_random_cost
	r["per_level_cost"] = per_level_cost
	r["per_level_random_cost"] = per_level_random_cost
	return &r
}
