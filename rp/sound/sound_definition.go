package sound

import (
	"path/filepath"
	"strings"

	g_util "github.com/respectZ/glowstone/util"
)

const (
	FORMAT_VERSION string = "1.20.20"
)

const (
	SOUND_CATEGORY_AMBIENT string = "ambient"
	SOUND_CATEGORY_BLOCK   string = "block"
	SOUND_CATEGORY_BUCKET  string = "bucket"
	SOUND_CATEGORY_HOSTILE string = "hostile"
	SOUND_CATEGORY_MUSIC   string = "music"
	SOUND_CATEGORY_NEUTRAL string = "neutral"
	SOUND_CATEGORY_PLAYER  string = "player"
	SOUND_CATEGORY_RECORD  string = "record"
	SOUND_CATEGORY_UI      string = "ui"
	SOUND_CATEGORY_WEATHER string = "weather"
)

type SoundDefinition struct {
	FormatVersion    string                 `json:"format_version"`
	SoundDefinitions map[string]*Definition `json:"sound_definitions"`
}

type Definition struct {
	UseLegacyMaxDistance bool     `json:"use_legacy_max_distance,omitempty"`
	Category             string   `json:"category"`
	MaxDistance          *float64 `json:"max_distance,omitempty"`
	MinDistance          *float64 `json:"min_distance,omitempty"`

	// Can be array of string or array of sound
	Sounds []interface{} `json:"sounds"`
}

type Sound struct {
	// When set to true, the sound will be loaded when the game is running low on memory
	LoadOnLowMemory bool `json:"load_on_low_memory,omitempty"`

	// When set to true, the sound will be streamed from disk
	Stream bool `json:"stream,omitempty"`

	// When set to true, the sound will be 3D
	Is3D bool `json:"is_3d,omitempty"`

	// Path to the sound file
	Name string `json:"name"`

	// Volume of the sound
	Volume float64 `json:"volume,omitempty"`

	// Pitch of the sound
	Pitch float64 `json:"pitch,omitempty"`

	// Chance of the sound being played
	Weight int `json:"weight,omitempty"`
}

func New() *SoundDefinition {
	return &SoundDefinition{
		FormatVersion:    FORMAT_VERSION,
		SoundDefinitions: make(map[string]*Definition),
	}
}

// Loads a sound_definitions.json file
func Load(src string) (*SoundDefinition, error) {
	var sound SoundDefinition
	err := g_util.LoadJSON(src, &sound)
	if err != nil {
		return nil, err
	}
	return &sound, nil
}

func (s *SoundDefinition) Encode() ([]byte, error) {
	return g_util.MarshalJSON(s)
}

// Add a sound to the sound_definitions.json file
//
// Example:
//
//		definition, sound := sound_definition.AddSound("mob.skeleton.say", "sounds/mob/skeleton/say1")
//	 definition.Category = sound_definition.SOUND_CATEGORY_HOSTILE
//	 // Add more sound to the 'mob.skeleton.say' definition
//	 definition.AddSound("mob.skeleton.say2", nil, nil)
func (s *SoundDefinition) AddSound(name string, path string) (*Definition, *Sound) {
	if s.SoundDefinitions == nil {
		s.SoundDefinitions = make(map[string]*Definition)
	}
	r := &Definition{
		Sounds: make([]interface{}, 0),
	}

	// Replace double backslashes with single backslashes
	path = strings.Replace(path, "\\\\", "\\", -1)

	sound := &Sound{
		Name:   path,
		Volume: 1.0,
		Pitch:  1.0,
	}
	r.Sounds = append(r.Sounds, sound)
	s.SoundDefinitions[name] = r
	return r, sound
}

func (s *SoundDefinition) RemoveSound(name string) {
	delete(s.SoundDefinitions, name)
}

func (s *SoundDefinition) IsEmpty() bool {
	return len(s.SoundDefinitions) == 0
}

func (s *SoundDefinition) Save(pathToRP string) error {
	// Prevent overwriting
	d, err := Load(filepath.Join(pathToRP, "sounds", "sound_definitions.json"))
	if err == nil {
		for k, v := range s.SoundDefinitions {
			d.SoundDefinitions[k] = v
		}
		s = d
	}
	data, err := s.Encode()
	if err != nil {
		return err
	}
	return g_util.Writefile(filepath.Join(pathToRP, "sounds", "sound_definitions.json"), data)
}

// Add a sound to the definition.
//
// Example:
//
//		definition, sound := sound_definition.AddSound("mob.skeleton.say", "sounds/mob/skeleton/say1")
//	 definition.Category = sound_definition.SOUND_CATEGORY_HOSTILE
//	 // Add more sound to the 'mob.skeleton.say' definition
//	 definition.AddSound("mob.skeleton.say2", nil, nil)
func (d *Definition) AddSound(path string, volume *float64, pitch *float64) *Sound {
	r := &Sound{
		Name:   path,
		Volume: 1.0,
		Pitch:  1.0,
	}
	if volume != nil {
		r.Volume = *volume
	}
	if pitch != nil {
		r.Pitch = *pitch
	}
	d.Sounds = append(d.Sounds, r)
	return r
}
