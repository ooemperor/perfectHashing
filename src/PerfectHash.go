package src

import (
	"fmt"
)

/*
IPerfectHash interfaces that specifies the functionalities for the perfect hash algorithm.
*/
type IPerfectHash interface {
	Join(table1 *ResultSet, table2 *ResultSet) ([]*PerfectHashResult, error)

	/*
		returns a pointer to the correlated ResultEntry after the HashJoin
		expects that the ResultEntry provided is from the first array.
	*/
	GetRelatedEntry(entry *ResultEntry, table1 *ResultSet, table2 *ResultSet) (*ResultEntry, error)
}

/*
PerfectHash is the struct that defines the necessary structures for executing the hash.
*/
type PerfectHash struct {
	// []*any should be a array that contains the pointers to the result set
	// should be the size of the result space of the hash function
	// should not use slices but rather a well-defined array
	JoinType string // left or inner at the moment
}

/*
PerfectHashResult defines the object in which we store the result of a successful join.
*/
type PerfectHashResult struct {
	Entry1 *ResultEntry
	Entry2 *ResultEntry
}

/*
Join implements the interface but does nothing meaningful yet
TODO: implement this function
*/
func (ph *PerfectHash) Join(table1 *ResultSet, table2 *ResultSet) ([]*PerfectHashResult, error) {
	if ph.JoinType == "" {
		ph.JoinType = "left"
	}

	var output []*PerfectHashResult
	// calculate the hashes for table2
	hashArray2, err := table2.GetHashArray()

	if err != nil {
		return nil, err
	}

	for i := range table1.Entries {
		tmpEntry := table1.Entries[i]
		pos, err := tmpEntry.GetPosition()
		if err != nil {
			return nil, err
		}

		relatedEntry := hashArray2[pos]

		if relatedEntry == nil {
			// we did not found a match in the second table
			// this is mainly used in left joins and comparable (not inner joins)
			fmt.Println("no match found")

			if ph.JoinType == "left" {
				// only append to the output if the JoinType is left
				output = append(output, &PerfectHashResult{tmpEntry, relatedEntry})
			}

		} else {
			output = append(output, &PerfectHashResult{tmpEntry, relatedEntry})
		}
	}
	return output, nil
}

/*
GetRelatedEntry fetches the related entry from the other table
*/
func (ph *PerfectHash) GetRelatedEntry(entry *ResultEntry, table1 *ResultSet, table2 *ResultSet) (*ResultEntry, error) {

	// calculate the hashes for table2
	hashArray2, err := table2.GetHashArray()

	if err != nil {
		return nil, err
	}
	pos, err := entry.GetPosition()
	relatedEntry := hashArray2[pos]

	return relatedEntry, nil
}
