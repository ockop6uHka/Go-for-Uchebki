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
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsInf(ax, 0) || math.IsInf(ay, 0) || math.IsInf(bx, 0) || math.IsInf(by, 0) ||
				math.IsInf(cx, 0) || math.IsInf(cy, 0) || math.IsInf(dx, 0) || math.IsInf(dy, 0) {
				// Проверка на + и - бесконечности. Их пропускаем
				continue
			}
			fmt.Printf("<polygon points= '%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// ищем угловую точку (x, y) ячейки (i, j)
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)
	// вычисляем высоту поверхности z
	z := f(x, y)
	// изометрически проецируем (x, y, z) на двумерную канву SVG( sx, sy)
	sx := canvasWidth/2 + (x-y)*cos30*xyScale
	sy := canvasHeight/2 + (x+y)*sin30*xyScale - z*zScale
	return sx, sy
}

// func f(x, y float64) float64 {
// 	r := math.Hypot(x, y) // расстояние от начала координат
// 	return math.Atan(r) - r
// }

func f(x, y float64) float64 {
	return math.Sin(x) * math.Cos(y)
}
