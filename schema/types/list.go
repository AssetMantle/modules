package types

type List interface {
	GetList() []interface{}

	Search(func()) int

	Apply(func()) List
	Add(...interface{}) List
	Remove(...interface{}) List
	Mutate(...interface{}) List
}
