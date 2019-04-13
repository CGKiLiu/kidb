package memdb

import (
	"testing"
	dt "kidb/src/datatype"
	"fmt"
)

	func TestDBPutOne(t *testing.T){
	db := NewDB()
	key := dt.NewSlice([]byte("Liu"))
	value := dt.NewSlice([]byte("Qi"))
	db.Put(key, value)
}

func TestDBGetOne(t *testing.T){
	db := NewDB()
	key := dt.NewSlice([]byte("Liu"))
	value := dt.NewSlice([]byte("Qi"))
	db.Put(key, value)

	var res = db.Get(key)
	fmt.Println(string(res.Data()))
}

func TestGetAndPutMulti(t *testing.T){
	db := NewDB()
	for i := 0; i < 10; i++{
		key := dt.NewSlice([]byte(fmt.Sprint(i)))
		value := dt.NewSlice([]byte(fmt.Sprint(100-i)))
		db.Put(key, value)
	}

	for i:= 10; i >=0; i--{
		key := dt.NewSlice([]byte(fmt.Sprint(i)))
		value := db.Get(key)
		if value != nil{
			fmt.Print(string(value.Data()))
			fmt.Println("")
		}

	}

}
