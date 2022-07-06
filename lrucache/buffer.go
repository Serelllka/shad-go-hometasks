package lrucache

type CircularBuffer struct {
	cap               int
	size              int
	buffer            []int
	firstInd, lastInd int
}

func NewBuffer(cap int) (cb CircularBuffer) {
	return CircularBuffer{
		cap:      cap,
		size:     0,
		buffer:   make([]int, cap),
		firstInd: 0,
		lastInd:  0,
	}
}

func (buf *CircularBuffer) nextIndex(index int) (newIndex int) {
	index++
	if index == buf.cap {
		index = 0
	}
	return index
}

func (buf *CircularBuffer) prevIndex(index int) (newIndex int) {
	if index == 0 {
		index = buf.cap
	}
	index--
	return index
}

func (buf *CircularBuffer) Append(items ...int) {
	if buf.cap == 0 {
		return
	}

	for _, item := range items {
		buf.buffer[buf.lastInd] = item
		if buf.size < buf.cap {
			buf.size += 1
		} else {
			buf.firstInd = buf.nextIndex(buf.firstInd)
		}
		buf.lastInd = buf.nextIndex(buf.lastInd)
	}
}

func (buf *CircularBuffer) Out() (output []int) {
	output = make([]int, 0, buf.size)
	for _, item := range buf.linearizeIndices() {
		output = append(output, buf.buffer[item])
	}
	return output
}

func (buf *CircularBuffer) Front() (item int, out bool) {
	if len(buf.buffer) == 0 {
		return 0, false
	}

	return buf.buffer[buf.prevIndex(buf.lastInd)], true
}

func (buf *CircularBuffer) Back() (item int, out bool) {
	if len(buf.buffer) == 0 {
		return 0, false
	}

	return buf.buffer[buf.firstInd], true
}

func (buf *CircularBuffer) linearizeIndices() (output []int) {
	if len(buf.buffer) == 0 {
		return []int{}
	}

	output = make([]int, 0, buf.size)
	index := buf.firstInd

	if buf.size < buf.cap {
		for index != buf.lastInd {
			output = append(output, index)
			index = buf.nextIndex(index)
		}
	} else {
		tmp := index
		for {
			output = append(output, index)
			index = buf.nextIndex(index)
			if index == tmp {
				break
			}
		}
	}
	return
}
