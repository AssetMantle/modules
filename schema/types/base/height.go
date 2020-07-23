package base

import "github.com/persistenceOne/persistenceSDK/schema/types"

type height struct {
	Height int64
}

var _ types.Height = (*height)(nil)

func (height height) Get() int64 { return height.Height }
func (height height) IsGreaterThan(Height types.Height) bool {
	return height.Get() > Height.Get()
}
func NewHeight(Height int64) types.Height {
	return height{Height: Height}
}
