package carver

import (
	"image"
	_ "image/jpeg"
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
	{
		name:   "5x10.png",
		width:  5,
		height: 10,
		vseam:  []int{1, 1, 1, 2, 1, 2, 3, 2, 1, 0},
		hseam:  []int{1, 2, 2, 1, 1},
		eRow:   2,
		energy: []float64{1000, 0, 190, 234, 1000},
	},
	{
		name:   "16x16.png",
		width:  16,
		height: 16,
		vseam:  []int{7, 7, 6, 7, 6, 5, 4, 3, 4, 5, 4, 5, 6, 7, 7, 6},
		hseam:  []int{7, 8, 8, 7, 6, 5, 4, 3, 2, 2, 2, 2, 1, 2, 1, 1},
		eRow:   1,
		energy: []float64{1000, 312, 262, 268, 332, 169, 215, 117, 300, 247, 265,
			263, 138, 372, 214, 1000},
	},
	/*
		{
			name:   "10x10.jpg",
			width:  10,
			height: 10,
			vseam:  []int{7, 8, 7, 6, 5, 6, 5, 6, 7, 6},
			hseam:  []int{1, 2, 3, 2, 2, 3, 3, 2, 1, 1},
			eRow:   3,
			energy: []float64{1000, 242, 182, 333, 235, 194, 213, 442, 288, 1000},
		},
	*/
}

func loadTestFile(t *testing.T, tf TestFile) Carver {
	f, err := os.Open(path.Join(testPath, tf.name))
	if err != nil {
		t.Errorf("Could not open test file %s", tf.name)
		return nil
	}
	im, _, err := image.Decode(f)
	if err != nil {
		t.Errorf("Could not decode test file %s", tf.name)
		return nil
	}
	c, err := NewArrayCarver(im)
	if err != nil {
		t.Errorf("%s: %s", tf.name, err)
		return nil
	}
	return c
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCarvers(t *testing.T) {
	for i := range testFiles {
		ti := testFiles[i]
		c := loadTestFile(t, ti)
		if c == nil {
			return
		}
		if c.Width() != ti.width || c.Height() != ti.height {
			t.Errorf("carver: invalid dimensions for %s", ti.name)
		}
		for i, te := range ti.energy {
			e, err := c.Energy(i, ti.eRow)
			if err != nil {
				t.Errorf("%s: Energy: %s", ti.name, err)
			}

			// other test client rounds energy, so check to within 0.5
			d := math.Abs(e - te)
			if d > 0.5 {
				t.Errorf("%s: invalid energy %f (expecting %f)", ti.name, e, te)
			}
		}
		vseam, err := c.VSeam()
		if err != nil {
			t.Errorf("%s: error finding V. Seam", ti.name)
		}
		if !equals(vseam, ti.vseam) {
			t.Errorf("%s: vseam mismatch", ti.name)
			t.Errorf("received:  %v", vseam)
			t.Errorf("expecting: %v", ti.vseam)
		}
	}
}

func loadFirstArrayCarver(t *testing.T) *ArrayCarver {
	f, err := os.Open(path.Join(testPath, testFiles[0].name))
	if err != nil {
		t.Errorf("Could not open default test file: %s", testFiles[0].name)
	}
	im, _, err := image.Decode(f)
	if err != nil {
		t.Errorf("Could not decode default test file %s", testFiles[0].name)
	}
	c, err := NewArrayCarver(im)
	if err != nil {
		t.Errorf("%s: err", testFiles[0].name)
	}
	return c
}

func TestEnergy(t *testing.T) {
	c := loadFirstArrayCarver(t)
	_, err := c.Energy(-1, 1)
	_, err2 := c.Energy(1, -1)
	if err != ErrInvalid || err2 != ErrInvalid {
		t.Error("out of bounds check fail")
	}
	_, err = c.Energy(c.Width(), 1)
	_, err2 = c.Energy(1, c.Height())
	if err != ErrInvalid || err2 != ErrInvalid {
		t.Error("out of bounds check fail")
	}
	e, err := c.Energy(0, 1)
	if err != nil || e != MaxEnergy {
		t.Errorf("border energy must equal MaxEnergy (%f)", MaxEnergy)
	}
	e, err = c.Energy(1, 0)
	if err != nil || e != MaxEnergy {
		t.Errorf("border energy must equal MaxEnergy (%f)", MaxEnergy)
	}
	e, err = c.Energy(c.Width()-1, 1)
	if err != nil || e != MaxEnergy {
		t.Errorf("border energy must equal MaxEnergy (%f)", MaxEnergy)
	}
	e, err = c.Energy(1, c.Height()-1)
	if err != nil || e != MaxEnergy {
		t.Errorf("border energy must equal MaxEnergy (%f)", MaxEnergy)
	}
}

func TestArraryVerifySeam(t *testing.T) {
	c := loadFirstArrayCarver(t)
	err := c.verifySeam(testFiles[0].hseam, true)
	err2 := c.verifySeam(testFiles[0].vseam, false)
	if err != nil || err2 != nil {
		t.Error("seam test failed")
	}
	err = c.verifySeam(testFiles[0].hseam, false)
	err2 = c.verifySeam(testFiles[0].vseam, true)
	if err == nil || err2 == nil {
		t.Error("seam length test failed")
	}
	tseam := []int{0, 1, 4, 2, 0}
	err = c.verifySeam(tseam, true)
	if err == nil {
		t.Error("seam continuity test fail")
	}
}
