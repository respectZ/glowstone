package component

// Code generated by scrapper/block.ts;

// Makes your block into a custom crafting table which enables the crafting table UI and the ability to craft recipes. Experimental toggles required: Holiday Creator Features (in format versions before 1.19.50).
type CraftingTable struct {
	// Defines the tags recipes should define to be crafted on this table. Limited to 64 tags. Each tag is limited to 64 characters.
	CraftingTags []string `json:"crafting_tags,omitempty"`

	// Specifies the language file key that maps to what text will be displayed in the UI of this table. If the string given can not be resolved as a loc string, the raw string given will be displayed. If this field is omitted, the name displayed will default to the name specified in the "display_name" component. If this block has no "display_name" component, the name displayed will default to the name of the block.
	TableName string `json:"table_name,omitempty"`
}