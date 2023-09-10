//Special pythagorean triplet

package main

import "fmt"

func findSpecialPythagoreanTriplet(n int) (int, int, int) {
	for a := 1; a <= n/3; a++ {
		for b := a; b <= (n-a)/2; b++ {
			c := n - a - b
			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return -1, -1, -1
}

func main() {
	var n int
	fmt.Print("Enter the sum value: ")
	fmt.Scan(&n)
	a, b, c := findSpecialPythagoreanTriplet(n)
	if a != -1 && b != -1 && c != -1 {
		fmt.Printf("The Special Pythagorean Triplet for sum %d is (%d, %d, %d)\n", n, a, b, c)
	} else {
		fmt.Printf("No Special Pythagorean Triplet found for sum %d\n", n)
	}
}

// Output
// Enter the sum value: 150
// The Special Pythagorean Triplet for sum 150 is (25, 60, 65)
