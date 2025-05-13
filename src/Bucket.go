package src

import "sort"

type Bucket struct {
	Keys        []string // used to store the values
	bucketIndex uint32
}

/*
Size returns the size of the Keys
*/
func (bucket *Bucket) Size() int {
	return len(bucket.Keys)
}

/*
Insert a key into the Bucket
*/
func (bucket *Bucket) Insert(key string) {
	bucket.Keys = append(bucket.Keys, key)
}

/*
BuildBuckets returns an array of Buckets given by the specific size
*/
func BuildBuckets(size uint32) []*Bucket {
	output := make([]*Bucket, size)
	for i := range output {
		output[i] = &Bucket{bucketIndex: uint32(i)}
	}
	return output
}

/*
SortBuckets sorts the buckets by size descending
*/
func SortBuckets(buckets []*Bucket) {
	//sort the buckets
	sort.Slice(buckets, func(i, j int) bool {
		return buckets[i].Size() > buckets[j].Size()
	})
}
