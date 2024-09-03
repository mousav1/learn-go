package main

import (
	"errors"
	"fmt"
)

type divisionError struct {
	err      string
	dividend float64
	divisor  float64
}

func (e *divisionError) Error() string {
	return fmt.Sprintf("cannot divide %f by %f: %s", e.dividend, e.divisor, e.err)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, &divisionError{
			err:      "division by zero",
			dividend: dividend,
			divisor:  divisor,
		}
	}
	return dividend / divisor, nil
}

func main() {
	dividend := 10.0
	divisor := 0.0

	result, err := divide(dividend, divisor)
	if err != nil {
		var divErr *divisionError
		if errors.As(err, &divErr) {
			fmt.Println("error:", divErr.err)
			return
		}
		fmt.Println(err)
		return
	}

	fmt.Printf("Result of division: %0.2f\n", result)
}
