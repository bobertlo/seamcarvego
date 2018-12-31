package main

import (
	"fmt"
	"github.com/bobertlo/seamcarvego/internal/carver"
	"image"
	"image/png"
	"os"
	"strconv"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: seamcarvego [input] [output] [cols] [rows]")
}

func main() {
	if len(os.Args) != 5 {
		usage()
		os.Exit(1)
	}

	// parse dimension arguments
	cols, err := strconv.Atoi(os.Args[3])
	rows, err2 := strconv.Atoi(os.Args[4])
	if err != nil || err2 != nil {
		fmt.Fprintln(os.Stderr, "invalid dimension arguments")
		usage()
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to decode %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}

	if rows > img.Bounds().Dy() || cols > img.Bounds().Dx() {
		fmt.Println("invalid target dimensions")
		os.Exit(1)
	}

	c, err := carver.NewArrayCarver(img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initializing carver: %s\n", err)
		os.Exit(1)
	}

	for c.Width() > cols || c.Height() > rows {
		if c.Height() > rows {
			s, err := c.HSeam()
			if err != nil {
				panic("carver error")
			}
			c.HRemoveSeam(s)
		}
		if c.Width() > cols {
			s, err := c.VSeam()
			if err != nil {
				panic("carver error")
			}
			c.VRemoveSeam(s)
		}
	}

	w, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening output file %s: %s\n", os.Args[2], err)
		os.Exit(1)
	}
	png.Encode(w, c.Img())
}
