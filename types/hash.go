package types

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Hash [32]uint8

func (h Hash) IsEmpty() bool {
	for i := 0; i < 32; i++ {
		if h[i] != 0 {
			return false
		}
	}
	return true
}

func HashFromByte(b []byte) Hash {

	if len(b) != 32 {
		msg := fmt.Sprintf("the byte length should be 32, given %d", len(b))
		panic(msg)
	}

	var value [32]uint8
	for i := 0; i <32; i++{
		value[i] = b[i]
	}

	return Hash(value)
}

func(h Hash) ToString() string {

	return hex.EncodeToString(h.ToSlice())
}

func(h Hash) ToSlice() []byte {

	b:= make([]byte, 32)
	for i := 0; i < 32; i++{
		b[i] = h[i]
	}
	return b
}

func RandomBytes(size int)[]byte {

	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() Hash{

	return HashFromByte(RandomBytes(32))
}
