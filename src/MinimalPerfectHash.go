package src

import (
	"fmt"
	"log"
)

/*
MinimalPerfectHash is the struct that defines the necessary structures for executing the hash.
*/
type MinimalPerfectHash struct {
	// []*any should be a array that contains the pointers to the result set
	// should be the size of the result space of the hash function
	// should not use slices but rather a well-defined array
	JoinType string // left or inner at the moment
}

/*
Join implements the interface but does nothing meaningful yet
*/
func (ph *MinimalPerfectHash) Join(table1 *ResultSet, table2 *ResultSet) ([]*PerfectHashResult, error) {
	return nil, nil
}

/*
GetRelatedEntry fetches the related entry from the other table
*/
func (ph *MinimalPerfectHash) GetRelatedEntry(entry *ResultEntry, table1 *ResultSet, table2 *ResultSet) (*ResultEntry, error) {
	return nil, nil
}

/*
Build constructs the Minimal PHF
*/
func (ph *MinimalPerfectHash) Build(table1 *ResultSet, table2 *ResultSet) {
	bucketCount := uint32(100)
	buckets := BuildBuckets(bucketCount)

	// Phase 1: insert the value into the Buckets
	for _, entry := range table2.Entries {
		// sort the entry into the buckets
		tmpVal, err := entry.GetPosition()

		if err != nil {
			log.Fatalf("Error in GetRelatedEntry %v", err)
		}

		bucketIndex := tmpVal % bucketCount
		tmpBucket := buckets[bucketIndex]
		tmpBucket.Insert(entry.Value)
	}
	// sort buckets
	for i := range buckets {
		fmt.Print(buckets[i].Size())
	}
	SortBuckets(buckets)

	for i := range buckets {
		fmt.Print(buckets[i].Size())
	}
}
