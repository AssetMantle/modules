package ids

type SplitID interface {
	ID
	GetOwnableID() ID
	IsSplitID()
}
