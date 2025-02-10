package main

import "fmt"

func reverString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return result

}

func palindroma() {
	var palabra string

	fmt.Print("Escribe una palabra: ")
	fmt.Scan(&palabra)

	if palabra == reverString(palabra) {
		fmt.Println("Es palindroma")
	} else {
		fmt.Println("No es palindroma")
	}
}
