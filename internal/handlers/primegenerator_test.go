package handlers

import (
	"reflect"
	"testing"
)

func TestGeneratePrimes_NegativeInput(t *testing.T) {
	input := -1
	expected := "invalid input -1"
	actual, err := GeneratePrimes(input)
	if actual != nil {
		t.Errorf("actual should be nil")
	}
	if expected != err.Error() {
		t.Errorf("actual %s", err)
	}
}

func TestGeneratePrimes_ZeroInput(t *testing.T) {
	input := 0
	expected := "invalid input 0"
	actual, err := GeneratePrimes(input)
	if actual != nil {
		t.Errorf("actual should be nil")
	}
	if expected != err.Error() {
		t.Errorf("actual %s", err)
	}
}

func TestGeneratePrimes_OneInput(t *testing.T) {
	input := 1
	expected := "invalid input 1"
	actual, err := GeneratePrimes(input)
	if actual != nil {
		t.Errorf("actual should be nil")
	}
	if expected != err.Error() {
		t.Errorf("actual %s", err)
	}
}

func TestGeneratePrimes_TenInput(t *testing.T) {
	input := 10
	expected := []int{2, 3, 5, 7}
	actual, err := GeneratePrimes(input)
	if err != nil {
		t.Errorf("error should be nil")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v", actual)
	}
}
