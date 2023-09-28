package animation

type Animation interface {
	// Encode
	Encode() ([]byte, error)
	// Create a new animation
	NewAnimation(string) *animationData
	// Get the animation data
	GetAnimation(string) *animationData
	// Remove animation data
	RemoveAnimation(string)
}
