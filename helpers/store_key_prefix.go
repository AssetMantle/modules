package helpers

type StoreKeyPrefix interface {
	// TODO define types and use key and storeKey
	GenerateStoreKey(key []byte) []byte
}
