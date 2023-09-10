package main

import "fmt"

func main() {
var prev, current, sum, limit int
prev, current, sum = 0, 1, 0

fmt.Print("Enter the upper limit of the Fibonacci sequence: ")
fmt.Scan(&limit)

for prev+current < limit {
next := prev + current
fmt.Println(next)
if next%2 == 0 {
sum += next
}
prev, current = current, next
}

fmt.Println("The sum of the even-valued terms is", sum)
}

// //Output
// Enter the upper limit of the Fibonacci sequence: 40
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
// The sum of the even-valued terms is 44