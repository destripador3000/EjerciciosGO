package main

import "fmt"

func restar(a int, b int) {
	fmt.Print("Escribe dos números separados por un espacio: ")
	fmt.Scan(&a, &b)
	result := a - b
	fmt.Print(result)

}
func sumar(a int, b int) {
	fmt.Print("Escribe dos números separados por un espacio: ")
	fmt.Scan(&a, &b)
	result := a + b
	fmt.Print(result)

}
func multiplicar(a int, b int) {
	fmt.Print("Escribe dos números separados por un espacio: ")
	fmt.Scan(&a, &b)
	result := a * b
	fmt.Print(result)

}
func dividir(a int, b int) {
	fmt.Print("Escribe dos números separados por un espacio: ")
	fmt.Scan(&a, &b)
	result := a / b
	fmt.Print(result)

}

func main() {
	var a, b int
	var opc int
	fmt.Print("Que operación matematica quieres hacer:\n1)sumar\n2)restar\n3)multiplicar\n4)dividir\n")
	fmt.Scan(&opc)

	if opc == 1 {
		sumar(a, b)
	} else if opc == 2 {

		restar(a, b)
	} else if opc == 3 {

		multiplicar(a, b)
	} else if opc == 4 {

		dividir(a, b)
	}

}
