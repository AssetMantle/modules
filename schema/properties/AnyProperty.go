package properties

type AnyProperty interface {
	Get() Property
	Property
}
