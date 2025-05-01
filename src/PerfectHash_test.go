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

	r1 := ResultSet{}
	r2 := ResultSet{}

	// result, err := hasher.Join(&r1, &r2)
	_, err := hasher.Join(&r1, &r2)

	if err != nil {
		t.Fatalf("Error encountered while joining two tables that should not Produce Error: %v", err)
	}

	// TODO: check the validity of the result

}

/*
TestPerfectHashJoinWithError Test if the join implementation is done with error when one is expected
TODO: change the content of the test to a proper content after the implementation
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
TODO: change the content of the test to a proper content after the implementation
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

}

/*
TestPerfectHashGetRelatedEntryNotFound tests if we receive the expected output if there is no Entry found for the PerfectHash
TODO: change the content of the test to a proper content after the implementation
*/
func TestPerfectHashGetRelatedEntryNotFound(t *testing.T) {
	/*
		This method is not yet implemented since this still needs to be done

		1) set up the perfect hash
		2) construct the two ResultSets that we want to join
		3) select an entry from the second one, where we should not be able to find a match
		4) check that we successfully receive an error since we do not find an entry
	*/
}
