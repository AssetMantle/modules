package ids

type OrderID interface {
	ID
	IsOrderID()
}
