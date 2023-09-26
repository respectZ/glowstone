package component

// Code generated by scrapper/item/main.ts;

// The icon item component determines the icon to represent the item in the UI and elsewhere. In format versions prior to 1.20.0, this component requires the 'Holiday Creator Features' experimental toggle.
type Record struct {
  // Signal strength for comparator blocks to use from 1 - 13.
  ComparatorSignal int `json:"comparator_signal,omitempty"`

  // Duration of sound event in seconds float value.
  Duration float64 `json:"duration,omitempty"`

  // Sound event types:  13, cat, blocks, chirp, far, mall, mellohi, stal, strad, ward, 11, wait, pigstep, otherside, 5, relic. The value may be one listed below.
  SoundEvent string `json:"sound_event,omitempty"`

}

