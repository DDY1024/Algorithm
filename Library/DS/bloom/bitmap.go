package main

// 字节 bitmap 实现
type Bitmap struct {
	data []byte
	size uint64
}

// (a + b - 1) / b
func NewBitMap(size uint64) *Bitmap {
	size = (size + 7) / 8 * 8 // 对齐校准
	bitmap := &Bitmap{
		size: size,
		data: make([]byte, size/8, size/8),
	}
	return bitmap
}

func NewBitMapFromData(data []byte) *Bitmap {
	bitmap := &Bitmap{
		size: uint64(len(data)) * 8,
		data: data,
	}
	return bitmap
}

// 8: 0 ~ 7
func (b *Bitmap) Set(pos uint64) bool {
	if pos >= b.size {
		return false
	}

	b.data[pos>>3] |= 1 << (pos & 0x07)
	return true
}

func (b *Bitmap) Unset(pos uint64) bool {
	if pos >= b.size {
		return false
	}

	b.data[pos>>3] &^= 1 << (pos & 0x07)
	return true
}

func (b *Bitmap) IsSet(pos uint64) bool {
	if pos >= b.size {
		return false
	}

	if b.data[pos>>3]&(1<<(pos&0x07)) > 0 {
		return true
	}

	return false
}

func (b *Bitmap) Resize(size uint64) {
	size = (size + 7) / 8 * 8
	if b.size == size {
		return
	}
	data := make([]byte, size/8, size/8)
	copy(data, b.data) // 1. dst，src
	b.data = data
	b.size = size
}

func (b *Bitmap) Size() uint64 {
	return b.size
}

func (b *Bitmap) Clear() {
	b.data = make([]byte, b.size/8, b.size/8)
}

func (b *Bitmap) Data() []byte {
	return b.data
}
