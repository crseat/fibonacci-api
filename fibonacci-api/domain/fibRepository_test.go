package domain

import "testing"

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
		t.Error("Invalid sequence returned by FindBy")
	}
}
