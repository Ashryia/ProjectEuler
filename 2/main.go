package main

import "fmt"

func main() {
	answer := 0

	for  value1, value2, value3 := 1,1,0; value3 < 4000000; value3 = value1+value2 {
		if (value3%2 == 0){
			answer += value3
		}
		value1, value2 = value2, value3
	}

        fmt.Printf("The answer is: %d\n", answer)

}
