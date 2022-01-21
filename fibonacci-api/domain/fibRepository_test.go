package domain

import (
	"math/big"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestFibRepositoryMap_UpdateFib(t *testing.T) {
	repo := NewFibRepository()
	testSequence := Sequence{
		Fib:      *big.NewInt(-1),
		Duration: -1,
		Algo:     "math",
		Input:    3,
		Status:   "incomplete",
	}

	// Set then update sequence.
	repo.Sequences[1] = testSequence
	repo.UpdateFib(1, *big.NewInt(2), 99)

	// Check to make sure it's been updated correctly
	updatedSequence := repo.Sequences[1]
	comparisonSequence := Sequence{
		Fib:      *big.NewInt(2),
		Duration: 99,
		Algo:     "math",
		Input:    3,
		Status:   "complete",
	}

	if reflect.DeepEqual(updatedSequence, comparisonSequence) {
		t.Error("Sequence was not updated correctly")
	}
}

func TestFibRepositoryMap_FindBy(t *testing.T) {
	repo := NewFibRepository()
	testSequence := Sequence{
		Fib:      *big.NewInt(2),
		Duration: 99,
		Algo:     "math",
		Input:    3,
		Status:   "complete",
	}

	// Set then find sequence.
	repo.Sequences[4] = testSequence
	foundSequence, err := repo.FindBy(4)
	if err != nil {
		t.Error("Error was returned while finding sequence: ", err)
		return
	}
	if reflect.DeepEqual(foundSequence, testSequence) {
		t.Error("Invalid sequence returned by FindBy Want:", testSequence, "Got:", *foundSequence)
	}
}

func TestFibRepositoryMap_CalculateFib(t *testing.T) {
	wg := &sync.WaitGroup{}
	//CalculateFib(input int, algo string)
	repo := NewFibRepository()
	sequence := Sequence{
		Fib:      *big.NewInt(-1),
		Duration: -1,
		Algo:     "math",
		Input:    13,
		Status:   "incomplete",
		Id:       -1,
	}
	response, err := repo.CalculateFib(sequence, wg)
	if err != nil {
		t.Error("Error was returned while calling CalculateFib: ", err)
		return
	}
	if response.Id != 1 {
		t.Error("Invalid identifier returned from CalculateFib. Want:", 1, "Got:", id)
	}
}

func TestFibRepositoryMap_IterateFib(t *testing.T) {
	wg := &sync.WaitGroup{}
	repo := NewFibRepository()
	sequence := Sequence{
		Fib:      *big.NewInt(-1),
		Duration: -1,
		Algo:     "iterate",
		Input:    65,
		Status:   "incomplete",
		Id:       -1,
	}
	result, err := repo.CalculateFib(sequence, wg)
	if err != nil {
		t.Error("Error was returned while calling IterateFib: ", err)
		return
	}
	//Need to sleep to give go routine time to finish
	time.Sleep(1 * time.Second)
	updatedSequence, err := repo.FindBy(result.Id)
	if updatedSequence.Fib.Cmp(big.NewInt(10610209857723)) != 0 {
		t.Error("Invalid result for iterateFib calculation. Want:", 10610209857723, "Got:", result)
	}
}

func TestFibRepositoryMap_MathFib(t *testing.T) {
	wg := &sync.WaitGroup{}
	repo := NewFibRepository()
	sequence := Sequence{
		Fib:      *big.NewInt(-1),
		Duration: -1,
		Algo:     "math",
		Input:    65,
		Status:   "incomplete",
		Id:       -1,
	}
	result, err := repo.CalculateFib(sequence, wg)
	if err != nil {
		t.Error("Error was returned while calling MathFib: ", err)
		return
	}
	//Need to sleep to give go routine time to finish
	time.Sleep(1 * time.Second)
	updatedSequence, err := repo.FindBy(result.Id)
	if updatedSequence.Fib.Cmp(big.NewInt(10610209857723)) != 0 {
		t.Error("Invalid result for mathFib calculation. Want:", 10610209857723, "Got:", updatedSequence.Fib)
	}
}

func TestFibRepositoryMap_RecurseFib(t *testing.T) {
	wg := &sync.WaitGroup{}
	repo := NewFibRepository()
	sequence := Sequence{
		Fib:      *big.NewInt(-1),
		Duration: -1,
		Algo:     "recursive",
		Input:    65,
		Status:   "incomplete",
		Id:       -1,
	}
	result, err := repo.CalculateFib(sequence, wg)
	if err != nil {
		t.Error("Error was returned while calling MathFib: ", err)
		return
	}
	//Need to sleep to give go routine time to finish
	time.Sleep(1 * time.Second)
	updatedSequence, err := repo.FindBy(result.Id)
	if updatedSequence.Fib.Cmp(big.NewInt(10610209857723)) != 0 {
		t.Error("Invalid result for mathFib calculation. Want:", 10610209857723, "Got:", updatedSequence.Fib)
	}
}
