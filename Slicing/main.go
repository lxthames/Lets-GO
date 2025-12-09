package main

import (
	"fmt"
)

func main() {
	demoBasicSlicing()
	demoSharedStorage()
	demoAppendWithSubslice()
	demoFullSliceExpression()
	demoCopy()
	demoArraySliceConversion()
	demoStringByteRune()
}

// ----------------------------------------
// 1. Basic slicing
// ----------------------------------------

func demoBasicSlicing() {
	fmt.Println("== demoBasicSlicing ==")

	x := []string{"a", "b", "c", "d"}

	y := x[:2]  // from 0 to 2: ["a", "b"]
	z := x[1:]  // from 1 to end: ["b", "c", "d"]
	d := x[1:3] // from 1 to 3: ["b", "c"]
	e := x[:]   // full slice: ["a", "b", "c", "d"]

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println()
}

// ----------------------------------------
// 2. Shared storage when slicing
// ----------------------------------------

func demoSharedStorage() {
	fmt.Println("== demoSharedStorage ==")

	x := []string{"a", "b", "c", "d"}
	y := x[:2] // ["a", "b"]
	z := x[1:] // ["b", "c", "d"]

	// All of these modify the same backing array.
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"

	fmt.Println("x:", x) // [x y z d]
	fmt.Println("y:", y) // [x y]
	fmt.Println("z:", z) // [y z d]
	fmt.Println()
}

// ----------------------------------------
// 3. append + subslices (shared capacity)
// ----------------------------------------

func demoAppendWithSubslice() {
	fmt.Println("== demoAppendWithSubslice ==")

	x := []string{"a", "b", "c", "d"} // len=4, cap=4
	y := x[:2]                        // len=2, cap=4

	fmt.Println("cap(x), cap(y):", cap(x), cap(y))

	// y has capacity 4, so appending will still write into
	// x's underlying array (no reallocation yet).
	y = append(y, "z")

	fmt.Println("x:", x) // underlying "c" becomes "z"
	fmt.Println("y:", y) // ["a", "b", "z"]
	fmt.Println()
}

// ----------------------------------------
// 4. Full slice expression to limit capacity
// ----------------------------------------

func demoFullSliceExpression() {
	fmt.Println("== demoFullSliceExpression ==")

	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d") // len=4, cap=5

	// Full slice expression:
	// y := x[start:end:capLimit]
	// len(y) = end - start
	// cap(y) = capLimit - start

	y := x[:2:2]  // start=0, end=2, capLimit=2 -> len=2, cap=2
	z := x[2:4:4] // start=2, end=4, capLimit=4 -> len=2, cap=2

	fmt.Println("x:", x)
	fmt.Println("len/cap y:", len(y), cap(y))
	fmt.Println("len/cap z:", len(z), cap(z))

	// These appends will *force* new backing arrays,
	// because y and z are already at full capacity.
	y = append(y, "i", "j", "k")
	z = append(z, "y")

	fmt.Println("after appends:")
	fmt.Println("x:", x) // still uses original backing array
	fmt.Println("y:", y) // separate backing array now
	fmt.Println("z:", z) // separate backing array now
	fmt.Println()
}

// ----------------------------------------
// 5. copy: making independent slices
// ----------------------------------------

func demoCopy() {
	fmt.Println("== demoCopy ==")

	x := []int{1, 2, 3, 4}
	y := make([]int, 4)

	n := copy(y, x) // copy as many as min(len(y), len(x))
	fmt.Println("y:", y, "copied:", n)

	// Copy subset
	y2 := make([]int, 2)
	n2 := copy(y2, x) // copy first 2
	fmt.Println("y2:", y2, "copied:", n2)

	// Copy from the middle
	y3 := make([]int, 2)
	copy(y3, x[2:]) // copy {3, 4}
	fmt.Println("y3:", y3)

	// Overlapping copy (still allowed)
	x2 := []int{1, 2, 3, 4}
	n3 := copy(x2[:3], x2[1:])            // copy [2,3,4] over first 3 positions
	fmt.Println("x2:", x2, "copied:", n3) // [2 3 4 4] 3
	fmt.Println()
}

// ----------------------------------------
// 6. Arrays <-> slices
// ----------------------------------------

func demoArraySliceConversion() {
	fmt.Println("== demoArraySliceConversion ==")

	// Array to slice shares memory
	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:] // full slice
	y := xArray[:2]
	z := xArray[2:]

	fmt.Println("xArray:", xArray)
	fmt.Println("xSlice:", xSlice)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	xArray[0] = 10
	fmt.Println("after modifying xArray[0] = 10")
	fmt.Println("xArray:", xArray)
	fmt.Println("xSlice:", xSlice)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// Slice to array creates a new array (copies)
	s := []int{1, 2, 3, 4}
	a := [4]int(s) // full copy
	small := [2]int(s)

	s[0] = 99

	fmt.Println("slice s:", s)
	fmt.Println("array a:", a)
	fmt.Println("array small:", small)

	// NOTE: converting to bigger array size than len(s) will panic at runtime
	// so we do NOT do something like: big := [5]int(s)
	fmt.Println()
}

// ----------------------------------------
// 7. Strings, bytes, runes, UTF-8
// ----------------------------------------

func demoStringByteRune() {
	fmt.Println("== demoStringByteRune ==")

	s := "Hello there"
	b := s[6] // single byte (0-based)

	fmt.Println("s:", s)
	fmt.Println("byte at index 6:", b, "as char:", string(b))

	// Slicing strings (byte-based indices)
	s2 := s[4:7] // "o t"
	s3 := s[:5]  // "Hello"
	s4 := s[6:]  // "there"

	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	// Example with emoji (multi-byte rune)
	sEmoji := "Hello ☀️"
	fmt.Println("sEmoji:", sEmoji)
	fmt.Println("len(sEmoji):", len(sEmoji), "(bytes)")

	// Converting string to []byte and []rune
	bs := []byte(sEmoji)
	rs := []rune(sEmoji)

	fmt.Println("[]byte(sEmoji):", bs)
	fmt.Println("[]rune(sEmoji):", rs)

	// rune/byte to string
	var r rune = 'x'
	var c byte = 'y'
	fmt.Println("string(rune('x')):", string(r))
	fmt.Println("string(byte('y')):", string(c))

	// Common bug: int -> string (code point, not digits)
	var num int = 65
	strFromInt := string(num)
	fmt.Println("string(65):", strFromInt, "(this is 'A', not \"65\")")

	fmt.Println()
}
