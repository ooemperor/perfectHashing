package src

import "hash/crc32"

type IResultEntry interface {
	GetPosition() (uint32, error)
}

/*
ResultEntry will be the object used to store a single row
this may be replaced by proper types later in this project
*/
type ResultEntry struct {
	Value string
}

/*
GetPosition retruns the position in a 2^32 array or initial idea for bucket mapping
*/
func (entry *ResultEntry) GetPosition() (uint32, error) {
	return entry.Hash(0xD5828281)
}

/*
Hash generates a number based on the seed polynom
*/
func (entry *ResultEntry) Hash(seed uint32) (uint32, error) {
	crc32q := crc32.MakeTable(seed) // maybe this polynom needs to be adjusted later
	return crc32.Checksum([]byte(entry.Value), crc32q), nil
}
