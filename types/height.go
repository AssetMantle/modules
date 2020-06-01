package types

type Height interface {
	String() string

	IsGraterThat(Height) bool
}

type BaseHeight struct {
	Height int
}

var _ Height = (*BaseHeight)(nil)

func (baseHeight BaseHeight) String() string           {}
func (baseHeight BaseHeight) IsGraterThat(Height) bool {}
