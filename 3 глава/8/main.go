package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

// MandelbrotStrategy определяет интерфейс стратегии для функции Мандельброта
type MandelbrotStrategy interface {
	Generate(z complex128) color.Color
}

// NormalMandelbrotStrategy реализует стратегию для обычных чисел с плавающей запятой
type NormalMandelbrotStrategy struct{}

func (s NormalMandelbrotStrategy) Generate(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:
				return color.RGBA{66, 30, 15, 255}
			default:
				return palettes[n%16]
			}
		}
	}
	return color.Black
}

// BigFloatMandelbrotStrategy реализует стратегию для чисел с плавающей запятой с произвольной точностью
type BigFloatMandelbrotStrategy struct{}

func (s BigFloatMandelbrotStrategy) Generate(z complex128) color.Color {
	const iterations = 20
	const contrast = 15
	zR := new(big.Float).SetFloat64(real(z))
	zI := new(big.Float).SetFloat64(imag(z))
	var vR, vI = new(big.Float), new(big.Float)
	for i := uint8(0); i < iterations; i++ {
		vR2, vI2 := new(big.Float), new(big.Float)
		vR2.Mul(vR, vR).Sub(vR2, new(big.Float).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := new(big.Float).Mul(vR, vR).Add(vI, new(big.Float).Mul(vI, vI))
		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			switch {
			case i > 5:
				return color.RGBA{66, 30, 15, 255}
			default:
				return palettes[i%16]
			}
		}
	}
	return color.Black
}

// MandelbrotRenderer представляет рендерер для изображения фрактала Мандельброта
type MandelbrotRenderer struct {
	Strategy MandelbrotStrategy
}

// NewMandelbrotRenderer создает новый экземпляр MandelbrotRenderer с заданной стратегией
func NewMandelbrotRenderer(strategy MandelbrotStrategy) *MandelbrotRenderer {
	return &MandelbrotRenderer{Strategy: strategy}
}

// Generate создает изображение фрактала Мандельброта с использованием стратегии
func (r *MandelbrotRenderer) Generate(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			y := float64(py)/float64(height)*(ymax-ymin) + ymin
			z := complex(x, y)
			img.Set(px, py, r.Strategy.Generate(z))
		}
	}
	return img
}

// SaveToFile сохраняет изображение в файл
func (r *MandelbrotRenderer) SaveToFile(filename string, img *image.RGBA) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

var palettes = [...]color.RGBA{
	{66, 30, 15, 255},
	{25, 7, 26, 255},
	{9, 1, 47, 255},
	{4, 4, 73, 255},
	{0, 7, 100, 255},
	{12, 44, 138, 255},
	{24, 82, 177, 255},
	{57, 125, 209, 255},
	{134, 181, 229, 255},
	{211, 236, 248, 255},
	{241, 233, 191, 255},
	{248, 201, 95, 255},
	{255, 170, 0, 255},
	{204, 128, 0, 255},
	{153, 87, 0, 255},
	{106, 52, 3, 255},
}

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
)

func main() {
	width, height := 1024, 1024

	normalStrategy := NormalMandelbrotStrategy{}
	normalRenderer := NewMandelbrotRenderer(normalStrategy)
	normalImg := normalRenderer.Generate(width, height)
	normalRenderer.SaveToFile("mandelbrot_normal.png", normalImg)

	bigFloatStrategy := BigFloatMandelbrotStrategy{}
	bigFloatRenderer := NewMandelbrotRenderer(bigFloatStrategy)
	bigFloatImg := bigFloatRenderer.Generate(width, height)
	bigFloatRenderer.SaveToFile("mandelbrot_bigfloat.png", bigFloatImg)
}
