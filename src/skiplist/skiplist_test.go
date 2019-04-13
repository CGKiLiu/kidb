package skiplist

import (
	"testing"
	"kidb/src/comparator"
	"fmt"
)

func TestSkipListInsert(t *testing.T) {
	cmp := comparator.InternalKeyComparator{}
	skl := NewSkipList(cmp)
	for i:=0; i < 1024; i++{
		skl.Insert([]byte("a"+fmt.Sprint(i)))
	}
	skl.print()
}

func TestSkipListFind(t *testing.T){
	skl := NewSkipList(comparator.InternalKeyComparator{})
	for i:=0; i < 11; i++{
		skl.Insert([]byte("a"+fmt.Sprint(i)))
	}
	for i:=0; i < 14; i++{
		findResult := skl.Find([]byte("a"+fmt.Sprint(i)))
		if findResult != nil{
			findResult.Print()
		}
	}
}