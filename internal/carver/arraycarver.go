package carver

import (
	"image"
)

type ArrayCarver struct {
}

func NewArrayCarver(file string) (*ArrayCarver, error) {
	return nil, nil
}

func (*ArrayCarver) Img() *image.Image {
	return nil
}

func (*ArrayCarver) Height() int {
	return 0
}

func (*ArrayCarver) Width() int {
	return 0
}

func (*ArrayCarver) Energy(y, x int) (int, error) {
	return 0, nil
}

func (*ArrayCarver) HSeam() ([]int, error) {
	return nil, nil
}

func (*ArrayCarver) VSeam() ([]int, error) {
	return nil, nil
}

func (*ArrayCarver) HRemoveSeam([]int) error {
	return nil
}

func (*ArrayCarver) VRemoveSeam([]int) error {
	return nil
}
