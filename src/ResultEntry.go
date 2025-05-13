package src

import "hash/crc32"

type IResultEntry interface {
	/*
		TODO: check the return type in the final implementation
	*/
	GetPosition() (uint32, error)
}

/*
ResultEntry will be the object used to store a single row
this may be replaced by proper types later in this project
*/
type ResultEntry struct {
	// TODO: replace this with proper array and not slices if possible
	Value string
}

func (entry *ResultEntry) GetPosition() (uint32, error) {
	crc32q := crc32.MakeTable(0xD5828281) // maybe this polynom needs to be adjusted later
	return crc32.Checksum([]byte(entry.Value), crc32q), nil
}
