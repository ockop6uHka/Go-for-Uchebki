package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	xMin, yMin, xMax, yMax = -2, -2, +2, +2
	imgWidth, imgHeight    = 1024, 1024
)

// NewtonFractalGenerator представляет генератор фрактала Ньютона
type NewtonFractalGenerator struct {
	img *image.RGBA
}

// NewNewtonFractalGenerator создает новый генератор фрактала Ньютона
func NewNewtonFractalGenerator() *NewtonFractalGenerator {
	return &NewtonFractalGenerator{img: image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))}
}

// Generate создает изображение фрактала Ньютона
func (nf *NewtonFractalGenerator) Generate() {
	for py := 0; py < imgHeight; py++ {
		for px := 0; px < imgWidth; px++ {
			x := float64(px)/imgWidth*(xMax-xMin) + xMin
			y := float64(py)/imgHeight*(yMax-yMin) + yMin
			z := complex(x, y)

			root, iterations := nf.calculateNewton(z)

			color := nf.colorForRoot(root, iterations)
			nf.img.Set(px, py, color)
		}
	}
}

// SaveToFile сохраняет изображение в файл
func (nf *NewtonFractalGenerator) SaveToFile(filename string) {
	file, err := os.Create(filename)
	nf.handleError(err)
	defer file.Close()

	err = png.Encode(file, nf.img)
	nf.handleError(err)
}

// handleError обрабатывает ошибку и выводит сообщение при ее наличии
func (nf *NewtonFractalGenerator) handleError(err error) {
	if err != nil {
		fmt.Println("Произошла ошибка:", err)
	}
}

// calculateNewton вычисляет корень и количество итераций для фрактала Ньютона
func (nf *NewtonFractalGenerator) calculateNewton(z complex128) (int, int) {
	const (
		iterations = 200
		threshold  = 1e-6
	)

	for i := 0; i < iterations; i++ {
		z = z - (cmplx.Pow(z, 3)-1)/(3*cmplx.Pow(z, 2))
		for root := 0; root < 3; root++ {
			if cmplx.Abs(z-cmplx.Exp(complex(0, 2*math.Pi*float64(root)/3))) < threshold {
				return root, i
			}
		}
	}
	return -1, iterations
}

// colorForRoot возвращает цвет в зависимости от корня и количества итераций
func (nf *NewtonFractalGenerator) colorForRoot(root, iterations int) color.Color {
	const maxColorValue = 255

	if root == 0 {
		return color.RGBA{maxColorValue - uint8(iterations), 0, 0, 255}
	} else if root == 1 {
		return color.RGBA{0, maxColorValue - uint8(iterations), 0, 255}
	} else if root == 2 {
		return color.RGBA{0, 0, maxColorValue - uint8(iterations), 255}
	}

	return color.Black
}

func main() {
	nfGenerator := NewNewtonFractalGenerator()
	nfGenerator.Generate()
	nfGenerator.SaveToFile("newton_fractal.png")
}
