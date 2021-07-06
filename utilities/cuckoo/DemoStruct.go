package cuckoo


type id struct {
	IDString string `json:"idString"`
}


func newId(idstring string) *id {
	return &id{
		IDString: idstring,
	}
}
