package main

import (
	"fmt"
	"math"
)

const (
	canvasWidth, canvasHeight = 600, 320                  // размер канвы в пикселях
	cells                     = 100                       // количество ячеек сетки
	xyRange                   = 30.0                      // диапазон осей (-xyRange..+ xyRange)
	xyScale                   = canvasWidth / 2 / xyRange // пикселей в единице x или y
	zScale                    = canvasHeight * 0.4        // пикселей в единице z
	angle                     = math.Pi / 6               // углы осей x, y (= 30 градусов)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30 degree), cos(30 degree)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height = '%d'>", canvasWidth, canvasHeight)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			polygon := generatePolygon(i, j)
			if polygon[0].x == 0 && polygon[0].y == 0 {
				continue
			}
			fmt.Printf("<polygon points= '%g,%g %g,%g %g,%g %g,%g' style='fill:%s' />\n",
				polygon[0].x, polygon[0].y,
				polygon[1].x, polygon[1].y,
				polygon[2].x, polygon[2].y,
				polygon[3].x, polygon[3].y,
				polygon[0].color)
		}
	}
	fmt.Println("</svg>")
}

type point struct {
	x, y  float64
	color string
}

func generatePolygon(i, j int) [4]point {
	var polygon [4]point
	for k := 0; k < 4; k++ {
		x, y, color := getCornerColor(i, j, k)
		if math.IsInf(x, 0) || math.IsInf(y, 0) {
			// Проверка на + и - бесконечности. Их пропускаем
			return [4]point{}
		}
		polygon[k] = point{x, y, color}
	}
	return polygon
}

func getCornerColor(i, j, k int) (float64, float64, string) {
	// ищем угловую точку (x, y) ячейки (i, j)
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)
	// вычисляем высоту поверхности z
	z := f(x, y)
	// изометрически проецируем (x, y, z) на двумерную канву SVG( sx, sy)
	sx := canvasWidth/2 + (x-y)*cos30*xyScale
	sy := canvasHeight/2 + (x+y)*sin30*xyScale - z*zScale
	color := getColor(z)
	return sx, sy, color
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // расстояние от начала координат
	return math.Sin(r) / r
}

func getColor(z float64) string {
	if z > 0 {
		return "#ff0000"
	}
	return "#0000ff"
}
