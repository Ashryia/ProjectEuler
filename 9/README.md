# Special Pythagorean triplet
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which, 
a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000. 
Find the product abc. 

# Brainstorming

a < b < c

a^2 + b^2 = c^2 

a + b + c = 1000

What is a / b / c? 

How would I do this brute forcey? 

a could equal 1
b could equal 2 (a + 1)
c could equal 997 ---- 1000 minus a + b

Does a squared plus b squared equal c squared? No

Increase B ++ , decrease C ++, check again until out of numbers.
Then increase A ++, set b to A + 1, set C to (1000 - A + B), and loop again

Does 
(a * a ) + ( b * b ) = (c * c) ? No

Increase B through all available options
