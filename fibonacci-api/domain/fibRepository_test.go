package domain

import (
	"sync"
	"testing"
)

func TestFibRepositoryMap_UpdateFib(t *testing.T) {
	repo := NewFibRepository()
	testSequence := Sequence{
		Fib:      -1,
		Duration: -1,
		Algo:     "math",
		Input:    3,
		Status:   "incomplete",
	}

	// Set then update sequence.
	repo.Sequences[1] = testSequence
	repo.UpdateFib(1, 2, 99)

	// Check to make sure it's been updated correctly
	updatedSequence := repo.Sequences[1]
	comparisonSequence := Sequence{
		Fib:      2,
		Duration: 99,
		Algo:     "math",
		Input:    3,
		Status:   "complete",
	}

	if updatedSequence != comparisonSequence {
		t.Error("Sequence was not updated correctly")
	}
}

func TestFibRepositoryMap_FindBy(t *testing.T) {
	repo := NewFibRepository()
	testSequence := Sequence{
		Fib:      2,
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
	if *foundSequence != testSequence {
		t.Error("Invalid sequence returned by FindBy Want:", testSequence, "Got:", *foundSequence)
	}
}

func TestFibRepositoryMap_CalculateFib(t *testing.T) {
	wg := &sync.WaitGroup{}
	//CalculateFib(input int, algo string)
	repo := NewFibRepository()
	id, err := repo.CalculateFib(13, "", wg)
	if err != nil {
		t.Error("Error was returned while calling CalculateFib: ", err)
		return
	}
	if id != 1 {
		t.Error("Invalid identifier returned from CalculateFib. Want:", 1, "Got:", id)
	}
}

func TestFibRepositoryMap_IterateFib(t *testing.T) {
	wg := &sync.WaitGroup{}
	repo := NewFibRepository()
	testInput := int64(13)
	result, err := repo.IterateFib(testInput, wg)
	if err != nil {
		t.Error("Error was returned while calling IterateFib: ", err)
		return
	}
	if result != 144 {
		t.Error("Invalid result for fib calculation. Want:", 144, "Got:", result)
	}
}

func TestFibRepositoryMap_MathFib(t *testing.T) {
	wg := &sync.WaitGroup{}
	repo := NewFibRepository()
	testInput := int64(13)
	result, err := repo.MathFib(testInput, wg)
	if err != nil {
		t.Error("Error was returned while calling MathFib: ", err)
		return
	}
	if result != 144 {
		t.Error("Invalid result for fib calculation, Want:", 144, "Got:", result)
	}
}

func TestFibRepositoryMap_RecurseFib(t *testing.T) {
	wg := &sync.WaitGroup{}
	repo := NewFibRepository()
	testInput := int64(13)
	// Adjust for 0 indexing
	tempInput := testInput - 1
	var numMap = map[int64]int64{0: 0, 1: 1}
	result := repo.RecurseFib(tempInput, numMap, wg)
	if result != 144 {
		t.Error("Invalid result for fib calculation. Want:", 144, "Got:", result)
	}
}
