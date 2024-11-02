package main

import (
	"fmt"
	"time"
)

// Calculates layup sequence in O(n)
// Memoizes any seen result to avoid exponential growth while retaining
// the simplicity of a naturally recursive solution
func layupSequence(n int, mem map[int]int) int {
	if elem, ok := mem[n]; ok {
		return elem
	} else {
		// Minor optimization here, instead of running:
		// mem[n] = layupSequence(n-1, mem) +/- layupSequence(n-2, mem)
		// Checks for memoized value at each step which reduces iterations in half
		if _, ok := mem[n-1]; !ok {
			mem[n-1] = layupSequence(n-1, mem)
		}
		if _, ok := mem[n-2]; !ok {
			mem[n-2] = layupSequence(n-2, mem)
		}
		if n%2 == 0 {
			mem[n] = mem[n-1] + mem[n-2]
		} else {
			mem[n] = mem[n-1] - mem[n-2]
		}
	}
	return mem[n]
}

// Runs the function at intervals of multiples of 10 to generate data for the graph
// Includes solution for N:10000 Result:-7655024931378849993
func main() {
	for i := 1; i <= 1000000; i *= 10 {
		start := time.Now()
		var mem = map[int]int{
			1: 1,
			2: 2,
		}
		ans := layupSequence(i, mem)
		elapsed := time.Since(start)

		fmt.Printf(
			"Runtime:%f\nN:%d\nResult:%d\n\n",
			elapsed.Seconds(),
			i,
			ans,
		)
	}
}
