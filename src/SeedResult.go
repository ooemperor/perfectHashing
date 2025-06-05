package src

import "sync"

/*
SeedResult synchronized obejct to enable multithreaded search for seeds that solve our search for the PerfectHashFunction
*/
type SeedResult struct {
	mu     sync.Mutex
	Result uint32
}

/*
Set the result of the seed.
*/
func (sr *SeedResult) Set(res uint32) {
	sr.mu.Lock()
	sr.Result = res
	sr.mu.Unlock()
}

/*
Get returns the current result of the SeedResult
*/
func (sr *SeedResult) Get() (res uint32) {
	return sr.Result
}
