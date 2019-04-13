package memdb

import (
	"kidb/src/skiplist"
	dt "kidb/src/datatype"

	"encoding/binary"
	"kidb/src/comparator"
)

type DB struct{
	_comparator comparator.Comparator
	_table *skiplist.SkipList

}

func NewDB() *DB {
	db := &DB{
		_comparator: comparator.DefaultComparator,
		_table:      skiplist.NewSkipList(comparator.DefaultComparator),
	}
	return db
}

//Get return the corresponding value, if the key-value pair is no exist, return nil
func (db *DB) Get(key *dt.Slice) *dt.Slice{
	node := db._table.Find(key.Data())
	if node == nil{
		return nil
	}
	keyValueSlice := dt.NewSlice(node.Get())
	keyValueData := keyValueSlice.Data()
	var keySize = binary.BigEndian.Uint32(keyValueData[0:4])
	//var valueSize = binary.BigEndian.Uint32(keyValueData[4+keySize: 8+keySize])
	var value = keyValueData[8+keySize:]
	return dt.NewSlice(value)
}

//Put put key-value into DB
func (db *DB) Put(key *dt.Slice, value *dt.Slice){

	// Format of an entry is concatenation of:
	// key_size		: int32 of key.Size()
	// key bytes	: [key_size]byte
	// value_size	: int32 of value.Size()
	// value bytes	: [value_size]byte

	keySize := key.Size()
	valueSize := value.Size()
	encodeLen := int32(4) + keySize + int32(4) + valueSize
	buf := make([]byte, encodeLen)
	var keySizeByte = make([]byte, 4)
	binary.BigEndian.PutUint32(keySizeByte, uint32(keySize))
	var valueSizeByte = make([]byte, 4)
	binary.BigEndian.PutUint32(valueSizeByte, uint32(valueSize))

	copy(buf[0:4], keySizeByte)
	copy(buf[4:4+keySize], key.Data())
	copy(buf[4+keySize: 8+keySize], valueSizeByte)
	copy(buf[8+keySize:], value.Data())

	db._table.Insert(buf)
}