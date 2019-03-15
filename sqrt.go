package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x > 0 {
		z := 1.0
		// First guess
		z -= (z*z - x) / (2 * z)
		// Iterate until change is very small
		for zNew, delta := z, z; delta > 0.00000001; z = zNew {
			zNew -= (zNew*zNew - x) / (2 * zNew)
			delta = z - zNew
		}
		return z
	} else if x == float64(0) {
		return x
	} else {
		return math.NaN()
	}
}
func main() {
	fmt.Println(Sqrt(10))
	fmt.Println(math.Sqrt(10))
}
