package base

import "github.com/persistenceOne/persistenceSDK/schema/types"

var _ types.Error =(*CustomError)(nil)

func (m *CustomError) Error() string {
	return m.ProtoError
}
