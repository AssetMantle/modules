package types

type Ownable interface {
	GetOwnerID() ID
	GetOwnableID() ID
}
