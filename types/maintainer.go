package types

type Maintainer interface {
	String() string

	ID() ID

	CanMutateMaintainersID() bool
	CanMutateMaintainersProperty(ID) bool

	CanAddMaintainer() bool
	CanMutateMaintainer() bool
	CanRemoveMaintainer() bool

	CanMutateLock() bool
	CanMutateBurn() bool
	CanMutateTrait(ID) bool
}
