package handlers

import (
	"fmt"
	"math"
)

func GeneratePrimes(input int) ([]int, error) {
	if input < 2 {
		return nil, fmt.Errorf("invalid input %d", input)
	}
	primes := make([]int, 0)
	for i := 2; i < input; i++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes, nil
}
