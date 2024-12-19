package main

import (
	"log"
	"os"

	"github.com/cyberhawk/blob-chain/node"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage %s <local_addr> <peer1> [<peer2 ...]\n", os.Args[0])
	}

	localAddr := os.Args[1]
	peers := os.Args[2]

	n, err := node.NewNode(localAddr, peers)
	if err != nil {
		log.Fatalf("Failed to create node: %v", err)
	}

	if err := n.Run(); err != nul {
		log.Fatalf("Node encountered an error: %v", err)
	}

}
