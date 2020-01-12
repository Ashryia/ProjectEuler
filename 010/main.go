package main

import "fmt"

func main() {
	counter := 2
	answer := 0
	for ; counter < 2000000; counter++{
		if isPrime(counter){
			answer += counter
		}
	}
	fmt.Printf("The answer is %d\n", answer)
}

func isPrime(val int) bool {
        max := val/2
        div := 2
        for ;div <= max ; div,max=div+1,val/div {
                if val%div == 0{
                        return false
                }
        }
        return true
}



