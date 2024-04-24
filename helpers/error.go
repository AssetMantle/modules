package helpers

type Error interface {
	Wrapf(string, ...interface{}) error
	Is(error) bool
	error
}
