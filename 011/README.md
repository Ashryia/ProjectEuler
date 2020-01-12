#Largest product in a grid

In the 20x20 grid below, four numbers along a diagonal line have been marked in red. 
08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70
67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21
24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72
21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95
78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48

Numbers are 26, 63, 78, 14. It's row 7, 8, 9, 10. 

The product of these numbers is 26 x 63 x 78 x 14 = 1788696.

What is the greatest product of four adjacent numbers in the same direction (up, down, left, right, or diagonally) in the 20x20 grid. 

# Brainstorming

If I were doing this manually, I would go through in each direction: right, up, diagonal down right, diagonal down left. Up doesnt matter, left doesn't matter. 

How do I store this grid? array of arrays. 20 arrays each containing 20 values. 
structure: {{1,2,3,4},{1,2,3,4},{1,2,3,4}}

Right:
Access value 1 as: struct[0][0]
Access value 2 as: struct[0][1]
Access value 3 as: struct[0][2]
Access value 4 as: struct[0][3]

I can either take the entire grid of [20][20]int , or I can take it a row at a time. I think I should take the whole grid. 
for each element in the array, take that element and grab the next four. If four more do not exist, go to the next row until we're out of rows. 

While there are rows, 
for each row, takes the first value, then the next three values and multiply them. 


---------------------------------

Down:
Access value 1 as: struct[0][0]
Access value 2 as: struct[1][0]
Access value 3 as: struct[2][0]
Access value 4 as: struct[3][0]


---------------------------------
Diagonal down:
start at struct[20][0]
Try to go down (a + 1)(b + 1) . Each value must be under 20 
21/1 - not a thing, so go up. 
19/0 then 20,1 then 21,2 - not a thing so go up
18,0 then 19,1 then 20,2 then 21,3 - not a thing so go up
17,0 then 18,1 then 19,2 then 20,3 - YAY it's four valid values. 
Move up the grid until you get to 0,0. Then we move across the grid. 


20/0, 21/1, 22/2, 23/3 -- EXIT
19/0, 20/1, 21/2       -- EXIT
18/0, 19/1, 20/2,      -- EXIT
17/0, 18/1, 19/2, 20/3 -- yes
16/0, 17/1, 18/2, 19/3 
17/1, 18/2, 19/3, 20/4 --
15/0, 16/1, 17/2, 18/3
16/1, 17/2, 18/3, 19/4
17/2, 18/3, 19/4, 20/5

if a + 4 



---------------------------------
Diagonal up:
start at struct[0][0]

0/0, -1/2           -- EXIT
1/0, 0/1, -1/2      -- EXIT
2/0, 1/1, 0/2, -1/3 -- EXIT
3/0, 2/1, 1/2, 0/3
4/0, 3/1, 2/2, 1/3     ---yes
3/1, 2/2, 1/3, 0/4 
2/2, 1/3, 0/4, -1/5 -- EXIT


a goes dow and needs to be above 0
b goes up and needs to be below 20
