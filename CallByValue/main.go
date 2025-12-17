package main

import "fmt"

// 1) Passing a nil *int: you cannot make the caller's pointer non-nil
func failedUpdateNil(g *int) {
	x := 10
	g = &x // only updates the local copy of the pointer
}

// To change the caller's pointer variable, pass **int (pointer-to-pointer)
func setPtr(pp **int) {
	x := 10
	*pp = &x // updates the caller's pointer
}

// 2) Reassigning a *int parameter does NOT affect the caller's variable
func failedReassign(px *int) {
	x2 := 20
	px = &x2 // only changes local copy of the pointer
}

// Dereferencing DOES update the caller's value
func updateValue(px *int) {
	*px = 20
}

// 3) Maps: mutations are visible to the caller
func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

// 4) Slices: element updates visible; append does not update caller's len unless returned
func modSliceNoReturn(s []int) {
	for i := range s {
		s[i] *= 2
	}
	s = append(s, 99) // may write into shared backing array, but caller's len won't change
}

// Correct pattern: return the updated slice header
func modSliceReturn(s []int) []int {
	for i := range s {
		s[i] *= 2
	}
	s = append(s, 99)
	return s
}

func main() {
	fmt.Println("=== Pointer: nil cannot be made non-nil via *T parameter ===")
	var f *int // nil
	failedUpdateNil(f)
	fmt.Println("after failedUpdateNil, f =", f)

	setPtr(&f)
	fmt.Println("after setPtr(&f), f =", *f)

	fmt.Println("\n=== Pointer: reassignment vs dereference ===")
	x := 10
	failedReassign(&x)
	fmt.Println("after failedReassign(&x), x =", x) // still 10

	updateValue(&x)
	fmt.Println("after updateValue(&x), x =", x) // now 20

	fmt.Println("\n=== Map: mutations are visible ===")
	m := map[int]string{1: "first", 2: "second"}
	modMap(m)
	fmt.Println("map after modMap:", m)

	fmt.Println("\n=== Slice: elements change, but append doesn't change caller's len (unless returned) ===")
	// Make a slice with extra capacity so append can stay in the same backing array.
	s := make([]int, 3, 5)
	s[0], s[1], s[2] = 1, 2, 3
	fmt.Printf("before: s=%v len=%d cap=%d\n", s, len(s), cap(s))

	modSliceNoReturn(s)
	fmt.Printf("after modSliceNoReturn: s=%v len=%d cap=%d\n", s, len(s), cap(s))

	// The appended value may exist in the backing array but is beyond the caller's len.
	fmt.Printf("resliced to cap (shows backing array): %v\n", s[:cap(s)])

	// Correct: return the new slice header if you need the length update.
	s2 := []int{1, 2, 3}
	fmt.Printf("\nbefore: s2=%v len=%d cap=%d\n", s2, len(s2), cap(s2))
	s2 = modSliceReturn(s2)
	fmt.Printf("after modSliceReturn: s2=%v len=%d cap=%d\n", s2, len(s2), cap(s2))
}
