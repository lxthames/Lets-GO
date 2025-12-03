package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {

	// 1. Slice literal (no size specified)
	x := []int{10, 20, 30}
	fmt.Println("1) slice literal x:", x)

	// 2. Sparse slice literal
	y := []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println("2) sparse slice literal y:", y)
	// Produces: [1 0 0 0 0 4 6 0 0 0 100 15]

	// 3. Slice of slices (multidimensional slice)
	var grid [][]int
	grid = append(grid, []int{1, 2, 3})
	grid = append(grid, []int{4, 5, 6})
	fmt.Println("3) slice of slices grid:", grid)

	// 4. Creating a slice without literal → zero value is nil
	var s []int
	fmt.Println("4) s is nil?", s == nil)
	fmt.Println("   s:", s)

	// 5. Indexing slices (same rules as arrays)
	nums := []int{100, 200, 300}
	fmt.Println("5) nums[0]:", nums[0])
	// fmt.Println(nums[3])  // panics: index out of range

	// 6. Slices are NOT comparable
	// The following would NOT compile:
	// fmt.Println(nums == x)

	// 7. Comparing slices using slices.Equal (Go 1.21+)
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	c := []int{1, 2, 3, 4, 5, 6}
	str := []string{"a", "b", "c"}
	_ = str
	fmt.Println("7) slices.Equal(a, b):", slices.Equal(a, b)) // true
	fmt.Println("   slices.Equal(a, c):", slices.Equal(a, c)) // false

	// This does NOT compile (different element types)
	// fmt.Println(slices.Equal(a, str))

	// 8. slices.EqualFunc example (custom comparison)
	people1 := []string{"Bob", "alice", "JOHN"}
	people2 := []string{"bob", "Alice", "john"}

	equalCaseInsensitive := slices.EqualFunc(people1, people2,
		func(a, b string) bool {
			return normalize(a) == normalize(b)
		})

	fmt.Println("8) EqualFunc (case-insensitive):", equalCaseInsensitive)
}

func normalize(s string) string {
	return strings.ToLower(s)
}
