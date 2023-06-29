package helpers

type Record interface {
	GetKey() Key
	GetMappable() Mappable
}
