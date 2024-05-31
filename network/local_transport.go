package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddress
	consumeCh chan RPC
	lock      sync.RWMutex
	peers     map[NetAddress]*LocalTransport
}

func NewLocalTransport(addr NetAddress) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddress]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	peer, ok := tr.(*LocalTransport)
	if !ok {
		return fmt.Errorf("invalid transport type")
	}

	t.peers[tr.Addr()] = peer
	return nil
}

func (t *LocalTransport) SendMessage(to NetAddress, payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s: could not send message to %s", t.addr, to)
	}

	peer.consumeCh <- RPC{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}

func (t *LocalTransport) Addr() NetAddress {
	return t.addr
}
