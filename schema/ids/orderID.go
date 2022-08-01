package ids

type OrderID interface {
	ID
	GetHashID() HashID
}
