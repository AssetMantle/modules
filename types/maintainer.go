package types

type Maintainer interface {
	Name() string

	ID() ID

	CanMutateMaintainersProperty(ID) bool

	CanAddMaintainer() bool
	CanRemoveMaintainer() bool
	CanMutateMaintainer() bool

	CanMutateLock() bool
	CanMutateBurn() bool
	CanMutateTrait(ID) bool
}
