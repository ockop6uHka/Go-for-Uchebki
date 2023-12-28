package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xMin, yMin, xMax, yMax = -2, -2, +2, +2
	imgWidth, imgHeight    = 1024, 1024
	supersampleFactor      = 2
)

func main() {
	img := generateImage()
	png.Encode(os.Stdout, img) // Игнор ошибок
}

func generateImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, imgWidth*supersampleFactor, imgHeight*supersampleFactor))

	for py := 0; py < imgHeight; py++ {
		for px := 0; px < imgWidth; px++ {
			for sdy := 0; sdy < supersampleFactor; sdy++ {
				for sdx := 0; sdx < supersampleFactor; sdx++ {
					x, y := calculateCoordinates(px, py, sdx, sdy)
					z := complex(x, y)
					img.Set(px*supersampleFactor+sdx, py*supersampleFactor+sdy, calculateMandelbrotColor(z))
				}
			}
		}
	}

	return img
}

func calculateCoordinates(px, py, sdx, sdy int) (float64, float64) {
	x := float64(px*supersampleFactor+sdx)/float64(imgWidth*supersampleFactor)*(xMax-xMin) + xMin
	y := float64(py*supersampleFactor+sdy)/float64(imgHeight*supersampleFactor)*(yMax-yMin) + yMin
	return x, y
}

func calculateMandelbrotColor(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			red := uint8(real(v) * 128)
			green := uint8(imag(v) * 128)
			blue := 255 - contrast*n
			return color.RGBA{red, green, blue, 255}
		}
	}

	return color.Black
}
