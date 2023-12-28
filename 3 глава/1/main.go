package main

import (
	"fmt"
	"math"
)

const (
	canvasWidth, canvasHeight = 600, 320                  // размер канвы в пикселях
	cells                     = 100                       // количество ячеек сетки
	xyrange                   = 30.0                      // диапазон осей (-xyrange..+ xyrange)
	xyscale                   = canvasWidth / 2 / xyrange // пикселей в единице х или у
	zscale                    = canvasHeight * 0.4        // пикселей в единице z
	angle                     = math.Pi / 6               // углы осей х, у (= 30 градусов)
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
	// ищем угловую точку (х, у) ячейки (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// вычисляем высоту поверхности z
	z := f(x, y)
	// изометрически проецируем (x, y, z) на двумерную канву SVG( sx, sy)
	sx := canvasWidth/2 + (x-y)*cos30*xyscale
	sy := canvasHeight/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // расстояние от начала координат
	return math.Sin(r) / r
}
