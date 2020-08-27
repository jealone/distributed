package distributed

import (
	"hash/crc32"
)

type Partition interface {
	Partition([]byte, int) int
}

type CheckSum func([]byte) uint32

func ModularHashIEEE() *ModularHash {
	return NewModularHash(crc32.ChecksumIEEE)
}

func NewModularHash(check CheckSum) *ModularHash {
	return &ModularHash{
		checkSum: check,
	}
}

type ModularHash struct {
	checkSum CheckSum
}

func (part *ModularHash) Partition(key []byte, N int) int {
	return int(part.checkSum(key)) % N
}

type KetamaHash struct {
}
