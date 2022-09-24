package utils

import (
	"errors"
	"sync/atomic"
)

// UnboundedChan is an unbounded chan.
// In is used to write without blocking, which supports multiple writers.
// and Out is used to read, which supports multiple readers.
// You can close the in channel if you want.
type UnboundedChan[T any] struct {
	bufCount int64
	In       chan<- T       // channel for write
	Out      <-chan T       // channel for read
	buffer   *RingBuffer[T] // buffer
}

// Len returns len of In plus len of Out plus len of buffer.
// It is not accurate and only for your evaluating approximate number of elements in this chan,
// see https://github.com/smallnest/chanx/issues/7.
func (c UnboundedChan[T]) Len() int {
	return len(c.In) + c.BufLen() + len(c.Out)
}

// BufLen returns len of the buffer.
// It is not accurate and only for your evaluating approximate number of elements in this chan,
// see https://github.com/smallnest/chanx/issues/7.
func (c UnboundedChan[T]) BufLen() int {
	return int(atomic.LoadInt64(&c.bufCount))
}

// NewUnboundedChan creates the unbounded chan.
// in is used to write without blocking, which supports multiple writers.
// and out is used to read, which supports multiple readers.
// You can close the in channel if you want.
func NewUnboundedChan[T any](initCapacity int) *UnboundedChan[T] {
	return NewUnboundedChanSize[T](initCapacity, initCapacity, initCapacity)
}

// NewUnboundedChanSize is like NewUnboundedChan, but you can set initial capacity for In, Out, Buffer.
func NewUnboundedChanSize[T any](initInCapacity, initOutCapacity, initBufCapacity int) *UnboundedChan[T] {
	in := make(chan T, initInCapacity)
	out := make(chan T, initOutCapacity)
	ch := UnboundedChan[T]{In: in, Out: out, buffer: NewRingBuffer[T](initBufCapacity)}

	go process(in, out, &ch)

	return &ch
}

func process[T any](in, out chan T, ch *UnboundedChan[T]) {
	defer close(out)
loop:
	for {
		val, ok := <-in
		if !ok { // in is closed
			break loop
		}

		// make sure values' order
		// buffer has some values
		if atomic.LoadInt64(&ch.bufCount) > 0 {
			ch.buffer.Write(val)
			atomic.AddInt64(&ch.bufCount, 1)
		} else {
			// out is not full
			select {
			case out <- val:
				continue
			default:
			}

			// out is full
			ch.buffer.Write(val)
			atomic.AddInt64(&ch.bufCount, 1)
		}

		for !ch.buffer.IsEmpty() {
			select {
			case val, ok := <-in:
				if !ok { // in is closed
					break loop
				}
				ch.buffer.Write(val)
				atomic.AddInt64(&ch.bufCount, 1)

			case out <- ch.buffer.Peek():
				ch.buffer.Pop()
				atomic.AddInt64(&ch.bufCount, -1)
				if ch.buffer.IsEmpty() && ch.buffer.size > ch.buffer.initialSize { // after burst
					ch.buffer.Reset()
					atomic.StoreInt64(&ch.bufCount, 0)
				}
			}
		}
	}

	// drain
	for !ch.buffer.IsEmpty() {
		out <- ch.buffer.Pop()
		atomic.AddInt64(&ch.bufCount, -1)
	}

	ch.buffer.Reset()
	atomic.StoreInt64(&ch.bufCount, 0)
}

var ErrIsEmpty = errors.New("ringbuffer is empty")

// RingBuffer is a ring buffer for common types.
// It is never full and always grows if it will be full.
// It is not thread-safe(goroutine-safe) so you must use the lock-like synchronization primitive to use it in multiple writers and multiple readers.
type RingBuffer[T any] struct {
	buf         []T
	initialSize int
	size        int
	r           int // read pointer
	w           int // write pointer
}

func NewRingBuffer[T any](initialSize int) *RingBuffer[T] {
	if initialSize <= 0 {
		panic("initial size must be great than zero")
	}
	// initial size must >= 2
	if initialSize == 1 {
		initialSize = 2
	}

	return &RingBuffer[T]{
		buf:         make([]T, initialSize),
		initialSize: initialSize,
		size:        initialSize,
	}
}

func (r *RingBuffer[T]) Read() (T, error) {
	var t T
	if r.r == r.w {
		return t, ErrIsEmpty
	}

	v := r.buf[r.r]
	r.r++
	if r.r == r.size {
		r.r = 0
	}

	return v, nil
}

func (r *RingBuffer[T]) Pop() T {
	v, err := r.Read()
	if err == ErrIsEmpty { // Empty
		panic(ErrIsEmpty.Error())
	}

	return v
}

func (r *RingBuffer[T]) Peek() T {
	if r.r == r.w { // Empty
		panic(ErrIsEmpty.Error())
	}

	v := r.buf[r.r]
	return v
}

func (r *RingBuffer[T]) Write(v T) {
	r.buf[r.w] = v
	r.w++

	if r.w == r.size {
		r.w = 0
	}

	if r.w == r.r { // full
		r.grow()
	}
}

func (r *RingBuffer[T]) grow() {
	var size int
	if r.size < 1024 {
		size = r.size * 2
	} else {
		size = r.size + r.size/4
	}

	buf := make([]T, size)

	copy(buf[0:], r.buf[r.r:])
	copy(buf[r.size-r.r:], r.buf[0:r.r])

	r.r = 0
	r.w = r.size
	r.size = size
	r.buf = buf
}

func (r *RingBuffer[T]) IsEmpty() bool {
	return r.r == r.w
}

// Capacity returns the size of the underlying buffer.
func (r *RingBuffer[T]) Capacity() int {
	return r.size
}

func (r *RingBuffer[T]) Len() int {
	if r.r == r.w {
		return 0
	}

	if r.w > r.r {
		return r.w - r.r
	}

	return r.size - r.r + r.w
}

func (r *RingBuffer[T]) Reset() {
	r.r = 0
	r.w = 0
	r.size = r.initialSize
	r.buf = make([]T, r.initialSize)
}
