package src

import (
	"hash/crc32"
	"testing"
)

type MockResultEntry struct {
	value string
}

func (entry *MockResultEntry) GetPosition() (uint32, error) {
	crc32q := crc32.MakeTable(0xD5828281)
	return crc32.Checksum([]byte("Hello world"), crc32q), nil
}

/*
TestIResultEntryHash tests if the hashing returns the expected result
TODO: change the content of the test to a proper content after the implementation
*/
func TestResultEntryGetPosition(t *testing.T) {
	rEntry := MockResultEntry{value: "test"}
	result, err := rEntry.GetPosition()

	if err != nil {
		t.Fatalf("Error %v occured when fetching the position", err.Error())
	}

	if result != 694472804 {
		t.Fatalf("Result does not match the expected 694472804 with result: %v", result)
	}
}
