package carver

import (
	"image"
)

// MaxEnergy is the default energy value used for borders (which would
// be otherwise undefined)
const MaxEnergy = 1000.0

// Carver is an interface for seam carvers
type Carver interface {
	Img() image.Image
	Height() int
	Width() int
	Energy(int, int) (float64, error)
	HSeam() ([]int, error)
	VSeam() ([]int, error)
	HRemoveSeam([]int) error
	VRemoveSeam([]int) error
}

// New creates a new seam carver (currently, using the ArrayCarver
// implementation.)
func New(img image.Image) (Carver, error) {
	return NewArrayCarver(img)
}
