package main

import (
	"fmt"
)

func main() {
	demoUint8AndInt32()
	demoByteAndRune()
	demoStringBytesVsRunes()
	demoIterateString()
}

// ----------------------------------------
// 1. uint8 vs int32
// ----------------------------------------

func demoUint8AndInt32() {
	fmt.Println("== demoUint8AndInt32 ==")

	var u uint8 // 8-bit unsigned, 0 to 255
	var i int32 // 32-bit signed, -2147483648 to 2147483647

	u = 0
	fmt.Println("uint8 min:", u)

	u = 255
	fmt.Println("uint8 max:", u)

	// i can hold negative and positive values
	i = -2147483648
	fmt.Println("int32 min:", i)

	i = 2147483647
	fmt.Println("int32 max:", i)

	// Uncommenting these will NOT compile (overflow):
	// u = 256       // too big for uint8
	// i = 2147483648 // too big for int32

	fmt.Println()
}

// ----------------------------------------
// 2. byte vs rune
// ----------------------------------------

func demoByteAndRune() {
	fmt.Println("== demoByteAndRune ==")

	var b byte // alias for uint8 (raw byte)
	var r rune // alias for int32 (Unicode code point)

	b = 'A' // ASCII, one byte
	r = '你' // Chinese character, one rune (multiple bytes in UTF-8)

	fmt.Printf("byte b = %v, as char = %c\n", b, b)
	fmt.Printf("rune r = %v, as char = %c\n", r, r)

	fmt.Println("byte is an alias for uint8, rune is an alias for int32")
	fmt.Println()
}

// ----------------------------------------
// 3. String: bytes vs runes
// ----------------------------------------

func demoStringBytesVsRunes() {
	fmt.Println("== demoStringBytesVsRunes ==")

	s := "Hello ☀️"

	fmt.Println("string:", s)
	fmt.Println("len(s) (bytes):", len(s))

	bs := []byte(s)
	rs := []rune(s)

	fmt.Println("[]byte(s):", bs)
	fmt.Println("len([]byte(s)):", len(bs))

	fmt.Println("[]rune(s):", rs)
	fmt.Println("len([]rune(s)) (runes / code points):", len(rs))

	fmt.Println()
}

// ----------------------------------------
// 4. Iterating string: by bytes vs by runes
// ----------------------------------------

func demoIterateString() {
	fmt.Println("== demoIterateString ==")

	s := "Go🙂語"

	fmt.Println("string:", s)
	fmt.Println("-- byte-based loop (wrong for characters) --")
	for i := 0; i < len(s); i++ {
		fmt.Printf("i=%d byte=%d char=%q\n", i, s[i], s[i])
	}

	fmt.Println("-- rune-based loop (correct for characters) --")
	for i, r := range s {
		fmt.Printf("byteIndex=%d rune=%d char=%q\n", i, r, r)
	}

	fmt.Println()
}
