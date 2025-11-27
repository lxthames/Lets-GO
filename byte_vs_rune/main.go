package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// 1. Using byte for ASCII characters
	var c byte = 'A'
	fmt.Println("1. Byte example (ASCII only):")
	fmt.Println("Byte value:", c)
	fmt.Printf("As character: %c\n\n", c)

	// 2. Using rune for Unicode characters
	var r rune = '你'
	fmt.Println("2. Rune example (Unicode):")
	fmt.Println("Rune value:", r)
	fmt.Printf("As character: %c\n\n", r)

	// 3. Why byte is wrong for Unicode (multi-byte UTF-8)
	fmt.Println("3. Byte slice of \"你\":")
	b := []byte("你")
	fmt.Println("Byte slice:", b) // multiple bytes
	fmt.Println()

	// Correct usage: rune slice
	fmt.Println("Correct rune slice of \"你\":")
	rs := []rune("你")
	fmt.Println("Rune slice:", rs)
	fmt.Printf("As character: %c\n\n", rs[0])

	// 4. Iterating over a string using bytes (WRONG for Unicode)
	s := "Hello 你好"
	fmt.Println("4. Incorrect iteration using bytes:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) // breaks Unicode
	}
	fmt.Println("\n")

	// Correct iteration using runes
	fmt.Println("Correct iteration using runes:")
	for _, rr := range s {
		fmt.Printf("%c ", rr)
	}
	fmt.Println("\n")

	// 5. Byte slice vs rune slice
	fmt.Println("5. Byte slice vs Rune slice:")
	data := []byte{0x48, 0x49, 0x50}
	fmt.Println("Byte slice (raw data):", data)

	text := []rune("Hello 世界")
	fmt.Println("Rune slice (text characters):", text)
	fmt.Printf("Characters: ")
	for _, ch := range text {
		fmt.Printf("%c ", ch)
	}
	fmt.Println("\n")

	// 6. Length differences
	str := "你好"
	fmt.Println("6. Length differences for \"你好\":")
	fmt.Println("Byte length:", len(str))
	fmt.Println("Rune length:", utf8.RuneCountInString(str))
}
