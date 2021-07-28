package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPrimeGeneratorError_Error(t *testing.T) {
	actual := &PrimeGeneratorError{
		target: rand.Int(),
	}
	expected := fmt.Sprintf("pge: invalid input %d", actual.target)
	assert.EqualError(t, actual, expected)
}

func TestPrimeGeneratorError_Unwrap(t *testing.T) {
	expected := fmt.Errorf("wraped error")
	actual := &PrimeGeneratorError{
		err: expected,
	}
	assert.EqualError(t, actual.Unwrap(), expected.Error())
}

func TestPrimeGeneratorError_New(t *testing.T) {
	target := rand.Int()
	wrappedErr := fmt.Errorf("wraped error")
	expectedError := fmt.Sprintf("pge: invalid input %d", target)
	actual := NewPrimeGeneratorError(target, wrappedErr)
	assert.EqualError(t, actual, expectedError)
	assert.EqualError(t, actual.Unwrap(), wrappedErr.Error())
}
