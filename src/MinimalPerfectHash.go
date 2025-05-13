package src

import (
	"log"
	"math/rand"
)

/*
MinimalPerfectHash is the struct that defines the necessary structures for executing the hash.
*/
type MinimalPerfectHash struct {
	// []*any should be a array that contains the pointers to the result set
	// should be the size of the result space of the hash function
	// should not use slices but rather a well-defined array
	JoinType    string // left or inner at the moment
	BucketCount uint32
	SeedArray   []uint32
	Results     []*ResultEntry
}

/*
Join implements the interface but does nothing meaningful yet
*/
func (mph *MinimalPerfectHash) Join(table1 *ResultSet, table2 *ResultSet) ([]*PerfectHashResult, error) {
	var result []*PerfectHashResult
	for _, entry1 := range table1.Entries {
		relatedEntry, _ := mph.GetRelatedEntry(entry1)
		result = append(result, &PerfectHashResult{entry1, relatedEntry})
	}
	return result, nil
}

/*
GetRelatedEntry fetches the related entry from the other table
*/
func (mph *MinimalPerfectHash) GetRelatedEntry(entry *ResultEntry) (*ResultEntry, error) {
	// fetch the bucket
	bucketIndex, _ := entry.GetPosition()
	bucketIndex = bucketIndex % mph.BucketCount
	seed := mph.SeedArray[bucketIndex]
	hash, _ := entry.Hash(seed)
	pos := hash % uint32(len(mph.Results))

	return mph.Results[pos], nil
}

/*
Build constructs the Minimal PHF and stores it in the struct
*/
func (mph *MinimalPerfectHash) Build(table2 *ResultSet) error {
	buckets := BuildBuckets(mph.BucketCount)
	seedArray := make([]uint32, mph.BucketCount)

	// Phase 1: insert the value into the Buckets
	for _, entry := range table2.Entries {
		// sort the entry into the buckets
		tmpVal, err := entry.GetPosition()

		if err != nil {
			log.Fatalf("Error in GetRelatedEntry %v", err)
		}

		bucketIndex := tmpVal % mph.BucketCount
		tmpBucket := buckets[bucketIndex]
		tmpBucket.Insert(entry)
	}

	// Phase 2: Sort the buckets
	SortBuckets(buckets)

	// Phase 3: Init the result output array
	resultArray := make([]*ResultEntry, len(table2.Entries))
	resultBitMap := make([]bool, len(table2.Entries))

	// Phase 4: Start by hashing the buckets
	for _, bucket := range buckets {
		// go over the buckets
		bucketSuccess := false
		for !bucketSuccess {
			seed := rand.Uint32()
			tmpBitMap := make([]bool, len(table2.Entries))
			collision := false
			for _, entry := range bucket.Keys {

				hash, _ := entry.Hash(seed)
				posCand := hash % uint32(len(table2.Entries))

				// check the current and resultBitMap for collisions
				if tmpBitMap[posCand] == false && resultBitMap[posCand] == false {
					// map the tmpBitMap to true and continue
					tmpBitMap[posCand] = true
				} else {
					// we break the current seed loop
					collision = true
					break
				}
			}

			if !collision {
				// we found a perfectly good mapping
				bucketSuccess = true
				seedArray[bucket.bucketIndex] = seed

				for _, entry := range bucket.Keys {
					hash, _ := entry.Hash(seed)
					pos := hash % uint32(len(table2.Entries))
					resultBitMap[pos] = true
					resultArray[pos] = entry
				}

				// now update the resultLoop

			} else {
				// we continue with the outer loop and a new seed and try again
				continue
			}
		}
	}
	mph.SeedArray = seedArray
	mph.Results = resultArray

	return nil
}
