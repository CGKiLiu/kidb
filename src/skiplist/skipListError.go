package skiplist
type skipNodeError struct {
	errorStr string
}

func (snError *skipNodeError) Error() string {
	return snError.errorStr
}
