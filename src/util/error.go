package arena

type blockAllocError struct {
	errorStr string
}

func (bAllocErr *blockAllocError) Error() string {
	return "errorStr"
}
