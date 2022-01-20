//Package dto implements the data transfer objects and validates the input
package dto

import (
	"fibonacci-api/errs"
)

type NewRequest struct {
	Algorithm string
	Id        int64
}

//Validate the algorithm that was passed in.
func (r NewRequest) Validate() *errs.AppError {
	panic("implement me")
}
