package carver

import (
	"image"
)

const MaxEnergy = 1000.0

type Carver struct{
	img image.Image
}

func New(img image.Image) *Carver {
	c := &Carver {
		img: img,
	}
	return c
}
