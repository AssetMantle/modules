package types

type Assets interface {
	String() string

	Asset(ID) Asset

	Remove(Asset) error
	Add(Asset) error
}
