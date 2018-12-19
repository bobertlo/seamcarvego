package carver

import (
	"image"
)

const MaxEnergy = 1000.0

type Carver interface {
	Img() *image.Image
	Height() int
	Width() int
	HSeam() ([]int, error)
	VSeam() ([]int, error)
	HRemoveSeam([]int) error
	VRemoveSeam([]int) error
	Energy(int, int) (int, error)
}

func New(file string) (*Carver, error) {
	return nil, nil
}
