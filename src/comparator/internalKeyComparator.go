package comparator

import (
	"bytes"
	"encoding/binary"
)

//InternalKeyComparator ...
type InternalKeyComparator struct {
}

//Compare return 0 if a==b, -1 if a < b , 1 if a > b
func (ikeyComp InternalKeyComparator) Compare(a, b []byte) int {
	keySize := binary.BigEndian.Uint32(a[0:4])
	key := a[4:4+keySize]
	return bytes.Compare(key, b)
}

var DefaultComparator = &InternalKeyComparator{}