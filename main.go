package main

import "fmt"

func factorial(f int) float64 {
	fac := 1
	for i:= 1; i < f+1; i++ {
		fac = fac * i
	}
	return float64(fac)
}

func main()  {
	var euler float64
	var limit int

	fmt.Scan(&limit)

	for i:= 0; i <= limit; i++ {
		euler += 1/factorial(i)
	}

	fmt.Print(euler)
}
