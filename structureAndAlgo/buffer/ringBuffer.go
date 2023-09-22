package buffer

type RingBuffer[T any] struct {
	buf        []T
	readPoint  int
	writePoint int
	isFull     bool
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buf: make([]T, size),
	}
}

func (r *RingBuffer[T]) Write(data []T) int {
	if len(data) == 0 || r.Writeable() == 0 {
		return 0
	}

	var writed int

	if r.writePoint >= r.readPoint {
		// 버퍼 마지막 까지 쓸수있는 범위
		writeableToEnd := len(r.buf) - r.writePoint
		writed = min(writeableToEnd, len(data))
	} else {
		writed = min(r.Writeable(), len(data))
	}

	copy(r.buf[r.writePoint:], data[:writed])
	r.writePoint += writed
	r.writePoint %= len(r.buf)

	// isFull?
	if writed > 0 && r.writePoint == r.readPoint {
		r.isFull = true
	}

	remain := len(data) - writed
	if remain > 0 && r.Writeable() > 0 {
		writed += r.Write(data[writed:])
	}
	return writed
}

func (r *RingBuffer[T]) Read(count int) []T {
	if r.ReadAble() == 0 || count == 0 {
		return []T{}
	}
	readCount := min(count, r.ReadAble())
	rst := make([]T, readCount)

	// 읽기 시작포인트 + 읽을려는 갯수가 최대 용량을 넘을경
	if r.readPoint+readCount >= len(r.buf) {
		// 읽기 시작 포인트와 버퍼크기의 차이
		remainUntilEnd := len(r.buf) - r.readPoint
		copy(rst, r.buf[r.readPoint:])
		r.readPoint = 0

		//나머지
		remain := readCount - remainUntilEnd
		copy(rst[remainUntilEnd:], r.buf[:remain])

		r.readPoint += remain
	} else {
		copy(rst, r.buf[r.readPoint:r.readPoint+readCount])
		r.readPoint += readCount
	}
	r.isFull = false
	return rst
}

func (r *RingBuffer[T]) ReadAble() int {
	if r.isFull {
		return len(r.buf)
	}
	if r.writePoint < r.readPoint {
		return len(r.buf) - r.readPoint + r.writePoint
	}
	return r.writePoint - r.readPoint
}

func (r *RingBuffer[T]) Writeable() int {
	return len(r.buf) - r.ReadAble()
}
