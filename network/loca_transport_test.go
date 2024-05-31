package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	trA := NewLocalTransport("A")
	trB := NewLocalTransport("B")

	trA.Connect(trB)
	trB.Connect(trA)

	// Verifica se os peers foram conectados corretamente
	assert.NotNil(t, trA.Peers(trB.Addr()))
	assert.NotNil(t, trB.Peers(trA.Addr()))
}

func TestSendMessage(t *testing.T) {
	trA := NewLocalTransport("A")
	trB := NewLocalTransport("B")

	trA.Connect(trB)
	trB.Connect(trA)

	msg := []byte("kaizoku ni orewa naru")
	assert.Nil(t, trA.SendMessage(trB.Addr(), msg))

	rpc := <-trB.Consume()
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, trA.Addr())
}

