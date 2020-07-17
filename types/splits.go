package types

type Splits interface {
	GetID() ID

	Get(ID) Split
	GetList() []Split

	Fetch(ID) Splits
	Add(Split) Splits
	Remove(Split) Splits
	Mutate(Split) Splits
}
