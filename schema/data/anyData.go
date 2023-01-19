package data

type AnyData interface {
	Data
	Get() Data
	IsAnyData()
}
