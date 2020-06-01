package types

type Height interface {
	Count() int
	IsGraterThat(Height) bool
}

type BaseHeight struct {
	Height int
}

var _ Height = (*BaseHeight)(nil)

func (baseHeight BaseHeight) Count() int { return baseHeight.Height }
func (baseHeight BaseHeight) IsGraterThat(height Height) bool {
	return baseHeight.Count() > height.Count()
}
