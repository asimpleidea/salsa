package salsa

import "errors"

var (
	// ErrOptions represents an error that occurs when the options provided by
	// the user contains invalid values.
	ErrOptions = errors.New("provided options are not valid")
	// ErrIteratorFinished specifies that the iterator arrived at the end of
	// its internal list and thus could not retrieve any more results.
	ErrIteratorFinished = errors.New("iterator has finished")
	// ErrUknownIndex is returned when the provided or returned index is not
	// known or supported.
	ErrUknownIndex = errors.New("unknown or unsupported index")
	// ErrServerSide is returned when the request for source was successfully
	// sent but could not complete because of a server side error.
	ErrServerSide = errors.New("server side error")
	// ErrClientSide is returned when the request for source was successfully
	// sent but could not complete because of an error on user's request.
	ErrClientSide = errors.New("client side error")
)

type sauceError struct {
	msg string
	err error
}

func (se *sauceError) Error() string {
	return se.msg
}

func (se *sauceError) Unwrap() error {
	return se.err
}
