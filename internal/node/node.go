package node

type Node struct {
	repo      Repository
	service   Service
	transport Transport
}

func NewNode(localAddr string, peers []string) (*Node, error) {
	repo := NewInMemoryRepository(peers)
	transportStub := &udpTransport{localAddr: localAddr}
	service := NewService(repo, transportStub)
	transport := NewUDPTransport(localAddr, service)

	transportStub.service = service

	n := &Node{
		repo:      repo,
		service:   service,
		transport: transport,
	}

	return n, nil
}

func (n *Node) Run() error {
	return n.transport.StartListening()
}
