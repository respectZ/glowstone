package animation_controller

type AnimationController interface {
	// Encode
	Encode() ([]byte, error)

	// Create a new animation controller
	//
	// Example:
	// 	animation_controller := glowstone.NewAnimationController("controller.animation.player")
	New(name string) AnimationControllerData
	// Get the animation controller data
	Get(string) AnimationControllerData
	// Remove animation controller data
	Remove(string)
}

type AnimationControllerData interface {
	// Get state
	GetState(string) *animationControllerState
	// Set state
	SetState(string, *animationControllerState)
	// Set initial state
	SetInitialState(string)
	// New state
	NewState(string) *animationControllerState
}
