package types

type Filter struct {
	AllOf    []Filter    `json:"all_of,omitempty"` // AND
	AnyOf    []Filter    `json:"any_of,omitempty"` // OR
	Test     string      `json:"test"`
	Subject  string      `json:"subject,omitempty"`
	Domain   string      `json:"domain,omitempty"`
	Operator string      `json:"operator,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}
