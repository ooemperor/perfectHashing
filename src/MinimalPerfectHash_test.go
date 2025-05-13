package src

import "testing"

/*
TestMinimalPerfectHashJoin Test if the join implementation is done correctly
*/
func TestMinimalPerfectHashJoin(t *testing.T) {
	/*
		This method is not yet implemented since this still needs to be done

		1) set up the perfect hash
		2) construct the two ResultSets that we want to join
		3) join the two result sets
		4) check if we have any errors found
		5) check if the results actually looks good
	*/

	hasher := MinimalPerfectHash{BucketCount: 100}

	e1 := ResultEntry{Value: "v1"}
	e2 := ResultEntry{Value: "v2"}
	e3 := ResultEntry{Value: "v4"}
	e4 := ResultEntry{Value: "v1"}
	e5 := ResultEntry{Value: "v2"}

	r1 := ResultSet{Entries: []*ResultEntry{&e1, &e2, &e3}}
	r2 := ResultSet{Entries: []*ResultEntry{&e4, &e5}}

	err := hasher.Build(&r2)
	if err != nil {
		t.Fatalf("Error while building two tables that should not Produce Error: %v", err)
	}

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
func TestMinimalPerfectHashJoinWithError(t *testing.T) {
	/*
		This method is not yet implemented since this still needs to be done

		1) set up the perfect hash
		2) construct the two ResultSets that we want to join
		3) join the two result sets
		4) check if the result is nil as expected
	*/

	hasher := MinimalPerfectHash{BucketCount: 100}

	r1 := ResultSet{}
	r2 := ResultSet{}

	err := hasher.Build(&r2)
	if err != nil {
		t.Fatalf("Error while building two tables that should not Produce Error: %v", err)
	}

	// result, err := hasher.Join(&r1, &r2)
	result, _ := hasher.Join(&r1, &r2)

	if result != nil {
		t.Fatalf("Join produces non nil result while it should not: %v", result)
	}

}

/*
TestPerfectHashGetRelatedEntry
*/
func TestMinimalPerfectHashGetRelatedEntry(t *testing.T) {
	/*
		This method is not yet implemented since this still needs to be done

		1) set up the perfect hash
		2) construct the two ResultSets that we want to join
		3) select an entry from the second one, where we should be able to find a match
		4) check for any errors during the join operation
		5) check if the related Entry is actually the one that we expect.
	*/

	hasher := MinimalPerfectHash{BucketCount: 100}

	e1 := ResultEntry{Value: "v1"}
	//e2 := ResultEntry{Value: "v2"}
	//e3 := ResultEntry{Value: "v4"}
	e4 := ResultEntry{Value: "v1"}
	e5 := ResultEntry{Value: "v2"}

	//r1 := ResultSet{Entries: []*ResultEntry{&e1, &e2, &e3}}
	r2 := ResultSet{Entries: []*ResultEntry{&e4, &e5}}

	err := hasher.Build(&r2)
	if err != nil {
		t.Fatalf("Error while building two tables that should not Produce Error: %v", err)
	}

	relatedEntry, err := hasher.GetRelatedEntry(&e1)

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
