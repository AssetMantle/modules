package types

type Height interface {
	String() string

	IsGraterThat(Height) bool
}
