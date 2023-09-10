package main

import (
"fmt"
"strconv"
)

func isPalindrome(n int) bool {
str := strconv.Itoa(n)
length := len(str)
for i := 0; i < length/2; i++ {
if str[i] != str[length-i-1] {
return false
}
}
return true
}

func largestPalindromeProduct(start, end int) (int, int, int) {
largestPalindrome := 0
var multiplicand1, multiplicand2 int
for i := end; i >= start; i-- {
for j := i; j >= start; j-- {
product := i * j
if product < largestPalindrome {
break
}
if isPalindrome(product) && product > largestPalindrome {
largestPalindrome = product
multiplicand1 = i
multiplicand2 = j
}
}
}
return largestPalindrome, multiplicand1, multiplicand2
}

func main() {
var start, end int
fmt.Print("Enter the start value: ")
fmt.Scan(&start)
fmt.Print("Enter the end value: ")
fmt.Scan(&end)

result, multiplicand1, multiplicand2 := largestPalindromeProduct(start, end)
fmt.Printf("The largest palindrome product between %d and %d is: %d\n", start, end, result)
fmt.Printf("The multiplicands are: %d and %d\n", multiplicand1, multiplicand2)
}

// Output
// Enter the start value: 100
// Enter the end value: 999
// The largest palindrome product between 100 and 999 is: 906609
// The multiplicands are: 993 and 913