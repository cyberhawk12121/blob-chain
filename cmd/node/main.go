package main

import (
	"log"
	"os"

	"github.com/cyberhawk/blob-chain/internal/node"
)

func main() {
	//
	if len(os.Args) < 3 {
		log.Fatalf("Usage %s <local_addr> <peer1> [<peer2 ...]\n", os.Args[0])
	}

	localAddr := os.Args[1] // first one is the local address, this node is running on and listening
	peers := os.Args[2:]    // All the rest are the peers

	n, err := node.NewNode(localAddr, peers)
	if err != nil {
		log.Fatalf("Failed to create node: %v", err)
	}

	if err := n.Run(); err != nil {
		log.Fatalf("Node encountered an error: %v", err)
	}

	// Keep the program running so it can continue listening on the port.
	select {}
}
