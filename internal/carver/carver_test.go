package carver

import (
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
}

var testFiles = []TestFile{
	TestFile{
		name: "5x10.png",
		width: 5,
		height: 10,
		vseam: []int{ 0, 1, 1, 2, 1, 2, 3, 2, 1, 1 },
		hseam: []int{ 1, 2, 2, 1, 1 },
	},
	TestFile{
		name:   "16x16.png",
		width:  16,
		height: 16,
		vseam:  []int{7, 7, 6, 7, 6, 5, 4, 3, 4, 5, 4, 5, 6, 7, 7, 6},
		hseam:  []int{7, 8, 8, 7, 6, 5, 4, 3, 2, 2, 2, 2, 1, 2, 1, 1},
	},
}

func loadTestFile(t *testing.T, tf TestFile) *Carver {
	name := path.Join(testPath, tf.name)
	c, err := New(name)
	if err != nil {
		t.Errorf("Could not load test file %s", name)
	}
	return c
}

func TestCarvers(t *testing.T) {
}
