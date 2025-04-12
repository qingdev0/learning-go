// Creating a slice with make
package main

import (
	"fmt"
	"slices"
)

func main() {
	// 1. Uninitialized slice
	var s []string
	fmt.Printf("Uninitialized: %v, nil: %t, empty: %t\n", s, s == nil, len(s) == 0)

	// 2. Create slice with make
	s = make([]string, 3)
	fmt.Printf("Initialized with make: %v, len: %d, cap: %d\n", s, len(s), cap(s))

	// 3. Set values
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Printf("Set values: %v, len: %d, cap: %d\n", s, len(s), cap(s))

	// 4. Append elements (grows as needed)
	s = append(s, "d")
	fmt.Printf("Appended \"d\": %v, len: %d, cap: %d\n", s, len(s), cap(s))
	s = append(s, "e", "f")
	fmt.Printf("Appended \"e\", \"f\": %v, len: %d, cap: %d\n", s, len(s), cap(s))

	// 5. Copy slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Printf("Copied: %v, len: %d, cap: %d\n", c, len(c), cap(c))

	// 6. Slice operations
	fmt.Printf("Sliced 2:3: %v, len: %d, cap: %d\n", s[2:3], len(s[2:3]), cap(s[2:3]))
	fmt.Printf("Sliced 2:5: %v, len: %d, cap: %d\n", s[2:5], len(s[2:5]), cap(s[2:5]))
	fmt.Printf("Sliced 0:5: %v, len: %d, cap: %d\n", s[:5], len(s[:5]), cap(s[:5]))

	// 7. Initialize slice with values
	t1 := []string{"z", "y", "y"}
	fmt.Printf("Initialized with values: %v, len: %d, cap: %d\n", t1, len(t1), cap(t1))

	// 8. Compare slices
	t2 := []string{"y", "y", "z"}
	if slices.Equal(t1, t2) {
		fmt.Printf("Slices are equal: %v == %v\n", t1, t2)
	} else {
		fmt.Printf("Slices are not equal: %v != %v\n", t1, t2)
	}
}
