package component

// Code generated by scrapper/entityBp.ts;

// Causes the mob to ignore attackable targets for a given duration.  
type AdmireItem struct {
  // Duration, in seconds, for which mob won't admire items if it was hurt
  CooldownAfterBeingAttacked int `json:"cooldown_after_being_attacked,omitempty"`
  // Duration, in seconds, that the mob is pacified.
  Duration *int `json:"duration,omitempty"`
}
