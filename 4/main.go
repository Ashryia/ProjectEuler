package main

import (
	"fmt"
	"strconv"
	)

func main() {
	maxValue := 999 * 999
	minValue := 100 * 100
	answer := 0
	for ; maxValue > minValue ; maxValue-- {
		if isPalindrome(maxValue) {
			if isProduct(maxValue){
				answer = maxValue
				fmt.Printf("The answer is %d\n", answer)
				return
			}
		}
	}
	fmt.Printf("Did not find an answer, answer is %d\n", answer)	
}
func isPalindrome(num int) bool {
	result := strconv.Itoa(num)
	length := len(result) - 1
	start := 0
	correct := false
	for ; start < length && start != length; start, length = start + 1, length -1 {
        	if result[start] == result[length]{
			correct = true
		}else{
			return false
		}
	}
	return correct
}
func isProduct(val int) bool {
	max := val/100
	threeDigMax := 999
	div := 100
	for ; div < max && div < threeDigMax; div,max= div+1, val/div {
		if val%div == 0 && ( (val/div) > 100 && (val/div) < 999 ) {
			return true
		}
	}
	return false
}
