package btree

func readVarInt(buf []byte) (uint64, int) {
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

func bytesToInt(bytes []byte) uint64 {
	var result uint64
	for _, b := range bytes {
		result = (result << 8) | uint64(b)
	}
	return result
}
