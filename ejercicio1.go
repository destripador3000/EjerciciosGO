package main

import "fmt"

func ciclo() {
	listaCosas := []string{"Carro", "Moto", "Avión", "Bus"}
	/*Clausura Range la i representa la posición del elemento y n representa el elemento en si.*/

	/*Si se quiere imprimir solo una de las 2 se puede omitir una de las variables poniendo _ */
	for i, n := range listaCosas {
		fmt.Println(i, n)

	}

}
