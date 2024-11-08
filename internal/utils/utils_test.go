package utils

import (
	"testing"
)

func TestVarInt(t *testing.T) {
	var testVal uint64 = 0b1111101000
	buf := []byte{0b10000111, 0b01101000}
	decodedVal, bytesRead := ReadVarInt(buf)

	if testVal != decodedVal {
		t.Errorf(`Bytes Read = %d
		Decoded Value = %b 
		Actual Value = %b`,
			bytesRead, decodedVal, testVal)
	}
}

func TestBytesToInt(t *testing.T) {
	var testVal uint64 = 0b1111101000
	//
	buf := []byte{0b11, 0b11101000}
	decodedVal := BytesToInt(buf)

	if testVal != decodedVal {
		t.Errorf(`
		Decoded Value = %b 
		Actual Value = %b`,
			decodedVal, testVal)
	}
}
