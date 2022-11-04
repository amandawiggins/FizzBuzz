package main

import (
	"fmt"
	"strconv"
)

func main() {
	fizzbuzz()
}

func fizzbuzz() {
	for i := 1; i < 101; i++ {
		var output string

		if i%3 == 0 {
			output = output + "fizz"
		}

		if i%5 == 0 {
			output = output + "buzz"
		}

		if output == "" {
			output = strconv.Itoa(i)
		}

		fmt.Println(output)
	}
}
