package main

import (
	"fmt"
	"math"
)

func main() {
	var a, vo, so, t float64
	fmt.Printf("Please, enter aceleration\n")
	fmt.Scan(&a)
	fmt.Printf("Please, enter initial velocity\n")
	fmt.Scan(&vo)
	fmt.Printf("Please, enter initial displacement\n")
	fmt.Scan(&so)
	fn := GenDisplaceFn(a, vo, so)

	fmt.Printf("Please, enter time\n")
	fmt.Scan(&t)
	fmt.Println("The displacement after", t, "seconds:", fn(t))
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return ((1.0/2.0)*a*(math.Pow(t, 2)) + vo*t + so)
	}
}
