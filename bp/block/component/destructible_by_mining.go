package component

// Code generated by scrapper/block.ts;

// Describes the destructible by mining properties for this block. If set to true, the block will take the default number of seconds to destroy. If set to false, this block is indestructible by mining. If the component is omitted, the block will take the default number of seconds to destroy.
type DestructibleByMining struct {
	// Sets the number of seconds it takes to destroy the block with base equipment. Greater numbers result in greater mining times.
	SecondsToDestroy float64 `json:"seconds_to_destroy,omitempty"`
}
