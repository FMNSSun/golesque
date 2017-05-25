package golesque

import "container/list"

// A buffer is a structure backed by an array that has space reserved both
// at the front and at the back to allow for append and prepend without too
// much reallocation and copying. Of course, once the space at the back or
// at the front is exceeded reallocation and copying becomes necessary.
type Buffer struct {
	buf        []*GLSQObj
	startIndex int // where is the start?
	length     int // current length.
	spaceFront int // when reallocating this will indicate how much extra space to allocate at the front
	spaceBack  int // when reallocating this will indicate how much extra space to allocate at the back
	maxLength  int // maximum length supported with current space
}

func NewBuffer(expectedSize int, spaceFront int, spaceBack int) *Buffer {
	bufSize := expectedSize + spaceFront + spaceBack

	buf := &Buffer{
		buf:        make([]*GLSQObj, bufSize),
		startIndex: spaceFront,
		length:     expectedSize,
		spaceFront: spaceFront,
		spaceBack:  spaceBack,
		maxLength:  bufSize,
	}

	return buf
}

func (buf *Buffer) ToList() *list.List {
	ls := list.New()
	for i := 0; i < buf.length; i++ {
		ls.PushBack(buf.buf[i+buf.startIndex])
	}

	return ls
}

func (buf *Buffer) Prepend(obj *GLSQObj) {
	if buf.startIndex == 0 {
		panic("not enough space")
		//TODO: reallocate
		return
	}

	buf.startIndex--
	buf.buf[buf.startIndex] = obj
	buf.length++
}

func (buf *Buffer) Append(obj *GLSQObj) {
	if buf.length >= buf.maxLength {
		panic("not enough space")
		//TODO: reallocate
		return
	}

	idx := buf.startIndex + buf.length
	buf.buf[idx] = obj
	buf.length++
}
