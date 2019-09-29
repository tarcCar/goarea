package goarea

import "math"

//PI  valor de pi
const PI = 3.1416

//Circ calcula o valor da area
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

//Rect calcula o valor da area
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianEquilatero(base, altura float64) float64 {
	return (base * altura) / 2
}
