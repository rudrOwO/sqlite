package utils

import "log"

func ReadVarInt(buf []byte) (uint64, int) {
	var result uint64
	for i, b := range buf {
		result <<= 7
		result |= uint64(b & 0x7f)
		if b&0x80 == 0 {
			return result, i + 1
		}
	}
	return result, 0
}

func BytesToInt(bytes []byte) uint64 {
	var result uint64
	for _, b := range bytes {
		result = (result << 8) | uint64(b)
	}
	return result
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
