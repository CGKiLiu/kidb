package arena

type sizeT uint64

const blockSize sizeT = 4096

type bufferBlock struct {
	block  []byte
	off    sizeT
	remain sizeT
	next   *bufferBlock
}

func newBlock() *bufferBlock {
	bufBlock := &bufferBlock{
		block:  make([]byte, blockSize),
		off:    0,
		remain: blockSize,
		next:   nil,
	}
	return bufBlock
}

func (bufBlock *bufferBlock) String() string {
	return string(bufBlock.block[:bufBlock.off])
}

func (bufBlock *bufferBlock) blockAlloc(bytes sizeT) ([]byte, error) {
	if bytes > bufBlock.remain {
		return nil, &blockAllocError{
			errorStr: "Insufficient remaining space",
		}
	}
	bufBlock.off += bytes
	bufBlock.remain -= bytes
	return bufBlock.block[bufBlock.off-bytes : bufBlock.off], nil
}

//Arena ...
type Arena struct {
	head   *bufferBlock
	tail   *bufferBlock
	blocks uint64
}

func newArena() *Arena {
	arena := &Arena{
		head:   nil,
		tail:   nil,
		blocks: uint64(0),
	}
	return arena
}

//Alloc ...
func (arena *Arena) Alloc(bytes sizeT) ([]byte, error) {
	if arena.tail == nil {
		arena.tail = newBlock()
		arena.head = arena.tail
	}
	if arena.tail.remain < bytes {
		return arena.tail.blockAlloc(bytes)
	} else {
		newBlock := newBlock()
		arena.tail.next = newBlock
		arena.tail = newBlock
		return arena.tail.blockAlloc(bytes)
	}
}
