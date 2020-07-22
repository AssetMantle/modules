package types

type Orders interface {
	GetID() ID

	Get(ID) Order
	GetList() []Order

	Fetch(ID) Orders
	Add(Order) Orders
	Remove(Order) Orders
	Mutate(Order) Orders
}
