package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for i :=0; i < 10 ; i++ {
		z -= (z*z - x) / (2*z)
	} 

	return z
}

func main() {
	fmt.Printf("Square root of %d is %.5f \n", 1, Sqrt(1))
	fmt.Printf("Square root of %d is %.5f \n", 2, Sqrt(2))
	fmt.Printf("Square root of %d is %.5f \n", 3, Sqrt(3))
	fmt.Printf("Square root of %d is %.5f \n", 4, Sqrt(4))
	fmt.Printf("Square root of %d is %.5f \n", 5, Sqrt(5))
	fmt.Printf("Square root of %d is %.5f \n", 6, Sqrt(6))
	fmt.Printf("Square root of %d is %.5f \n", 7, Sqrt(7))
	fmt.Printf("Square root of %d is %.5f \n", 8, Sqrt(8))
	fmt.Printf("Square root of %d is %.5f \n", 9, Sqrt(9))
	fmt.Printf("Square root of %d is %.5f \n", 10, Sqrt(10))
	fmt.Printf("Square root of %d is %.5f \n", 25, Sqrt(25))
}
