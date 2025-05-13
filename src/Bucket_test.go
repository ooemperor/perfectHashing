package src

import "testing"

/*
TestBucketEmptySize Checks if the bucket returns 0 as size if no entry present
*/
func TestBucketEmptySize(t *testing.T) {
	bucket := Bucket{}

	size := bucket.Size()

	if size != 0 {
		t.Fatalf("Size of bucket expected 0, got %d", size)
	}
}

/*
TestBucketInsertAndSize tests the setup of single bucket and the insert
*/
func TestBucketInsertAndSize(t *testing.T) {
	rEntry := ResultEntry{Value: "test"}

	bucket := Bucket{}

	size := bucket.Size()

	if size != 0 {
		t.Fatalf("Size of bucket expected 0, got %d", size)
	}
	bucket.Insert(&rEntry)

	size = bucket.Size()

	if size != 1 {
		t.Fatalf("Size of bucket expected 1, got %d", size)
	}
}

/*
TestBuildBuckets tests the build of the buckets
*/
func TestBuildBuckets(t *testing.T) {
	buckets := BuildBuckets(100)
	if len(buckets) != 100 {
		t.Fatalf("Buckets size expected 100, got %d", len(buckets))
	}
}

/*
TestBuildBuckets tests the build of the buckets then tests if the sorting is working properly
*/
func TestBuildBucketsAndSort(t *testing.T) {
	buckets := BuildBuckets(100)

	if len(buckets) != 100 {
		t.Fatalf("Buckets size expected 100, got %d", len(buckets))
	}

	bucket := buckets[1]
	bucket.Insert(&ResultEntry{Value: "test"})

	if buckets[1] != bucket {
		t.Fatalf("Buckets %v expected , got %v", *buckets[1], *bucket)
	}

	SortBuckets(buckets)

	if buckets[1] == bucket {
		t.Fatalf("Buckets %v not expected", *buckets[1])
	}

	if buckets[0] != bucket {
		t.Fatalf("Buckets %v expected , got %v", *buckets[1], *bucket)
	}
}
