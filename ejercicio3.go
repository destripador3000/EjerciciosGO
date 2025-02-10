package main

import "fmt"

func fizzBuzz() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%v Fizz Buzz \n", i)

		} else if i%3 == 0 {
			fmt.Printf("%v Fizz \n", i)

		} else if i%5 == 0 {
			fmt.Printf("%v Buzz \n", i)
		} else {
			fmt.Printf("%v \n", i)
		}

	}

}
