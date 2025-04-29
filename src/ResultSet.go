package src

type IResultSet interface {
	/*
		Get the array of pointers to the result set for the given hash function
	*/
	GetHashArray() ([]*ResultSet, error)
}

/*
ResultSet defines the object used to store all the results for a given table.
this may be replaced by proper types later in this project
*/
type ResultSet struct {
	// TODO: replace the slices here with proper array if possible
	entries []*ResultEntry
}
