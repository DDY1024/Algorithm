package other

const (
	mask = 0x07
)

type Bitmap struct {
	data []byte
	size int
}

func NewBitMap(size int) *Bitmap {
	size = (size + 8 - 1) / 8 * 8
	bitmap := &Bitmap{
		size: size,
		data: make([]byte, size/8),
	}
	return bitmap
}

func NewBitMapFromData(data []byte) *Bitmap {
	bitmap := &Bitmap{
		size: len(data) * 8,
		data: data,
	}
	return bitmap
}

func (b *Bitmap) Set(pos int) bool {
	if pos >= b.size {
		return false
	}
	b.data[pos>>3] |= 1 << (pos & mask)
	return true
}

func (b *Bitmap) Unset(pos int) bool {
	if pos >= b.size {
		return false
	}
	b.data[pos>>3] &^= 1 << (pos & mask)
	return true
}

func (b *Bitmap) IsSet(pos int) bool {
	if pos >= b.size {
		return false
	}
	return b.data[pos>>3]&(1<<(pos&mask)) > 0
}

func (b *Bitmap) Resize(size int) {
	size = (size + 7) / 8 * 8
	if b.size == size {
		return
	}
	data := make([]byte, size/8, size/8)
	copy(data, b.data)
	b.data = data
	b.size = size
}

func (b *Bitmap) Size() int {
	return b.size
}

func (b *Bitmap) Clear() {
	b.data = make([]byte, b.size>>3)
}
