package main

import (
	"github.com/yplog/peerkat-relay/pkg/relay"
	"log"
)

func main() {
	log.Println("Starting relay node...")
	relayNode := relay.New()

	relayNode.Start()
}
