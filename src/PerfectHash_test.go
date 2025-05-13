package src

import "testing"

/*
TestPerfectHashJoin Test if the join implementation is done correctly
TODO: change the content of the test to a proper content after the implementation
*/
func TestPerfectHashJoin(t *testing.T) {
	/*
		This method is not yet implemented since this still needs to be done

		1) set up the perfect hash
		2) construct the two ResultSets that we want to join
		3) join the two result sets
		4) check if we have any errors found
		5) check if the results actually looks good
	*/

	hasher := PerfectHash{}

	e1 := ResultEntry{Value: "v1"}
	e2 := ResultEntry{Value: "v2"}
	e3 := ResultEntry{Value: "v4"}
	e4 := ResultEntry{Value: "v1"}
	e5 := ResultEntry{Value: "v2"}

	r1 := ResultSet{Entries: []*ResultEntry{&e1, &e2, &e3}}
	r2 := ResultSet{Entries: []*ResultEntry{&e4, &e5}}

	// result, err := hasher.Join(&r1, &r2)
	result, err := hasher.Join(&r1, &r2)

	if err != nil {
		t.Fatalf("Error encountered while joining two tables that should not Produce Error: %v", err)
	}

	if len(result) != 3 {
		t.Fatalf("Result length is wrong: %v", len(result))
	}

}

/*
TestPerfectHashJoinWithError Test if the join implementation is done with error when one is expected
*/
func TestPerfectHashJoinWithError(t *testing.T) {
	/*
		This method is not yet implemented since this still needs to be done

		1) set up the perfect hash
		2) construct the two ResultSets that we want to join
		3) join the two result sets
		4) check if the result is nil as expected
	*/

	hasher := PerfectHash{}

	r1 := ResultSet{}
	r2 := ResultSet{}

	// result, err := hasher.Join(&r1, &r2)
	result, _ := hasher.Join(&r1, &r2)

	if result != nil {
		t.Fatalf("Join produces non nil result while it should not: %v", result)
	}

}

/*
TestPerfectHashGetRelatedEntry
*/
func TestPerfectHashGetRelatedEntry(t *testing.T) {
	/*
		This method is not yet implemented since this still needs to be done

		1) set up the perfect hash
		2) construct the two ResultSets that we want to join
		3) select an entry from the second one, where we should be able to find a match
		4) check for any errors during the join operation
		5) check if the related Entry is actually the one that we expect.
	*/

	hasher := PerfectHash{}

	e1 := ResultEntry{Value: "v1"}
	e2 := ResultEntry{Value: "v2"}
	e3 := ResultEntry{Value: "v4"}
	e4 := ResultEntry{Value: "v1"}
	e5 := ResultEntry{Value: "v2"}

	r1 := ResultSet{Entries: []*ResultEntry{&e1, &e2, &e3}}
	r2 := ResultSet{Entries: []*ResultEntry{&e4, &e5}}

	relatedEntry, err := hasher.GetRelatedEntry(&e1, &r1, &r2)

	if err != nil {
		t.Fatalf("Error encountered while joining two tables that should not Produce Error: %v", err)
	}

	if relatedEntry == nil {
		t.Fatalf("relatedEntry should not be nil when it should have been returned")
	}

	if relatedEntry.Value != "v1" {
		t.Fatalf("relatedEntry.value should be 'v1' when it should have been returned")
	}

	if relatedEntry != &e4 {
		t.Fatalf("pointer to the expected result is incorrect, expected %v, got %v", &e4, &relatedEntry)
	}

}

/*
TestPerfectHashGetRelatedEntryNotFound tests if we receive the expected output if there is no Entry found for the PerfectHash
*/
func TestPerfectHashGetRelatedEntryNotFound(t *testing.T) {
	/*
		This method is not yet implemented since this still needs to be done

		1) set up the perfect hash
		2) construct the two ResultSets that we want to join
		3) select an entry from the second one, where we should not be able to find a match
		4) check that we successfully receive an error since we do not find an entry
	*/

	hasher := PerfectHash{}

	e1 := ResultEntry{Value: "v1"}
	e2 := ResultEntry{Value: "v2"}
	e3 := ResultEntry{Value: "v4"}
	e4 := ResultEntry{Value: "v1"}
	e5 := ResultEntry{Value: "v2"}

	r1 := ResultSet{Entries: []*ResultEntry{&e1, &e2, &e3}}
	r2 := ResultSet{Entries: []*ResultEntry{&e4, &e5}}

	relatedEntry, err := hasher.GetRelatedEntry(&e3, &r1, &r2)

	if err != nil {
		t.Fatalf("Error encountered while joining two tables that should not Produce Error: %v", err)
	}

	if relatedEntry != nil {
		t.Fatalf("relatedEntry should be nil but got different value %v", relatedEntry)
	}
}
