package carver

import (
	"image"
	_ "image/png"
	"math"
	"os"
	"path"
	"testing"
)

const testPath = "../../test/"

type TestFile struct {
	name   string
	width  int
	height int
	hseam  []int
	vseam  []int
	energy []float64
	eRow   int
}

var testFiles = []TestFile{
	TestFile{
		name:   "5x10.png",
		width:  5,
		height: 10,
		vseam:  []int{0, 1, 1, 2, 1, 2, 3, 2, 1, 1},
		hseam:  []int{1, 2, 2, 1, 1},
		eRow:   2,
		energy: []float64{1000, 0, 190, 234, 1000},
	},
	TestFile{
		name:   "16x16.png",
		width:  16,
		height: 16,
		vseam:  []int{7, 7, 6, 7, 6, 5, 4, 3, 4, 5, 4, 5, 6, 7, 7, 6},
		hseam:  []int{7, 8, 8, 7, 6, 5, 4, 3, 2, 2, 2, 2, 1, 2, 1, 1},
		eRow:   1,
		energy: []float64{1000, 312, 262, 268, 332, 169, 215, 117, 300, 247, 265,
			263, 138, 372, 214, 1000},
	},
}

func loadTestFile(t *testing.T, tf TestFile) Carver {
	f, err := os.Open(path.Join(testPath, tf.name))
	if err != nil {
		t.Errorf("Could not open test file %s", tf.name)
	}
	im, _, err := image.Decode(f)
	if err != nil {
		t.Errorf("Could not decode test file %s", tf.name)
	}
	c, err := New(im)
	if err != nil {
		t.Errorf("%s: err", tf.name)
	}
	return c
}

func TestCarvers(t *testing.T) {
	for i := range testFiles {
		ti := testFiles[i]
		c := loadTestFile(t, ti)
		if c.Width() != ti.width || c.Height() != ti.height {
			t.Errorf("carver: invalid dimensions for %s", ti.name)
		}
		for i, te := range ti.energy {
			e, err := c.Energy(i, ti.eRow)
			if err != nil {
				t.Errorf("%s: Energy: %s", ti.name, err)
			}
			d := math.Abs(e - te)
			if d > 0.5 {
				t.Errorf("%s: invalid energy %f (expecting %f)", ti.name, e, te)
			}
		}
	}
}

func loadDefaultCarver(t *testing.T) Carver {
	f, err := os.Open(path.Join(testPath, testFiles[0].name))
	if err != nil {
		t.Errorf("Could not open default test file: %s", testFiles[0].name)
	}
	im, _, err := image.Decode(f)
	if err != nil {
		t.Errorf("Could not decode default test file %s", testFiles[0].name)
	}
	c, err := New(im)
	if err != nil {
		t.Errorf("%s: err", testFiles[0].name)
	}
	return c
}

func TestEnergy(t *testing.T) {
	c := loadDefaultCarver(t)
	_, err := c.Energy(-1,1)
	_, err2 := c.Energy(1,-1)
	if err != ErrInvalid || err2 != ErrInvalid {
		t.Error("out of bounds check fail")
	}
	_, err = c.Energy(c.Width(), 1)
	_, err2 = c.Energy(1, c.Height())
	if err != ErrInvalid || err2 != ErrInvalid {
		t.Error("out of bounds check fail")
	}
	e, err := c.Energy(0,1)
	if err != nil || e != MaxEnergy {
		t.Errorf("border energy must equal MaxEnergy (%f)", MaxEnergy)
	}
	e, err = c.Energy(1,0)
	if err != nil || e != MaxEnergy {
		t.Errorf("border energy must equal MaxEnergy (%f)", MaxEnergy)
	}
	e, err = c.Energy(c.Width()-1,1)
	if err != nil || e != MaxEnergy {
		t.Errorf("border energy must equal MaxEnergy (%f)", MaxEnergy)
	}
	e, err = c.Energy(1,c.Height()-1)
	if err != nil || e != MaxEnergy {
		t.Errorf("border energy must equal MaxEnergy (%f)", MaxEnergy)
	}
}
