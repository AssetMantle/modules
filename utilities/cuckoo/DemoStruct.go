package cuckoo

type ID struct {
	IDString string `json:"idString"`
}

func NewID(idString string) *ID {
	return &ID{
		IDString: idString,
	}
}
