package types

type HasImmutables interface {
	GetImmutables() Immutables
	GetMaintainersID() ID
}
