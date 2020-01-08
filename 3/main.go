package main

import "fmt"

func main() {

	val := 600851475143
	div,answer := 2,0
	max := val/2

	if isPrime(val){
		answer = val
	}else {
		for ;div < max; div,max=div+1,val/div {
			if val%div == 0{
				if isPrime(val/div){
					answer = val/div
					return
				}else {
					if isPrime(div){
						answer = div
					}
				}
			}
		}
	}


        fmt.Printf("The answer is: %d\n", answer)

}

// Use this to figure out if something is a prime number
func isPrime(val int) bool {
	max := val/2
	div := 2
	for ;div < max; div,max=div+1,val/div {
		if val%div == 0{
			return false
		}
	}
	return true
}
