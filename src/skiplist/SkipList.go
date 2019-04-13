package skiplist

import (
	"kidb/src/comparator"
	"sync/atomic"
	"math/rand"
	"fmt"
	"sync"
)
//Key key store in skipList
type Key []byte

const sklMaxHeight int32 = 12

//SkipList key-value
type SkipList struct {
	mu sync.RWMutex
	maxHeight int32       //max height of skipList
	head      *SkipNode //head of skipList
	size      int       //size of skipList
	cmp       comparator.Comparator
}

//NewSkipList Create a new SkipList
func NewSkipList(cmp comparator.Comparator) *SkipList {
	newSkipList := &SkipList{
		maxHeight: 0,
		head:      newSkipNode([]byte(""), sklMaxHeight),
		size:      0,
		cmp:       cmp,
	}
	return newSkipList
}

func (skl *SkipList) GetMaxHeight() int32{
	skl.mu.RLock()
	defer skl.mu.RUnlock()
	maxHeight := atomic.LoadInt32(&skl.maxHeight)
	return maxHeight
}

func (skl *SkipList) randomHeight() int32{
	const kBranching = 3
	height := int32(1)
	for height < sklMaxHeight && rand.Int31()%kBranching == 0{
		height += 1
	}
	return height
}

func (skl *SkipList) keyIsAfterNode(key Key, node *SkipNode) bool {
	return node != nil && skl.cmp.Compare(key, node.key) > 0
}

func (skl *SkipList) findGreaterOrEqual(key Key, prev []*SkipNode) *SkipNode{
	skl.mu.RLock()
	defer skl.mu.RUnlock()

	gNode := skl.head
	level := skl.GetMaxHeight()-1
	for true{
		next := gNode.GetNext(level)
		if skl.keyIsAfterNode(key, next){
			gNode = next
		}else{
			if prev != nil {prev[level] = gNode}
			if level == 0{
				return next
			}else{
				level -= 1
			}
		}
	}
	return nil
}

//Insert insert a new node into skipList
func (skl *SkipList) Insert(key Key) {
	skl.mu.Lock()
	skl.mu.Unlock()
	curHeight := skl.randomHeight()

	prev := make([]*SkipNode, sklMaxHeight)
	if curHeight > skl.GetMaxHeight(){
		for i := skl.GetMaxHeight(); i < curHeight; i++ {
			prev[i] = skl.head
		}
		atomic.StoreInt32(&skl.maxHeight, curHeight)
	}

	gNode := skl.findGreaterOrEqual(key, prev)

	if gNode == nil || (skl.cmp.Compare(key, gNode.key) < 0){
		newNode := newSkipNode(key, curHeight)
		for i:= int32(curHeight-1); i >= 0; i--{
			newNode.SetNext(i, prev[i].GetNext(i))
			prev[i].SetNext(i, newNode)
		}
	}
	//skl.print()
}

func (skl* SkipList) Find(key Key) *SkipNode{
	skl.mu.Lock()
	skl.mu.Unlock()
	firNode := skl.head
	var curNode *SkipNode = nil
	for curNode = firNode.GetNext(0); skl.cmp.Compare(curNode.Get(), key) < 0; curNode = curNode.GetNext(0){

	}
	if skl.cmp.Compare(curNode.Get(), key)==0{
		return curNode
	}else{
		return nil
	}
}

func (skl *SkipList) print(){
	//height := skl.GetMaxHeight()
	for i := sklMaxHeight-1 ; i>= 0; i--{
		firNode := skl.head
		fmt.Printf("%2d  | ", i)
		for curNode := firNode.GetNext(i); curNode!=nil ;curNode = curNode.GetNext(i){
			fmt.Print(curNode.key[:])
			fmt.Print("["+fmt.Sprint(len(curNode.next))+"] -----> ")
		}
		fmt.Println("nil")
		if i == 0 { break }
	}
}

//SkipNode store data in skipnode
type SkipNode struct {
	mu sync.RWMutex
	key  Key         //key of this node
	next []*SkipNode //skipNode* 切片，用来保存该结点每一层的下一个节点的指针
}

//newSkipNode Create a new skipNode
func newSkipNode(k Key, height int32) *SkipNode {
	node := &SkipNode{
		key:  k,
		next: make([]*SkipNode, height),
	}
	for i := range node.next {
		node.next[i] = nil
	}
	return node
}

func (node *SkipNode) Get() []byte{
	node.mu.RLock()
	defer node.mu.RUnlock()
	return node.key
}

func (node *SkipNode) Print(){
	fmt.Printf("key: "+string(node.key)+" level: %v \n", len(node.next))
}

//Next 返回当前结点的下一个结点
func (node *SkipNode) GetNext(n int32) *SkipNode {
	node.mu.RLock()
	defer node.mu.RUnlock()
	if n >= int32(len(node.next)) {
		skipNodeErr := &skipNodeError{
			errorStr: "GetNext: out of range!"+fmt.Sprintf("range: 0 - %v, access: %v", len(node.next)-1, n),
		}
		panic(skipNodeErr)
	} else {
		return node.next[n]
	}
}

//Set 设置下一个结点
func (node *SkipNode) SetNext(n int32, nxtNode *SkipNode) error {
	node.mu.Lock()
	defer node.mu.Unlock()
	if n >= int32(len(node.next)) {
		return &skipNodeError{
			errorStr: "setNext: out of range!"+fmt.Sprintf("range: 0 - %v, access: %v", len(node.next)-1, n),
		}
	} else {
		node.next[n] = nxtNode
		return nil
	}
}


