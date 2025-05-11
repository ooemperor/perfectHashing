package src

import (
	"testing"
)

/*
TestResultSetGetHashArray checks if we receive a proper output for the TestResultSetGetHashArray
*/
func TestResultSetGetHashArray(t *testing.T) {
	rEntry := ResultEntry{value: "test"}
	var entries []*ResultEntry

	entries = append(entries, &rEntry)

	rSet := ResultSet{entries: entries}

	result, err := rSet.GetHashArray()

	if err != nil {
		t.Fatalf("Error %v occured when retreiving the HashArray", err)
	}

	if len(result) != 4294967296 {
		t.Fatalf("GetHashArray result does have an incorrect length")
	}
}

/*
TestResultSetGetHashArrayEmptyInput checks if we receive an expected error for the TestResultSetGetHashArrayEmptyInput
*/
func TestResultSetGetHashArrayEmptyInput(t *testing.T) {
	var entries []*ResultEntry
	rSet := ResultSet{entries: entries}

	_, err := rSet.GetHashArray()

	if err != nil {
		t.Fatalf("We did get an error we expected, since we passed an empty entries. ")
	}

}
