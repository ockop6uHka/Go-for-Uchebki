package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xMin, yMin, xMax, yMax = -2, -2, +2, +2
		imgWidth, imgHeight    = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	for py := 0; py < imgHeight; py++ {
		y := float64(py)/float64(imgHeight)*(yMax-yMin) + yMin
		for px := 0; px < imgWidth; px++ {
			x := float64(px)/float64(imgWidth)*(xMax-xMin) + xMin
			z := complex(x, y)
			// Точка (px, py) представляет комплексное значение z
			img.Set(px, py, calculateMandelbrotColor(z))
		}
	}
	png.Encode(os.Stdout, img) // Игнор ошибок
}

func calculateMandelbrotColor(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			red := uint8(real(v) * 128)   // Преобразование части вещественной части в цвет
			green := uint8(imag(v) * 128) // Преобразование части мнимой части в цвет
			blue := 255 - contrast*n      // Интенсивность цвета
			return color.RGBA{red, green, blue, 255}
		}
	}

	return color.Black
}
