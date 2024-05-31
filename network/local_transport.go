package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct{
	localAddr	NetAddress
	consumeCh	chan	RPC
	lock	sync.RWMutex
	peers	map[NetAddress]*LocalTransport
}

func NewLocalTransport(localAddr NetAddress) *LocalTransport {

	return &LocalTransport{
		localAddr: localAddr,
		consumeCh: make(chan RPC, 1024),
		peers: make(map[NetAddress]*LocalTransport),
	}
} 

func (t *LocalTransport) Consume() <-chan RPC {

	return t.consumeCh
}

func (t *LocalTransport) Connect(tr *LocalTransport) error {

	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr

	return nil
}

func(t *LocalTransport) SendMessage(destiny NetAddress, payload []byte) error {

	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[destiny]
	if !ok {
		return fmt.Errorf("%s: could not set message to address %s", t.localAddr,  destiny)
	}

	peer.consumeCh <- RPC{
		From: t.localAddr,
		Payload: payload,
	}

	return nil
}

func (t *LocalTransport) Addr() NetAddress {
	return t.localAddr
}


