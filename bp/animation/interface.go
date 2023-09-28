package animation

type Animation interface {
	// Encode
	Encode() ([]byte, error)
	// Create a new animation
	//
	// Example:
	// 	animation := glowstone.NewAnimation("animation.player.hold")
	New(string) *animationData
	// Get the animation data
	Get(string) *animationData
	// Remove animation data
	Remove(string)
}
