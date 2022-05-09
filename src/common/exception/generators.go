package exception

func NewError(err error) error {
	return &Exception{
		err:   err,
		stack: populateStack(),
	}
}

// Error wrapper for postgres queries.
func NewPQError(err error, query string, args []any) error {
	return &Exception{
		err:   err,
		stack: populateStack(),
		query: query,
		args:  args,
	}
}
