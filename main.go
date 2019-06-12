package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

func dist(a, b [2]int) float64 {
	return math.Sqrt(sq(a[0]-b[0]) + sq(a[1]-b[1]))
}

func sq(x int) float64 {
	return float64(x * x)
}

func colors(c []color.Color) {
	for i := range c {
		c[i] = color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255}
	}
}

func points(p [][2]int) {
	for i := range p {
		p[i] = [2]int{rand.Intn(width), rand.Intn(height)}
	}
}

var width, height, sites int = 500, 500, 500

func main() {
	set := make([][2]int, sites)
	points(set)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	cols := make([]color.Color, len(set))
	colors(cols)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			bestP := 0
			bestD := dist([2]int{x, y}, set[0])
			for i := 1; i < len(set); i++ {
				currD := dist([2]int{x, y}, set[i])
				if currD < bestD {
					bestD = currD
					bestP = i
				}
			}

			img.Set(x, y, cols[bestP])
		}
		fmt.Println(x)
	}

	err := save(img, "img.png")
	if err != nil {
		panic(err)
	}
}

func save(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
