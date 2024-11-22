/*
Based on:

	https://www.sqlite.org/datatype3.html
	https://www.sqlite.org/fileformat.html#varint
	https://www.sqlite.org/fileformat.html#record_format

Datatypes correspondence:

	VARINT => unint64
	INTEGER => int64?
	NULL => nil
	TEXT => string
	BLOB => []byte
*/
package dataformat

type DeserializedTypes interface {
	int64 | float64 | string
}

func DeserializeVarint(bytes []byte) (uint64, uint16) {
	var result uint64
	for i, b := range bytes {
		result <<= 7
		result |= uint64(b & 0x7f)
		if b&0x80 == 0 {
			return result, uint16(i + 1)
		}
	}
	return result, 0
}

func DeserializeInteger(bytes []byte) (result uint64) { // FIX  Should be int64
	for _, b := range bytes {
		result = (result << 8) | uint64(b)
	}
	return result
}

func DeserializeFloat(bytes []byte) (result float64) {
	// TODO
	return result
}