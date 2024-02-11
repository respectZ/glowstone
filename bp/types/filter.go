package types

type Filter struct {
	AllOf    []Filter `json:"all_of,omitempty"` // AND
	AnyOf    []Filter `json:"any_of,omitempty"` // OR
	Test     string   `json:"test,omitempty"`
	Subject  string   `json:"subject,omitempty"`
	Domain   string   `json:"domain,omitempty"`
	Operator string   `json:"operator,omitempty"`

	// Check other family
	OtherWithFamilies string `json:"other_with_families,omitempty"`

	Value interface{} `json:"value"`
}
