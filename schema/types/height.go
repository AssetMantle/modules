package types

type Height interface {
	Get() int64
	IsGreaterThan(Height) bool
}
