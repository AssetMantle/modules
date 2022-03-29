package types

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
)

type List interface {
	GetList() []traits.Listable

	Search(func()) int

	Apply(func()) List
	Add(...interface{}) List
	Remove(...interface{}) List
	Mutate(...interface{}) List
}
