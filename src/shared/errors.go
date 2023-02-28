package shared

type Error struct {
	Err  error
	Path string
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func NewError(err error, path string) *Error {
	return &Error{
		Err:  err,
		Path: path,
	}
}
