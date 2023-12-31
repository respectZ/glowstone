package component

// Code generated by scrapper/item/main.ts;

// Projectile item component. Projectile items shoot out, like an arrow. In format versions prior to 1.20.10, this component requires the 'Holiday Creator Features' experimental toggle.
type Projectile struct {
  // Defines the time a projectile needs to charge in order to critically hit
  MinimumCriticalPower float64 `json:"minimum_critical_power,omitempty"`

  // The entity to be fired as a projectile
  ProjectileEntity interface{} `json:"projectile_entity,omitempty"`

}


