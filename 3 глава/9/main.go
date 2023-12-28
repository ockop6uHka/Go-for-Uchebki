package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"strconv"
	"sync"
)

// FractalParams хранит параметры для генерации фрактала
type FractalParams struct {
	CenterX float64
	CenterY float64
	Scale   float64
}

// FractalType определяет тип фрактала
type FractalType int

const (
	Mandelbrot FractalType = iota
	// Добавьте другие типы фракталов при необходимости
)

// FractalGenerator генерирует изображение фрактала
type FractalGenerator struct {
	fractalType FractalType
}

// NewFractalGenerator создает новый генератор фрактала с указанным типом
func NewFractalGenerator(fractalType FractalType) *FractalGenerator {
	return &FractalGenerator{fractalType: fractalType}
}

// Generate генерирует изображение фрактала с заданными параметрами
func (gen *FractalGenerator) Generate(params FractalParams, width, height int) image.Image {
	fractalImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// Логика генерации фрактала
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			// Вычисление комплексного числа в соответствии с параметрами
			realPart := params.CenterX + (float64(px)/float64(width)-0.5)*params.Scale
			imagPart := params.CenterY + (float64(py)/float64(height)-0.5)*params.Scale
			z := complex(realPart, imagPart)

			// Применение алгоритма фрактала
			fractalImg.Set(px, py, gen.calculateColor(z))
		}
	}

	return fractalImg
}

// calculateColor возвращает цвет для конкретной точки фрактала
func (gen *FractalGenerator) calculateColor(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, 110, 0, 255}
		}
	}
	return color.Black
}

// FractalHandler обрабатывает запросы для генерации фрактала
type FractalHandler struct {
	fractalGenerator *FractalGenerator
	paramsMutex      sync.Mutex
}

// NewFractalHandler создает новый обработчик фрактала
func NewFractalHandler(fractalType FractalType) *FractalHandler {
	return &FractalHandler{fractalGenerator: NewFractalGenerator(fractalType)}
}

// SetParams устанавливает параметры фрактала
func (handler *FractalHandler) SetParams(params FractalParams) {
	handler.paramsMutex.Lock()
	defer handler.paramsMutex.Unlock()
	// Устанавливаем новые параметры фрактала
	// Можно добавить валидацию параметров, если нужно
	// Например, чтобы не допустить слишком больших значений
	// или проверить, что Scale > 0 и т.д.
}

// Handle запрос обрабатывает запрос на генерацию фрактала с текущими параметрами
func (handler *FractalHandler) Handle(w http.ResponseWriter, r *http.Request) {
	handler.paramsMutex.Lock()
	defer handler.paramsMutex.Unlock()

	// Получаем значения параметров из запроса
	query := r.URL.Query()
	centerX, _ := strconv.ParseFloat(query.Get("centerX"), 64)
	centerY, _ := strconv.ParseFloat(query.Get("centerY"), 64)
	scale, _ := strconv.ParseFloat(query.Get("scale"), 64)

	// Устанавливаем новые параметры фрактала
	handler.SetParams(FractalParams{CenterX: centerX, CenterY: centerY, Scale: scale})

	// Генерация изображения фрактала
	fractalImg := handler.fractalGenerator.Generate(FractalParams{CenterX: centerX, CenterY: centerY, Scale: scale}, 800, 600)

	// Отправка изображения клиенту
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, fractalImg)
}

func main() {
	http.HandleFunc("/fractal", NewFractalHandler(Mandelbrot).Handle)
	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
