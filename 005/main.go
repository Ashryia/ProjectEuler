package main

import "fmt"

func main() {

	answer := 2520
	for divisible := false; divisible != true; answer++ {
		for div,isAnswer := 2,false; div <= 20 ; div++ {
			if answer%div == 0 {
				isAnswer = true
			}else{
				isAnswer = false
			}
			if !isAnswer{
				break
			}
			if div == 20 {
				divisible = true
			}
		}
	}
	fmt.Println(answer-1)

}
