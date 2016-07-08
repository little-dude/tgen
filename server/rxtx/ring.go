package rxtx

// kind of ring buffer implementation
// inspired by https://github.com/zfjagann/golang-ring/blob/master/ring.go
type Ring struct {
	In       chan []*RawPacket
	Out      chan []*RawPacket
	head     int
	tail     int
	len      int
	capacity int
	buff     [][]*RawPacket
	done     chan bool
}

func (r *Ring) run() {
	defer func() {
		close(r.Out)
		r.done <- true
	}()

loop:
	for {
		// give priority to writes, and only read from the ring buffer if there
		// is nothing to write
		select {
		case data := <-r.In: // data incoming: add it to the buffer and continue
			if data == nil { // r.In is closed, exit
				break loop
			}
			r.set(data)
		default: // no data incoming, let see if there is something to read and if someone wants to read it
			if r.tail < r.head { // there is something to read
				select {
				case r.Out <- r.peek(): // a goroutine is reading from r.Out
					r.get()
				default: // nobody is reading from r.Out, do nothing and continue
				}
			} else { // nothing to read, the buffer is empty
				// wait for data to come, so that we do not consume CPU
				data := <-r.In
				if data == nil { // r.In is closed, exit
					break loop
				}
				r.set(data)
			}
		}
	}
	for r.head >= r.tail {
		r.Out <- r.get()
	}
}

func (r *Ring) peek() (v []*RawPacket) {
	return r.buff[r.tail%r.len]
}

func (r *Ring) set(v []*RawPacket) {
	if r.head-r.tail == r.len-1 {
		r.resize(r.len * 4)
	}
	r.head = r.head + 1
	r.buff[r.head%r.len] = v
}

func (r *Ring) get() (v []*RawPacket) {
	v = r.buff[r.tail%r.len]
	r.tail = r.tail + 1
	// shrinking is expensive for big buffers, we don't do it too often
	if r.len > r.capacity && r.head-r.tail <= r.len/10 {
		r.resize(r.len / 5)
	}
	return v
}

func (r *Ring) resize(size int) {
	newbuf := make([][]*RawPacket, size)
	t := r.tail % r.len
	h := r.head % r.len
	// note: extend is normally called before t == h
	if t >= h {
		copy(newbuf, r.buff[t:])
		copy(newbuf[r.len-t:], r.buff[:h+1])
		r.head = r.len - t + h
	} else {
		copy(newbuf, r.buff[t:h+1])
		r.head = h - t
	}
	r.buff = newbuf
	r.len = size
	r.tail = 0
}

func NewRingBuf(capacity int) *Ring {
	r := Ring{
		buff:     make([][]*RawPacket, capacity),
		capacity: capacity,
		len:      capacity,
		head:     -1,
		tail:     0,
		In:       make(chan []*RawPacket, capacity/10),
		Out:      make(chan []*RawPacket, capacity/10),
		done:     make(chan bool),
	}
	go r.run()
	return &r
}

func (r *Ring) Close() {
	close(r.In)
	<-r.done
}

func (r *Ring) Kill() {
	close(r.In)
	for _ = range r.Out {
	}
	<-r.done
}
