package ids

type SplitID interface {
	ID
	GetOwnableID() OwnableID
	IsSplitID()
	SplitIDString() string
}
