package errors

type Error interface {
	Wrapf(string, ...interface{}) error

	error
}
