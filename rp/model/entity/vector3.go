package entity

import g_util "github.com/respectZ/glowstone/util"

type Vector3 []float64

type IVector3 interface {
	UnmarshalJSON([]byte) error

	New(float64, float64, float64) *Vector3
	X() float64
	Y() float64
	Z() float64
}

func (v *Vector3) UnmarshalJSON(data []byte) error {
	var temp []float64
	if err := g_util.UnmarshalJSON(data, &temp); err != nil {
		return err
	}
	*v = temp
	return nil
}

func (v *Vector3) New(x, y, z float64) *Vector3 {
	vector3 := &Vector3{x, y, z}
	*v = *vector3
	return vector3
}

func (v *Vector3) X() float64 {
	return (*v)[0]
}

func (v *Vector3) Y() float64 {
	return (*v)[1]
}

func (v *Vector3) Z() float64 {
	return (*v)[2]
}
