package schema

type Height interface {
	Get() int64
	IsGreaterThan(Height) bool
}

type height struct {
	Height int64
}

var _ Height = (*height)(nil)

func (height height) Get() int64 { return height.Height }
func (height height) IsGreaterThan(Height Height) bool {
	return height.Get() > Height.Get()
}
func NewHeight(Height int64) Height {
	return height{Height: Height}
}
