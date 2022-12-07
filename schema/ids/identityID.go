package ids

type IdentityID interface {
	ID
	GetHashID() ID
	IsIdentityID()
}
