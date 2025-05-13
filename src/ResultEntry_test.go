package src

import (
	"testing"
)

/*
TestIResultEntryHash tests if the hashing returns the expected result
*/
func TestResultEntryGetPosition(t *testing.T) {
	rEntry := ResultEntry{Value: "test"}
	result, err := rEntry.GetPosition()

	if err != nil {
		t.Fatalf("Error %v occured when fetching the position", err.Error())
	}

	if result != 3704742031 {
		t.Fatalf("Result does not match the expected 3704742031 with result: %v", result)
	}
}
