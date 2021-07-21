package cuckoo

type id struct {
	IDString string `json:"idString"`
}

func newId(idString string) *id {
	return &id{
		IDString: idString,
	}
}
