package main

import (
	"time"

	"github.com/c0utin/laygochain/network"
)

func main(){
	
	trLocal := network.NewLocalTransport("LOCAL") 

	// demonstration
	trRemote := network.NewLocalTransport("REMOTE") 

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func(){
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("hello from REMOTE"))
			time.Sleep(3 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	server := network.NewServer(opts)
	server.Start()

}

