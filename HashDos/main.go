// man.go
//
// This program demonstrates and explains two design choices in Go's map
// implementation that help defend against Hash DoS (hash-based denial-of-service)
// attacks:
//
// 1. Each map uses a randomized hash seed.
// 2. The iteration order of a map (for-range over map) is deliberately made
//    to vary between iterations.
//
// These properties together make it much harder for an attacker to:
//   - Predict how keys will be placed into buckets.
//   - Rely on deterministic behavior to construct worst-case collision patterns.
//
// NOTE: The details of Go's runtime implementation are not part of the language
// spec and can change between versions. The important guarantees are:
//   - You must NOT rely on map iteration order.
//   - Map hashing is not stable or predictable across runs.

package main

import "fmt"

func main() {
	demoMapIterationOrder()
	explainConcepts()
}

// demoMapIterationOrder prints the order of keys for the same map multiple times.
// You will see that the order may differ between runs, and can even differ
// between loops in the same execution.
func demoMapIterationOrder() {
	fmt.Println("Demonstrating non-deterministic map iteration order:\n")

	m := map[string]int{
		"alpha":   1,
		"beta":    2,
		"gamma":   3,
		"delta":   4,
		"epsilon": 5,
	}

	// We iterate over the same map multiple times.
	// The Go runtime is allowed to (and typically does) vary the order.
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d: ", i)
		for k := range m {
			fmt.Printf("%s ", k)
		}
		fmt.Println()
	}

	fmt.Println("\nObserve that the key order is not guaranteed.")
	fmt.Println("You must not rely on any specific ordering when ranging over a map.\n")
}

// explainConcepts prints a textual explanation of what is going on and why
// Go behaves this way.
func explainConcepts() {
	fmt.Println("Explanation:")
	fmt.Println("------------")
	fmt.Println(`
1. Randomized hash per map
   -------------------------
   Internally, Go's map implementation uses a hash function to decide which
   bucket each key belongs to. To make targeted collision attacks harder,
   the runtime mixes in a random value (a hash seed) every time a map is
   created.

   That means:
     - The same set of keys may be distributed differently across maps.
     - The exact layout of the map in memory is not predictable from the
       outside.
     - An attacker cannot easily craft many keys that all collide into the
       same bucket in a controlled way.

   This design is one of the defenses against Hash DoS attacks, where an
   attacker sends a large number of specially chosen keys that cause
   excessive hash collisions and force the server to spend a lot of CPU
   time resolving them.

2. Non-deterministic iteration order
   ---------------------------------
   When you iterate over a map with:

       for k, v := range myMap {
           ...
       }

   Go does not guarantee any particular order for the keys. In fact, the
   runtime intentionally varies the iteration order between loops (and
   between program runs).

   This has two important consequences:
     - You should NEVER rely on map iteration order in your program logic.
     - It becomes harder for an attacker to exploit assumptions about
       predictable internal structure or ordering.

3. Why this helps against Hash DoS
   --------------------------------
   A Hash DoS (hash-based Denial of Service) attack tries to exploit the
   worst-case behavior of hash tables: when many different keys all end up
   in the same bucket, operations like lookup and insert degrade from
   average O(1) to O(n), dramatically increasing CPU usage.

   By:
     - Randomizing the hash seed per map, and
     - Varying iteration order,

   Go makes it much more difficult for an attacker to:
     - Predict exactly how keys will be placed into buckets.
     - Construct a stable, reproducible collision pattern across different
       runs or deployments.

   In practice, these properties raise the bar significantly for mounting
   effective Hash DoS attacks against Go programs using maps.`,
	)
}
