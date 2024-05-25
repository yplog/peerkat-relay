package main

import (
	"github.com/yplog/peerkat-relay/pkg/relay"
)

func main() {
	relayNode := relay.New()

	relayNode.Start()
}
