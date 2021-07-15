package traits

type Sortable interface {
	Compare(Sortable) int
}
