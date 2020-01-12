package main

import (
	"fmt"
)

func main() {
	// holds the upperbound
	upperbound := 1000

	// while the value is less than the upper bound, check if the value is a multiple of 3 or 5
	value := 1
	answer := 0
	for value < upperbound {
		if (value%3 == 0) || (value%5 == 0) {
			//fmt.Printf("%d is a multiple of 3 or 5\n", value)
			answer = answer + value
		}
		value++
	}

	fmt.Printf("The answer is: %d\n", answer)
}
