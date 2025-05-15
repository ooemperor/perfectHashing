package src

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"testing"
)

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
		t.Fatalf("relatedEntry.value should be 'v1' when it should have been returned but is %v", relatedEntry.Value)
	}

	if relatedEntry != &e4 {
		t.Fatalf("pointer to the expected result is incorrect, expected %v, got %v", &e4, &relatedEntry)
	}

}

func BenchmarkMinimalPerfectHash_Build(b *testing.B) {
	tests := []struct {
		bucketCount int
		entryCount  int
		threadCount int
	}{
		{1000, 1000, 1},
		{10000, 1000, 1},
		{100000, 1000, 1},
		{10000, 10000, 1},
		{100000, 10000, 1},
		{100000, 100000, 1},
		{1000, 1000, 2},
		{10000, 1000, 2},
		{100000, 1000, 2},
		{10000, 10000, 2},
		{100000, 10000, 2},
		{100000, 100000, 2},
		{1000, 1000, 4},
		{10000, 1000, 4},
		{100000, 1000, 4},
		{10000, 10000, 4},
		{100000, 10000, 4},
		{100000, 100000, 4},
		{1000, 1000, 8},
		{10000, 1000, 8},
		{100000, 1000, 8},
		{10000, 10000, 8},
		{100000, 10000, 8},
		{100000, 100000, 8},
		{1000, 1000, 16},
		{10000, 1000, 16},
		{100000, 1000, 16},
		{10000, 10000, 16},
		{100000, 10000, 16},
		{100000, 100000, 16},
		{1000, 1000, 32},
		{10000, 1000, 32},
		{100000, 1000, 32},
		{10000, 10000, 32},
		{100000, 10000, 32},
		{100000, 100000, 32},
	}

	for _, tt := range tests {
		b.Run(fmt.Sprintf("Buckets-%v-Entries-%v-Threads-%v", tt.bucketCount, tt.entryCount, tt.threadCount), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rs := generateResultSet(tt.entryCount)
				mph := &MinimalPerfectHash{
					BucketCount:  uint32(tt.bucketCount),
					ThreadsCount: uint32(tt.threadCount),
				}
				if err := mph.Build(rs); err != nil {
					b.Fatalf("Build failed: %v", err)
				}
			}
		})
	}
}

func BenchmarkMinimalPerfectHash_Build_Entries(b *testing.B) {
	bucketCount := 10000
	threadCount := 4

	for val := range 100 {
		entries := 1000 * val
		b.Run(fmt.Sprintf("Buckets-%v-Entries-%v-Threads-%v-Factor-%v", bucketCount, entries, threadCount, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rs := generateResultSet(entries)
				mph := &MinimalPerfectHash{
					BucketCount:  uint32(bucketCount),
					ThreadsCount: uint32(threadCount),
				}
				if err := mph.Build(rs); err != nil {
					b.Fatalf("Build failed: %v", err)
				}
				// fmt.Println(len(mph.Results))
			}
		})
	}
}

func BenchmarkMinimalPerfectHash_Build_Buckets(b *testing.B) {

	entryCount := 100000
	threadCount := 4

	for val := range 100 {
		buckets := val * 1000
		b.Run(fmt.Sprintf("Buckets-%v-Entries-%v-Threads-%v-Factor-%v", buckets, entryCount, threadCount, 1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rs := generateResultSet(entryCount)
				mph := &MinimalPerfectHash{
					BucketCount:  uint32(buckets),
					ThreadsCount: uint32(threadCount),
				}
				if err := mph.Build(rs); err != nil {
					b.Fatalf("Build failed: %v", err)
				}
				// fmt.Println(len(mph.Results))
			}
		})
	}
}

func BenchmarkMinimalPerfectHash_Build_Factors(b *testing.B) {
	tt := struct {
		bucketCount int
		entryCount  int
		threadCount int
	}{10000, 100000, 4}

	var factor float64
	factor = 1
	for _ = range 10 {
		factor += 0.1
		b.Run(fmt.Sprintf("Buckets-%v-Entries-%v-Threads-%v-Factor-%f", tt.bucketCount, tt.entryCount, tt.threadCount, factor), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rs := generateResultSet(tt.entryCount)
				mph := &MinimalPerfectHash{
					BucketCount:       uint32(tt.bucketCount),
					ThreadsCount:      uint32(tt.threadCount),
					ResultSpaceFactor: factor,
				}
				if err := mph.Build(rs); err != nil {
					b.Fatalf("Build failed: %v", err)
				}
				// fmt.Println(len(mph.Results))
			}
		})
	}
}

func BenchmarkMinimalPerfectHash_Build_Big(b *testing.B) {
	tests := []struct {
		bucketCount int
		entryCount  int
		threadCount int
	}{
		{100000, 500000, 8},
	}

	for _, tt := range tests {
		b.Run(fmt.Sprintf("Buckets-%v-Entries-%v-Threads-%v", tt.bucketCount, tt.entryCount, tt.threadCount), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rs := generateResultSet(tt.entryCount)
				mph := &MinimalPerfectHash{
					BucketCount:  uint32(tt.bucketCount),
					ThreadsCount: uint32(tt.threadCount),
				}
				if err := mph.Build(rs); err != nil {
					b.Fatalf("Build failed: %v", err)
				}
			}
		})
	}
}

/*
BenchmarkMinimalPerfectHashLookup_100k Benchmarks the lookup of 100k values
*/
func BenchmarkMinimalPerfectHashLookup_100k_Factor_1(b *testing.B) {
	// build the resultset once to not slow down the Benchmark
	r := generateResultSet(100000)
	hasher := MinimalPerfectHash{BucketCount: 100000}
	_ = hasher.Build(r)

	for b.Loop() {
		randomVal := rand.Uint32N(100000)
		re := ResultEntry{Value: strconv.Itoa(int(randomVal))}
		_, _ = hasher.GetRelatedEntry(&re)
	}
}

func generateResultSet(size int) *ResultSet {
	r := &ResultSet{Entries: []*ResultEntry{}}
	for i := range size {
		re := ResultEntry{Value: strconv.Itoa(i)}

		r.Entries = append(r.Entries, &re)
	}
	return r
}
