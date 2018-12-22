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
	r := (float64(ar) - float64(br)) / 257
	g := (float64(ag) - float64(bg)) / 257
	b := (float64(ab) - float64(bb)) / 257
	return math.Pow(r, 2) + math.Pow(g, 2) + math.Pow(b, 2);
}

func (a *ArrayCarver) Energy(x, y int) (float64, error) {
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

func (a *ArrayCarver) verifySeam(seam []int, h bool) error {
	var max int
	var length int
	if h {
		max = a.Height()-1
		length = a.Width()
	} else {
		max = a.Width()-1
		length = a.Height()
	}

	if len(seam) != length {
		return ErrInvalid
	}

	var prev int = seam[0]
	for i, n := range(seam) {
		if n < 0 || n > max {
			return ErrInvalid
		}
		if i > 0 {
			d := n - prev
			if d > 1 || d < -1 {
				return ErrInvalid
			}
		}
		prev = n
	}

	return nil
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
