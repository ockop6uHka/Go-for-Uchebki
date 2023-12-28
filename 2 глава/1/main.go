package main

import (
	"fmt"
)

type Temperature struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

const (
	AbsoluteZeroCelsius = -273.15
	FreezingCelsius     = 0
	BoilingCelsius      = 100
	AbsoluteZeroKelvin  = -273.15
)

func NewTemperatureFromCelsius(celsius float64) Temperature {
	return Temperature{
		Celsius:    celsius,
		Fahrenheit: celsiusToFahrenheit(celsius),
		Kelvin:     celsiusToKelvin(celsius),
	}
}

func NewTemperatureFromFahrenheit(fahrenheit float64) Temperature {
	celsius := fahrenheitToCelsius(fahrenheit)
	return Temperature{
		Celsius:    celsius,
		Fahrenheit: fahrenheit,
		Kelvin:     celsiusToKelvin(celsius),
	}
}

func NewTemperatureFromKelvin(kelvin float64) Temperature {
	celsius := kelvinToCelsius(kelvin)
	return Temperature{
		Celsius:    celsius,
		Fahrenheit: celsiusToFahrenheit(celsius),
		Kelvin:     kelvin,
	}
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.2f°C, %.2f°F, %.2fK", t.Celsius, t.Fahrenheit, t.Kelvin)
}

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius - AbsoluteZeroKelvin
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin + AbsoluteZeroKelvin
}

func main() {
	celsiusTemperature := NewTemperatureFromCelsius(10)
	fahrenheitTemperature := NewTemperatureFromFahrenheit(10)
	kelvinTemperature := NewTemperatureFromKelvin(10)

	fmt.Println("Celsius:", celsiusTemperature)
	fmt.Println("Fahrenheit:", fahrenheitTemperature)
	fmt.Println("Kelvin:", kelvinTemperature)
}
