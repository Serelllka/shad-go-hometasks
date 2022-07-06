package lrucache

type lrucache struct {
	cap     int
	buff    CircularBuffer
	content map[int]int
}

func (c *lrucache) Get(key int) (value int, status bool) {
	value, status = c.content[key]
	return
}

func (c *lrucache) Set(key, value int) {
	if c.cap == 0 {
		return
	}

	if _, ok := c.content[key]; ok {
		c.content[key] = value
		return
	}

	if len(c.content) == c.cap {
		if item, ok := c.buff.Front(); ok {
			delete(c.content, item)
		}
	}

	c.content[key] = value
	c.buff.Append(key)
}

func (c *lrucache) Range(f func(key, value int) bool) {
	for _, index := range c.content {

	}
	panic("implement me")
}

func (c *lrucache) Clear() {
	panic("implement me")
}

func New(cap int) Cache {
	return &lrucache{
		cap:     cap,
		buff:    NewBuffer(cap),
		content: make(map[int]int),
	}
}
