package ids

type AnyID interface {
	Get() ID
	ID
}
