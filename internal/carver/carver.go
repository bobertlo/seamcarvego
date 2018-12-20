package carver

import (
	"image"
)

const MaxEnergy = 1000.0

type Carver interface {
	Img() image.Image
	Height() int
	Width() int
	Energy(int, int) (int, error)
	HSeam() ([]int, error)
	VSeam() ([]int, error)
	HRemoveSeam([]int) error
	VRemoveSeam([]int) error
}

func New(img image.Image) (Carver, error) {
	return NewArrayCarver(img)
}
