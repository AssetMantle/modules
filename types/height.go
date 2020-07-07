package types

type Height interface {
	Get() int64
	IsGraterThat(Height) bool
}

type height struct {
	Height int64
}

var _ Height = (*height)(nil)

func (height height) Get() int64 { return height.Height }
func (height height) IsGraterThat(Height Height) bool {
	return height.Get() > Height.Get()
}
func NewHeight(Height int64) Height {
	return height{Height: Height}
}
