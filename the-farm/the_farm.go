package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood returns the amount of fodder per cow.
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	amount, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return amount * factor / float64(cows), nil
}

// ValidateInputAndDivideFood checks that the number of cows is positive,
// then returns the result of DivideFood.
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows < 1 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// ValidateNumberOfCows returns an InvalidCowsError when passed a non-positive
// number of cows.
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	}
	if cows == 0 {
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	}
	return nil
}
