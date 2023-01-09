package ids

type AnyOwnableID interface {
	OwnableID
	Get() OwnableID
	IsAnyOwnableID()
}
