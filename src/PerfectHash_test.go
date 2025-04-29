package src

import "testing"

/*
TestPerfectHashJoin Test if the join implementation is done correclty
TODO: change the content of the test to a proper content after the implementation
*/
func TestPerfectHashJoin(t *testing.T) {
	/*
		This method is not yet implemented since this still needs to be done

		1) set up the perfect hash
		2) construct the two ResultSets that we wanna join
		3) join the two result sets
		4) check if we have any errors found
		5) check if the results actually looks good
	*/

}

/*
TestPerfectHashGetRelatedEntry
TODO: change the content of the test to a proper content after the implementation
*/
func TestPerfectHashGetRelatedEntry(t *testing.T) {
	/*
		This method is not yet implemented since this still needs to be done

		1) set up the perfect hash
		2) construct the two ResultSets that we wanna join
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
		2) construct the two ResultSets that we wanna join
		3) select an entry from the second one, where we should not be able to find a match
		4) check that we successfully receive an error since we do not find an entry
	*/
}
