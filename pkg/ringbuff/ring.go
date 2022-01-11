package ringbuff

type RingBuff struct {
	size int      // current size
	max  int      // max size
	head int      // start index
	buff []string // buffer
}

func NewRingBuff(size int) *RingBuff {
	b := make([]string, size)
	return &RingBuff{
		size: 0,
		head: 0,
		max:  size,
		buff: b,
	}

}

func (rb *RingBuff) GetSize() int {
	return rb.size
}

func (rb *RingBuff) IsEmpty() bool {
	return rb.size == 0
}

func (rb *RingBuff) Get(index int) (string, bool) {
	if index > rb.size {
		// can't be true
		return "", false
	}

	i := (rb.head + index) % rb.max

	return rb.buff[i], true
}

func (rb *RingBuff) Add(s string) {
	i := (rb.head + rb.size) % rb.max
	rb.buff[i] = s

	if rb.size >= rb.max {
		rb.head = (rb.head + 1) % rb.max
	} else {
		rb.size++
	}
}
