package main

import "fmt"

func main() {
	// 1. Basic array declaration (fixed size, zero-initialized)
	var a [3]int
	fmt.Println("1) basic array a:", a) // [0 0 0]

	// 2. Array literal with initial values
	var b = [3]int{10, 20, 30}
	fmt.Println("2) initialized array b:", b)

	// 3. Sparse array literal
	// This creates: [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
	var c = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println("3) sparse array c:", c)

	// 4. Using ... to let Go infer the length
	var d = [...]int{10, 20, 30}
	fmt.Println("4) inferred-length array d:", d, "len(d):", len(d))

	// 5. Array comparison with == and !=
	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println("5) x:", x, "y:", y, "x == y?", x == y)

	// NOTE: The following would NOT compile because lengths are different:
	// var z = [4]int{1, 2, 3, 4}
	// fmt.Println(x == z) // compile-time error: mismatched types [3]int and [4]int

	// 6. Simulating multidimensional arrays
	var m [2][3]int
	m[0][0] = 1
	m[0][1] = 2
	m[0][2] = 3
	m[1][0] = 4
	m[1][1] = 5
	m[1][2] = 6
	fmt.Println("6) 2x3 multidimensional array m:", m)

	// 7. Reading and writing elements (indexing)
	var e = [3]int{}
	e[0] = 100
	e[1] = 200
	e[2] = 300
	fmt.Println("7) array e:", e)
	fmt.Println("   e[0]:", e[0], "e[1]:", e[1], "e[2]:", e[2])

	// Out-of-bounds access:
	// e[3] = 999      // compile-time error or panic at runtime, depending on index form
	// fmt.Println(e[-1]) // compile-time error: invalid index

	// 8. len() with arrays
	fmt.Println("8) len(e):", len(e))

	// 9. Array length is part of the type
	var arr3 [3]int
	// var arr4 [4]int
	fmt.Printf("9) type of arr3 is %T\n", arr3)
	// fmt.Printf("type of arr4 is %T\n", arr4)

	// You cannot assign arrays of different lengths to each other:
	// arr3 = arr4 // compile-time error: cannot use arr4 (type [4]int) as type [3]int

	// 10. You cannot use a variable for array length (must be a constant)
	const size = 5
	var fixed = [size]int{1, 2, 3, 4, 5}
	fmt.Println("10) array with const size:", fixed)

	// This would NOT compile:
	// n := 5
	// var bad = [n]int{} // error: non-constant array bound

	// 11. Functions and array types (no generic “any length” array function)
	fmt.Println("11) passing arrays to functions:")
	demoArrayParam([3]int{1, 2, 3})
	// demoArrayParam([4]int{1, 2, 3, 4}) // compile-time error: mismatched types [4]int and [3]int

	// This shows that a function taking [3]int CANNOT accept [4]int, even though both are "arrays of int".

	// 1-column, 3-row array

	var z [3][1]int

	z[0][0] = 10
	z[1][0] = 20
	z[2][0] = 30

	fmt.Println("12)1-column, 3-row array:", z)
}

func demoArrayParam(a [3]int) {
	fmt.Println("    demoArrayParam got:", a)
}
