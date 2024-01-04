package entity

import (
	g_util "github.com/respectZ/glowstone/util"
)

type Bone struct {
	Name          string         `json:"name"`
	Pivot         []float64      `json:"pivot"`
	Cubes         ICubes         `json:"cubes"`
	Binding       string         `json:"binding,omitempty"`
	Debug         bool           `json:"debug,omitempty"`
	Inflate       float64        `json:"inflate,omitempty"`
	Locators      ILocators      `json:"locators,omitempty"`
	Mirror        bool           `json:"mirror,omitempty"`
	Parent        string         `json:"parent,omitempty"`
	PolyMesh      interface{}    `json:"poly_mesh,omitempty"` // todo: figure out what this is
	RenderGroupId int            `json:"render_group_id,omitempty"`
	Rotation      []*float64     `json:"rotation,omitempty"`
	TextureMeshes []*TextureMesh `json:"texture_meshes,omitempty"`
}

type bone_parse struct {
	Name          string         `json:"name"`
	Pivot         []float64      `json:"pivot"`
	Cubes         *Cubes         `json:"cubes"`
	Binding       string         `json:"binding,omitempty"`
	Debug         bool           `json:"debug,omitempty"`
	Inflate       float64        `json:"inflate,omitempty"`
	Locators      *Locators      `json:"locators,omitempty"`
	Mirror        bool           `json:"mirror,omitempty"`
	Parent        string         `json:"parent,omitempty"`
	PolyMesh      interface{}    `json:"poly_mesh,omitempty"` // todo: figure out what this is
	RenderGroupId int            `json:"render_group_id,omitempty"`
	Rotation      []*float64     `json:"rotation,omitempty"`
	TextureMeshes []*TextureMesh `json:"texture_meshes,omitempty"`
}

type Bones []*Bone

type IBones interface {
	Add(...*Bone)
	All() []*Bone
	Clear()
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	New() *Bone

	// Get the bone with the given name.
	//
	// name: The name of the bone.
	//
	// return: The bone with the given name.
	//
	// Example:
	//
	//  bone := bones.FindByName("body")
	FindByName(string) *Bone

	// Get the bone with the given parent name.
	//
	// parent: The name of the parent bone.
	//
	// return: The bone with the given parent name.
	//
	// Example:
	//
	//  bone := bones.FindByParent("body")
	FindByParent(string) *Bone
}

func (b *Bones) Add(s ...*Bone) {
	*b = append(*b, s...)
}

func (b *Bones) All() []*Bone {
	return *b
}

func (b *Bones) Clear() {
	*b = make(Bones, 0)
}

func (b *Bones) IsEmpty() bool {
	return len(*b) == 0
}

func (b *Bones) Size() int {
	return len(*b)
}

func (b *Bones) UnmarshalJSON(data []byte) error {
	var temp []*bone_parse
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*b = make(Bones, 0)
	for _, bone := range temp {
		*b = append(*b, &Bone{
			Name:          bone.Name,
			Pivot:         bone.Pivot,
			Cubes:         bone.Cubes,
			Binding:       bone.Binding,
			Debug:         bone.Debug,
			Inflate:       bone.Inflate,
			Locators:      bone.Locators,
			Mirror:        bone.Mirror,
			Parent:        bone.Parent,
			PolyMesh:      bone.PolyMesh,
			RenderGroupId: bone.RenderGroupId,
			Rotation:      bone.Rotation,
			TextureMeshes: bone.TextureMeshes,
		})
	}
	return nil
}

func (b *Bones) New() *Bone {
	bone := &Bone{}
	*b = append(*b, bone)
	return bone
}

func (b *Bones) FindByName(name string) *Bone {
	for _, bone := range *b {
		if bone.Name == name {
			return bone
		}
	}
	return nil
}

func (b *Bones) FindByParent(parent string) *Bone {
	for _, bone := range *b {
		if bone.Parent == parent {
			return bone
		}
	}
	return nil
}
