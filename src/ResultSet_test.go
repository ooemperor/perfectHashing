package src

import (
	"testing"
)

type MockResultSet struct {
	entries []*MockResultEntry
}

func (set *MockResultSet) GetHashArray() ([1024]*MockResultEntry, error) {
	var output [1024]*MockResultEntry

	for _, resultSet := range set.entries {
		_, err := resultSet.GetPosition()
		if err != nil {
			return output, err
		}
		output[0] = resultSet
	}
	return output, nil
}

/*
TestResultSetGetHashArray checks if we receive a proper output for the TestResultSetGetHashArray
TODO: change the content of the test to a proper content after the implementation
*/
func TestResultSetGetHashArray(t *testing.T) {
	rEntry := MockResultEntry{value: "test"}
	var entries []*MockResultEntry

	entries = append(entries, &rEntry)

	rSet := MockResultSet{entries: entries}

	result, err := rSet.GetHashArray()

	if err != nil {
		t.Fatalf("Error %v occured when retreiving the HashArray", err)
	}

	if len(result) != 1024 {
		t.Fatalf("GetHashArray result does have an incorrect length")
	}
}
