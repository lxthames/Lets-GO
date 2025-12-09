package main

import (
	"fmt"
)

func main() {
	// 1) make([]int, 5) → length 5, capacity 5
	x := make([]int, 5)
	fmt.Println("1) x := make([]int, 5)")
	printSlice("x", x)
	fmt.Println("   x[0]..x[4] are valid and all zero-initialized\n")

	// Beginner mistake: thinking append will "fill" those 5 slots
	fmt.Println("2) Append to x (length 5) with x = append(x, 10)")
	x = append(x, 10)
	printSlice("x after append", x)
	fmt.Println("   Notice: 10 is added AFTER the 5 zeros, not replacing them.")
	fmt.Println("   Now len=6, cap likely doubled (10 in this example)\n")

	// 2) make with length 5 and capacity 10
	y := make([]int, 5, 10)
	fmt.Println("3) y := make([]int, 5, 10)")
	printSlice("y", y)
	fmt.Println("   y has 5 elements (all zero), but room (capacity) for 10.")
	fmt.Println("   You can index y[0]..y[4], and append up to 5 more without realloc.\n")

	// Show appending to y
	y = append(y, 1, 2, 3)
	printSlice("y after append 1,2,3", y)
	fmt.Println("   Still same backing array until length exceeds capacity.\n")

	// 3) make with length 0 and capacity 10
	z := make([]int, 0, 10)
	fmt.Println("4) z := make([]int, 0, 10)")
	printSlice("z", z)
	fmt.Println("   z is non-nil, len=0, cap=10. You CANNOT do z[0] yet (out of range).")
	fmt.Println("   But you CAN append safely without realloc.\n")

	// Append to z
	z = append(z, 5, 6, 7, 8)
	printSlice("z after append 5,6,7,8", z)
	fmt.Println("   Now len=4, cap still 10. You used 4 of the reserved slots.\n")

	// 4) Invalid example (in comments) – capacity < length
	fmt.Println("5) Invalid make examples (won't compile, shown as comments):")
	fmt.Println(`   // bad := make([]int, 5, 3)   // ❌ compile-time error`)
	fmt.Println("   If you somehow pass a smaller capacity via variable, it will panic at runtime.\n")
}

// Helper to print slice content, length, and capacity
func printSlice(name string, s []int) {
	fmt.Printf("   %s = %v  len=%d  cap=%d\n", name, s, len(s), cap(s))
}
