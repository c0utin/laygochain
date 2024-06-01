package types

import (
	"crypto/rand"
	"fmt"
)

type Hash [32]uint8

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

func RandomBytes(size int)[]byte {

	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() Hash{

	return HashFromByte(RandomBytes(32))
}
