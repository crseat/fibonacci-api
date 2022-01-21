//Package dto implements the data transfer objects and validates the input
package dto

import (
	"fibonacci-api/errs"
	"strconv"
)

type NewRequest struct {
	Algorithm string
	Input     int
	Id        int64
}

//ValidateInputNum validates and converts the input number that was passed in.
func (r NewRequest) ValidateInputNum(num string) (int, *errs.AppError) {
	inputNum, err := strconv.Atoi(num)
	if err != nil || inputNum < 1 || inputNum >= 100000 {
		appError := errs.NewValidationError("Please provide a valid input number. (Numbers greater than 0 and less than 100000 only)")
		return 0, appError
	}
	return inputNum, nil
}

//ValidateAlgo validates the algorithm that was passed in.
func (r NewRequest) ValidateAlgo() *errs.AppError {
	if r.Algorithm != "math" && r.Algorithm != "recursive" && r.Algorithm != "iterate" {
		return errs.NewValidationError("Please provide valid algorithm: math, recursive, iterate. Got: " + r.Algorithm)
	}
	return nil
}

//ValidateId validates and converts the identifier that was passed in.
func (r NewRequest) ValidateId(id string) (int64, *errs.AppError) {
	fibId, err := strconv.Atoi(id)
	if err != nil || fibId < 1 {
		appError := errs.NewValidationError("Please provide valid identifier. (Numbers greater than 0 only)")
		return 0, appError
	}
	return int64(fibId), nil
}
