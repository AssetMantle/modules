package ids

type MetaID interface {
	ID
	IsMetaID()
}
