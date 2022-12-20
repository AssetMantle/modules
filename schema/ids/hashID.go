package ids

type HashID interface {
	ID
	IsHashID()
	EncodedString() string
}
