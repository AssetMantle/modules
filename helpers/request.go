package helpers

type Request interface {
	Validate() error
}
