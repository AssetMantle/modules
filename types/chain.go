package types

type Chain interface {
	String() string
	GetID() ID
	GetTrustHeight() Height
}
