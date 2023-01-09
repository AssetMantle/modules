package ids

type OwnableID interface {
	ID
	ToAnyOwnableID() AnyOwnableID
}
