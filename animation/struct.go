package animation

import (
	"github.com/respectZ/glowstone/bp/animation"
)

type BPAnimation struct {
	Animation animation.Animation

	// Other stuff
	Dest string
}

// New creates a new BPAnimation. Dest doesn't need to start with "animations/"
//
// Example:
//
//	a := NewBP("player.animation.json")
//	b := NewBP("player/effect.animation.json")
func NewBP(dest string) *BPAnimation {
	return &BPAnimation{
		Animation: animation.New(),
		Dest:      dest,
	}
}

// LoadBP loads the BP from the given path. This is not setting the Dest.
//
// Example:
//
//	a, err := animation.Load("player.animation.json")
func LoadBP(src string) (*BPAnimation, error) {
	a, err := animation.Load(src)
	if err != nil {
		return nil, err
	}
	return &BPAnimation{
		Animation: a,
	}, nil
}

// Encode returns []byte of the BP.
//
// Example:
//
//	bp, err := e.Encode()
func (a *BPAnimation) Encode() ([]byte, error) {
	return a.Animation.Encode()
}
