package ids

type MaintainerID interface {
	ID
	GetClassificationID() ClassificationID
}
