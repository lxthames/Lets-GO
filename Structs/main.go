package main

import (
	"fmt"
)

// Named struct type
type person struct {
	name string
	age  int
	pet  string
}

// Types for comparison / conversion examples
type firstPerson struct {
	name string
	age  int
}

type secondPerson struct {
	name string
	age  int
}

type thirdPerson struct {
	age  int
	name string
}

type fourthPerson struct {
	firstName string
	age       int
}

type fifthPerson struct {
	name          string
	age           int
	favoriteColor string
}

func main() {
	demoZeroValueAndLiterals()
	demoFieldAccess()
	demoAnonymousStructs()
	demoCompareAndConvert()
}

// ---------------------------------------------------
// 1. Zero value and struct literals
// ---------------------------------------------------

func demoZeroValueAndLiterals() {
	fmt.Println("== demoZeroValueAndLiterals ==")

	// Zero value struct (all fields are zero value)
	var fred person
	fmt.Println("fred (zero value):", fred)

	// Zero value via empty literal (same as above)
	bob := person{}
	fmt.Println("bob (empty literal):", bob)

	// Positional literal (values in field declaration order)
	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println("julia (positional literal):", julia)

	// Keyed literal (field: value), can be in any order, can omit fields
	beth := person{
		age:  30,
		name: "Beth",
		// pet omitted → zero value ""
	}
	fmt.Println("beth (keyed literal):", beth)

	fmt.Println()
}

// ---------------------------------------------------
// 2. Field access with dot notation
// ---------------------------------------------------

func demoFieldAccess() {
	fmt.Println("== demoFieldAccess ==")

	p := person{}
	p.name = "Alice"
	p.age = 35
	p.pet = "parrot"

	fmt.Println("p:", p)
	fmt.Println("p.name:", p.name)
	fmt.Println("p.age:", p.age)
	fmt.Println("p.pet:", p.pet)

	fmt.Println()
}

// ---------------------------------------------------
// 3. Anonymous structs
// ---------------------------------------------------

func demoAnonymousStructs() {
	fmt.Println("== demoAnonymousStructs ==")

	// Anonymous struct type declared as a variable
	var anonPerson struct {
		name string
		age  int
		pet  string
	}

	anonPerson.name = "Charlie"
	anonPerson.age = 28
	anonPerson.pet = "dog"

	fmt.Println("anonPerson:", anonPerson)

	// Anonymous struct with literal
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println("anonymous pet struct:", pet)

	fmt.Println()
}

// ---------------------------------------------------
// 4. Comparing and converting struct types
// ---------------------------------------------------

func demoCompareAndConvert() {
	fmt.Println("== demoCompareAndConvert ==")

	f := firstPerson{
		name: "Bob",
		age:  50,
	}

	s := secondPerson{
		name: "Bob",
		age:  50,
	}

	fmt.Println("f (firstPerson):", f)
	fmt.Println("s (secondPerson):", s)

	// This does NOT compile (different named types):
	// fmt.Println(f == s)

	// But we CAN convert between firstPerson and secondPerson
	// because they have the same field names, order, and types.
	s2 := secondPerson(f)
	fmt.Println("s2 (secondPerson converted from f):", s2)

	// Conversions that are NOT allowed (examples shown as comments):

	// 1) Different field order:
	// type thirdPerson struct {
	//     age  int
	//     name string
	// }
	// t := thirdPerson(f) // compile error: cannot convert

	// 2) Different field names:
	// type fourthPerson struct {
	//     firstName string
	//     age       int
	// }
	// four := fourthPerson(f) // compile error: cannot convert

	// 3) Extra field:
	// type fifthPerson struct {
	//     name          string
	//     age           int
	//     favoriteColor string
	// }
	// five := fifthPerson(f) // compile error: cannot convert

	// --- Named struct vs anonymous struct ---

	var g struct {
		name string
		age  int
	}

	// Because g's fields (names, order, types) match firstPerson,
	// we can assign directly and compare with ==.
	g = f
	fmt.Println("g (anonymous struct):", g)
	fmt.Println("f == g ?", f == g)

	fmt.Println()
}
