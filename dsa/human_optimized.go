package main

import (
	"fmt"
	"strings"
)

func main() {
	humanV2(4)
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
	// Efficiently create and print the leg lines
	for i := 0; i < n; i++ {
		leftPos := 2*n - 2 - 2*i
		rightPos := 2*n + 2 + 2*i
		line := make([]byte, 4*n+1)
		for j := range line {
			if j == leftPos || j == rightPos {
				line[j] = '*'
			} else {
				line[j] = ' '
			}
		}
		fmt.Println(string(line))
	}
}
