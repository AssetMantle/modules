package types

type Owner interface {
	String() string

	ID() ID
	Share() Share
}
