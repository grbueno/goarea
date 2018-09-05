package goarea

import "math"

//Pi constante valor de Pi
const Pi = 3.1416

// Circ calcula area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula area do retangulo
// calcula a area e retorna o resultado
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return base * altura / 2
}
