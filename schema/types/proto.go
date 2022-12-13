package types

type Proto interface {
	Size() int
	MarshalTo(data []byte) (int, error)
	Unmarshal(dAtA []byte) error
}
