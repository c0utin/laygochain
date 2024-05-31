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

	assert.Equal(t, trA.peers[trB.localAddr], trB)
	assert.Equal(t, trB.peers[trA.localAddr], trA)


}

func TestSendMessage(t *testing.T) {
	trA := NewLocalTransport("A")
	trB := NewLocalTransport("B")

	trA.Connect(trB)
	trB.Connect(trA)

	msg := []byte("kaizoku ni orewa naru")
	assert.Nil(t, trA.SendMessage(trB.localAddr, msg)) 

	rpc := <-trB.Consume()
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, trA.localAddr)

}
