//Package domain handles all the business logic (how sequences are created, stored, and changed)
package domain

import (
	"fibonacci-api/dto"
	"fibonacci-api/errs"
	"math/big"
	"sync"
)

type Sequence struct {
	Fib      big.Int
	Duration int64
	Algo     string
	Input    int
	Status   string
	Id       int64
}

//ToNewResponseDto takes a Sequence object and converts it into an appropriate response to the client.
func (sequence Sequence) ToNewResponseDto() dto.NewResponse {
	return dto.NewResponse{
		Fib:      sequence.Fib,
		Duration: sequence.Duration,
		Algo:     sequence.Algo,
		Input:    sequence.Input,
		Status:   sequence.Status,
		Id:       sequence.Id,
	}
}

//FibRepository defines the interface for calculating and retrieving Sequence objects.
type FibRepository interface {
	CalculateFib(Sequence, *sync.WaitGroup) (Sequence, *errs.AppError)
	FindBy(int64) (*Sequence, *errs.AppError)
}
