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

func (a *ArrayCarver) toIndex(x, y int) int {
	return x + (y * a.Width())
}
func (*ArrayCarver) HSeam() ([]int, error) {
	return nil, nil
}

func (a *ArrayCarver) VSeam() ([]int, error) {
	w := a.Width()
	h := a.Height()
	distTo := make([]float64, w*h)
	edgeTo := make([]int, w*h)
	e := make([]float64, w*h)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			en, err := a.Energy(j,i)
			if err != nil {
				return nil, err
			}
			e[a.toIndex(j,i)] = en
		}
	}

	for i := 0; i < w; i++ {
		distTo[a.toIndex(i,0)] = MaxEnergy
	}

	for i := 1; i < h; i++ {
		for j := 0; j < w; j++ {
			ji := a.toIndex(j,i)
			distTo[ji] = math.Inf(1)
			if j > 0 {
				distTo[ji] = distTo[a.toIndex(j-1,i-1)] + e[ji]
				edgeTo[ji] = a.toIndex(j-1,i-1)
			}
			if distTo[a.toIndex(j,i-1)] + e[ji] < distTo[ji] {
				distTo[ji] = distTo[a.toIndex(j,i-1)] + e[ji]
				edgeTo[ji] = a.toIndex(j,i-1)
			}
			if j < w-1 && distTo[a.toIndex(j+1,i-1)] + e[ji] < distTo[ji] {
				distTo[ji] = distTo[a.toIndex(j+1,i-1)] + e[ji]
				edgeTo[ji] = a.toIndex(j+1,i-1)
			}
		}
	}

	mini := a.toIndex(0,h-1)
	min := distTo[mini]
	for i := mini+1; i < h*w; i++ {
		if distTo[i] < min {
			min = distTo[i]
			mini = i
		}
	}

	res := make([]int,h)
	n := mini
	for i := h-1; i >= 0; i-- {
		res[i] = n%w
		n = edgeTo[n]
	}

	return res, nil
}

func (*ArrayCarver) HRemoveSeam([]int) error {
	return nil
}

func (*ArrayCarver) VRemoveSeam([]int) error {
	return nil
}
