package src

/*
IPerfectHash interfaces that specifies the functionalities for the perfect hash algorithm.
*/
type IPerfectHash interface {
	Join(table1 *ResultSet, table2 *ResultSet) (*ResultSet, error)

	/*
		returns a pointer to the correlated ResultEntry after the HashJoin
		expects that the ResultEntry provided is from the first array.
	*/
	GetRelatedEntry(entry *ResultEntry) (*ResultEntry, error)
}

/*
PerfectHash is the struct that defines the necessary structures for executing the hash.
*/
type PerfectHash struct {
	// []*any should be a array that contains the pointers to the result set
	// should be the size of the result space of the hash function
	// should not use slices but rather a well-defined array
}
