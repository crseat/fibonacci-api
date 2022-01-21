package domain

import (
	"fibonacci-api/errs"
	"math"
	"math/big"
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
func (fibRepo FibRepositoryMap) UpdateFib(identifier int64, fibNumber big.Int, duration time.Duration) {
	tempSequence := fibRepo.Sequences[identifier]
	tempSequence.Fib = fibNumber
	tempSequence.Duration = duration.Microseconds()
	tempSequence.Status = "complete"
	fibRepo.Sequences[identifier] = tempSequence
}

// FindBy finds the sequence by its identifier and
func (fibRepo FibRepositoryMap) FindBy(identifier int64) (*Sequence, *errs.AppError) {
	sequence, sequencePresent := fibRepo.Sequences[identifier]
	if !sequencePresent {
		appError := errs.NewValidationError("There is no sequence with the given identifier")
		return nil, appError
	}
	response := Sequence{
		Fib:      sequence.Fib,
		Duration: sequence.Duration,
		Algo:     sequence.Algo,
		Input:    sequence.Input,
		Status:   sequence.Status,
		Id:       sequence.Id,
	}
	return &response, nil
}

// CalculateFib takes in the input and the algorithm and delegates the work to calculate the sequence to the proper function
func (fibRepo FibRepositoryMap) CalculateFib(sequence Sequence, wg *sync.WaitGroup) (Sequence, *errs.AppError) {
	incId()
	sequenceId := getId()
	sequence.Id = sequenceId
	fibRepo.Sequences[sequenceId] = sequence
	startTime := time.Now()
	if sequence.Algo == "iterate" {
		go fibRepo.IterateFib(int64(sequence.Input), sequenceId, startTime, wg)
	}
	if sequence.Algo == "recursive" {
		var numMap = map[int64]*big.Int{0: big.NewInt(0), 1: big.NewInt(1)}
		//account for zero indexing
		tempInput := sequence.Input - 1
		go fibRepo.RecurseFib(int64(tempInput), numMap, startTime, 1, wg)
	}
	if sequence.Algo == "math" {
		go fibRepo.MathFib(int64(sequence.Input), startTime, wg)
	}
	return sequence, nil
}

// IterateFib uses dynamic programming to iteratively calculate the sequence.
func (fibRepo FibRepositoryMap) IterateFib(input int64, id int64, startTime time.Time, wg *sync.WaitGroup) *errs.AppError {
	// Ensures graceful shutdown
	wg.Add(1)
	defer wg.Done()
	if input <= 1 {
		duration := time.Now().Sub(startTime)
		answer := *big.NewInt(input)
		fibRepo.UpdateFib(id, answer, duration)
		return nil
	}
	//Create our cache
	numDict := make([]big.Int, input+1)
	numDict[1] = *big.NewInt(1)
	//Build result iteratively
	for i := int64(2); i < input; i++ {
		numDict[i] = *new(big.Int).Add(&numDict[i-1], &numDict[i-2])
	}
	//Find time taken and update repo
	duration := time.Now().Sub(startTime)
	fibRepo.UpdateFib(id, numDict[input-1], duration)
	return nil
}

// RecurseFib uses recursion and memoization to recursively calculate the sequence.
func (fibRepo FibRepositoryMap) RecurseFib(input int64, numMap map[int64]*big.Int, startTime time.Time, recursionDepth int, wg *sync.WaitGroup) *big.Int {
	// Ensures graceful shutdown
	wg.Add(1)
	defer wg.Done()
	//memoization
	value, keyPresent := numMap[input]
	if keyPresent {
		return value
	}
	//recurse
	numMap[input] = new(big.Int).Add(fibRepo.RecurseFib(input-1, numMap, startTime, recursionDepth+1, wg), fibRepo.RecurseFib(input-2, numMap, startTime, recursionDepth+1, wg))
	//Find time taken and update repo
	if recursionDepth == 1 {
		duration := time.Since(startTime)
		fibRepo.UpdateFib(id, *numMap[input], duration)
	}
	return numMap[input]
}

//MathFib uses the golden ratio (Binet's formula) to calculate the nth fibonacci number.
func (fibRepo FibRepositoryMap) MathFib(input int64, startTime time.Time, wg *sync.WaitGroup) *errs.AppError {
	// Ensures graceful shutdown
	wg.Add(1)
	defer wg.Done()
	//need to adjust for zero indexing
	input -= 1
	//Define the golden ratio
	divisor := big.NewFloat(2)
	dividend := big.NewFloat(1 + math.Sqrt(5))
	var goldenRatio = new(big.Float).Quo(dividend, divisor)
	//raise the golden ratio by our input and divide by sqrt 5
	unrounded := new(big.Float).Quo(Pow(goldenRatio, uint64(input)), new(big.Float).Sqrt(big.NewFloat(5)))
	// Round and convert our answer
	delta := 0.5
	unrounded.Add(unrounded, big.NewFloat(delta))
	rounded, _ := unrounded.Int(nil)
	answer := rounded
	//Find time taken and update repo
	duration := time.Since(startTime)
	fibRepo.UpdateFib(id, *answer, duration)
	return nil
}

// NewFibRepository creates a new FibRepo.
func NewFibRepository() FibRepositoryMap {
	return FibRepositoryMap{
		Sequences: make(map[int64]Sequence),
	}
}

//Source: https://steemit.com/tutorial/@gopher23/power-and-root-functions-using-big-float-in-golang
func Pow(a *big.Float, e uint64) *big.Float {
	result := Zero().Copy(a)
	for i := uint64(0); i < e-1; i++ {
		result = Mul(result, a)
	}
	return result
}

func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

func Mul(a, b *big.Float) *big.Float {
	return Zero().Mul(a, b)
}
