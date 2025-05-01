package src

import (
	"fmt"
	"testing"
)

type MockResultSet struct {
	entries []*MockResultEntry
}

func (set *MockResultSet) GetHashArray() ([1024]*MockResultEntry, error) {
	var output [1024]*MockResultEntry

	if len(set.entries) == 0 {
		// no entries found, return an error
		return output, fmt.Errorf("the ResultSet does not have any entries")
	}

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

/*
TestResultSetGetHashArrayEmptyInput checks if we receive an expected error for the TestResultSetGetHashArrayEmptyInput
TODO: change the content of the test to a proper content after the implementation
*/
func TestResultSetGetHashArrayEmptyInput(t *testing.T) {
	var entries []*MockResultEntry
	rSet := MockResultSet{entries: entries}

	_, err := rSet.GetHashArray()

	if err == nil {
		t.Fatalf("We did not get an error we expected, since we passed an empty entries. ")
	}

}
