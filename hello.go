package main

import "fmt"

func main() {
	// DECLARACION DE CONSTANTES
	const pi float64 = 3.14
	const pi2 = 3.14
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	// DECLARACIÓN DE VARIABLES ENTERAS
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)
	// ZERO VALUES
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
	// AREA DE UN CUADRADO
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)
}