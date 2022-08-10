package ids

type OwnableID interface {
	ID
	IsOwnableID()
}
