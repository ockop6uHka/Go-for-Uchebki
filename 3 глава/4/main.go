package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var axisAngle = math.Pi / 6                                               // угол между осями x и y (30 градусов)
var sinAxisAngle, cosAxisAngle = math.Sin(axisAngle), math.Cos(axisAngle) // синус и косинус угла

const (
	canvasHeight = 320  // высота канвы в пикселях
	cells        = 100  // количество ячеек сетки
	xyRange      = 30.0 // диапазон осей (-xyRange..+ xyRange)
)

func main() {
	http.HandleFunc("/svg", handleSVG)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handleSVG(w http.ResponseWriter, r *http.Request) {
	width := 600
	height := 320
	strokeColor := "grey"

	if widthParam := r.URL.Query().Get("width"); widthParam != "" {
		widthInt, err := strconv.Atoi(widthParam)
		if err == nil && widthInt > 0 {
			width = widthInt
		}
	}

	if heightParam := r.URL.Query().Get("height"); heightParam != "" {
		heightInt, err := strconv.Atoi(heightParam)
		if err == nil && heightInt > 0 {
			height = heightInt
		}
	}

	if colorParam := r.URL.Query().Get("color"); colorParam != "" {
		strokeColor = colorParam
	}

	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", strokeColor, width, height)

	xyscale := float64(width) / 2 / xyRange
	zscale := float64(height) * 0.1 // пикселей в единице z

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, xyscale, zscale)
			bx, by := corner(i, j, width, xyscale, zscale)
			cx, cy := corner(i, j+1, width, xyscale, zscale)
			dx, dy := corner(i+1, j+1, width, xyscale, zscale)
			if math.IsInf(ax, 0) || math.IsInf(ay, 0) || math.IsInf(bx, 0) || math.IsInf(by, 0) ||
				math.IsInf(cx, 0) || math.IsInf(cy, 0) || math.IsInf(dx, 0) || math.IsInf(dy, 0) {
				// Пропускаем + и - бесконечности
				continue
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j, width int, xyscale, zscale float64) (float64, float64) {
	// Находим угловую точку (x, y) ячейки (i, j)
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z := f(x, y)
	// Изометрически проецируем (x, y, z) на двумерную канву SVG( sx, sy)
	sx := float64(width/2) + float64((x-y)*cosAxisAngle)*xyscale
	sy := canvasHeight/2 + float64((x+y)*sinAxisAngle)*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от начала координат
	return math.Sin(r) / r
}
