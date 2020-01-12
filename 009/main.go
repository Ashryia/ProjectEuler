package main

import "fmt"


//a could equal 1
//b could equal 2 (a + 1)
//c could equal 997 ---- 1000 minus a + b

//Does a squared plus b squared equal c squared? No

//Increase B ++ , decrease C ++, check again until out of numbers.
//Then increase A ++, set b to A + 1, set C to (1000 - A + B), and loop again

func main() {

	a := 1
	b := a + 1
	c := 1000 - (a + b)

	for ; ((a + b) <= 1000) && ((a*a) + (b*b) < (c*c)) ; a ++ {
		for ; ((a + b) <= 1000) && ((a*a) + (b*b) < (c*c)) ; b ++ {
			c = 1000 - (a + b)
			sum := (a * a ) + ( b * b )
			cval := (c * c )
			if sum == cval {
				fmt.Printf("The value of A is %d, the value of B is %d, the value of C is %d", a, b, c)
				answer := (a * b * c )
				fmt.Printf("The answer is: %d.", answer)
				return
			}
		} 
		b = a + 1
		c = 1000 - (a + b)
	}
}
