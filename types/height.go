package types

type Height interface {
	Get() int
	IsGreaterThan(Height) bool
}

type height struct {
	Height int
}

var _ Height = (*height)(nil)

func (height height) Get() int { return height.Height }
func (height height) IsGreaterThan(Height Height) bool {
	return height.Get() > Height.Get()
}
func NewHeight(Height int) Height {
	return height{Height: Height}
}
