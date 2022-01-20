package domain

import (
	"fibonacci-api/errs"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

type FibRepositoryMap struct {
	Sequences map[int64]Sequence
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
func (fibRepo FibRepositoryMap) UpdateFib(identifier int64, fibNumber int64, duration time.Duration) {
	tempSequence := fibRepo.Sequences[identifier]
	tempSequence.Fib = fibNumber
	tempSequence.Duration = duration
	tempSequence.Status = "complete"
	fibRepo.Sequences[identifier] = tempSequence
}

// FindBy finds the sequence by its identifier and
func (fibRepo FibRepositoryMap) FindBy(identifier int64) (*Sequence, *errs.AppError) {
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
func (fibRepo FibRepositoryMap) CalculateFib(input int, algo string, wg *sync.WaitGroup) (int64, *errs.AppError) {
	incId()
	sequenceId := getId()
	newSequence := Sequence{
		Fib:      -1,
		Duration: -1,
		Algo:     algo,
		Input:    input,
		Status:   "incomplete",
	}
	fibRepo.Sequences[sequenceId] = newSequence
	if algo == "iterate" {
		go fibRepo.IterateFib(int64(input), wg)
	}
	if algo == "recurse" {
		var numMap = map[int64]int64{0: 0, 1: 1}
		tempInput := input - 1
		go fibRepo.RecurseFib(int64(tempInput), numMap, wg)
	}
	if algo == "math" {
		go fibRepo.MathFib(int64(input), wg)
	}
	return sequenceId, nil
}

// IterateFib uses dynamic programming to iteratively calculate the sequence.
func (fibRepo FibRepositoryMap) IterateFib(input int64, wg *sync.WaitGroup) (int64, *errs.AppError) {
	// Ensures graceful shutdown
	wg.Add(1)
	defer wg.Done()
	if input <= 1 {
		return input, nil
	}
	//Create our cache
	numDict := make([]int64, input+1)
	numDict[1] = 1
	//Build result iteratively
	for i := int64(2); i < input; i++ {
		numDict[i] = numDict[i-1] + numDict[i-2]
	}
	return numDict[input-1], nil
}

// RecurseFib uses recursion and memoization to recursively calculate the sequence.
func (fibRepo FibRepositoryMap) RecurseFib(input int64, numMap map[int64]int64, wg *sync.WaitGroup) int64 {
	// Ensures graceful shutdown
	wg.Add(1)
	defer wg.Done()
	//memoization
	value, keyPresent := numMap[input]
	if keyPresent {
		return value
	}
	//recurse
	numMap[input] = fibRepo.RecurseFib(input-1, numMap, wg) + fibRepo.RecurseFib(input-2, numMap, wg)
	return numMap[input]
}

//MathFib uses the golden ratio (Binet's formula) to calculate the nth fibonacci number.
func (fibRepo FibRepositoryMap) MathFib(input int64, wg *sync.WaitGroup) (int64, *errs.AppError) {
	// Ensures graceful shutdown
	wg.Add(1)
	defer wg.Done()
	//need to adjust for zero indexing
	input -= 1
	//Define the golden ratio
	var goldenRatio = float64(1+math.Sqrt(5)) / 2
	//raise the golen ratio by our input and devide by sqrt 5
	unrounded := math.Pow(goldenRatio, float64(input)) / math.Sqrt(5)
	// Round and convert our answer
	answer := int64(math.Round(unrounded))
	return answer, nil
}

// NewFibRepository creates a new FibRepo.
func NewFibRepository() FibRepositoryMap {
	return FibRepositoryMap{
		Sequences: make(map[int64]Sequence),
	}
}
