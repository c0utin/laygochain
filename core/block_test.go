package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/c0utin/laygochain/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T){

	h := &Header{
		Version: 1,
		PrevBlock: types.RandomHash(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height: 1,
		Nonce: 696969,
	}


	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}

func TestBlock_Encode_Decode(t *testing.T) {

	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     989394,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)
}

func TestBlock_Hash(t *testing.T){

	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     989394,
		},
		Transactions: nil,
	}

	h := b.Hash()
	fmt.Println(h.ToString())
	assert.False(t, h.IsEmpty())
}









