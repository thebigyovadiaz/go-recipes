package main

import (
	"fmt"
)

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		output := ""
		if i%3 == 0 {
			output += "Fizz"
		}

		if i%5 == 0 {
			output += "Buzz"
		}

		if output == "" {
			fmt.Println(i)
		} else {
			fmt.Println(output)
		}
	}

	return
}
