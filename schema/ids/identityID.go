package ids

type IdentityID interface {
	ID
	GetHashID() HashID
	IsIdentityID()
}
