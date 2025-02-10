package main

import (
	"fmt"
	"strings"
)

func piramide(num int) {
	for i := 1; i < num+1; i++ {
		result := strings.Repeat("*", i)
		fmt.Println(result)

	}

}
func main() {
	piramide(7)
}
