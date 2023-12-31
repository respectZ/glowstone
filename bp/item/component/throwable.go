package component

// Code generated by scrapper/item/main.ts;

// Throwable item component. Throwable items, such as a snowball. In format versions prior to 1.20.10, this component requires the 'Holiday Creator Features' experimental toggle.
type Throwable struct {
  // Whether the item should use the swing animation when thrown.
  DoSwingAnimation bool `json:"do_swing_animation,omitempty"`

  // The scale at which the power of the throw increases
  LaunchPowerScale float64 `json:"launch_power_scale,omitempty"`

  // The maximum duration to draw a throwable item.
  MaxDrawDuration float64 `json:"max_draw_duration,omitempty"`

  // The maximum power to launch the throwable item.
  MaxLaunchPower float64 `json:"max_launch_power,omitempty"`

  // The minimum duration to draw a throwable item.
  MinDrawDuration float64 `json:"min_draw_duration,omitempty"`

  // Whether or not the power of the throw increases with duration charged. When true, The longer you hold, the more power it will have when released.
  ScalePowerByDrawDuration bool `json:"scale_power_by_draw_duration,omitempty"`

}


