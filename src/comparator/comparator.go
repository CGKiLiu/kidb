package comparator

//Comparator interface
type Comparator interface {
	Compare(a, b []byte) int
}

