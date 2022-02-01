package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	nbCows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.nbCows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	var fodderPerCow float64 = 0
	fodder, err := weightFodder.FodderAmount()

	if err != nil && err == ErrScaleMalfunction && fodder > 0 {
		fodderPerCow = fodder * 2 / float64(cows)
		return fodderPerCow, nil
	}

	if err != nil && err == ErrScaleMalfunction && fodder < 0 {
		return 0, errors.New("negative fodder")
	}

	if err != nil {
		return 0, err
	}

	if fodder < 0 {
		return 0, errors.New("negative fodder")
	}

	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	if cows < 0 {
		newErr := SillyNephewError{nbCows: cows}
		return 0, &newErr
	}

	fodderPerCow = fodder / float64(cows)

	return fodderPerCow, nil
}
