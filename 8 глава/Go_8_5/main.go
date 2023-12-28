package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"os"
	"sync"
)

func main() {
	imageResult := generateMandelbrotImage(8)
	saveImageToPNG(os.Stdout, imageResult)
}

func generateMandelbrotImage(workers int) *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup

	for workerID := 0; workerID < workers; workerID++ {
		wg.Add(1)
		minHeight := height * workerID / workers
		maxHeight := height * (workerID + 1) / workers
		if workerID == workers-1 {
			maxHeight = height
		}
		go func(minHeight, maxHeight int) {
			defer wg.Done()
			generateWorkerMandelbrot(img, minHeight, maxHeight, width, height, xmin, xmax, ymin, ymax)
		}(minHeight, maxHeight)
	}

	wg.Wait()
	return img
}

func generateWorkerMandelbrot(img *image.RGBA, minHeight, maxHeight, width, height int, xmin, xmax, ymin, ymax float64) {
	for py := minHeight; py < maxHeight; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, calculateMandelbrotColor(z))
		}
	}
}

func calculateMandelbrotColor(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func saveImageToPNG(wr io.Writer, img image.Image) {
	png.Encode(wr, img)
}
