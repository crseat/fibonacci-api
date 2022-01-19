//Package service defines and implements the actions available to be performed by the user.
package service

import (
	"fibonacci-api/domain"
	"fibonacci-api/dto"
	"fibonacci-api/errs"
	"sync"
	"time"
)

//FibService processes requests for new and existing sequences.
type FibService interface {
	NewSequence(dto.NewRequest, time.Time, *sync.WaitGroup) (*dto.NewResponse, *errs.AppError)
	FindById(req dto.NewRequest) (*dto.NewResponse, *errs.AppError)
}

type DefaultFibService struct {
	Repo domain.FibRepository
}

// NewSequence takes in a NewRequest dto and passes the information to the domain in order to process.
func (service DefaultFibService) NewSequence(req dto.NewRequest, startTime time.Time, wg *sync.WaitGroup) (*dto.NewResponse, *errs.AppError) {
	panic("implement me")
}

// FindById takes in a NewRequest, queries the repo for the sequence using the corresponding id, and then converts
// response into NewResponse
func (service DefaultFibService) FindById(req dto.NewRequest) (*dto.NewResponse, *errs.AppError) {
	targetSequence, err := service.Repo.FindBy(req.Id)
	if err != nil {
		return nil, err
	}
	response := targetSequence.ToNewResponseDto()
	return &response, nil
}

// NewFibonacciService creates new DefaultFibService using the passed in fibRepo
func NewFibonacciService(fibRepository domain.FibRepository) DefaultFibService {
	return DefaultFibService{fibRepository}
}
