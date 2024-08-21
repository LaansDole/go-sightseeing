package main

import (
	"fmt"
	"strings"
)

func main() {
	humanV2(3)
}

func humanV2(n int) {
	headV2(n)
	bodyV2(n)
	legV2(n)
}

func headV2(n int) {
	// Create the head once and reuse it
	space := strings.Repeat(" ", 2*n)
	line := space + "*"
	for i := 0; i < n; i++ {
		fmt.Println(line)
	}
}

func bodyV2(n int) {
	// Create the body line once and print it
	mid := 2 * n
	line := make([]byte, 4*n+1)
	for i := 0; i < len(line); i++ {
		if i == mid {
			line[i] = '+'
		} else if i%2 == 0 {
			line[i] = '*'
		} else {
			line[i] = ' '
		}
	}
	fmt.Println(string(line))
}

func legV2(n int) {
	leftPos := 2*n - 2
	rightPos := 2*n + 2
	lineLength := 4*n + 1
	line := make([]byte, lineLength)

	// Initialize the line with spaces
	for j := range line {
		line[j] = ' '
	}

	// Update positions for each iteration and print
	for i := 0; i < n; i++ {
		line[leftPos] = '*'
		line[rightPos] = '*'
		fmt.Println(string(line))

		// Reset the positions for the next iteration
		line[leftPos] = ' '
		line[rightPos] = ' '

		// Move positions inward
		leftPos -= 2
		rightPos += 2
	}
}
