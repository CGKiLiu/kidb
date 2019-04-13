package datatype

import (
	"bytes"
	"fmt"
)

//Slice
type Slice struct{
	_data []byte
	_size int32
}

func NewSlice(data []byte) *Slice{
	return &Slice{
		_data:data,
		_size:int32(len(data)),
	}
}

func (sl *Slice) Data() []byte{
	return sl._data
}

func (sl *Slice) Size() int32{
	return sl._size
}

func (sl *Slice) Compare(osl *Slice) int{
	return bytes.Compare(sl.Data(), osl.Data())
}

func (sl *Slice) Print(){
	if sl != nil{
		fmt.Print(string(sl._data))
	}
}

func Compare(a Slice, b Slice) int {
	return bytes.Compare(a.Data(), b.Data())
}

