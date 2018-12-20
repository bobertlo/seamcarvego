package carver

import (
	"image"
)

type ArrayCarver struct {
	img image.Image
}

func NewArrayCarver(img image.Image) (*ArrayCarver, error) {
	c := &ArrayCarver{
		img: img,
	}
	return c, nil
}

func (a *ArrayCarver) Img() image.Image {
	return a.img
}

func (a *ArrayCarver) Height() int {
	b := a.img.Bounds()
	return b.Max.Y - b.Min.Y
}

func (a *ArrayCarver) Width() int {
	b := a.img.Bounds()
	return b.Max.X - b.Min.X
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
