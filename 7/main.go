package main

import "fmt"

func main() {
	counter := 2
	primeCounter := 0
	for ; primeCounter < 10001 ; counter++ {
		if isPrime(counter){
			fmt.Println(counter)
			primeCounter++
		}
	}
	fmt.Println(counter-1, primeCounter)
}

func isPrime(val int) bool {
	max := val/2
	div := 2
	for ; div <= max ; div,max = div+1, val/div{
		if val%div == 0 {
			return false
		}
	}
	return true
}
