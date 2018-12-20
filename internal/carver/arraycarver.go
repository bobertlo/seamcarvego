package carver

import (
	"image"
	"image/color"
	"math"
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

func gradient(ca, cb color.Color) float64 {
	ar, ag, ab, _ := ca.RGBA()
	br, bg, bb, _ := cb.RGBA()
	r := float64((ar - br) >> 8)
	g := float64((ag - bg) >> 8)
	b := float64((ab - bb) >> 8)
	return math.Pow(r, 2) + math.Pow(g, 2) + math.Pow(b, 2);
}

func (a *ArrayCarver) Energy(y, x int) (float64, error) {
	h := a.Height()
	w := a.Width()

	if (x < 0 || y < 0 || x >= w || y >= h) {
		return 0, ErrInvalid
	}

	if (x == 0 || y == 0 || x == w-1 || y == h-1) {
		return MaxEnergy, nil
	}

	x1 := a.img.At(x-1, y)
	x2 := a.img.At(x+1, y)
	y1 := a.img.At(x, y-1)
	y2 := a.img.At(x, y+1)

	xg := gradient(x1, x2)
	yg := gradient(y1, y2)

	return math.Sqrt(xg + yg), nil
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
