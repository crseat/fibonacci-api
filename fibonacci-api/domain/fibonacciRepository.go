package domain

import (
	"fibonacci-api/errs"
	"sync/atomic"
)

type FibRepositoryMap struct {
	Fib      *int
	Duration *int64
	Algo     *string
	Input    *int
	Status   *string
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
func (FibRepo FibRepositoryMap) UpdateFib() {
	panic("implement me")
}

// FindBy finds the sequence by its identifier and
func (FibRepo FibRepositoryMap) FindBy(id int) (*Sequence, *errs.AppError) {
	panic("implement me")
}

// CalculateFib takes in the input and the algorithm and delegates the work to calculate the sequence to the proper function
func (FibRepo FibRepositoryMap) CalculateFib(input int, algo string) (*Sequence, *errs.AppError) {
	panic("implement me")
}

// IterateFib uses dynamic programming to iteratively calculate the sequence.
func (FibRepo FibRepositoryMap) IterateFib(input int) (*Sequence, *errs.AppError) {
	panic("implement me")
}

// RecurseFib uses recursion and memoization to recursively calculate the sequence.
func (FibRepo FibRepositoryMap) RecurseFib(input int) (*Sequence, *errs.AppError) {
	panic("implement me")
}

//MathFib uses the golden ratio (Binet's formula) to calculate the nth fibonacci number.
func (FibRepo FibRepositoryMap) MathFib(input int) (*Sequence, *errs.AppError) {
	panic("implement me")
}

// NewFibRepository creates a new FibRepo.
func NewFibRepository() FibRepositoryMap {
	return FibRepositoryMap{
		Fib:      new(int),
		Duration: new(int64),
		Algo:     new(string),
		Input:    new(int),
		Status:   new(string),
	}
}
