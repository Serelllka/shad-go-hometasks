package lrucache

type CircularBuffer struct {
	cap               int
	size              int
	buffer            []string
	firstInd, lastInd int
}

func NewBuffer(cap int) (cb CircularBuffer) {
	return CircularBuffer{
		cap:      cap,
		size:     0,
		buffer:   make([]string, cap),
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

func (buf *CircularBuffer) Append(item string) {
	buf.buffer[buf.lastInd] = item
	if buf.size < buf.cap {
		buf.size += 1
	} else {
		buf.firstInd = buf.nextIndex(buf.firstInd)
	}
	buf.lastInd = buf.nextIndex(buf.lastInd)
}

func (buf *CircularBuffer) Out() (output []string) {
	output = make([]string, buf.size)
	index := buf.firstInd
	for index != buf.lastInd {
		output = append(output, buf.buffer[index])
		index = buf.nextIndex(index)
	}
	return output
}

func (buf *CircularBuffer) Front() (item string) {
	return buf.buffer[buf.prevIndex(buf.lastInd)]
}

func (buf *CircularBuffer) Back() (item string) {
	return buf.buffer[buf.firstInd]
}

func (buf *CircularBuffer) linearizePointers() (output []int) {
	output = make([]int, buf.size)
	index := buf.firstInd
	if buf.size == 0 {
		return nil
	}

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
