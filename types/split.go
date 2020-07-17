package types

type Split interface {
	GetID() ID
	Ownable
	Transactional
}
