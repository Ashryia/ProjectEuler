package main

import "fmt"

func main() {

	squareOfSum, sumOfSquare := 0, 0

	for counter := 1 ; counter <= 100; counter++ {
		squareOfSum += counter
		sumOfSquare = sumOfSquare + (counter * counter)
	}
	
	squareOfSum = squareOfSum * squareOfSum

	answer := squareOfSum - sumOfSquare
	fmt.Println(answer)
}
