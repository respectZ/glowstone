package component

// Code generated by scrapper/entityBp.ts;

// Fires off a specified event when a block in the block list is broken within the sensor range.  
type BlockSensor struct {
  // List of blocks to watch for being broken to fire off a specified event. If a block is in multiple lists, multiple events will fire.
  OnBreak []interface{} `json:"on_break,omitempty"`
  // The maximum radial distance in which a specified block can be detected. The biggest radius is 32.0.
  SensorRadius float64 `json:"sensor_radius,omitempty"`
  // List of sources that break the block to listen for. If none are specified, all block breaks will be detected.
  Sources []interface{} `json:"sources,omitempty"`
}