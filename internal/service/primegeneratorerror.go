package service

import "fmt"

type PrimeGeneratorError struct {
	target int
	err    error
}

func NewPrimeGeneratorError(target int, err error) *PrimeGeneratorError {
	return &PrimeGeneratorError{
		target: target,
		err:    err,
	}
}

func (pge *PrimeGeneratorError) Error() string {
	return fmt.Sprintf("pge: invalid input %d", pge.target)
}

func (pge *PrimeGeneratorError) Unwrap() error {
	return pge.err
}
