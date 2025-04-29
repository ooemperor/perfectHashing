package src

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
	results []any
}
