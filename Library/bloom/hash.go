package main

import (
	"github.com/spaolacci/murmur3"
)

// hash 函数选取采用 murmur3 方便加偏移量扩展为多个 hash 函数
func Hash(data []byte) uint64 {
	return murmur3.Sum64(data)
}
