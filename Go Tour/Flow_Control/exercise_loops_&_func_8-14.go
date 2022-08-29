/*
	Que link : https://go.dev/tour/flowcontrol/8
*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// In here I have used the initial guess as x.
	z := x * 1.0
	// I have made the basic assumption of 10 iterations.
	fmt.Println("")
	zr := 1.0
	tc := 0
	for i := 0; i < 20; i++ {
		//fmt.Printf("For %2dth iteration, approximated value is %g\n",i+1,z)
		z -= (z*z - x) / (2 * z)
		ratio := math.Pow(10, float64(6))
		// Assumed a threshold value of 3.
		if zr-(math.Round(z*ratio)/ratio) < 0.000000001 && tc > 3 {
			fmt.Printf("Breaking the loop at %2dth iteration, with approximated value is %g\n", i+1, z)
			break
		} else if zr-(math.Round(z*ratio)/ratio) < 0.000000001 {
			tc++
		}
		zr = math.Round(z*ratio) / ratio
	}
	fmt.Println("")
	return z
}

func exercise_8_14() {
	i := 1
	for i = 1; i <= 2000; i++ {
		fmt.Printf("Sqrt(%2d) = %g\n", i, Sqrt(float64(i)))
	}
}
