package relay

import (
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/multiformats/go-multiaddr"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Relay struct {
	host   host.Host
	stopCh chan os.Signal
}

func New() *Relay {
	return &Relay{
		host:   nil,
		stopCh: make(chan os.Signal, 1),
	}
}

func (r *Relay) Start() {
	var err error
	r.host, err = libp2p.New(libp2p.EnableRelay())
	if err != nil {
		log.Fatalf("Failed to create relay node: %v", err)
	}

	relayAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", r.host.ID().String()))
	for _, addr := range r.host.Addrs() {
		log.Default().Println(addr.String())
	}

	log.Default().Printf("Relay node: %s\n", r.host.ID().String())
	log.Default().Printf("Relay address: %s\n", relayAddr.String())
	log.Default().Println("Relay complete address: ", r.host.Addrs()[0].Encapsulate(relayAddr).String())

	signal.Notify(r.stopCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-r.stopCh:
		r.Stop()
	}
}

func (r *Relay) Stop() {
	if r.host != nil {
		_ = r.host.Close()
	}
}
