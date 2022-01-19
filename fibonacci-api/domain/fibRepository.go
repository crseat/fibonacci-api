package domain

import (
	"fibonacci-api/errs"
	"sync/atomic"
	"time"
)

type FibRepositoryMap struct {
	Sequences map[int]Sequence
}

// Define and keep track of the sequence ids
//source: https://stackoverflow.com/questions/27917750/how-to-define-a-global-counter-in-golang-http-server
var id int64 = 0

// incId increments the number of the incrementing identifier and returns the new value
func incId() int64 {
	return atomic.AddInt64(&id, 1)
}

// getId returns the current value of the incrementing identifier
func getId() int64 {
	return atomic.LoadInt64(&id)
}

// UpdateFib updates the FibRepositoryMap after the sequence is done being calculated.
func (fibRepo FibRepositoryMap) UpdateFib(identifier int, fibNumber int64, duration time.Duration) {
	tempSequence := fibRepo.Sequences[identifier]
	tempSequence.Fib = fibNumber
	tempSequence.Duration = duration
	tempSequence.Status = "complete"
	fibRepo.Sequences[identifier] = tempSequence
}

// FindBy finds the sequence by its identifier and
func (fibRepo FibRepositoryMap) FindBy(identifier int) (*Sequence, *errs.AppError) {
	sequence := fibRepo.Sequences[identifier]
	response := Sequence{
		Fib:      sequence.Fib,
		Duration: sequence.Duration,
		Algo:     sequence.Algo,
		Input:    sequence.Input,
		Status:   sequence.Status,
	}
	return &response, nil
}

// CalculateFib takes in the input and the algorithm and delegates the work to calculate the sequence to the proper function
func (fibRepo FibRepositoryMap) CalculateFib(input int, algo string) (*Sequence, *errs.AppError) {
	panic("implement me")
}

// IterateFib uses dynamic programming to iteratively calculate the sequence.
func (fibRepo FibRepositoryMap) IterateFib(input int) (*Sequence, *errs.AppError) {
	panic("implement me")
}

// RecurseFib uses recursion and memoization to recursively calculate the sequence.
func (fibRepo FibRepositoryMap) RecurseFib(input int) (*Sequence, *errs.AppError) {
	panic("implement me")
}

//MathFib uses the golden ratio (Binet's formula) to calculate the nth fibonacci number.
func (fibRepo FibRepositoryMap) MathFib(input int) (*Sequence, *errs.AppError) {
	panic("implement me")
}

// NewFibRepository creates a new FibRepo.
func NewFibRepository() FibRepositoryMap {
	return FibRepositoryMap{
		Sequences: make(map[int]Sequence),
	}
}
