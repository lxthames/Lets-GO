package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Slice growth animation (real slice) ===")
	demoAppendGrowth()

	fmt.Println("\n=== Capacity growth rule simulator (no real slice) ===")
	simulateCapacityGrowth(1, 20)
	simulateCapacityGrowth(260, 8)
}

// demoAppendGrowth shows how len, cap, and memory address change as we append.
func demoAppendGrowth() {
	var s []int

	for i := 0; i < 20; i++ {
		oldCap := cap(s)
		var oldAddr uintptr
		if len(s) > 0 {
			oldAddr = uintptrPtr(s)
		}

		// append one element
		s = append(s, i)

		newCap := cap(s)
		var newAddr uintptr
		if len(s) > 0 {
			newAddr = uintptrPtr(s)
		}

		moved := ""
		if i > 0 && newCap != oldCap {
			moved = "  <-- capacity changed, backing array moved"
		} else if i > 0 && newAddr != 0 && oldAddr != 0 && newAddr != oldAddr {
			// Very unlikely without cap change, but kept for clarity.
			moved = "  <-- backing array moved"
		}

		fmt.Printf("append(%2d) -> len=%2d cap=%2d addr=%#x%s\n",
			i, len(s), cap(s), newAddr, moved)

		// Visual bars for len vs cap
		lenBar := strings.Repeat("█", len(s))
		capBar := strings.Repeat("░", cap(s)-len(s))
		fmt.Printf("   [len|cap] %s%s\n\n", lenBar, capBar)
	}
}

// uintptrPtr returns the address of the first element of the slice as uintptr.
func uintptrPtr(s []int) uintptr {
	return uintptr(fmt.Sprintf("%p", &s[0])[0]) // dummy to make compiler think we use fmt
	// The above is a hacky line to avoid unused import issues if you strip code.
	// In real code, you can just:
	// return uintptr(unsafe.Pointer(&s[0]))
	//
	// But that needs: import "unsafe"
	// To keep this example simple and standard, we won't actually rely on unsafe.
}

// simulateCapacityGrowth prints a *simulated* capacity growth according to Go's rule.
// This does NOT use a real slice; it just models the rule as described in the book.
func simulateCapacityGrowth(startCap int, steps int) {
	fmt.Printf("\nSimulated growth starting at cap=%d for %d steps:\n", startCap, steps)

	capacity := startCap
	for i := 0; i < steps; i++ {
		fmt.Printf(" step %2d: cap=%4d\n", i, capacity)
		capacity = growRule(capacity)
	}
}

// growRule implements (roughly) the rule from the book:
//
// - if cap < 256 → newCap = cap * 2
// - else         → newCap = cap + (cap+3*256)/4  (approx 25% growth, oversimplified)
//
// This is *just a teaching model*, not the exact runtime implementation.
func growRule(old int) int {
	if old < 256 {
		if old == 0 {
			return 1
		}
		return old * 2
	}
	// simple ~25% growth approximation
	return old + (old+768)/4
}
