package schema

type Chain interface {
	String() string
	GetID() ID
	GetTrustHeight() Height
}
