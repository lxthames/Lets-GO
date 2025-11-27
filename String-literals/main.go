package main

import (
	"fmt"
)

func main() {
	// 1. Basic interpreted string
	msg1 := "Hello, world!"
	fmt.Println("1.", msg1)

	// 2. Interpreted string with newline + double quotes
	msg2 := "Greetings and\n\"Salutations\""
	fmt.Println("2.", msg2)

	// 3. Unicode escape in interpreted string
	msg3 := "I \u2665 Go"
	fmt.Println("3.", msg3)

	// 4. Valid escape for double quotes (single quote escape is not allowed)
	msg4 := "\"This is a double quote inside a string\""
	fmt.Println("4.", msg4)

	// 5. Raw string literal (multi-line, no escapes processed)
	msg5 := `Greetings and
"Salutations"
This is line 3\nThis is not a newline — it's literal`
	fmt.Println("5.", msg5)

	// 6. Raw string for Windows file paths
	pathRaw := `C:\Users\John\Documents\file.txt`
	fmt.Println("6. Raw path:", pathRaw)

	// Interpreted version (requires escaping backslashes)
	pathInterp := "C:\\Users\\John\\Documents\\file.txt"
	fmt.Println("   Interpreted path:", pathInterp)

	// 7. Raw string for regex (clean)
	regexRaw := `^\d{3}-\d{2}-\d{4}$`
	fmt.Println("7. Raw regex:", regexRaw)

	// Interpreted regex (messy due to escaping)
	regexInterp := "^\\d{3}-\\d{2}-\\d{4}$"
	fmt.Println("   Interpreted regex:", regexInterp)
}
