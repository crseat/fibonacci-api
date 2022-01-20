//Package domain handles all the business logic (how sequences are created, stored, and changed)
package domain

import (
	"fibonacci-api/dto"
	"fibonacci-api/errs"
	"sync"
	"time"
)

type Sequence struct {
	Fib      int64
	Duration time.Duration
	Algo     string
	Input    int
	Status   string
}

//ToNewResponseDto takes a Sequence object and converts it into an appropriate response to the client.
func (sequence Sequence) ToNewResponseDto() dto.NewResponse {
	return dto.NewResponse{
		Fib:      sequence.Fib,
		Duration: sequence.Duration,
		Algo:     sequence.Algo,
		Input:    sequence.Input,
		Status:   sequence.Status,
	}
}

//FibRepository defines the interface for calculating and retrieving Sequence objects.
type FibRepository interface {
	CalculateFib(input int, algo string, wg *sync.WaitGroup) (int64, *errs.AppError)
	FindBy(id int64) (*Sequence, *errs.AppError)
}
