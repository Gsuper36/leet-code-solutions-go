// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-321))
}

func reverse(x int) int {
	var digits []int
	reversed := 0
	negative := x < 0

	if negative {
		x = x * -1
	}

	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}
	rank := len(digits) - 1
	for i, digit := range digits {
		reversed += digit * int(math.Pow(10, float64(rank-i)))
	}

	if negative {
		reversed = reversed * -1
	}

	return reversed
}
