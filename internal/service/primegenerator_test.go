package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrimes_NegativeInput(t *testing.T) {
	input := -1
	expected := NewPrimeGeneratorError(input, nil)
	actual, err := GeneratePrimes(input)
	assert.Nil(t, actual)
	assert.EqualError(t, err, expected.Error())
}

func TestGeneratePrimes_ZeroInput(t *testing.T) {
	input := 0
	expected := NewPrimeGeneratorError(input, nil)
	actual, err := GeneratePrimes(input)
	assert.Nil(t, actual)
	assert.EqualError(t, err, expected.Error())
}

func TestGeneratePrimes_OneInput(t *testing.T) {
	input := 1
	expected := NewPrimeGeneratorError(input, nil)
	actual, err := GeneratePrimes(input)
	assert.Nil(t, actual)
	assert.EqualError(t, err, expected.Error())
}

func TestGeneratePrimes_TenInput(t *testing.T) {
	input := 10
	expected := []int{2, 3, 5, 7}
	actual, err := GeneratePrimes(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
