//Package service defines and implements the actions available to be performed by the user.
package service

import (
	"fibonacci-api/domain"
	"fibonacci-api/dto"
	"fibonacci-api/errs"
	"math/big"
	"sync"
)

//FibService processes requests for new and existing sequences.
type FibService interface {
	NewSequence(dto.NewRequest, *sync.WaitGroup) (*dto.NewResponse, *errs.AppError)
	FindById(req dto.NewRequest) (*dto.NewResponse, *errs.AppError)
}

type DefaultFibService struct {
	Repo domain.FibRepository
}

// NewSequence takes in a NewRequest dto and passes the information to the domain in order to process.
func (service DefaultFibService) NewSequence(req dto.NewRequest, wg *sync.WaitGroup) (*dto.NewResponse, *errs.AppError) {
	sequence := domain.Sequence{
		Fib:      *big.NewInt(-1),
		Duration: -1,
		Algo:     req.Algorithm,
		Input:    req.Input,
		Status:   "incomplete",
	}
	newSequence, err := service.Repo.CalculateFib(sequence, wg)
	if err != nil {
		return nil, err
	}
	response := newSequence.ToNewResponseDto()
	return &response, nil
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
