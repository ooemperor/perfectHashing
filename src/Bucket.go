package src

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
