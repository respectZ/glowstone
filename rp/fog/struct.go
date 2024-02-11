package fog

import (
	"fmt"
	"reflect"

	g_util "github.com/respectZ/glowstone/util"
)

type Fog struct {
	FormatVersion string       `json:"format_version"`
	FogSettings   *FogSettings `json:"minecraft:fog_settings"`
}

type FogSettings struct {
	Description *Description `json:"description"`
	Volumetric  *Volumetric  `json:"volumetric,omitempty"` // The volumetric fog settings.
	Distance    *Distance    `json:"distance,omitempty"`   // The distance fog settings for different camera locations.
}

type Description struct {
	// The identifier for these fog settings. The identifier must include a namespace.
	Identifier string `json:"identifier"`
}

type Volumetric struct {
	// The density settings for different camera locations.
	Density *Density `json:"density,omitempty"`
	// The coefficient settings for the volumetric fog in different blocks.
	MediaCoefficients *MediaCoefficients `json:"media_coefficients,omitempty"`
}

type Density struct {
	// Fog density values as light passes through air blocks.
	Air *DensityData `json:"air,omitempty"`
	// Fog density values as light passes through lava blocks.
	Lava *DensityData `json:"lava,omitempty"`
	// Fog density values as light passes through lava blocks while the player has lava resistance.
	LavaResistance *DensityData `json:"lava_resistance,omitempty"`
	PowderSnow     *DensityData `json:"powder_snow,omitempty"`
	// Fog density values as light passes through water blocks.
	Water   *DensityData `json:"water,omitempty"`
	Weather *DensityData `json:"weather,omitempty"`
}

type DensityData struct {
	// The maximum amount of opaqueness that the ground fog will take on. A value from [0.0, 1.0].
	MaxDensity float64 `json:"max_density,omitempty"`
	// The height in blocks that the ground fog will become it's maximum density.
	MaxDensityHeight float64 `json:"max_density_height,omitempty"`
	// When set to true, the density will be uniform across all heights.
	Uniform bool `json:"uniform,omitempty"`
	// The height in blocks that the ground fog will be completely transparent and begin to appear. This value needs to be at least 1 higher than 'max_density_height'.
	ZeroDensityHeight float64 `json:"zero_density_height,omitempty"`
}

type MediaCoefficients struct {
	// Fog coefficient values while light passes through air.
	Air *MediaCoefficientData `json:"air,omitempty"`
	// Fog coefficient values while light passes through water.
	Water *MediaCoefficientData `json:"water,omitempty"`
	// Fog coefficient values while light passes through clouds.
	Cloud *MediaCoefficientData `json:"cloud,omitempty"`
}

type MediaCoefficientData struct {
	// Proportion of light that is absorbed (lost) per block.
	Absorption string `json:"absorption,omitempty"`
	// Proportion of light that is scattered per block.
	Scattering string `json:"scattering,omitempty"`
}

type Distance struct {
	// The fog settings when the camera is in the air.
	Air *FogData `json:"air,omitempty"`
	// The fog settings when the camera is in lava.
	Lava *FogData `json:"lava,omitempty"`
	// The fog settings when the camera is in lava and the player has the lava resistance effect active.
	LavaResistance *FogData `json:"lava_resistance,omitempty"`
	// The fog settings when the camera is inside a Powder Snow block.
	PowderSnow *FogData `json:"powder_snow,omitempty"`
	// The fog settings when the camera is in water.
	Water *FogData `json:"water,omitempty"`
	// The fog settings for when the camera is in the air with active weather (rain, snow, etc..).
	Weather *FogData `json:"weather,omitempty"`
}

type FogData struct {
	// The color that the fog will take on.
	FogColor string `json:"fog_color,omitempty"`
	// The distance from the player that the fog will begin to appear. 'fog_start' must be less than or equal to 'fog_end'.
	FogStart float64 `json:"fog_start,omitempty"`
	// The distance from the player that the fog will become fully opaque. 'fog_end' must be greater than or equal to 'fog_start'.
	FogEnd float64 `json:"fog_end,omitempty"`
	// Determines how distance value is used. Fixed distance is measured in blocks. Dynamic distance is multiplied by the current render distance.
	RenderDistanceType string `json:"render_distance_type,omitempty"`
	// Additional fog data which will slowly transition to the distance fog of current biome.
	TransitionFog *TransitionFog `json:"transition_fog,omitempty"`
}

type TransitionFog struct {
	//  Initial fog that will slowly transition into water distance fog of the biome when player goes into water.
	InitFog *FogData `json:"init_fog,omitempty"`
	// Total amount of time takes to complete fog transition.
	MaxSeconds float64 `json:"max_seconds,omitempty"`
	// The progress of fog transition after 'mid_seconds' seconds.
	//
	// Allowed values: 0.0 - 1.0
	MidPercent float64 `json:"mid_percent,omitempty"`
	// The time takes to reach certain progress('mid_percent') of fog transition.
	MidSeconds float64 `json:"mid_seconds,omitempty"`
	// The minimum progress of fog transition.
	//
	// Allowed values: 0.0 - 1.0
	MinPercent float64 `json:"min_percent,omitempty"`
}

var FORMAT_VERSION = "1.19.70"

func New(identifier string) *Fog {
	return &Fog{
		FormatVersion: FORMAT_VERSION,
		FogSettings: &FogSettings{
			Description: &Description{
				Identifier: identifier,
			},
			Distance: &Distance{
				Air: &FogData{
					TransitionFog: &TransitionFog{
						InitFog: &FogData{},
					},
				},
				Lava: &FogData{
					TransitionFog: &TransitionFog{
						InitFog: &FogData{},
					},
				},
				LavaResistance: &FogData{
					TransitionFog: &TransitionFog{
						InitFog: &FogData{},
					},
				},
				PowderSnow: &FogData{
					TransitionFog: &TransitionFog{
						InitFog: &FogData{},
					},
				},
				Water: &FogData{
					TransitionFog: &TransitionFog{
						InitFog: &FogData{},
					},
				},
				Weather: &FogData{
					TransitionFog: &TransitionFog{
						InitFog: &FogData{},
					},
				},
			},
			Volumetric: &Volumetric{
				Density: &Density{
					Air:            &DensityData{},
					Lava:           &DensityData{},
					LavaResistance: &DensityData{},
					PowderSnow:     &DensityData{},
					Water:          &DensityData{},
					Weather:        &DensityData{},
				},
				MediaCoefficients: &MediaCoefficients{
					Air:   &MediaCoefficientData{},
					Water: &MediaCoefficientData{},
					Cloud: &MediaCoefficientData{},
				},
			},
		},
	}
}

func NewSimpleAir(identifier string, r uint8, g uint8, b uint8, start float64, end float64) *Fog {
	// Conver rgb to hex
	color := fmt.Sprintf("#%02x%02x%02x", r, g, b)
	return &Fog{
		FormatVersion: FORMAT_VERSION,
		FogSettings: &FogSettings{
			Description: &Description{
				Identifier: identifier,
			},
			Distance: &Distance{
				Air: &FogData{
					FogColor:           color,
					FogStart:           start,
					FogEnd:             end,
					RenderDistanceType: "fixed",
				},
			},
		},
	}
}

func Load(src string) (*Fog, error) {
	e := &Fog{}
	err := g_util.LoadJSON(src, e)
	return e, err
}

func isZero(val interface{}) bool {
	rv := reflect.ValueOf(val)

	switch rv.Kind() {
	case reflect.Ptr:
		return rv.IsNil() || isZero(rv.Elem().Interface())
	case reflect.Struct:
		numFields := rv.NumField()
		for i := 0; i < numFields; i++ {
			if !isZero(rv.Field(i).Interface()) {
				return false
			}
		}
		return true
	default:
		zero := reflect.Zero(rv.Type())
		return rv.Interface() == zero.Interface()
	}
}

func setToNil(val interface{}) {
	rv := reflect.ValueOf(val).Elem()

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			if isZero(field.Interface()) {
				field.Set(reflect.Zero(field.Type()))
			} else {
				setToNil(field.Interface())
			}
		} else if field.Kind() == reflect.Struct {
			setToNil(field.Addr().Interface())
			if isZero(field.Interface()) {
				field.Set(reflect.Zero(field.Type()))
			}
		}
	}
}

func (f *Fog) Encode() ([]byte, error) {
	setToNil(f)
	return g_util.MarshalJSON(f)
}

func (f *Fog) GetIdentifier() string {
	return f.FogSettings.Description.Identifier
}
