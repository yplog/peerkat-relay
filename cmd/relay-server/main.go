package main

import (
	"github.com/yplog/peerkat-relay/pkg/relay"
	"log"
)

func main() {
	log.Default().Println("Starting relay node...")
	relayNode := relay.New()

	relayNode.Start()
}
