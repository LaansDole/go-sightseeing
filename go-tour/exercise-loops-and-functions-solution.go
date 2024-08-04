package main

import (
	"fmt"
	"strconv"
)

func TruncateDecimal(num float64) float64 {
	formattedNum := fmt.Sprintf("%.10f", num)
	truncatedNum, _ := strconv.ParseFloat(formattedNum, 64)

	return truncatedNum
}

func SqrtHelper(z, x float64) float64 {
	z -= (z*z - x) / (2 * z)
	return z
}

func Sqrt(x float64) float64 {
	z, sqrtX := 1.0, 1.0
	count := 0

	if z == x {
		return z
	}

	for count < 10 {
		z = SqrtHelper(z, x)
		if TruncateDecimal(sqrtX) == TruncateDecimal(z) {
			return sqrtX
		} else {
			sqrtX = z
		}
		fmt.Printf("Guessing Sqrt %v: %v\n", count, z)

		count += 1
	}

	return z
}

func main() {
	fmt.Println(Sqrt(100))
}
