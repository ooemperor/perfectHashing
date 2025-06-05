package src

import "fmt"

/*
IResultSet is the interface for the below implemented ResultSet
*/
type IResultSet interface {
	/*
		Get the array of pointers to the result set for the given hash function
	*/
	GetHashArray() ([]*ResultEntry, error)
	GetLenght() (uint32, error)
}

/*
ResultSet defines the object used to store all the results for a given table.
this may be replaced by proper types later in this project
*/
type ResultSet struct {
	Entries []*ResultEntry
}

/*
GetHashArray is the function that constructs an array of given length for the Result
*/
func (rs *ResultSet) GetHashArray() ([]*ResultEntry, error) {
	localArray := make([]*ResultEntry, 4294967296)

	for i := range rs.Entries {
		hash, err := rs.Entries[i].GetPosition()

		if err != nil {
			return nil, err
		}

		if localArray[hash] != nil {
			return nil, fmt.Errorf("duplicate hash %d", hash)
		}
		localArray[hash] = rs.Entries[i]
	}

	return localArray, nil
}
