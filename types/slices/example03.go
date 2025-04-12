// Creating and manipulating 2D slices
package main

import "fmt"

func main() {
	// Create a 2D slice with outer length of 3
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Printf("2D slice: %v\n", twoD)
	for i, inner := range twoD {
		fmt.Printf("  Row %d: %v, len: %d, cap: %d\n", i, inner, len(inner), cap(inner))
	}
}
