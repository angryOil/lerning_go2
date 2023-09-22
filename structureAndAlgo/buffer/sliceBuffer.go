package buffer

type SliceBuffer[T any] struct {
	buf []T
}

func NewSliceBuffer[T any]() *SliceBuffer[T] {
	return &SliceBuffer[T]{}
}

func (sb *SliceBuffer[T]) Write(data []T) {
	sb.buf = append(sb.buf, data...)
}

func (sb *SliceBuffer[T]) ReadAble() int {
	return len(sb.buf)
}

func (sb *SliceBuffer[T]) Read(count int) []T {
	readCount := min(count, sb.ReadAble())
	rst := make([]T, readCount)
	copy(rst, sb.buf[:readCount])
	sb.buf = sb.buf[readCount:]
	return rst
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
