package entity

import g_util "github.com/respectZ/glowstone/util"

type MinecraftGeometry struct {
	Description *Description `json:"description"`
	Bones       IBones       `json:"bones,omitempty"`
	Cape        string       `json:"cape,omitempty"`
}

type MinecraftGeometries []*MinecraftGeometry

type IMinecraftGeometry interface {
	UnmarshalJSON([]byte) error

	Add(...*MinecraftGeometry)
	All() []*MinecraftGeometry
	Clear()
	IsEmpty() bool
	Size() int

	New(string) *MinecraftGeometry

	FindByIdentifier(string) *MinecraftGeometry
}

type minecraftGeometry_parse struct {
	Description *Description `json:"description"`
	Bones       *Bones       `json:"bones,omitempty"`
	Cape        string       `json:"cape,omitempty"`
}

func (m *MinecraftGeometries) UnmarshalJSON(data []byte) error {
	var temp []*minecraftGeometry_parse
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	(*m) = make(MinecraftGeometries, 0)
	for _, geometry := range temp {
		*m = append(*m, &MinecraftGeometry{
			Description: geometry.Description,
			Bones:       geometry.Bones,
			Cape:        geometry.Cape,
		})
	}
	return nil
}

func (m *MinecraftGeometries) Add(s ...*MinecraftGeometry) {
	*m = append(*m, s...)
}

func (m *MinecraftGeometries) All() []*MinecraftGeometry {
	return *m
}

func (m *MinecraftGeometries) Clear() {
	*m = make(MinecraftGeometries, 0)
}

func (m *MinecraftGeometries) IsEmpty() bool {
	return len(*m) == 0
}

func (m *MinecraftGeometries) Size() int {
	return len(*m)
}

func (m *MinecraftGeometries) New(identifier string) *MinecraftGeometry {
	return &MinecraftGeometry{
		Description: &Description{
			Identifier: identifier,
		},
	}
}

func (m *MinecraftGeometries) FindByIdentifier(identifier string) *MinecraftGeometry {
	for _, geometry := range *m {
		if geometry.Description.Identifier == identifier {
			return geometry
		}
	}
	return nil
}
