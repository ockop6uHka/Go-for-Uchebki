package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64

type Feet float64
type Metre float64

type Pound float64
type Kilo float64

type Converter struct {
	Value float64
}

func (c Converter) ToCelsius() Celsius       { return Celsius(c.Value) }
func (c Converter) ToFahrenheit() Fahrenheit { return Fahrenheit(c.Value) }

func (f Converter) ToFeet() Feet   { return Feet(f.Value) }
func (f Converter) ToMetre() Metre { return Metre(f.Value) }

func (p Converter) ToPound() Pound { return Pound(p.Value) }
func (p Converter) ToKilo() Kilo   { return Kilo(p.Value) }

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conversion error: %v\n", err)
			os.Exit(1)
		}

		converter := Converter{Value: t}

		fmt.Printf("%.2f = %.2f°C, %.2f°F\n", converter.Value, converter.ToCelsius(), converter.ToFahrenheit())
		fmt.Printf("%.2f = %.2f feet, %.2f meters\n", converter.Value, converter.ToFeet(), converter.ToMetre())
		fmt.Printf("%.2f = %.2f pounds, %.2f kilograms\n", converter.Value, converter.ToPound(), converter.ToKilo())
		fmt.Println() // Добавлен этот вызов для вывода пустой строки после каждой итерации
	}

	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}
