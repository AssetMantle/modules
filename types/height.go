package types

type Height interface {
	Count() int
	IsGraterThat(Height) bool
}

type BaseHeight struct {
	Height int
}

var _ Height = (*BaseHeight)(nil)

func (baseHeight BaseHeight) Count() int               {}
func (baseHeight BaseHeight) IsGraterThat(Height) bool {}
