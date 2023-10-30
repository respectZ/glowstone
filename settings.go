package glowstone

type Settings struct {
	MinifyJSON    bool
	FormatVersion *SettingsFormatVersion
}

type SettingsFormatVersion struct {
	BlockBP  string
	ItemBP   string
	EntityBP string

	// RP

	EntityRP     string
	BlocksRP     [3]int
	AttachableRP string
}
