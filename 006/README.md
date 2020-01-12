# Sum square difference

The sum of the squares of the first ten natural numbers is 385. 
(1^2) + (2^2) + ... + (10^2) = 385

The square of the sum of the first en natural numbers is 3025.

(1 + 2 + ... + 10 )^2 

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is:
3025 - 385 = 2640. 

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum. 

## Brainstorming

First, get the "square of the sum". This is the bigger number. Use a for loop to get the sum. Then square it. 
Next, get the "sum of the squares". This is the smaller number. Use a for loop to square each number, then sum it. 

Could use the same for loop. 
For value, while value is less than max value, increment value {
	squareofsum += value
	sumofSquare = sumofSquare + (value * value)
}
